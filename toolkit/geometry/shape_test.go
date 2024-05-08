package geometry_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/geometry"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewShape(t *testing.T) {
	Convey("TestNewShape", t, func() {
		shape := geometry.NewShape[int](
			geometry.NewPoint(3, 0),
			geometry.NewPoint(3, 1),
			geometry.NewPoint(3, 2),
			geometry.NewPoint(3, 3),
			geometry.NewPoint(4, 3),
		)
		fmt.Println(shape)
		points := shape.Points()
		count := shape.PointCount()
		So(count, ShouldEqual, 5)
		So(points[0], ShouldEqual, geometry.NewPoint(3, 0))
		So(points[1], ShouldEqual, geometry.NewPoint(3, 1))
		So(points[2], ShouldEqual, geometry.NewPoint(3, 2))
		So(points[3], ShouldEqual, geometry.NewPoint(3, 3))
		So(points[4], ShouldEqual, geometry.NewPoint(4, 3))
	})
}

func TestNewShapeWithString(t *testing.T) {
	Convey("TestNewShapeWithString", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{
			"###X###",
			"###X###",
			"###X###",
			"###XX##",
		}, 'X')

		points := shape.Points()
		count := shape.PointCount()
		So(count, ShouldEqual, 5)
		So(points[0], ShouldEqual, geometry.NewPoint(3, 0))
		So(points[1], ShouldEqual, geometry.NewPoint(3, 1))
		So(points[2], ShouldEqual, geometry.NewPoint(3, 2))
		So(points[3], ShouldEqual, geometry.NewPoint(3, 3))
		So(points[4], ShouldEqual, geometry.NewPoint(4, 3))
	})
}

func TestShape_Points(t *testing.T) {
	Convey("TestShape_Points", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{
			"###X###",
			"##XXX##",
		}, 'X')

		points := shape.Points()
		So(points[0], ShouldEqual, geometry.NewPoint(3, 0))
		So(points[1], ShouldEqual, geometry.NewPoint(2, 1))
		So(points[2], ShouldEqual, geometry.NewPoint(3, 1))
		So(points[3], ShouldEqual, geometry.NewPoint(4, 1))
	})
}

func TestShape_PointCount(t *testing.T) {
	Convey("TestShape_PointCount", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{
			"###X###",
			"##XXX##",
		}, 'X')

		So(shape.PointCount(), ShouldEqual, 4)
	})
}

func TestShape_String(t *testing.T) {
	Convey("TestShape_String", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{
			"###X###",
			"##XXX##",
		}, 'X')

		str := shape.String()

		So(str, ShouldEqual, "[[3 0] [2 1] [3 1] [4 1]]\n# X #\nX X X")
	})
}

func TestShape_Search(t *testing.T) {
	Convey("TestShape_Search", t, func() {
		shape := geometry.NewShapeWithString[int]([]string{
			"###X###",
			"##XXX##",
			"###X###",
		}, 'X')

		shapes := shape.ShapeSearch(
			geometry.WithShapeSearchDeduplication(),
			geometry.WithShapeSearchDesc(),
		)
		So(len(shapes), ShouldEqual, 1)
		for _, shape := range shapes {
			So(shape.PointCount(), ShouldEqual, 5)
		}
	})

}
