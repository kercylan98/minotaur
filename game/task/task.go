package task

import (
	"time"
)

// NewTask 生成任务
func NewTask(options ...Option) *Task {
	task := new(Task)
	for _, option := range options {
		option(task)
	}
	return task.refreshTaskStatus()
}

// Task 是对任务信息进行描述和处理的结构体
type Task struct {
	Type            string        `json:"type,omitempty"`             // 任务类型
	Status          Status        `json:"status,omitempty"`           // 任务状态
	Cond            Condition     `json:"cond,omitempty"`             // 任务条件
	CondValue       map[any]any   `json:"cond_value,omitempty"`       // 任务条件值
	Counter         int64         `json:"counter,omitempty"`          // 任务要求计数器
	CurrCount       int64         `json:"curr_count,omitempty"`       // 任务当前计数
	CurrOverflow    bool          `json:"curr_overflow,omitempty"`    // 任务当前计数是否允许溢出
	Deadline        time.Time     `json:"deadline,omitempty"`         // 任务截止时间
	StartTime       time.Time     `json:"start_time,omitempty"`       // 任务开始时间
	LimitedDuration time.Duration `json:"limited_duration,omitempty"` // 任务限时
}

// IsComplete 判断任务是否已完成
func (slf *Task) IsComplete() bool {
	return slf.Status == StatusComplete
}

// IsFailed 判断任务是否已失败
func (slf *Task) IsFailed() bool {
	return slf.Status == StatusFailed
}

// IsReward 判断任务是否已领取奖励
func (slf *Task) IsReward() bool {
	return slf.Status == StatusReward
}

// ReceiveReward 领取任务奖励，当任务状态为已完成时，才能领取奖励，此时返回 true，并且任务状态变更为已领取奖励
func (slf *Task) ReceiveReward() bool {
	if slf.Status != StatusComplete {
		return false
	}
	slf.Status = StatusReward
	return true
}

// IncrementCounter 增加计数器的值，当 incr 为负数时，计数器的值不会发生变化
//   - 如果需要溢出计数器，可通过 WithOverflowCounter 设置可溢出的任务计数器
func (slf *Task) IncrementCounter(incr int64) *Task {
	if incr < 0 {
		return slf
	}
	slf.CurrCount += incr
	if !slf.CurrOverflow && slf.CurrCount > slf.Counter {
		slf.CurrCount = slf.Counter
	}
	return slf.refreshTaskStatus()
}

// DecrementCounter 减少计数器的值，当 decr 为负数时，计数器的值不会发生变化
func (slf *Task) DecrementCounter(decr int64) *Task {
	if decr < 0 {
		return slf
	}
	slf.CurrCount -= decr
	if slf.CurrCount < 0 {
		slf.CurrCount = 0
	}
	return slf.refreshTaskStatus()
}

// AssignConditionValueAndRefresh 分配条件值并刷新任务状态
func (slf *Task) AssignConditionValueAndRefresh(key, value any) *Task {
	if slf.Cond == nil {
		return slf
	}
	if _, exist := slf.Cond[key]; !exist {
		return slf
	}
	if slf.CondValue == nil {
		slf.CondValue = make(map[any]any)
	}
	slf.CondValue[key] = value
	return slf.refreshTaskStatus()
}

// AssignConditionValueAndRefreshByCondition 分配条件值并刷新任务状态
func (slf *Task) AssignConditionValueAndRefreshByCondition(condition Condition) *Task {
	if slf.Cond == nil {
		return slf
	}
	if slf.CondValue == nil {
		slf.CondValue = make(map[any]any)
	}
	for k, v := range condition {
		if _, exist := slf.Cond[k]; !exist {
			continue
		}
		slf.CondValue[k] = v
	}
	return slf.refreshTaskStatus()
}

// ResetStatus 重置任务状态
//   - 该函数会将任务状态重置为已接受状态后，再刷新任务状态
//   - 当任务条件变更，例如任务计数要求为 10，已经完成的情况下，将任务计数要求变更为 5 或 20，此时任务状态由于是已完成或已领取状态，不会自动刷新，需要调用该函数刷新任务状态
func (slf *Task) ResetStatus() *Task {
	slf.Status = StatusAccept
	return slf.refreshTaskStatus()
}

// refreshTaskStatus 刷新任务状态
func (slf *Task) refreshTaskStatus() *Task {
	curr := time.Now()
	if (!slf.StartTime.IsZero() && curr.Before(slf.StartTime)) || (!slf.Deadline.IsZero() && curr.After(slf.Deadline)) || slf.Status >= StatusComplete {
		return slf
	}
	slf.Status = StatusComplete

	if slf.Counter > 0 && slf.CurrCount < slf.Counter {
		slf.Status = StatusAccept
		return slf
	}
	if slf.Cond != nil {
		for k, v := range slf.Cond {
			if v != slf.CondValue[k] {
				slf.Status = StatusAccept
				return slf
			}
		}
	}
	if !slf.Deadline.IsZero() && slf.Status == StatusAccept {
		if slf.Deadline.After(curr) {
			slf.Status = StatusFailed
			return slf
		}
	}
	if slf.LimitedDuration > 0 && slf.Status == StatusAccept {
		if curr.Sub(slf.StartTime) > slf.LimitedDuration {
			slf.Status = StatusFailed
			return slf
		}
	}
	return slf
}
