package activity

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"time"
)

// Option 活动选项
type Option[Type, ID generic.Basic] func(*Activity[Type, ID])

// WithUpcomingTime 设置活动预告时间
func WithUpcomingTime[Type, ID generic.Basic](t time.Time) Option[Type, ID] {
	return func(activity *Activity[Type, ID]) {
		activity.tl.AddState(stateUpcoming, t)
	}
}

// WithStartTime 设置活动开始时间
func WithStartTime[Type, ID generic.Basic](t time.Time) Option[Type, ID] {
	return func(activity *Activity[Type, ID]) {
		activity.tl.AddState(stateStarted, t)
	}
}

// WithEndTime 设置活动结束时间
func WithEndTime[Type, ID generic.Basic](t time.Time) Option[Type, ID] {
	return func(activity *Activity[Type, ID]) {
		activity.tl.AddState(stateEnded, t)
	}
}

// WithExtendedShowTime 设置延长展示时间
func WithExtendedShowTime[Type, ID generic.Basic](t time.Time) Option[Type, ID] {
	return func(activity *Activity[Type, ID]) {
		activity.tl.AddState(stateExtendedShowEnded, t)
	}
}

// WithLoop 设置活动循环，时间间隔小于等于 0 表示不循环
//   - 当活动状态展示结束后，会根据该选项设置的时间间隔重新开始
func WithLoop[Type, ID generic.Basic](interval time.Duration) Option[Type, ID] {
	return func(activity *Activity[Type, ID]) {
		if interval <= 0 {
			interval = 0
		}
		activity.loop = interval
	}
}

// WithLazy 设置活动数据懒加载
//   - 该选项仅用于全局数据，默认情况下，活动全局数据会在活动注册时候加载，如果设置了该选项，则会在第一次获取数据时候加载
func WithLazy[Type, ID generic.Basic](lazy bool) Option[Type, ID] {
	return func(activity *Activity[Type, ID]) {
		activity.lazy = lazy
	}
}
