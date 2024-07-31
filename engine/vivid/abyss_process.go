package vivid

import "github.com/kercylan98/minotaur/engine/prc"

type AbyssProcess interface {
	OnInitialize(system *ActorSystem)
	prc.Process
}
