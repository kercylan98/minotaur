package vivid

type Message struct {
	Sender   ActorId
	Receiver ActorId
	Command  any
	Params   []any
	Results  []any
}
