package super_test

import (
	"github.com/kercylan98/minotaur/utils/super"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMatch(t *testing.T) {
	Convey("TestMatch", t, func() {
		So(super.Match[int, string](1).
			Case(1, "a").
			Case(2, "b").
			Default("c"), ShouldEqual, "a")
		So(super.Match[int, string](2).
			Case(1, "a").
			Case(2, "b").
			Default("c"), ShouldEqual, "b")
		So(super.Match[int, string](3).
			Case(1, "a").
			Case(2, "b").
			Default("c"), ShouldEqual, "c")
	})
}
