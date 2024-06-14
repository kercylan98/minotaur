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
	var core *_ActorCore
	var ok bool
	defer func() {
		if reason := recover(); reason != nil {
			ctx.GetSystem().GetLogger().Error("ActorPanic", log.Any("reason", reason))
			core.supervisorExec(core, ctx.GetMessage(), reason)
		}
	}()

	actorCtx := ctx.GetContext()
	if core, ok = actorCtx.(*_ActorCore); ok && core == nil {
		actor.OnReceive(ctx) // OnInit
		return
	}

	// 空闲时间重置
	switch ctx.GetMessage().(type) {
	case OnTerminate:
	default:
		core.refreshIdleTimeout()
		defer core.refreshIdleTimeout()
	}

	behavior := actorCtx.matchBehavior(ctx.GetMessage())
	if behavior == nil {
		actor.OnReceive(ctx)
		return
	}

	behavior.onHandler(ctx, ctx.GetMessage())
}
