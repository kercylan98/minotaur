package g2d

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/g2d/shape"
	"testing"
)

func TestMatrixShape(t *testing.T) {
	var m = [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	s := shape.NewShape()
	s.AddPoints(shape.NewPointWithArrays([][2]int{{0, 0}, {0, 1}, {1, 1}, {2, 1}}...)...)
	fmt.Println(s)
	result := MatrixShapeSearchWithYX(m, []*shape.Shape{s}, func(val int) bool {
		return val == 1
	})
	fmt.Println(len(result))
}
