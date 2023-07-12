package task

type (
	RefreshTaskCountEvent      func(taskType int, increase int64)
	RefreshTaskChildCountEvent func(taskType int, key any, increase int64)
)

var (
	refreshTaskCountEventHandles      = make(map[int]RefreshTaskCountEvent)
	refreshTaskChildCountEventHandles = make(map[int]RefreshTaskChildCountEvent)
)

// RegRefreshTaskCount 注册任务计数刷新事件
func RegRefreshTaskCount(taskType int, handler RefreshTaskCountEvent) {
	refreshTaskCountEventHandles[taskType] = handler
}

// OnRefreshTaskCount 触发任务计数刷新事件
func OnRefreshTaskCount(taskType int, increase int64) {
	if handler, ok := refreshTaskCountEventHandles[taskType]; ok {
		handler(taskType, increase)
	}
}

// RegRefreshTaskChildCount 注册任务子计数刷新事件
func RegRefreshTaskChildCount(taskType int, handler RefreshTaskChildCountEvent) {
	refreshTaskChildCountEventHandles[taskType] = handler
}

// OnRefreshTaskChildCount 触发任务子计数刷新事件
func OnRefreshTaskChildCount(taskType int, key any, increase int64) {
	if handler, ok := refreshTaskChildCountEventHandles[taskType]; ok {
		handler(taskType, key, increase)
	}
}
