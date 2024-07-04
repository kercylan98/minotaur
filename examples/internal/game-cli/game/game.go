package game

import (
	"game/game/views/login"
	"github.com/kercylan98/minotaur/core/vivid"
)

type Game struct {
	*vivid.ActorSystem
}

func Run() {
	login.Login()
}
