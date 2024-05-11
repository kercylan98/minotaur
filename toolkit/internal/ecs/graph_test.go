package ecs

import (
	"fmt"
	"testing"
)

func TestQueryComponent(t *testing.T) {

	var graph graphNode

	graph.AddArchetype(1, 2, 3)

	fmt.Println(graph)
}
