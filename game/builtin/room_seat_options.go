package builtin

import "github.com/kercylan98/minotaur/game"

type RoomSeatOption[PlayerID comparable, Player game.Player[PlayerID]] func(seat *RoomSeat[PlayerID, Player])

// WithRoomSeatAutoManage 通过自动管理的方式创建带有座位号的房间
//   - 默认情况下需要自行维护房间用户的座位号信息
//   - 自动管理模式下，将注册房间的 Room.RegPlayerJoinRoomEvent 和 Room.RegPlayerLeaveRoomEvent 事件以便玩家在加入或者离开时维护座位信息
func WithRoomSeatAutoManage[PlayerID comparable, Player game.Player[PlayerID]]() RoomSeatOption[PlayerID, Player] {
	return func(seatRoom *RoomSeat[PlayerID, Player]) {
		seatRoom.autoMode.Do(func() {
			seatRoom.RegPlayerJoinRoomEvent(seatRoom.onJoinRoom)
			seatRoom.RegPlayerLeaveRoomEvent(seatRoom.onLeaveRoom)
		})
	}
}

// WithRoomSeatFillIn 通过补位的方式创建带有座位号的房间
//   - 默认情况下玩家离开座位不会影响其他玩家
//   - 补位情况下，靠前的玩家离开座位将有后方玩家向前补位
func WithRoomSeatFillIn[PlayerID comparable, Player game.Player[PlayerID]]() RoomSeatOption[PlayerID, Player] {
	return func(seatRoom *RoomSeat[PlayerID, Player]) {
		seatRoom.fillIn = true
	}
}
