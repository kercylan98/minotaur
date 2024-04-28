package ecs

import (
	"fmt"
	"testing"
)

func TestQueryComponent(t *testing.T) {

	var g = &graph{
		next: make(map[ComponentId]*graph),
		prev: make(map[ComponentId]*graph),
	}

	g.generate([]ComponentId{1, 2, 3}, 0)

	g.Print()
	fmt.Println()
}
