package task

import (
	"time"
)

// Option 任务选项
type Option func(task *Task)

// WithType 设置任务类型
func WithType(taskType string) Option {
	return func(task *Task) {
		task.Type = taskType
	}
}

// WithCondition 设置任务完成条件，当满足条件时，任务状态为完成
//   - 任务条件值需要变更时可通过 Task.AssignConditionValueAndRefresh 方法变更
//   - 当多次设置该选项时，后面的设置会覆盖之前的设置
func WithCondition(condition Condition) Option {
	return func(task *Task) {
		if condition == nil {
			return
		}
		if task.Cond == nil {
			task.Cond = condition
			return
		}
		for k, v := range condition {
			task.Cond[k] = v
		}
	}
}

// WithCounter 设置任务计数器，当计数器达到要求时，任务状态为完成
//   - 一些场景下，任务计数器可能会溢出，此时可通过 WithOverflowCounter 设置可溢出的任务计数器
//   - 当多次设置该选项时，后面的设置会覆盖之前的设置
//   - 如果需要初始化计数器的值，可通过 initCount 参数设置
func WithCounter(counter int64, initCount ...int64) Option {
	return func(task *Task) {
		task.Counter = counter
		if len(initCount) > 0 {
			task.CurrCount = initCount[0]
			if task.CurrCount < 0 {
				task.CurrCount = 0
			} else if task.CurrCount > task.Counter {
				task.CurrCount = task.Counter
			}
		}
	}
}

// WithOverflowCounter 设置可溢出的任务计数器，当计数器达到要求时，任务状态为完成
//   - 当多次设置该选项时，后面的设置会覆盖之前的设置
//   - 如果需要初始化计数器的值，可通过 initCount 参数设置
func WithOverflowCounter(counter int64, initCount ...int64) Option {
	return func(task *Task) {
		task.Counter = counter
		task.CurrOverflow = true
		if len(initCount) > 0 {
			task.CurrCount = initCount[0]
			if task.CurrCount < 0 {
				task.CurrCount = 0
			}
		}
	}
}

// WithDeadline 设置任务截止时间，超过截至时间并且任务未完成时，任务状态为失败
func WithDeadline(deadline time.Time) Option {
	return func(task *Task) {
		task.Deadline = deadline
	}
}

// WithLimitedDuration 设置任务限时，超过限时时间并且任务未完成时，任务状态为失败
func WithLimitedDuration(start time.Time, duration time.Duration) Option {
	return func(task *Task) {
		task.StartTime = start
		task.LimitedDuration = duration
	}
}
