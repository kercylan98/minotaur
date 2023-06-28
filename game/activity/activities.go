package activity

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/offset"
	"github.com/kercylan98/minotaur/utils/timer"
)

const (
	TickerMark = "ACTIVITY_SYSTEM_TICKER"
)

const (
	tickerStart     = "MINOTAUR_ACTIVITY_%d_START_c8JDh23!df"
	tickerStop      = "MINOTAUR_ACTIVITY_%d_STOP_3d!fj2@#d"
	tickerShowStart = "MINOTAUR_ACTIVITY_%d_SHOW_START_3d#f@2@#d"
	tickerShowStop  = "MINOTAUR_ACTIVITY_%d_SHOW_STOP_68!fj11#d"
	tickerNewDay    = "MINOTAUR_ACTIVITY_%d_NEW_DAY_3d!fd8@_d"
)

var (
	ticker         *timer.Ticker
	activityOffset *offset.Time
	activities     map[int64]activityInterface
)

// NoneData 空数据活动
type NoneData struct{}

func init() {
	SetTicker(nil)
	activityOffset = offset.GetGlobal()
	activities = map[int64]activityInterface{}
}

// Register 注册活动
func Register[PlayerID comparable, ActivityData, PlayerData any](activity *Activity[PlayerID, ActivityData, PlayerData]) {
	activity.Register()
}

// Get 获取活动
func Get[PlayerID comparable, ActivityData, PlayerData any](activityId int64) *Activity[PlayerID, ActivityData, PlayerData] {
	return activities[activityId].(*Activity[PlayerID, ActivityData, PlayerData])
}

// HasActivity 检查活动是否存在
func HasActivity(activityId int64) bool {
	return hash.Exist(activities, activityId)
}

// SetOffsetTime 设置活动的时间偏移时间
func SetOffsetTime(time *offset.Time) {
	activityOffset = time
	for _, activity := range activities {
		activity.stopTicker()
		activity.refreshTicker(false)
	}
}

// SetTicker 通过指定特定的定时器取代默认的定时器
//   - 当传入的定时器为 nil 时，将使用默认的定时器
func SetTicker(t *timer.Ticker) {
	if ticker != nil {
		if ticker.Mark() == TickerMark {
			ticker.Release()
		}
		for _, activity := range activities {
			activity.stopTicker()
		}
	}
	if ticker = t; ticker == nil {
		ticker = timer.GetTicker(30, timer.WithMark(TickerMark))
		for _, activity := range activities {
			activity.refreshTicker(false)
		}
	}
}
