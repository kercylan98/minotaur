package cluster

import "github.com/kercylan98/minotaur/engine/vivid"

type ActorContext interface {
	vivid.ActorContext
	Cluster() *ActorSystem
}

func newActorContext(system *ActorSystem, ctx vivid.ActorContext) *actorContext {
	return &actorContext{
		ActorContext: ctx,
	}
}

type actorContext struct {
	vivid.ActorContext
	system *ActorSystem
}

func (a *actorContext) Cluster() *ActorSystem {
	return a.system
}
