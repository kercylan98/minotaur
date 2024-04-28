package geometry_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
)

func ExampleGetShapeCoverageAreaWithPoint() {
	// # # #
	// # X #
	// # X X

	var points []geometry.Point[int]
	points = append(points, geometry.NewPoint(1, 1))
	points = append(points, geometry.NewPoint(2, 1))
	points = append(points, geometry.NewPoint(2, 2))

	left, right, top, bottom := geometry.GetShapeCoverageAreaWithPoint(points...)
	fmt.Println(fmt.Sprintf("left: %v, right: %v, top: %v, bottom: %v", left, right, top, bottom))

	// left: 1, right: 2, top: 1, bottom: 2
}

func ExampleGetShapeCoverageAreaWithPos() {
	// # # #    0 1 2
	// # X #    3 4 5
	// # X X    6 7 8

	left, right, top, bottom := geometry.GetShapeCoverageAreaWithPos(3, 4, 7, 8)
	fmt.Println(fmt.Sprintf("left: %v, right: %v, top: %v, bottom: %v", left, right, top, bottom))

	// left: 1, right: 2, top: 1, bottom: 2
}

func ExampleCoverageAreaBoundless() {
	// # # #
	// # X #
	// # X X

	//   ↓

	// X #
	// X X

	left, right, top, bottom := geometry.CoverageAreaBoundless(1, 2, 1, 2)
	fmt.Println(fmt.Sprintf("left: %v, right: %v, top: %v, bottom: %v", left, right, top, bottom))

	// left: 0, right: 1, top: 0, bottom: 1
}
