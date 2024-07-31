package vivid

import "github.com/kercylan98/minotaur/engine/prc"

type AbyssProcess interface {
	Initialize(system *ActorSystem)

	prc.UnboundProcess
}
