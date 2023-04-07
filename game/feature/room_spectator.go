package feature

// RoomSpectator 房间观众席接口定义
type RoomSpectator[P Player] interface {
	Room[P]
	// GetSpectatorMaximum 获取观众人数上限
	GetSpectatorMaximum() int
	// JoinSpectator 加入观众席
	JoinSpectator(player P) error
	// LeaveSpectator 离开观众席
	LeaveSpectator(guid int64) error
	// GetSpectatorPlayer 获取特定观众席玩家
	GetSpectatorPlayer(guid int64) P
	// GetSpectatorPlayers 获取观众席玩家
	GetSpectatorPlayers() map[int64]P
}
