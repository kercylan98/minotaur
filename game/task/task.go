package task

// NewTask 新建任务
func NewTask(id int64, options ...Option) *Task {
	task := &Task{
		id: id,
	}
	for _, option := range options {
		option(task)
	}
	if task.count > task.done {
		task.count = task.done
	}
	task.Add(0)
	return task
}

type Task struct {
	id     int64 // 任务ID
	count  int   // 任务计数
	done   int   // 任务完成计数
	reward bool  // 是否已领取奖励
}

// Reset 重置任务
func (slf *Task) Reset() {
	slf.count = 0
	slf.reward = false
	switch slf.GetState() {
	case StateDone:
		OnTaskDoneEvent(slf)
	}
}

// Add 增加任务计数
func (slf *Task) Add(count int) {
	if count != 0 {
		slf.count += count
		if slf.count < 0 {
			slf.count = 0
		} else if slf.count > slf.done {
			slf.count = slf.done
		}
	}
	switch slf.GetState() {
	case StateDone:
		OnTaskDoneEvent(slf)
	}
}

// GetState 获取任务状态
func (slf *Task) GetState() State {
	if slf.count >= slf.done {
		if slf.reward {
			return StateReward
		}
		return StateDone
	}
	return StateDoing
}

// Reward 返回是否领取过奖励，并且设置任务为领取过奖励的状态
func (slf *Task) Reward() bool {
	reward := slf.reward
	slf.reward = true
	return reward
}
