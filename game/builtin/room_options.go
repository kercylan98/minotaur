package builtin

import "minotaur/game"

// RoomOption 房间构建可选项
type RoomOption[PlayerID comparable, Player game.Player[PlayerID]] func(room *Room[PlayerID, Player])

// WithRoomPlayerLimit 限制房间的玩家数量上限
func WithRoomPlayerLimit[PlayerID comparable, Player game.Player[PlayerID]](playerLimit int) RoomOption[PlayerID, Player] {
	return func(room *Room[PlayerID, Player]) {
		room.playerLimit = playerLimit
	}
}
