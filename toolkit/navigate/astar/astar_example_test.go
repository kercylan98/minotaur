package astar_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/geometry"
	"github.com/kercylan98/minotaur/toolkit/navigate/astar"
)

type Graph struct {
	geometry.FloorPlan
}

type Entity struct {
	geometry.Vector2
}

func (g *Graph) GetNodeId(node *Entity) int {
	x, y := node.GetXY()
	a, b := int(x), int(y)
	mergedNumber := (a << 16) | (b & 0xFFFFFFFF)
	return mergedNumber
}

func (g *Graph) GetNeighbours(t *Entity) []*Entity {
	var neighbours []*Entity
	for _, direction := range geometry.Direction2D4() {
		next := geometry.CalcOffsetInDirection2D(t.Vector2, direction, 1)
		if g.FloorPlan.IsFree(next) {
			neighbours = append(neighbours, &Entity{
				Vector2: next,
			})
		}
	}
	return neighbours
}

func ExampleFind() {
	graph := Graph{
		FloorPlan: geometry.FloorPlan{
			"===========",
			"X XX  X   X",
			"X  X   XX X",
			"X XX      X",
			"X     XXX X",
			"X XX  X   X",
			"X XX  X   X",
			"===========",
		},
	}

	paths := astar.Find[int, *Entity](
		&graph,
		&Entity{geometry.NewVector2(1, 1)},
		&Entity{geometry.NewVector2(8, 6)},
		// 曼哈顿距离
		func(a, b *Entity) float64 {
			return a.Sub(b.Vector2).Length()
		},
		func(a, b *Entity) float64 {
			return a.Sub(b.Vector2).Length()
		},
	)

	for _, path := range paths {
		graph.Put(path.Vector2, '.')
	}

	fmt.Println(graph)

	// Output:
	// ===========
	// X.XX  X   X
	// X. X   XX X
	// X.XX .....X
	// X.....XXX.X
	// X XX  X  .X
	// X XX  X ..X
	// ===========
}
