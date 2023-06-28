package activity

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/times"
	"time"
)

// NewActivity 创建一个新的活动
//   - id: 活动 ID
//   - period: 活动周期
func NewActivity[PlayerID comparable, ActivityData, PlayerData any](id int64, period times.Period, options ...Option[PlayerID, ActivityData, PlayerData]) *Activity[PlayerID, ActivityData, PlayerData] {
	activity := &Activity[PlayerID, ActivityData, PlayerData]{
		id:     id,
		data:   newData[PlayerID, ActivityData, PlayerData](),
		period: period,
	}

	for _, option := range options {
		option(activity)
	}
	return activity
}

// Activity 活动
type Activity[PlayerID comparable, ActivityData, PlayerData any] struct {
	id                    int64
	data                  *Data[PlayerID, ActivityData, PlayerData]
	period                times.Period
	beforeShow, afterShow time.Time

	playerDataLoadHandle func(activity *Activity[PlayerID, ActivityData, PlayerData], playerId PlayerID) PlayerData

	startEventHandles        []StartEventHandle[PlayerID, ActivityData, PlayerData]
	stopEventHandles         []StopEventHandle[PlayerID, ActivityData, PlayerData]
	startShowEventHandles    []StartShowEventHandle[PlayerID, ActivityData, PlayerData]
	stopShowEventHandles     []StopShowEventHandle[PlayerID, ActivityData, PlayerData]
	playerJoinEventHandles   []PlayerJoinEventHandle[PlayerID, ActivityData, PlayerData]
	playerLeaveEventHandles  []PlayerLeaveEventHandle[PlayerID, ActivityData, PlayerData]
	newDayEventHandles       []NewDayEventHandle[PlayerID, ActivityData, PlayerData]
	playerNewDayEventHandles []PlayerNewDayEventHandle[PlayerID, ActivityData, PlayerData]
}

// GetStart 获取活动开始时间
func (slf *Activity[PlayerID, ActivityData, PlayerData]) GetStart() time.Time {
	return slf.period.Start()
}

// GetEnd 获取活动结束时间
func (slf *Activity[PlayerID, ActivityData, PlayerData]) GetEnd() time.Time {
	return slf.period.End()
}

// GetPeriod 获取活动周期
func (slf *Activity[PlayerID, ActivityData, PlayerData]) GetPeriod() times.Period {
	return slf.period
}

// StoreData 通过特定的存储函数存储活动数据
func (slf *Activity[PlayerID, ActivityData, PlayerData]) StoreData(handle func(data *Data[PlayerID, ActivityData, PlayerData])) {
	handle(slf.data)
}

// GetActivityData 获取活动数据
func (slf *Activity[PlayerID, ActivityData, PlayerData]) GetActivityData() ActivityData {
	return slf.data.Data
}

// GetPlayerData 获取玩家数据
func (slf *Activity[PlayerID, ActivityData, PlayerData]) GetPlayerData(playerId PlayerID) PlayerData {
	return slf.data.PlayerData[playerId]
}

// ChangeTime 修改活动时间
func (slf *Activity[PlayerID, ActivityData, PlayerData]) ChangeTime(period times.Period) {
	slf.period = period
	slf.refreshTicker(false)
}

// ChangeBeforeShowTime 修改活动开始前的展示时间
func (slf *Activity[PlayerID, ActivityData, PlayerData]) ChangeBeforeShowTime(beforeShow time.Time) {
	slf.beforeShow = beforeShow
	slf.refreshTicker(false)
}

// ChangeAfterShowTime 修改活动结束后的展示时间
func (slf *Activity[PlayerID, ActivityData, PlayerData]) ChangeAfterShowTime(afterShow time.Time) {
	slf.afterShow = afterShow
	slf.refreshTicker(false)
}

// IsShowAndOpen 判断活动是否展示并开启
func (slf *Activity[PlayerID, ActivityData, PlayerData]) IsShowAndOpen() bool {
	return slf.IsShow() && slf.IsOpen()
}

// IsShow 判断活动是否展示
//   - 活动展示时，活动并非一定开启。常用于活动的预告和结束后的成果展示
//   - 当活动没有设置展示时间时，活动展示与活动开启一致
//   - 如果活动仅设置了活动开始前的展示时间，则活动展示将从开始前的展示时间持续到活动结束
//   - 如果活动仅设置了活动结束后的展示时间，则活动展示将从活动开始持续到结束后的展示时间
func (slf *Activity[PlayerID, ActivityData, PlayerData]) IsShow() bool {
	if slf.beforeShow.IsZero() && slf.afterShow.IsZero() {
		return slf.IsOpen()
	} else if slf.beforeShow.IsZero() {
		return slf.IsOpen() && slf.afterShow.After(activityOffset.Now())
	} else if slf.afterShow.IsZero() {
		return slf.IsOpen() && slf.beforeShow.Before(activityOffset.Now())
	}
	return times.NewPeriod(slf.beforeShow, slf.afterShow).IsOngoing(activityOffset.Now())
}

// IsOpen 判断活动是否开启
func (slf *Activity[PlayerID, ActivityData, PlayerData]) IsOpen() bool {
	current := activityOffset.Now()
	return slf.period.IsOngoing(current)
}

// IsInvalid 判断活动是否无效
func (slf *Activity[PlayerID, ActivityData, PlayerData]) IsInvalid() bool {
	current := activityOffset.Now()
	if slf.beforeShow.IsZero() && !slf.afterShow.IsZero() {
		return current.After(slf.afterShow) || current.Equal(slf.afterShow)
	} else if !slf.beforeShow.IsZero() && slf.afterShow.IsZero() {
		end := slf.GetEnd()
		return current.After(end) || current.Equal(end)
	} else if !slf.beforeShow.IsZero() && !slf.afterShow.IsZero() {
		end := slf.GetEnd()
		return current.After(end) || current.Equal(end)
	} else {
		end := slf.GetEnd()
		return current.After(end) || current.Equal(end)
	}
}

// Register 将该活动进行注册
func (slf *Activity[PlayerID, ActivityData, PlayerData]) Register() {
	if slf.IsInvalid() {
		return
	}
	if !generic.IsNil(slf.data) {
		activities[slf.id] = slf
	}
	slf.refreshTicker(true)
}

// Join 设置玩家加入活动
//   - 通常玩家应该在上线的时候加入活动
func (slf *Activity[PlayerID, ActivityData, PlayerData]) Join(playerId PlayerID) {
	if slf.playerDataLoadHandle != nil {
		playerData := slf.playerDataLoadHandle(slf, playerId)
		if !generic.IsNil(playerData) {
			slf.data.PlayerData[playerId] = playerData
		}
	}
	current := activityOffset.Now()
	slf.OnPlayerJoinEvent(playerId)
	slf.playerNewDay(playerId, current)
}

// Leave 设置玩家离开活动
//   - 当玩家离开活动时，玩家的活动数据将会被清除
func (slf *Activity[PlayerID, ActivityData, PlayerData]) Leave(playerId PlayerID) {
	slf.OnPlayerLeaveEvent(playerId)
	delete(slf.data.PlayerData, playerId)
}

// RegStartEvent 注册活动开始事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegStartEvent(handle StartEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.startEventHandles = append(slf.startEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnStartEvent() {
	for _, handle := range slf.startEventHandles {
		handle(slf)
	}
}

// RegStopEvent 注册活动结束事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegStopEvent(handle StopEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.stopEventHandles = append(slf.stopEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnStopEvent() {
	slf.stopTicker()
	for _, handle := range slf.stopEventHandles {
		handle(slf)
	}
}

// RegStartShowEvent 注册活动开始展示事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegStartShowEvent(handle StartShowEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.startShowEventHandles = append(slf.startShowEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnStartShowEvent() {
	for _, handle := range slf.startShowEventHandles {
		handle(slf)
	}
}

// RegStopShowEvent 注册活动结束展示事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegStopShowEvent(handle StopShowEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.stopShowEventHandles = append(slf.stopShowEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnStopShowEvent() {
	for _, handle := range slf.stopShowEventHandles {
		handle(slf)
	}
}

// RegPlayerJoinEvent 注册玩家加入活动事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegPlayerJoinEvent(handle PlayerJoinEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.playerJoinEventHandles = append(slf.playerJoinEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnPlayerJoinEvent(playerId PlayerID) {
	for _, handle := range slf.playerJoinEventHandles {
		handle(slf, playerId)
	}
}

// RegPlayerLeaveEvent 注册玩家离开活动事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegPlayerLeaveEvent(handle PlayerLeaveEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.playerLeaveEventHandles = append(slf.playerLeaveEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnPlayerLeaveEvent(playerId PlayerID) {
	for _, handle := range slf.playerLeaveEventHandles {
		handle(slf, playerId)
	}
}

// RegNewDayEvent 注册新的一天事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegNewDayEvent(handle NewDayEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.newDayEventHandles = append(slf.newDayEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnNewDayEvent() {
	for _, handle := range slf.newDayEventHandles {
		handle(slf)
	}
}

// RegPlayerNewDayEvent 注册玩家新的一天事件
func (slf *Activity[PlayerID, ActivityData, PlayerData]) RegPlayerNewDayEvent(handle PlayerNewDayEventHandle[PlayerID, ActivityData, PlayerData]) {
	slf.playerNewDayEventHandles = append(slf.playerNewDayEventHandles, handle)
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) OnPlayerNewDayEvent(playerId PlayerID) {
	for _, handle := range slf.playerNewDayEventHandles {
		handle(slf, playerId)
	}
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) refreshTicker(init bool) {
	var (
		current = activityOffset.Now()
		start   = slf.period.Start()
		end     = slf.period.End()
	)

	if slf.IsOpen() {
		if init {
			slf.OnStartEvent()
		}
		if distance := end.Sub(current); distance > 0 {
			ticker.After(fmt.Sprintf(tickerStop, slf.id), distance, slf.OnStopEvent)
		}
		ticker.After(fmt.Sprintf(tickerNewDay, slf.id), times.GetNextDayInterval(current), slf.newDay)
	} else {
		if distance := start.Sub(current); distance > 0 {
			ticker.After(fmt.Sprintf(tickerNewDay, slf.id), times.GetNextDayInterval(start), slf.newDay)
			ticker.After(fmt.Sprintf(tickerStart, slf.id), distance, slf.OnStartEvent)
		}
		if end.Before(current) {
			slf.OnStopEvent()
		}
	}
	if slf.IsShow() {
		if init {
			slf.OnStartShowEvent()
		}
		if distance := slf.afterShow.Sub(current); distance > 0 {
			ticker.After(fmt.Sprintf(tickerShowStop, slf.id), distance, slf.OnStopShowEvent)
		}
	} else {
		if distance := slf.beforeShow.Sub(current); distance > 0 {
			ticker.After(fmt.Sprintf(tickerShowStart, slf.id), distance, slf.OnStartShowEvent)
		}
		if slf.afterShow.Before(current) {
			slf.OnStopShowEvent()
		}
	}

}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) stopTicker() {
	ticker.StopTimer(fmt.Sprintf(tickerStart, slf.id))
	ticker.StopTimer(fmt.Sprintf(tickerStop, slf.id))
	ticker.StopTimer(fmt.Sprintf(tickerShowStart, slf.id))
	ticker.StopTimer(fmt.Sprintf(tickerShowStop, slf.id))
	ticker.StopTimer(fmt.Sprintf(tickerNewDay, slf.id))
}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) newDay() {
	current := activityOffset.Now()
	ticker.After(fmt.Sprintf(tickerNewDay, slf.id), times.GetNextDayInterval(current), slf.newDay)

	last := time.Unix(slf.data.LastNewDay, 0)
	if !times.IsSameDay(last, times.GetToday(current)) {
		slf.OnNewDayEvent()
	}

	for playerId := range slf.data.PlayerData {
		slf.playerNewDay(playerId, current)
	}

}

func (slf *Activity[PlayerID, ActivityData, PlayerData]) playerNewDay(playerId PlayerID, current time.Time) {
	last := time.Unix(slf.data.PlayerLastNewDay[playerId], 0)
	if !times.IsSameDay(last, times.GetToday(current)) {
		slf.OnPlayerNewDayEvent(playerId)
	}
}
