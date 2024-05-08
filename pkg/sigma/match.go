package sigma

import (
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

// Operation Operational data of event
func Operation(attr,operation,value string,event Event) bool {

	// if not contain this attribute then return false
	if !event.CheckEventFieldName(attr) {
		return false
	}

	strategy := StrategyFactory(operation)

	if strategy == nil {

		return false
	}
	ok := strategy.Execute(event.FieldByName(attr),value)

	return ok
}


