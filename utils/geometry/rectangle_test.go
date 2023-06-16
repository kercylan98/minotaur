package geometry_test

import (
	"github.com/kercylan98/minotaur/utils/geometry"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetShapeCoverageAreaWithCoordinateArray(t *testing.T) {
	Convey("TestGetShapeCoverageAreaWithCoordinateArray", t, func() {
		var points []geometry.Point[int]
		points = append(points, geometry.NewPoint(1, 1))
		points = append(points, geometry.NewPoint(2, 1))
		points = append(points, geometry.NewPoint(2, 2))

		left, right, top, bottom := geometry.GetShapeCoverageAreaWithCoordinateArray(points...)

		So(left, ShouldEqual, 1)
		So(right, ShouldEqual, 2)
		So(top, ShouldEqual, 1)
		So(bottom, ShouldEqual, 2)
	})
}

func TestGetShapeCoverageAreaWithPos(t *testing.T) {
	Convey("TestGetShapeCoverageAreaWithPos", t, func() {
		left, right, top, bottom := geometry.GetShapeCoverageAreaWithPos(3, 4, 7, 8)

		So(left, ShouldEqual, 1)
		So(right, ShouldEqual, 2)
		So(top, ShouldEqual, 1)
		So(bottom, ShouldEqual, 2)
	})
}
