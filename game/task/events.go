package task

type (
	TaskDoneEventHandle func(eventId int64, task *Task)
)

var taskDoneEventHandles = make(map[int64]map[int64]TaskDoneEventHandle)

// RegTaskDoneEvent 注册任务完成事件
func RegTaskDoneEvent(id int64, eventId int64, handle TaskDoneEventHandle) {
	events, exist := taskDoneEventHandles[id]
	if !exist {
		events = map[int64]TaskDoneEventHandle{}
		taskDoneEventHandles[id] = events
	}
	events[eventId] = handle
}

// UnRegTaskDoneEvent 取消注册任务完成事件
func UnRegTaskDoneEvent(id int64, eventId int64) {
	events, exist := taskDoneEventHandles[id]
	if exist {
		delete(events, eventId)
	}
}

// OnTaskDoneEvent 任务完成事件
func OnTaskDoneEvent(task *Task) {
	for eventId, handle := range taskDoneEventHandles[task.id] {
		handle(eventId, task)
	}
}
