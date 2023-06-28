package activity_test

import (
	"github.com/kercylan98/minotaur/game/activity"
	"github.com/kercylan98/minotaur/utils/offset"
	"github.com/kercylan98/minotaur/utils/times"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestActivity_IsInvalid(t *testing.T) {
	Convey("TestActivity_IsInvalid", t, func() {
		offsetTime := offset.NewTime(-time.Now().Sub(time.Date(2023, 06, 28, 13, 0, 0, 0, time.Local)))
		activity.SetOffsetTime(offsetTime)
		t.Log(offsetTime.Now())
		act := activity.NewActivity[int, activity.NoneData, activity.NoneData](1,
			times.NewPeriod(
				times.DateWithHMS(2023, 06, 21, 0, 0, 0),
				times.DateWithHMS(2023, 06, 22, 0, 0, 0),
			),
		)
		So(act.IsInvalid(), ShouldBeTrue)

		act = activity.NewActivity[int, activity.NoneData, activity.NoneData](1,
			times.NewPeriod(
				times.DateWithHMS(2023, 06, 28, 0, 0, 0),
				times.DateWithHMS(2023, 06, 29, 0, 0, 0),
			),
		)
		So(act.IsInvalid(), ShouldBeFalse)

		act = activity.NewActivity[int, activity.NoneData, activity.NoneData](1,
			times.NewPeriod(
				times.DateWithHMS(2023, 06, 26, 0, 0, 0),
				times.DateWithHMS(2023, 06, 28, 0, 0, 0),
			),
			activity.WithAfterShowTime[int, activity.NoneData, activity.NoneData](times.Day),
		)
		So(act.IsInvalid(), ShouldBeFalse)
	})
}
