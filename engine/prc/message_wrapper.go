package prc

type MessageWrapper struct {
	Sender   *ProcessId
	Receiver *ProcessId
	Message  Message
	Seq      uint64
}

func WrapMessage(sender, receiver *ProcessId, message Message, seq uint64) *MessageWrapper {
	return &MessageWrapper{
		Sender:   sender,
		Receiver: receiver,
		Message:  message,
		Seq:      seq,
	}
}

func UnwrapMessage(wrapper Message) (sender, receiver *ProcessId, message Message, seq uint64) {
	w, ok := wrapper.(*MessageWrapper)
	if !ok {
		return nil, nil, message, 0
	}
	return w.Sender, w.Receiver, w.Message, w.Seq
}
