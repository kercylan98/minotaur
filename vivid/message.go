package vivid

type Message struct {
	Sender   ActorId
	Receiver ActorId
	Command  string
	Params   []any
	Results  []any
}
