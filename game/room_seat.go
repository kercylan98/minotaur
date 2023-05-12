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
	// GetSeatInfo 获取座位信息，空缺的位置将为空
	GetSeatInfo() []*PlayerID
	// GetSeatInfoMap 以map的方式获取座位号
	GetSeatInfoMap() map[int]PlayerID
	// GetSeatInfoMapVacancy 以map的方式获取座位号，空缺的位置将被保留为nil
	GetSeatInfoMapVacancy() map[int]*PlayerID
	// GetSeatInfoWithPlayerIDMap 获取座位信息，将以玩家ID作为key
	GetSeatInfoWithPlayerIDMap() map[PlayerID]int
}
