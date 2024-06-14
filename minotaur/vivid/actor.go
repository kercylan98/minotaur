package vivid

import "github.com/kercylan98/minotaur/toolkit/log"

type Actor interface {
	OnReceive(ctx MessageContext)
}

type FreeActor[T any] struct {
	actor T
}

func (i *FreeActor[T]) OnReceive(ctx MessageContext) {

}

func (i *FreeActor[T]) GetActor() T {
	return i.actor
}

func onReceive(actor Actor, ctx MessageContext) {
	defer func() {
		if reason := recover(); reason != nil {
			ctx.GetSystem().GetLogger().Error("ActorPanic", log.Any("reason", reason))
			actorCtx := ctx.GetContext()
			actorCtx.supervisorExec(ctx.GetContext().(*_ActorCore), ctx.GetMessage(), reason)
		}
	}()

	actorCtx, _ := ctx.GetContext().(*_ActorCore)
	if actorCtx == nil {
		actor.OnReceive(ctx)
		return
	}

	behavior := actorCtx.matchBehavior(ctx.GetMessage())
	if behavior == nil {
		actor.OnReceive(ctx)
		return
	}

	behavior.onHandler(ctx, ctx.GetMessage())
}
