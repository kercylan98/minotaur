package game

// Actor 表示游戏中的对象，具有唯一标识符
type Actor interface {
	// SetGuid 设置对象的唯一标识符
	//  - 需要注意的是该函数不应该主动执行，否则可能产生意想不到的情况
	SetGuid(guid int64)
	// GetGuid 获取对象的唯一标识符
	GetGuid() int64
}
