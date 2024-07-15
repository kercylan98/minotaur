package vivid

type DeadLetterEvent struct {
	Sender   ActorRef
	Receiver ActorRef
	Message  Message
}
