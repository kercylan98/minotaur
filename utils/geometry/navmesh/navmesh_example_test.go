package navmesh_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
	"github.com/kercylan98/minotaur/utils/geometry/navmesh"
	"github.com/kercylan98/minotaur/utils/maths"
)

func ExampleNavMesh_FindPath() {
	fp := geometry.FloorPlan{
		"=================================",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"X                               X",
		"=================================",
	}

	var walkable []geometry.Shape[int]
	walkable = append(walkable,
		geometry.NewShape(
			geometry.NewPoint(5, 5),
			geometry.NewPoint(15, 5),
			geometry.NewPoint(15, 15),
			geometry.NewPoint(5, 15),
		),
		geometry.NewShape(
			geometry.NewPoint(15, 5),
			geometry.NewPoint(25, 5),
			geometry.NewPoint(25, 15),
			geometry.NewPoint(15, 15),
		),
		geometry.NewShape(
			geometry.NewPoint(15, 15),
			geometry.NewPoint(25, 15),
			geometry.NewPoint(25, 25),
			geometry.NewPoint(15, 25),
		),
	)

	for _, shape := range walkable {
		for _, edge := range shape.Edges() {
			sx, bx := maths.MinMax(edge.GetStart().GetX(), edge.GetEnd().GetX())
			sy, by := maths.MinMax(edge.GetStart().GetY(), edge.GetEnd().GetY())

			for x := sx; x <= bx; x++ {
				for y := sy; y <= by; y++ {
					fp.Put(geometry.NewPoint[int](int(x), int(y)), '+')
				}
			}
		}
	}

	nm := navmesh.NewNavMesh(walkable, 0)
	path := nm.FindPath(
		geometry.NewPoint(6, 6),
		geometry.NewPoint(18, 24),
	)
	for _, point := range path {
		fp.Put(geometry.NewPoint(point.GetX(), point.GetY()), 'G')
	}

	fmt.Println(fp)

	// Output:
	// =================================
	// X                               X
	// X                               X
	// X                               X
	// X                               X
	// X    +++++++++++++++++++++      X
	// X    +G        +         +      X
	// X    +         +         +      X
	// X    +         +         +      X
	// X    +         +         +      X
	// X    +         +         +      X
	// X    +         +         +      X
	// X    +         +         +      X
	// X    +         +         +      X
	// X    +         +         +      X
	// X    ++++++++++G++++++++++      X
	// X              +         +      X
	// X              +         +      X
	// X              +         +      X
	// X              +         +      X
	// X              +         +      X
	// X              +         +      X
	// X              +         +      X
	// X              +         +      X
	// X              +  G      +      X
	// X              +++++++++++      X
	// X                               X
	// X                               X
	// =================================
}
