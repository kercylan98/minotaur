package vivid

type Actor interface {
	OnReceive(ctx ActorContext)
}

type FunctionalActor func(ctx ActorContext)

func (f FunctionalActor) OnReceive(ctx ActorContext) {
	f(ctx)
}
