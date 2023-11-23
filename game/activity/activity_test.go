package activity_test

import (
	"github.com/kercylan98/minotaur/game/activity"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

type ActivityData struct {
	players []string
}

type PlayerData struct {
	info string
}

func TestRegTypeByGlobalData(t *testing.T) {
	controller := activity.DefineGlobalDataActivity[int, int, *ActivityData](1, func(activityId int, data *activity.DataMeta[*ActivityData]) {
		data.Data.players = append(data.Data.players, "temp")
	})

	activity.RegUpcomingEvent(1, func(activityId int) {
		t.Log(controller.GetGlobalData(activityId).players)
		t.Log("即将开始")
	})

	activity.RegStartedEvent(1, func(activityId int) {
		t.Log("开始")
	})

	activity.RegEndedEvent(1, func(activityId int) {
		t.Log(controller.GetGlobalData(activityId).players)
		t.Log("结束")
	})

	activity.RegExtendedShowStartedEvent(1, func(activityId int) {
		t.Log("延长展示开始")
	})

	activity.RegExtendedShowEndedEvent(1, func(activityId int) {
		t.Log("延长展示结束")
	})

	activity.RegNewDayEvent(1, func(activityId int) {
		t.Log("新的一天")
	})

	now := time.Now()

	if err := activity.LoadOrRefreshActivity(1, 1,
		activity.WithUpcomingTime[int, int](now.Add(1*time.Second)),
		activity.WithStartTime[int, int](now.Add(2*times.Second)),
		activity.WithEndTime[int, int](now.Add(3*times.Second)),
		activity.WithExtendedShowTime[int, int](now.Add(4*times.Second)),
		activity.WithLoop[int, int](3*times.Second),
	); err != nil {
		t.Fatal(err)
	}

	time.Sleep(times.Week)
}
