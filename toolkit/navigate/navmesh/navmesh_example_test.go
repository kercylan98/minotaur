package navmesh_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/geometry"
	"github.com/kercylan98/minotaur/toolkit/maths"
	"github.com/kercylan98/minotaur/toolkit/navigate/navmesh"
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

	var walkable []geometry.Polygon
	walkable = append(walkable,
		geometry.NewPolygon(
			geometry.NewPoint(5, 5),
			geometry.NewPoint(15, 5),
			geometry.NewPoint(15, 15),
			geometry.NewPoint(5, 15),
		),
		geometry.NewPolygon(
			geometry.NewPoint(15, 5),
			geometry.NewPoint(25, 5),
			geometry.NewPoint(25, 15),
			geometry.NewPoint(15, 15),
		),
		geometry.NewPolygon(
			geometry.NewPoint(15, 15),
			geometry.NewPoint(25, 15),
			geometry.NewPoint(25, 25),
			geometry.NewPoint(15, 25),
		),
	)

	for _, shape := range walkable {
		for _, edge := range shape.GetEdges() {
			sx, bx := maths.MinMax(edge[0].GetX(), edge[1].GetX())
			sy, by := maths.MinMax(edge[0].GetY(), edge[1].GetY())

			for x := sx; x <= bx; x++ {
				for y := sy; y <= by; y++ {
					fp.Put(geometry.NewPoint(x, y), '+')
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
