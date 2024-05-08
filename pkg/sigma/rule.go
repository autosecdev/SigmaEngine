package sigma

import (
	"fmt"
	"strings"
)

// Rule is a struct that contains information about the Sigma rule
type Rule struct {
	Title          string
	ID             string
	Status         string
	Description    string
	References     []string
	Author         string
	Date           string
	Modified       string
	Tags           []string
	LogSource      LogSource
	Detection      map[string]interface{}
	FalsePositives []string
	Level          string
}

// LogSource is a struct that contains information about the log source
type LogSource struct {
	Product  string
	Category string
}

// Match the event
func (r *Rule) Match(detection interface{}, key string, event *Event) {
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
			r.Match(v, fullKey, event)
		}
	case []interface{}:
		for _, v := range actual {
			r.Match(v, key, event)
		}
	case string:
		//fmt.Printf("Match found for field '%s' with operation '%s' at attr is '%s' on value '%s'\n", field, operation, attr, actual)
		event.Matched = Operation(attr,operation,actual,*event)
	default:
		fmt.Printf("Unknown type: %T\n", actual)
	}
}

// Extract the detection
func (r *Rule) Extract() map[string]interface{} {
	tx := make(map[string]interface{})
	for k, v := range r.Detection {
		if k != "condition" {
			tx[k] = v
		}
	}
	return tx
}