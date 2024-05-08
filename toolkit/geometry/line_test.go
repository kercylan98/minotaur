package geometry_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
	"testing"
)

func TestCalcLineSegmentIsIntersect(t *testing.T) {
	line1 := geometry.NewLineSegment(geometry.NewPoint(1, 1), geometry.NewPoint(3, 5))
	line2 := geometry.NewLineSegment(geometry.NewPoint(0, 5), geometry.NewPoint(3, 6))
	fmt.Println(geometry.CalcLineSegmentIsIntersect(line1, line2))
}
