package builtin

import "minotaur/game"

type WorldOption[PlayerID comparable, Player game.Player[PlayerID]] func(world *World[PlayerID, Player])

func WithWorldPlayerLimit[PlayerID comparable, Player game.Player[PlayerID]](playerLimit int) WorldOption[PlayerID, Player] {
	return func(world *World[PlayerID, Player]) {
		world.playerLimit = playerLimit
	}
}
