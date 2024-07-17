package vivid

type Actor interface {
	OnReceive(ctx ActorContext)
}
