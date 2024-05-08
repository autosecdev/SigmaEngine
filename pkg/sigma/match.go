package sigma

import (
	"fmt"
	"strings"
)

func init() {
	// on start initialization all
	registerStrategy("contains", &ContainsStrategy{})
	registerStrategy("endswith", &EndsWithStrategy{})
	registerStrategy("startswith", &StartsWithStrategy{})
}

// OperationStrategy strategy interface
type OperationStrategy interface {
	Execute(text string, subtext string) bool
}

// strategy registry tables
var strategyRegistry = make(map[string]OperationStrategy)

// registerStrategy
func registerStrategy(name string, strategy OperationStrategy) {
	strategyRegistry[name] = strategy
}

// StrategyFactory strategy factory function
func StrategyFactory(operation string) OperationStrategy {
	if strategy, ok := strategyRegistry[operation]; ok {
		return strategy
	}
	return nil
}

// ContainsStrategy contains string
type ContainsStrategy struct{}

func (s *ContainsStrategy) Execute(text string, subtext string) bool {
	return strings.Contains(text, subtext)
}

// EndsWithStrategy end with string
type EndsWithStrategy struct{}

func (s *EndsWithStrategy) Execute(text string, subtext string) bool {
	return strings.HasSuffix(text, subtext)
}

// StartsWithStrategy start with string
type StartsWithStrategy struct{}

func (s *StartsWithStrategy) Execute(text string, subtext string) bool {
	return strings.HasPrefix(text, subtext)
}

func Extract(detection map[string]interface{}) map[string]interface{} {
	tx := make(map[string]interface{})
	for k, v := range detection {
		if k != "condition" {
			tx[k] = v
		}
	}
	return tx
}

// Operation Operational data of event
func Operation(attr,operation,value string,event Event) bool {

	// if not contain this attribute then return false
	if !event.CheckEventFieldName(attr) {
		return false
	}

	ok := StrategyFactory(operation).Execute(event.FieldByName(attr),value)

	return ok
}

// Match the event
func Match(detection interface{}, key string, event *Event) {
	// Split the key only if it contains an operation modifier
	field := key
	operation := ""
	attr := ""
	// maybe the key is more |,example: process.command|contains|all or selection_base64|process.command|contains|all
	if strings.Contains(key, "|") {
		parts := strings.Split(key, "|")
		field = parts[0]
		if len(parts) > 2 && event.CheckEventFieldName(parts[1]) {
			operation = parts[2]
			attr = parts[1]
		} else {
			operation = parts[1]
			attr = parts[0]
		}
	}

	switch actual := detection.(type) {
	case map[string]interface{}:
		for k, v := range actual {
			fullKey := field + "|" + k
			Match(v, fullKey, event)
		}
	case []interface{}:
		for _, v := range actual {
			Match(v, key, event)
		}
	case string:
		fmt.Printf("Match found for field '%s' with operation '%s' at attr is '%s' on value '%s'\n", field, operation, attr, actual)
		if Operation(attr,operation,actual,*event) {
			fmt.Println("Matched!!!!","\t ", field)
		}
	default:
		fmt.Printf("Unknown type: %T\n", actual)
	}
}
