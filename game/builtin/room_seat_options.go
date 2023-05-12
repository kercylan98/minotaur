package builtin

import "github.com/kercylan98/minotaur/game"

type RoomSeatOption[PlayerID comparable, Player game.Player[PlayerID]] func(seat *RoomSeat[PlayerID, Player])

// WithRoomSeatFillIn 通过补位的方式创建带有座位号的房间
//   - 默认情况下玩家离开座位不会影响其他玩家
//   - 补位情况下，靠前的玩家离开座位将有后方玩家向前补位
func WithRoomSeatFillIn[PlayerID comparable, Player game.Player[PlayerID]]() RoomSeatOption[PlayerID, Player] {
	return func(seatRoom *RoomSeat[PlayerID, Player]) {
		seatRoom.fillIn = true
	}
}
