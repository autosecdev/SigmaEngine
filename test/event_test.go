package test

import (
	"fmt"
	"github.com/autosecdev/SigmaEngine/pkg/sigma"
	"testing"
)

func TestCheckEventFieldName(test *testing.T) {

	event := sigma.Event{}

	fmt.Println(event.CheckEventFieldName("CommandLine"))
}

func TestOperation(test *testing.T) {
	fmt.Println(sigma.StrategyFactory("contains").Execute("zhangsan","x"))
}

func TestFieldByName(test *testing.T) {
	event := sigma.Event{
		Image: "test",
		CommandLine: "test",
		ParentImage: "test",
	}

	fmt.Println(event.FieldByName("Imagae") == "a")
}