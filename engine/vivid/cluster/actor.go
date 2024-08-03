package cluster

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

type Actor interface {
	OnReceive(ctx ActorContext)
}

type FunctionalActor func(ctx ActorContext)

func (f FunctionalActor) OnReceive(ctx ActorContext) {
	f(ctx)
}

func newActor(system *ActorSystem, provider ActorProvider) *actor {
	return &actor{
		system:   system,
		provider: provider,
	}
}

// 集群内的 Actor，用于将 vivid.Actor 转换为 Actor
type actor struct {
	system   *ActorSystem
	provider ActorProvider
	ctx      *actorContext
	actor    Actor
}

func (a *actor) OnReceive(ctx vivid.ActorContext) {
	switch ctx.Message().(type) {
	case *vivid.OnLaunch:
		a.ctx = newActorContext(a.system, ctx)
		a.actor = a.provider.Provide()
		a.actor.OnReceive(a.ctx)
	default:
		a.actor.OnReceive(a.ctx)
	}
}
