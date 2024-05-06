package test

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/autosecdev/SigmaEngine/config"
	"github.com/autosecdev/SigmaEngine/pkg/sigma"
	"github.com/autosecdev/SigmaEngine/utils"
	"strings"
	"testing"
)

func TestGetRule(t *testing.T) {

	sigmaConfig := config.SigmaConfig{
		RuleDir: "/Users/fate/Downloads/sigma_all_rules/rules/linux/test/",
	}
	rules, err := sigma.GetRuleList(sigmaConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rule := range rules {

		for k, v := range Extract(rule.Detection) {

			Match(v, k)
		}
	}
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
func Match(detection interface{}, key string) {
	// Split the key only if it contains an operation modifier
	field := key
	operation := ""
	// maybe the key is more |,example: process.command|contains|all or selection_base64|process.command|contains|all
	if strings.Contains(key, "|") {
		parts := strings.Split(key, "|")
		field = parts[0]
		operation = parts[1]
	}

	switch actual := detection.(type) {
	case map[string]interface{}:
		for k, v := range actual {
			fullKey := field + "|" + k
			Match(v, fullKey)
		}
	case []interface{}:
		for _, v := range actual {
			Match(v, key)
		}
	case string:
		fmt.Printf("Match found for field '%s' with operation '%s' on value '%s'\n", field, operation, actual)
	default:
		fmt.Printf("Unknown type: %T\n", actual)
	}
}

func TestConvert(t *testing.T) {

	data := utils.ReadFile("/Users/fate/go/GoCode/SigmaEngine/test/data.json")

	jsonParsed, err := gabs.ParseJSON(data)

	if err != nil {

		fmt.Println(err)
		return
	}

	var ok bool
	event := sigma.Event{}

	// Get the event type
	event.Image, ok = jsonParsed.Path("process.executable").Data().(string)

	if !ok {
		fmt.Println("process.executable not found")
		return
	}

	args, ok := jsonParsed.Path("process.args").Data().([]interface{})

	if !ok {
		fmt.Println("process.args not found")
		return
	}

	argsList, err := utils.ToStringSlice(args)

	if err != nil {
		fmt.Println(err)
		return
	}

	event.CommandLine = strings.Join(argsList, " ")

	fmt.Println(event)
}
