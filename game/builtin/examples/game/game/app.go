package game

import (
	"minotaur/game"
	"minotaur/game/builtin"
	"minotaur/server"
)

var (
	Server *server.Server
	Game   *app
)

func init() {
	Server = server.New(server.NetworkTCP)
	Game = &app{
		World: builtin.NewWorld[int64, *Player](0),
	}
}

type app struct {
	game.World[int64, *Player]
}
