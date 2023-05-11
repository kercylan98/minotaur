package game

// RoomSeat 带有座位号的房间实现
type RoomSeat[PlayerID comparable, P Player[PlayerID]] interface {
	Room[PlayerID, P]
	// SetSeat 设置玩家座位号
	SetSeat(id PlayerID, seat int) error
	// GetSeat 获取玩家座位号
	GetSeat(id PlayerID) (int, error)
	// GetPlayerWithSeat 根据座位号获取玩家
	GetPlayerWithSeat(seat int) (P, error)
}
