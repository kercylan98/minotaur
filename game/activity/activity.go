package activity

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/timer"
	"reflect"
	"sync"
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
)

func init() {
	ticker = timer.GetTicker(10)
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

// LoadGlobalData 加载所有活动全局数据
func LoadGlobalData(handler func(activityType, activityId, data any)) {
	for _, f := range activityGlobalDataLoader {
		f(handler)
	}
}

// LoadEntityData 加载所有活动实体数据
func LoadEntityData(handler func(activityType, activityId, entityId, data any)) {
	for _, f := range activityEntityDataLoader {
		f(handler)
	}
}

// LoadOrRefreshActivity 加载或刷新活动
//   - 通常在活动配置刷新时候将活动通过该方法注册或刷新
func LoadOrRefreshActivity[Type, ID generic.Basic](activityType Type, activityId ID, options ...*Options) error {
	register, exist := activityRegister[activityType]
	if !exist {
		return fmt.Errorf("activity type %v not registered, activity %v registration failed", activityType, activityId)
	}

	activity := register(activityId, initOptions(options...)).(*Activity[Type, ID])
	if !activity.options.Tl.Check(true, stateLine...) {
		return fmt.Errorf("activity %v state timeline is invalid", activityId)
	}

	stateTrigger := map[byte]func(){
		stateUpcoming:          func() { OnUpcomingEvent(activity) },
		stateStarted:           func() { OnStartedEvent(activity) },
		stateEnded:             func() { OnEndedEvent(activity); OnExtendedShowStartedEvent(activity) },
		stateExtendedShowEnded: func() { OnExtendedShowEndedEvent(activity) },
	}

	for _, state := range stateLine {
		if activity.options.Tl.HasState(state) {
			activity.options.Tl.AddTriggerToState(state, stateTrigger[state])
			continue
		}
		for next := state; next <= stateLine[len(stateLine)-1]; next++ {
			if activity.options.Tl.HasState(next) {
				activity.options.Tl.AddTriggerToState(next, stateTrigger[state])
				break
			}
		}
	}

	activity.refresh()
	return nil
}

// Activity 活动描述
type Activity[Type, ID generic.Basic] struct {
	id      ID       // 活动 ID
	t       Type     // 活动类型
	options *Options // 活动选项
	state   byte     // 活动状态

	lazy         bool          // 是否懒加载
	tickerKey    string        // 定时器 key
	retention    time.Duration // 活动数据保留时间
	retentionKey string        // 保留定时器 key
	mutex        sync.RWMutex

	getLastNewDayTime func() time.Time // 获取上次新的一天的时间
	setLastNewDayTime func(time.Time)  // 设置上次新的一天的时间
	clearData         func()           // 清理活动数据
	initializeData    func()           // 初始化活动数据
}

func (slf *Activity[Type, ID]) refresh() {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	curr := time.Now()
	if slf.state = slf.options.Tl.GetStateByTime(curr); slf.state == stateUpcoming || (slf.state == stateStarted && !slf.options.Tl.HasState(stateUpcoming)) {
		ticker.StopTimer(slf.retentionKey)
		slf.initializeData()
	}

	for _, f := range slf.options.Tl.GetTriggerByState(slf.state) {
		if f != nil {
			f()
		}
	}

	next := slf.options.Tl.GetNextTimeByState(slf.state)
	if !next.IsZero() && next.After(curr) {
		ticker.After(slf.tickerKey, next.Sub(curr)+time.Millisecond*100, slf.refresh)
	} else {
		ticker.StopTimer(slf.tickerKey)
		ticker.StopTimer(fmt.Sprintf("activity:new_day:%d:%v", reflect.ValueOf(slf.t).Kind(), slf.id))
		if slf.options.Loop > 0 {
			slf.options.Tl.Move(slf.options.Loop * 2)
			ticker.After(slf.tickerKey, slf.options.Loop+time.Millisecond*100, slf.refresh)
			return
		}
		if slf.retention > 0 {
			ticker.After(slf.tickerKey, slf.retention, func() {
				ticker.StopTimer(slf.retentionKey)
				slf.clearData()
			})
		}
	}

}
