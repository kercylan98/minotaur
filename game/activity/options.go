package activity

import (
	"time"
)

type Option[PlayerID comparable, ActivityData, PlayerData any] func(activity *Activity[PlayerID, ActivityData, PlayerData])

// WithData 通过指定活动数据的方式来创建活动
func WithData[PlayerID comparable, ActivityData, PlayerData any](data *Data[PlayerID, ActivityData, PlayerData]) Option[PlayerID, ActivityData, PlayerData] {
	return func(activity *Activity[PlayerID, ActivityData, PlayerData]) {
		activity.data = data
	}
}

// WithActivityData 通过指定活动全局数据的方式来创建活动
//   - 该活动数据将会被作为活动的全局数据
//   - 默认情况下活动本身不包含任何数据
func WithActivityData[PlayerID comparable, ActivityData, PlayerData any](data ActivityData) Option[PlayerID, ActivityData, PlayerData] {
	return func(activity *Activity[PlayerID, ActivityData, PlayerData]) {
		activity.data.Data = data
	}
}

// WithPlayerDataLoadHandle 通过指定玩家数据加载函数的方式来创建活动
//   - 该函数将会在玩家数据加载时被调用
//   - 活动中的玩家数据将会被按需加载，只有在玩家加入活动时才会被加载
func WithPlayerDataLoadHandle[PlayerID comparable, ActivityData, PlayerData any](handle func(activity *Activity[PlayerID, ActivityData, PlayerData], playerId PlayerID) PlayerData) Option[PlayerID, ActivityData, PlayerData] {
	return func(activity *Activity[PlayerID, ActivityData, PlayerData]) {
		activity.playerDataLoadHandle = handle
	}
}

// WithBeforeShowTime 通过指定活动开始前的展示时间的方式来创建活动
func WithBeforeShowTime[PlayerID comparable, ActivityData, PlayerData any](showTime time.Duration) Option[PlayerID, ActivityData, PlayerData] {
	return func(activity *Activity[PlayerID, ActivityData, PlayerData]) {
		activity.beforeShow = activity.period.Start().Add(-showTime)
	}
}

// WithAfterShowTime 通过指定活动结束后的展示时间的方式来创建活动
func WithAfterShowTime[PlayerID comparable, ActivityData, PlayerData any](showTime time.Duration) Option[PlayerID, ActivityData, PlayerData] {
	return func(activity *Activity[PlayerID, ActivityData, PlayerData]) {
		activity.afterShow = activity.period.End().Add(showTime)
	}
}

// WithLastNewDay 通过指定活动最后触发新的一天的时间戳的方式来创建活动
func WithLastNewDay[PlayerID comparable, ActivityData, PlayerData any](lastNewDay int64) Option[PlayerID, ActivityData, PlayerData] {
	return func(activity *Activity[PlayerID, ActivityData, PlayerData]) {
		activity.data.LastNewDay = lastNewDay
	}
}

// WithPlayerLastNewDay 通过指定玩家最后触发新的一天的时间戳的方式来创建活动
func WithPlayerLastNewDay[PlayerID comparable, ActivityData, PlayerData any](playerLastNewDay map[PlayerID]int64) Option[PlayerID, ActivityData, PlayerData] {
	return func(activity *Activity[PlayerID, ActivityData, PlayerData]) {
		activity.data.PlayerLastNewDay = playerLastNewDay
	}
}
