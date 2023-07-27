package room

// Room 房间类似于简版的游戏世界(World)，不过没有游戏实体
type Room interface {
	// GetGuid 获取房间的唯一标识符
	GetGuid() int64
}
