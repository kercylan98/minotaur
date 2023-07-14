package stream_test

import (
	"github.com/kercylan98/minotaur/utils/stream"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func initMap() map[int]int {
	return map[int]int{
		1: 100,
		2: 200,
		3: 300,
	}
}

func TestWithMap(t *testing.T) {
	Convey("TestWithMap", t, func() {
		var s = initMap()
		var m = stream.WithMap(s).RandomDelete(1)
		So(m, ShouldNotBeNil)
		So(len(s), ShouldEqual, 2)
	})
}

func TestWithMapCopy(t *testing.T) {
	Convey("TestWithMapCopy", t, func() {
		var s = initMap()
		var m = stream.WithMapCopy(s).RandomDelete(1)
		So(m, ShouldNotBeNil)
		So(len(s), ShouldEqual, 3)
	})
}

func TestMap_Set(t *testing.T) {
	Convey("TestMap_Set", t, func() {
		var m = stream.WithMap(initMap()).Set(4, 400)
		So(m[4], ShouldEqual, 400)
	})
}

func TestMap_Filter(t *testing.T) {
	Convey("TestMap_Filter", t, func() {
		var m = stream.WithMap(initMap()).Filter(func(key int, value int) bool {
			return key == 1
		})
		So(len(m), ShouldEqual, 1)
	})
}

func TestMap_FilterKey(t *testing.T) {
	Convey("TestMap_FilterKey", t, func() {
		var m = stream.WithMap(initMap()).FilterKey(1)
		So(len(m), ShouldEqual, 2)
	})
}

func TestMap_FilterValue(t *testing.T) {
	Convey("TestMap_FilterValue", t, func() {
		var m = stream.WithMap(initMap()).FilterValue(100)
		So(len(m), ShouldEqual, 2)
	})
}

func TestMap_RandomKeep(t *testing.T) {
	Convey("TestMap_RandomKeep", t, func() {
		var m = stream.WithMap(initMap()).RandomKeep(1)
		So(len(m), ShouldEqual, 1)
	})
}

func TestMap_RandomDelete(t *testing.T) {
	Convey("TestMap_RandomDelete", t, func() {
		var m = stream.WithMap(initMap()).RandomDelete(1)
		So(len(m), ShouldEqual, 2)
	})
}

func TestMap_Distinct(t *testing.T) {
	Convey("TestMap_Distinct", t, func() {
		var m = stream.WithMap(map[int]int{
			1: 100,
			2: 200,
			3: 100,
		}).Distinct(func(key int, value int) bool {
			return value == 100
		})
		So(len(m), ShouldEqual, 1)
	})
}
