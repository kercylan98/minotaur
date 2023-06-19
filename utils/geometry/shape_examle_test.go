package geometry_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
)

func ExampleNewShape() {
	shape := geometry.NewShape[int](
		geometry.NewPoint(3, 0),
		geometry.NewPoint(3, 1),
		geometry.NewPoint(3, 2),
		geometry.NewPoint(3, 3),
		geometry.NewPoint(4, 3),
	)

	fmt.Println(shape)

	// Output:
	// [[3 0] [3 1] [3 2] [3 3] [4 3]]
	// X #
	// X #
	// X #
	// X X
}

func ExampleNewShapeWithString() {
	shape := geometry.NewShapeWithString[int]([]string{
		"###X###",
		"###X###",
		"###X###",
		"###XX##",
	}, 'X')

	fmt.Println(shape)

	// Output: [[3 0] [3 1] [3 2] [3 3] [4 3]]
	// X #
	// X #
	// X #
	// X X
}

func ExampleShape_Points() {
	shape := geometry.NewShapeWithString[int]([]string{
		"###X###",
		"##XXX##",
	}, 'X')

	points := shape.Points()

	fmt.Println(points)

	// Output:
	// [[3 0] [2 1] [3 1] [4 1]]
}

func ExampleShape_PointCount() {
	shape := geometry.NewShapeWithString[int]([]string{
		"###X###",
		"##XXX##",
	}, 'X')

	fmt.Println(shape.PointCount())

	// Output:
	// 4
}

func ExampleShape_String() {
	shape := geometry.NewShapeWithString[int]([]string{
		"###X###",
		"##XXX##",
	}, 'X')

	fmt.Println(shape)

	// Output:
	// [[3 0] [2 1] [3 1] [4 1]]
	// # X #
	// X X X
}

func ExampleShape_ShapeSearch() {
	shape := geometry.NewShapeWithString[int]([]string{
		"###X###",
		"##XXX##",
		"###X###",
	}, 'X')

	shapes := shape.ShapeSearch(
		geometry.WithShapeSearchDeduplication(),
		geometry.WithShapeSearchDesc(),
	)
	for _, shape := range shapes {
		fmt.Println(shape)
	}

	// Output:
	// [[3 0] [3 2] [2 1] [4 1] [3 1]]
	// # X #
	// X X X
	// # X #
}
