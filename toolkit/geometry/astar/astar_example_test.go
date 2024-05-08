package astar_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
	"github.com/kercylan98/minotaur/utils/geometry/astar"
)

type Graph struct {
	geometry.FloorPlan
}

func (slf Graph) Neighbours(point geometry.Point[int]) []geometry.Point[int] {
	neighbours := make([]geometry.Point[int], 0, 4)
	for _, direction := range geometry.DirectionUDLR {
		np := geometry.GetDirectionNextWithPoint(direction, point)
		if slf.IsFree(np) {
			neighbours = append(neighbours, np)
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

	paths := astar.Find[geometry.Point[int], int](graph, geometry.NewPoint(1, 1), geometry.NewPoint(8, 6), func(a, b geometry.Point[int]) int {
		return geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(a, b))
	}, func(a, b geometry.Point[int]) int {
		return geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(a, b))
	})

	for _, path := range paths {
		graph.Put(path, '.')
	}

	fmt.Println(graph)

	// Output:
	// ===========
	// X.XX  X   X
	// X. X   XX X
	// X.XX......X
	// X.... XXX.X
	// X XX  X ..X
	// X XX  X . X
	// ===========
}
