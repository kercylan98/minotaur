package activity

import (
	"github.com/kercylan98/minotaur/utils/offset"
	"github.com/kercylan98/minotaur/utils/timer"
)

type Option[PlayerID comparable, Data any, PlayerData any] func(activity *Activity[PlayerID, Data, PlayerData])

// WithOffsetTime 通过活动时间使用偏移时间的方式创建活动
func WithOffsetTime[PlayerID comparable, Data any, PlayerData any](offsetTime *offset.Time) Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.offset = offsetTime
	}
}

// WithGlobalOffsetTime 通过活动时间使用全局偏移时间的方式创建活动
func WithGlobalOffsetTime[PlayerID comparable, Data any, PlayerData any]() Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.offset = offset.GetGlobal()
	}
}

// WithTicker 通过使用特定的定时器创建活动
func WithTicker[PlayerID comparable, Data any, PlayerData any](ticker *timer.Ticker) Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.ticker = ticker
	}
}

// WithData 通过使用特定的活动数据创建活动
func WithData[PlayerID comparable, Data any, PlayerData any](data Data) Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.data = data
	}
}

// WithPlayerData 通过使用特定的玩家活动数据创建活动
func WithPlayerData[PlayerID comparable, Data any, PlayerData any](playerId PlayerID, playerData PlayerData) Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.playerData[playerId] = playerData
	}
}

// WithPlayerDataFromMap 通过使用 map 中的玩家活动数据创建活动
func WithPlayerDataFromMap[PlayerID comparable, Data any, PlayerData any](playerData map[PlayerID]PlayerData) Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.playerData = playerData
	}
}

// WithLastNewDayTime 通过使用特定的玩家最后触发新一天的时间数据创建活动
func WithLastNewDayTime[PlayerID comparable, Data any, PlayerData any](playerId PlayerID, newDayTime int64) Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.newDay[playerId] = newDayTime
	}
}

// WithLastNewDayTimeFromMap 通过使用 map 中的玩家最后触发新一天的时间数据创建活动
func WithLastNewDayTimeFromMap[PlayerID comparable, Data any, PlayerData any](newDayTime map[PlayerID]int64) Option[PlayerID, Data, PlayerData] {
	return func(activity *Activity[PlayerID, Data, PlayerData]) {
		activity.newDay = newDayTime
	}
}
