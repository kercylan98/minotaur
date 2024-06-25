package vivid

import "github.com/kercylan98/minotaur/minotaur/core"

type Module interface {
	OnLoad(support *ModuleSupport)
}

type TransportModule interface {
	Module

	ActorSystemAddress() core.Address
}
