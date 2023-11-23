package activity

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/kercylan98/minotaur/utils/times"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

const (
	stateClosed            = byte(iota) // 已关闭
	stateUpcoming                       // 即将开始
	stateStarted                        // 已开始
	stateEnded                          // 已结束
	stateExtendedShowEnded              // 延长展示结束
)

var (
	stateLine = []byte{stateClosed, stateUpcoming, stateStarted, stateEnded, stateExtendedShowEnded}

	ticker     *timer.Ticker
	tickerOnce sync.Once
	isDebug    atomic.Bool
)

func init() {
	ticker = timer.GetTicker(10)
}

// SetDebug 设置是否开启调试模式
func SetDebug(d bool) {
	isDebug.Store(d)
}

// SetTicker 设置自定义定时器，该方法必须在使用活动系统前调用，且只能调用一次
func SetTicker(size int, options ...timer.Option) {
	tickerOnce.Do(func() {
		if ticker != nil {
			ticker.Release()
		}
		ticker = timer.GetTicker(size, options...)
	})
}

// LoadOrRefreshActivity 加载或刷新活动
//   - 通常在活动配置刷新时候将活动通过该方法注册或刷新
func LoadOrRefreshActivity[Type, ID generic.Basic](activityType Type, activityId ID, options ...Option[Type, ID]) error {
	register := getControllerRegister(activityType)
	if register == nil {
		return errors.New("activity type not registered")
	}

	act, initFinishCallback := register(activityType, activityId)
	activity := act.(*Activity[Type, ID])
	for _, option := range options {
		option(activity)
	}
	initFinishCallback(activity)
	if !activity.tl.Check(true, stateLine...) {
		return errors.New("activity state timeline is invalid")
	}

	stateTrigger := map[byte]func(){
		stateUpcoming:          func() { OnUpcomingEvent(activity) },
		stateStarted:           func() { OnStartedEvent(activity) },
		stateEnded:             func() { OnEndedEvent(activity); OnExtendedShowStartedEvent(activity) },
		stateExtendedShowEnded: func() { OnExtendedShowEndedEvent(activity) },
	}
	for _, state := range stateLine {
		if activity.tl.HasState(state) {
			activity.tl.AddTriggerToState(state, stateTrigger[state])
			continue
		}
		for next := state; next <= stateLine[len(stateLine)-1]; next++ {
			if activity.tl.HasState(next) {
				activity.tl.AddTriggerToState(next, stateTrigger[state])
				break
			}
		}
	}

	activity.refresh()
	return nil
}

// LoadGlobalData 加载所有活动全局数据
//   - 一般用于持久化活动数据
func LoadGlobalData(handler func(activityType, activityId, data any)) {
	for _, f := range controllerGlobalDataReader {
		f(handler)
	}
}

// LoadEntityData 加载所有活动实体数据
//   - 一般用于持久化活动数据
func LoadEntityData(handler func(activityType, activityId, entityId, data any)) {
	for _, f := range controllerEntityDataReader {
		f(handler)
	}
}

// Activity 活动描述
type Activity[Type, ID generic.Basic] struct {
	id           ID                     // 活动 ID
	t            Type                   // 活动类型
	state        byte                   // 活动状态
	tl           *times.StateLine[byte] // 活动时间线
	loop         time.Duration          // 活动多久循环一次
	lazy         bool                   // 是否懒加载
	tickerKey    string                 // 定时器 key
	retention    time.Duration          // 活动数据保留时间
	retentionKey string                 // 保留定时器 key
	mutex        sync.RWMutex

	getLastNewDayTime func() time.Time // 获取上次新的一天的时间
	setLastNewDayTime func(time.Time)  // 设置上次新的一天的时间
}

func (slf *Activity[Type, ID]) refresh() {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	curr := time.Now()
	if slf.state = slf.tl.GetStateByTime(curr); slf.state == stateUpcoming {
		ticker.StopTimer(slf.retentionKey)
		resetActivityData(slf.t, slf.id)
	}

	for _, f := range slf.tl.GetTriggerByState(slf.state) {
		if f != nil {
			f()
		}
	}

	next := slf.tl.GetNextTimeByState(slf.state)
	if !next.IsZero() && next.After(curr) {
		ticker.After(slf.tickerKey, next.Sub(curr)+time.Millisecond*100, slf.refresh)
	} else {
		ticker.StopTimer(slf.tickerKey)
		ticker.StopTimer(fmt.Sprintf("activity:new_day:%d:%v", reflect.ValueOf(slf.t).Kind(), slf.id))
		if slf.loop > 0 {
			slf.tl.Move(slf.loop * 2)
			ticker.After(slf.tickerKey, slf.loop+time.Millisecond*100, slf.refresh)
			return
		}
		if slf.retention > 0 {
			ticker.After(slf.tickerKey, slf.retention, func() {
				ticker.StopTimer(slf.retentionKey)
				resetActivityData(slf.t, slf.id)
			})
		}
	}

}
