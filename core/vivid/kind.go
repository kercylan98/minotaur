package vivid

type Kind = string

func newActorKind(producer ActorProducer, options *ActorOptions) *kind {
	return &kind{
		producer: producer,
		options:  options,
	}
}

type kind struct {
	producer ActorProducer
	options  *ActorOptions
}
