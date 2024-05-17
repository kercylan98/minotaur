package vivid

import "github.com/panjf2000/ants/v2"

type ActorSystemExternal interface {
	// GetGoroutinePool 用于获取 goroutine 池
	GetGoroutinePool() *ants.Pool
}

type actorSystemExternal struct {
	system *ActorSystem
}

func (a *actorSystemExternal) init(system *ActorSystem) *actorSystemExternal {
	a.system = system
	return a
}

func (a *actorSystemExternal) GetGoroutinePool() *ants.Pool {
	return a.system.gp
}
