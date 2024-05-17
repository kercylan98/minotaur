package vivids

type ActorSystem interface {
	GetName() string

	Run() error

	Shutdown() error

	ActorOf(actor Actor, opts ...*ActorOptions) (ActorRef, error)

	GetActor() Query

	Tell(receiver ActorId, msg Message, opts ...MessageOption) error

	Ask(receiver ActorId, msg Message, opts ...MessageOption) (Message, error)
}
