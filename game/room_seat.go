package game

// RoomSeat 带有座位号的房间实现
type RoomSeat[PlayerID comparable, P Player[PlayerID]] interface {
	Room[PlayerID, P]
	// AddSeat 将玩家添加到座位号中
	AddSeat(id PlayerID)
	// AddSeatWithAssign 将玩家添加到座位号中，并分配特定的座位号
	AddSeatWithAssign(id PlayerID, seat int)
	// RemovePlayerSeat 移除玩家的座位号
	RemovePlayerSeat(id PlayerID)
	// RemoveSeat 移除特定座位号
	RemoveSeat(seat int)
	// SetSeat 设置玩家座位号，当玩家没有座位号时，将会返回错误信息
	//  - 如果座位号有其他玩家，他们的位置将互换
	SetSeat(id PlayerID, seat int) error
	// GetSeat 获取玩家座位号
	GetSeat(id PlayerID) (int, error)
	// GetPlayerIDWithSeat 根据座位号获取玩家ID
	GetPlayerIDWithSeat(seat int) (PlayerID, error)
	// GetSeatInfo 获取座位信息，空缺的位置将为空
	GetSeatInfo() []*PlayerID
	// GetSeatInfoMap 以map的方式获取座位号
	GetSeatInfoMap() map[int]PlayerID
	// GetSeatInfoMapVacancy 以map的方式获取座位号，空缺的位置将被保留为nil
	GetSeatInfoMapVacancy() map[int]*PlayerID
	// GetSeatInfoWithPlayerIDMap 获取座位信息，将以玩家ID作为key
	GetSeatInfoWithPlayerIDMap() map[PlayerID]int
	// GetNextSeat 获取下一个座位号，空缺的位置将会被跳过
	//  - 超出范围将返回-1
	//  - 当没有下一个座位号时将始终返回本身
	GetNextSeat(seat int) int
	// GetNextSeatVacancy 获取下一个座位号，空缺的位置将被保留
	//  - 超出范围将返回-1
	//  - 当没有下一个座位号时将始终返回本身
	GetNextSeatVacancy(seat int) int
}
