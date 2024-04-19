package ecs

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/mask"
	"strings"
)

type graph struct {
	mask mask.DynamicMask

	next map[ComponentId]*graph
	prev map[ComponentId]*graph
}

func (g *graph) generate(ids []ComponentId, exclude ComponentId) {
	if len(ids) == 0 {
		return
	}

	// 模拟一下这个函数，当 ids = [1, 2, 3] 时，exclude = 0
	// mask: 0 [Next: 1, 2, 3]
	//   - mask: 1 [Next: 2]
	//       - mask: 1, 2 [Next: 3]
	//           - mask: 1, 2, 3 [Next: ]
	//   - mask: 2 [Next: 1]
	//       - mask: 2, 1 [Next: 3]
	//           - mask: 2, 1, 3 [Next: ]
	//   - mask: 3 [Next: 1]
	//       - mask: 3, 1 [Next: 2]
	//           - mask: 3, 1, 2 [Next: ]

	for _, id := range ids {
		if id == exclude {
			continue
		}

		nextGraph := &graph{
			next: make(map[ComponentId]*graph),
			prev: make(map[ComponentId]*graph),
		}
		nextGraph.mask = g.mask.Clone()
		nextGraph.mask.Set(id)

		// Add next graph to current graph's next
		g.next[id] = nextGraph
		// Add current graph to next graph's prev
		nextGraph.prev[id] = g

		// Recursively generate graphs for remaining ids
		nextGraph.generate(ids, id)
	}

}

func (g *graph) Print() {
	var printGraph func(g *graph, depth int)
	printGraph = func(g *graph, depth int) {
		indent := strings.Repeat("  ", depth)
		for id, next := range g.next {
			fmt.Printf("%s%d ->\n", indent, id)
			printGraph(next, depth+1)
		}
	}
	printGraph(g, 0)
}
