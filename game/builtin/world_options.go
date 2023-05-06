package builtin

import "github.com/kercylan98/minotaur/game"

// WorldOption 世界构建可选项
type WorldOption[PlayerID comparable, Player game.Player[PlayerID]] func(world *World[PlayerID, Player])

// WithWorldPlayerLimit 限制世界的玩家数量上限
func WithWorldPlayerLimit[PlayerID comparable, Player game.Player[PlayerID]](playerLimit int) WorldOption[PlayerID, Player] {
	return func(world *World[PlayerID, Player]) {
		world.playerLimit = playerLimit
	}
}
