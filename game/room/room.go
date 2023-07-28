package room

// Room 游戏房间接口
type Room interface {
	// GetGuid 获取房间的唯一标识符
	GetGuid() int64
}
