package matrix

// Match3Item 三消成员接口定义
type Match3Item[Type comparable] interface {
	// SetGuid 设置guid
	SetGuid(guid int64)
	// GetGuid 获取guid
	GetGuid() int64
	// GetType 获取成员类型
	GetType() Type
	// Clone 克隆
	Clone() Match3Item[Type]
}
