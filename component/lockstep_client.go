package component

import "github.com/kercylan98/minotaur/game"

type LockstepClient[ID comparable] interface {
	game.Player[ID]
}
