package vivid

type ActorProducer func(options *ActorOptions) Actor

type Actor interface {
	OnReceive(ctx ActorContext)
}
