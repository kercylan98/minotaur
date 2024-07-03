package vivid

import (
	"github.com/kercylan98/minotaur/core"
)

type Module interface {
	OnLoad(support *ModuleSupport, hasTransportModule bool)
}

type TransportModule interface {
	Module

	ActorSystemAddress() core.Address
}

type PriorityModule interface {
	Module

	// Priority 模块的优先级，优先级越低越先加载，未实现的模块默认为 0
	Priority() int
}

type ShutdownModule interface {
	Module

	OnShutdown()
}

type KindHookModule interface {
	Module

	OnRegKind(kind Kind)
}
