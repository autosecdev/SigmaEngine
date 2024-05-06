package test

import (
	"fmt"
	"github.com/autosecdev/SigmaEngine/pkg/sigma"
	"testing"
)

func TestCheckEventFieldName(test *testing.T) {

	event := sigma.Event{}

	fmt.Println(event.CheckEventFieldName("Image"))
}
