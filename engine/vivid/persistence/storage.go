package persistence

// Storage 是用作 Actor 状态持久化的存储接口
type Storage interface {
	// Save 保存特定名称的状态快照及事件
	Save(name Name, snapshot Snapshot, events []Event) error

	// Load 加载特定名称的状态快照和事件
	Load(name Name) (snapshot Snapshot, events []Event, err error)
}
