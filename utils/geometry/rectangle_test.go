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

func TestCoverageAreaBoundless(t *testing.T) {
	Convey("TestCoverageAreaBoundless", t, func() {
		left, right, top, bottom := geometry.CoverageAreaBoundless(1, 2, 1, 2)

		So(left, ShouldEqual, 0)
		So(right, ShouldEqual, 1)
		So(top, ShouldEqual, 0)
		So(bottom, ShouldEqual, 1)
	})
}

func TestGenerateShapeOnRectangle(t *testing.T) {
	Convey("TestGenerateShapeOnRectangle", t, func() {
		var points []geometry.Point[int]
		points = append(points, geometry.NewPoint(1, 1))
		points = append(points, geometry.NewPoint(2, 1))
		points = append(points, geometry.NewPoint(2, 2))

		ps := geometry.GenerateShapeOnRectangle(points...)

		So(ps[0].GetX(), ShouldEqual, 0)
		So(ps[0].GetY(), ShouldEqual, 0)
		So(ps[0].GetData(), ShouldEqual, true)

		So(ps[1].GetX(), ShouldEqual, 1)
		So(ps[1].GetY(), ShouldEqual, 0)
		So(ps[1].GetData(), ShouldEqual, true)

		So(ps[2].GetX(), ShouldEqual, 0)
		So(ps[2].GetY(), ShouldEqual, 1)
		So(ps[2].GetData(), ShouldEqual, false)

		So(ps[3].GetX(), ShouldEqual, 1)
		So(ps[3].GetY(), ShouldEqual, 1)
		So(ps[3].GetData(), ShouldEqual, true)

	})
}
