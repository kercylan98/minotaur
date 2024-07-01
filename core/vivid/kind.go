package vivid

type Kind = string

type kind struct {
	producer ActorProducer
	options  *ActorOptions
}
