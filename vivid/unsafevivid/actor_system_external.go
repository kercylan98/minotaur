package unsafevivid

import "github.com/panjf2000/ants/v2"

type ActorSystemExternal struct {
	system *ActorSystem
}

func (a *ActorSystemExternal) init(system *ActorSystem) *ActorSystemExternal {
	a.system = system
	return a
}

func (a *ActorSystemExternal) GetGoroutinePool() *ants.Pool {
	return a.system.gp
}
