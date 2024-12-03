package datasheetv1

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit"
	"testing"
)

func TestParser(t *testing.T) {
	//input := "{name:string, age:int, list:[]string, scores:[3]int, config:map[string]string}"
	input := "{list:[]string, scores:[3]int}"

	parser := newParser(input)
	parsedType, err := parser.parse()
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	fmt.Println(string(toolkit.MarshalJSON(parsedType)))
}

func TestParseStructInfo(t *testing.T) {
	input := "{name:string, age:int, list:[]string, scores:[3]int, config:map[string]string}"

	result, err := ParseStructInfo(input)
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	fmt.Println(string(toolkit.MarshalJSON(result)))
}
