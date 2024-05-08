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

func MathTest(event sigma.Event) {

	sigmaConfig := config.SigmaConfig{
		RuleDir: "/Users/fate/Downloads/sigma_all_rules/rules/linux/process_creation/",
	}
	rules, err := sigma.LoadRuleAsList(sigmaConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rule := range rules {

		for k, v := range rule.Extract() {

			rule.Match(v, k,&event)

			if event.Matched {
				event.RuleTitle = rule.Title
				fmt.Println("匹配到规则：", event)
				return
			}
		}
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
	MathTest(event)
}
