package vivid

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
