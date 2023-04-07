package feature

// RoomReady 房间准备接口定义
type RoomReady[P Player] interface {
	Room[P]
	// Ready 设置玩家准备状态
	Ready(playerGuid int64, ready bool)
	// IsAllReady 是否全部玩家已准备
	IsAllReady() bool
	// GetReadyCount 获取已准备玩家数量
	GetReadyCount() int
	// GetUnready 获取未准备的玩家
	GetUnready() map[int64]P
}
