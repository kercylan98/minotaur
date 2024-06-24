package vivid

type ActorProducer func(options *ActorOptions) Actor

type OnReceiveFunc func(ctx ActorContext)

type Actor interface {
	OnReceive(ctx ActorContext)
}
