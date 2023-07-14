package stream_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/stream"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSlice_Filter(t *testing.T) {
	Convey("TestSlice_Filter", t, func() {
		d := []int{1, 2, 3, 4, 5}
		var s = stream.WithSlice(d).Reverse()
		fmt.Println(s)
		fmt.Println(d)
		So(len(s), ShouldEqual, 2)
	})
}
