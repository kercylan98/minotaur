package game

// Actor 表示游戏中的对象，具有唯一标识符
type Actor interface {
	// GetGuid 获取对象的唯一标识符
	GetGuid() int64
}
