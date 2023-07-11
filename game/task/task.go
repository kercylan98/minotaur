package task

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/offset"
	"time"
)

// NewTask 创建任务
func NewTask(id int64, taskType int, condition int64, options ...Option) *Task {
	task := &Task{
		id:        id,
		condition: condition,
		state:     StateAccept,
	}
	for _, option := range options {
		option(task)
	}
	if task.start.IsZero() {
		if task.offset != nil {
			task.start = task.offset.Now()
		} else {
			task.start = time.Now()
		}
	}
	for key := range task.childCount {
		if !hash.Exist(task.childCondition, key) {
			delete(task.childCount, key)
		}
	}
	if task.count == task.condition {
		task.state = StateFinish
	}
	return task
}

// Task 通用任务数据结构
type Task struct {
	id                       int64           // 任务ID
	taskType                 int             // 任务类型
	count                    int64           // 任务主计数
	condition                int64           // 任务完成需要的计数条件
	childCount               map[any]int64   // 任务子计数
	childCondition           map[any]int64   // 任务子计数条件
	state                    State           // 任务状态
	start                    time.Time       // 任务开始时间
	limitTime                time.Duration   // 任务限时
	fronts                   map[int64]*Task // 任务前置任务
	disableNotStartGetReward bool            // 禁止未开始的任务领取奖励

	offset *offset.Time // 任务偏移时间
}

// GetID 获取任务ID
func (slf *Task) GetID() int64 {
	return slf.id
}

// GetType 获取任务类型
func (slf *Task) GetType() int {
	return slf.taskType
}

// Reset 重置任务
func (slf *Task) Reset() {
	slf.count = 0
	slf.state = StateAccept
	for key := range slf.childCount {
		delete(slf.childCount, key)
	}
}

// GetFronts 获取前置任务
func (slf *Task) GetFronts() map[int64]*Task {
	return slf.fronts
}

// GetFrontsWithState 获取特定状态的前置任务
func (slf *Task) GetFrontsWithState(state State) map[int64]*Task {
	fronts := make(map[int64]*Task)
	for id, front := range slf.fronts {
		if front.GetState() == state {
			fronts[id] = front
		}
	}
	return fronts
}

// FrontsIsFinish 判断前置任务是否完成
func (slf *Task) FrontsIsFinish() bool {
	for _, front := range slf.fronts {
		state := front.GetState()
		if state == StateAccept || state == StateFail {
			return false
		}
	}
	return true
}

// GetReward 获取任务奖励
//   - 当任务状态为 StateFinish 时，调用 rewardHandle 函数
//   - 当任务状态不为 StateFinish 或奖励函数发生错误时，返回错误
func (slf *Task) GetReward(rewardHandle func() error) error {
	if !slf.IsStart() {
		return ErrTaskNotStart
	}
	switch slf.GetState() {
	case StateAccept:
		return ErrTaskNotFinish
	case StateReward:
		return ErrTaskRewardReceived
	case StateFail:
		return ErrTaskFail
	}
	if err := rewardHandle(); err != nil {
		return err
	}
	slf.state = StateReward
	return nil
}

// GetState 获取任务状态
func (slf *Task) GetState() State {
	return slf.state
}

// IsStart 判断任务是否开始
func (slf *Task) IsStart() bool {
	var current time.Time
	if slf.offset != nil {
		current = slf.offset.Now()
	} else {
		current = time.Now()
	}
	if current.Before(slf.start) {
		return false
	} else if slf.limitTime > 0 && current.Sub(slf.start) >= slf.limitTime {
		return false
	}
	return true
}

// SetCount 设置计数
func (slf *Task) SetCount(count int64) {
	if !slf.IsStart() || !slf.FrontsIsFinish() {
		return
	}
	slf.count = count
	if slf.count >= slf.condition {
		slf.count = slf.condition
	} else if slf.count < 0 {
		slf.count = 0
	}
	slf.refreshState()
}

// AddCount 增加计数
func (slf *Task) AddCount(count int64) {
	slf.SetCount(slf.count + count)
}

// GetCount 获取计数
func (slf *Task) GetCount() int64 {
	return slf.count
}

// GetCondition 获取计数条件
func (slf *Task) GetCondition() int64 {
	return slf.condition
}

// SetChildCount 设置子计数
func (slf *Task) SetChildCount(key any, count int64) {
	if !slf.IsStart() || !slf.FrontsIsFinish() || !hash.Exist(slf.childCondition, key) {
		return
	}
	if condition := slf.childCondition[key]; count > condition {
		count = condition
	} else if count < 0 {
		count = 0
	}
	slf.childCount[key] = count
	slf.refreshState()
}

// AddChildCount 增加子计数
func (slf *Task) AddChildCount(key any, count int64) {
	slf.SetChildCount(key, slf.childCount[key]+count)
}

// refreshState 刷新任务状态
func (slf *Task) refreshState() {
	slf.state = StateFinish
	if slf.count != slf.condition {
		slf.state = StateAccept
		return
	}
	for key, condition := range slf.childCondition {
		if slf.childCount[key] != condition {
			slf.state = StateAccept
			return
		}
	}
}
