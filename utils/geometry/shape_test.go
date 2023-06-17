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
	shape = append(shape, geometry.NewPoint(1, 3))
	geometry.ShapeStringHasBorder = true

	fmt.Println("形状：")
	fmt.Println(shape)

	shapes := shape.ShapeSearch(
		geometry.WithShapeSearchDesc(),
		geometry.WithShapeSearchDeduplication(),
		geometry.WithShapeSearchPointCountLowerLimit(3),
		geometry.WithShapeSearchRightAngle(),
	)

	for _, shape := range shapes {
		fmt.Println("搜索", shape.Points())
		fmt.Println(shape)
	}
}
