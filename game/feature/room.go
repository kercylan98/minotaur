package feature

// Room 游戏房间接口定义
type Room[P Player] interface {
	// GetGuid 获取房间 guid
	GetGuid() int64
	// JoinRoom 加入房间
	JoinRoom(player P) error
	// LeaveRoom 离开房间
	LeaveRoom(guid int64)
	// GetPlayerMaximum 获取游戏参与人上限
	GetPlayerMaximum() int
	// GetPlayer 获取特定玩家
	GetPlayer(guid int64) P
	// GetPlayers 获取所有玩家
	GetPlayers() map[int64]P
	// GetPlayerCount 获取玩家数量
	GetPlayerCount() int
	// IsExist 玩家是否存在
	IsExist(playerGuid int64) bool
}
