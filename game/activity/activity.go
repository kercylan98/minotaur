package activity

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/offset"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/kercylan98/minotaur/utils/times"
	"time"
)

func New[PlayerID comparable, Data any, PlayerData any](id int64, startTime, endTime time.Time, data Data, generatePlayerDataHandle func() PlayerData, options ...Option[PlayerID, Data, PlayerData]) *Activity[PlayerID, Data, PlayerData] {
	activity := &Activity[PlayerID, Data, PlayerData]{
		id:                       id,
		playerData:               make(map[PlayerID]PlayerData),
		newDay:                   map[PlayerID]int64{},
		online:                   map[PlayerID]struct{}{},
		data:                     data,
		generatePlayerDataHandle: generatePlayerDataHandle,
	}
	for _, option := range options {
		option(activity)
	}
	if activity.offset == nil {
		activity.offset = offset.NewTime(0)
	}
	if activity.ticker == nil {
		activity.ticker = timer.GetTicker(5)
	}
	activity.SetTime(startTime, endTime)
	return activity
}

type Activity[PlayerID comparable, Data any, PlayerData any] struct {
	id                       int64
	ticker                   *timer.Ticker
	offset                   *offset.Time
	startTime                time.Time
	endTime                  time.Time
	data                     Data
	playerData               map[PlayerID]PlayerData
	newDay                   map[PlayerID]int64
	online                   map[PlayerID]struct{}
	generatePlayerDataHandle func() PlayerData

	startEventHandles        []StartEventHandle[PlayerID, Data, PlayerData]
	finishEventHandles       []FinishEventHandle[PlayerID, Data, PlayerData]
	newDayEventHandles       []NewDayEventHandle[PlayerID, Data, PlayerData]
	playerNewDayEventHandles []PlayerNewDayEventHandle[PlayerID, Data, PlayerData]
}

// GetId 获取活动 ID
func (slf *Activity[PlayerID, Data, PlayerData]) GetId() int64 {
	return slf.id
}

// Join 设置玩家加入活动
func (slf *Activity[PlayerID, Data, PlayerData]) Join(playerId PlayerID) {
	if !hash.Exist(slf.playerData, playerId) {
		slf.playerData[playerId] = slf.generatePlayerDataHandle()
	}
	if slf.IsFirstDay() || !hash.Exist(slf.online, playerId) {
		return
	}
	nd := slf.newDay[playerId]
	if nd > 0 {
		ndt := time.Unix(nd, 0)
		now := slf.offset.Now()
		if !times.IsSameDay(ndt, now) {
			slf.onPlayerNewDayEvent(playerId)
			slf.newDay[playerId] = now.Unix()
		}
	}
}

// Leave 设置玩家离开活动
func (slf *Activity[PlayerID, Data, PlayerData]) Leave(playerId PlayerID) {
	delete(slf.online, playerId)
}

// SetTime 设置活动开始时间和结束时间
func (slf *Activity[PlayerID, Data, PlayerData]) SetTime(startTime, endTime time.Time) {
	// 如果结束时间小于开始时间
	if endTime.Before(startTime) {
		return
	}
	slf.startTime, slf.endTime = startTime, endTime
	if !slf.IsOpen() {
		slf.ticker.StopTimer(fmt.Sprintf("ACTIVITY_%d_END", slf.id))
		slf.ticker.StopTimer(fmt.Sprintf("ACTIVITY_%d_NEW_DAY", slf.id))
	} else {
		slf.ticker.After(fmt.Sprintf("ACTIVITY_%d_END", slf.id), slf.endTime.Sub(slf.startTime), func() {
			slf.onFinishEvent()
		})
		now := slf.offset.Now()
		next := times.GetToday(now).AddDate(0, 0, 1).Sub(now)
		slf.nd(next)
	}
}

// IsFirstDay 检查是否是活动第一天
func (slf *Activity[PlayerID, Data, PlayerData]) IsFirstDay() bool {
	return times.IsSameDay(slf.startTime, slf.offset.Now())
}

// IsLastDay 检查是否是活动最后一天
func (slf *Activity[PlayerID, Data, PlayerData]) IsLastDay() bool {
	return times.IsSameDay(slf.endTime, slf.offset.Now())
}

// IsOpen 检查活动是否开放
func (slf *Activity[PlayerID, Data, PlayerData]) IsOpen() bool {
	now := slf.offset.Now()
	return now.After(slf.startTime) && now.Before(slf.endTime)
}

// GetData 获取活动数据
func (slf *Activity[PlayerID, Data, PlayerData]) GetData() Data {
	return slf.data
}

// GetPlayerData 获取活动玩家数据
func (slf *Activity[PlayerID, Data, PlayerData]) GetPlayerData(playerId PlayerID) PlayerData {
	return slf.playerData[playerId]
}

// GetLastNewDayTime 获取玩家最后触发新一天的时间
func (slf *Activity[PlayerID, Data, PlayerData]) GetLastNewDayTime(playerId PlayerID) int64 {
	return slf.newDay[playerId]
}

// GetAllLastNewDayTime 获取所有玩家最后触发新一天的时间
func (slf *Activity[PlayerID, Data, PlayerData]) GetAllLastNewDayTime() map[PlayerID]int64 {
	return slf.newDay
}

// GetAllPlayerData 获取所有玩家的活动数据
func (slf *Activity[PlayerID, Data, PlayerData]) GetAllPlayerData() map[PlayerID]PlayerData {
	return slf.playerData
}

// GetStartTime 获取活动开始时间
func (slf *Activity[PlayerID, Data, PlayerData]) GetStartTime() time.Time {
	return slf.startTime
}

// GetEndTime 获取活动结束时间
func (slf *Activity[PlayerID, Data, PlayerData]) GetEndTime() time.Time {
	return slf.endTime
}

// GetEndOfDistanceTime 获取活动距离结束的时间
//   - 如果活动已经结束，将返回 <= 0 的数值
func (slf *Activity[PlayerID, Data, PlayerData]) GetEndOfDistanceTime() time.Duration {
	now := slf.offset.Now()
	return slf.endTime.Sub(now)
}

// GetPlayerCount 获取活动玩家数量
func (slf *Activity[PlayerID, Data, PlayerData]) GetPlayerCount() int {
	return len(slf.playerData)
}

// DeletePlayerData 删除特定玩家的活动数据
func (slf *Activity[PlayerID, Data, PlayerData]) DeletePlayerData(playerId PlayerID) {
	delete(slf.playerData, playerId)
}

// HasPlayer 检查活动中是否存在特定玩家
func (slf *Activity[PlayerID, Data, PlayerData]) HasPlayer(playerId PlayerID) bool {
	return hash.Exist(slf.playerData, playerId)
}

// RegStartEvent 活动开始时将立即执行被注册的事件处理函数
func (slf *Activity[PlayerID, Data, PlayerData]) RegStartEvent(handle StartEventHandle[PlayerID, Data, PlayerData]) {
	slf.startEventHandles = append(slf.startEventHandles, handle)
}

func (slf *Activity[PlayerID, Data, PlayerData]) onStartEvent() {
	for _, handle := range slf.startEventHandles {
		handle(slf)
	}
}

// RegFinishEvent 活动结束时将立即执行被注册的事件处理函数
func (slf *Activity[PlayerID, Data, PlayerData]) RegFinishEvent(handle FinishEventHandle[PlayerID, Data, PlayerData]) {
	slf.finishEventHandles = append(slf.finishEventHandles, handle)
}

func (slf *Activity[PlayerID, Data, PlayerData]) onFinishEvent() {
	for _, handle := range slf.finishEventHandles {
		handle(slf)
	}
}

// RegNewDayEvent 活动到达新的一天时将立即执行被注册的事件处理函数
func (slf *Activity[PlayerID, Data, PlayerData]) RegNewDayEvent(handle NewDayEventHandle[PlayerID, Data, PlayerData]) {
	slf.newDayEventHandles = append(slf.newDayEventHandles, handle)
}

func (slf *Activity[PlayerID, Data, PlayerData]) onNewDayEvent() {
	for _, handle := range slf.newDayEventHandles {
		handle(slf)
	}
}

// RegPlayerNewDayEvent 活动玩家到达新的一天时将立即执行被注册的事件处理函数
func (slf *Activity[PlayerID, Data, PlayerData]) RegPlayerNewDayEvent(handle PlayerNewDayEventHandle[PlayerID, Data, PlayerData]) {
	slf.playerNewDayEventHandles = append(slf.playerNewDayEventHandles, handle)
}

func (slf *Activity[PlayerID, Data, PlayerData]) onPlayerNewDayEvent(playerId PlayerID) {
	for _, handle := range slf.playerNewDayEventHandles {
		handle(slf, playerId)
	}
}

func (slf *Activity[PlayerID, Data, PlayerData]) nd(next time.Duration) {
	slf.ticker.After(fmt.Sprintf("ACTIVITY_%d_NEW_DAY", slf.id), next, func() {
		slf.onNewDayEvent()
		slf.nd(times.Day)
	})
}
