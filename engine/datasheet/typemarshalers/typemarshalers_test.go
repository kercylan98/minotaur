package typemarshalers_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/datasheet"
	"github.com/kercylan98/minotaur/engine/datasheet/typemarshalers"
	"testing"
)

func TestGo(t *testing.T) {
	p := datasheet.NewParser()
	v, err := p.ParseStruct("{ MyStruct | name: string, slice: []int, nested: { name: string }, array: [3]int, myMap: map[string]{name:string} }")
	if err != nil {
		panic(err)
	}

	result, err := typemarshalers.NewGo(typemarshalers.FunctionalGoConfigurator(func(config *typemarshalers.GoConfiguration) {
		config.WithTypePrefix()
	})).Marshal(v)
	if err != nil {
		panic(fmt.Errorf("err: %w\n%s", err, result))
	}

	fmt.Println(result)
}
