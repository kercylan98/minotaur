package prc

type MessageWrapper struct {
	Sender   *ProcessId
	Receiver *ProcessId
	Message  Message
}

func WrapMessage(sender, receiver *ProcessId, message Message) *MessageWrapper {
	return &MessageWrapper{
		Sender:   sender,
		Receiver: receiver,
		Message:  message,
	}
}

func UnwrapMessage(wrapper Message) (sender, receiver *ProcessId, message Message) {
	w, ok := wrapper.(*MessageWrapper)
	if !ok {
		return nil, nil, message
	}
	return w.Sender, w.Receiver, w.Message
}

func wrapDeliveryMessage(name string, data []byte, system bool, sender, receiver *ProcessId) *deliveryMessage {
	return &deliveryMessage{
		MessageType: name,
		MessageData: data,
		System:      system,
		Sender:      sender,
		Receiver:    receiver,
	}
}
