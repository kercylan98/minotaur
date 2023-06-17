package geometry_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
	"testing"
)

func TestShape_Search(t *testing.T) {
	var shape geometry.Shape[int]
	// 生成一个L形的shape
	shape = append(shape, geometry.NewPoint(1, 0))
	shape = append(shape, geometry.NewPoint(1, 1))
	shape = append(shape, geometry.NewPoint(1, 2))
	shape = append(shape, geometry.NewPoint(2, 2))

	fmt.Println(shape)

	shapes := shape.ShapeSearch(geometry.WithShapeSearchAsc(), geometry.WithShapeSearchDeduplication())

	for _, shape := range shapes {
		fmt.Println("图形", shape.Points())
		fmt.Println(shape)
	}
}
