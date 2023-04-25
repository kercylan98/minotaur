package builtin

type WorldOption[PlayerID comparable] func(world *World[PlayerID])

func WithWorldPlayerLimit[PlayerID comparable](playerLimit int) WorldOption[PlayerID] {
	return func(world *World[PlayerID]) {
		world.playerLimit = playerLimit
	}
}
