package room

import "github.com/kercylan98/minotaur/game"

type info[PlayerID comparable, P game.Player[PlayerID], R Room[PlayerID, P]] struct {
	room        R
	playerLimit int // 玩家人数上限, <= 0 表示无限制
}
