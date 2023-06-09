package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/synchronization"
)

// RoomOption 房间构建可选项
type RoomOption[PlayerID comparable, Player game.Player[PlayerID]] func(room *Room[PlayerID, Player])

// WithRoomSync 通过线程安全的方式创建房间
func WithRoomSync[PlayerID comparable, Player game.Player[PlayerID]]() RoomOption[PlayerID, Player] {
	return func(room *Room[PlayerID, Player]) {
		room.players = synchronization.NewMap[PlayerID, Player]()
	}
}

// WithRoomPlayerLimit 限制房间的玩家数量上限
func WithRoomPlayerLimit[PlayerID comparable, Player game.Player[PlayerID]](playerLimit int) RoomOption[PlayerID, Player] {
	return func(room *Room[PlayerID, Player]) {
		room.playerLimit = playerLimit
	}
}

// WithRoomNoMaster 设置房间为无主的
func WithRoomNoMaster[PlayerID comparable, Player game.Player[PlayerID]]() RoomOption[PlayerID, Player] {
	return func(room *Room[PlayerID, Player]) {
		room.noMaster = true
	}
}

// WithRoomKickPlayerCheckHandle 设置房间踢出玩家的检查处理函数
//   - 当没有设置该函数时，如果不是房主，将无法进行踢出
func WithRoomKickPlayerCheckHandle[PlayerID comparable, Player game.Player[PlayerID]](handle func(room *Room[PlayerID, Player], id, target PlayerID) error) RoomOption[PlayerID, Player] {
	return func(room *Room[PlayerID, Player]) {
		room.kickCheckHandle = handle
	}
}
