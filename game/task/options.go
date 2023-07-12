package task

import (
	"github.com/kercylan98/minotaur/utils/offset"
	"time"
)

type Option func(task *Task)

// WithChildCount 通过初始化子计数的方式创建任务
func WithChildCount(key any, childCount int64) Option {
	return func(task *Task) {
		if task.childCount == nil {
			task.childCount = make(map[any]int64)
		}
		if task.childCondition == nil {
			task.childCondition = make(map[any]int64)
		}
		task.childCount[key] = childCount
	}
}

// WithChild 通过指定子计数的方式创建任务
//   - 只有当子计数与主计数均达到条件时，任务才会完成
//   - 通常用于多条件的任务
func WithChild(key any, childCondition int64) Option {
	return func(task *Task) {
		if task.childCount == nil {
			task.childCount = make(map[any]int64)
		}
		if task.childCondition == nil {
			task.childCondition = make(map[any]int64)
		}
		task.childCondition[key] = childCondition
	}
}

// WithDisableNotStartGetReward 禁止未开始的任务领取奖励
func WithDisableNotStartGetReward() Option {
	return func(task *Task) {
		task.disableNotStartGetReward = true
	}
}

// WithCount 通过初始化计数的方式创建任务
func WithCount(count int64) Option {
	return func(task *Task) {
		task.SetCount(count)
	}
}

// WithStartTime 通过指定开始时间的方式创建任务
//   - 只有当时间在开始时间之后，任务才会开始计数
func WithStartTime(startTime time.Time) Option {
	return func(task *Task) {
		task.start = startTime
	}
}

// WithOffsetTime 通过指定偏移时间的方式创建任务
func WithOffsetTime(offset *offset.Time) Option {
	return func(task *Task) {
		task.offset = offset
	}
}

// WithLimitedTime 通过限时的方式创建任务
func WithLimitedTime(limitTime time.Duration) Option {
	return func(task *Task) {
		task.limitTime = limitTime
	}
}

// WithFront 通过指定任务前置任务的方式创建任务
//   - 当前置任务未完成时，当前任务不会开始计数
func WithFront(fronts ...*Task) Option {
	return func(task *Task) {
		if task.fronts == nil {
			task.fronts = make(map[int64]*Task)
		}
		for _, front := range fronts {
			task.fronts[front.GetID()] = front
		}
	}
}
