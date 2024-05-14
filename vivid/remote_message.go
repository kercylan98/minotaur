package vivid

// RemoteMessageEvent 远程消息事件
type RemoteMessageEvent struct {
	Message  RemoteMessage
	Callback func(err error) // 回调函数，即便是 err 为 nil 也会触发回调
}

// RemoteMessage 远程消息是内部转换的可序列化的消息，被用于 ActorSystem 之间的通信
type RemoteMessage struct {
	Opts     *MessageOptions // 消息选项
	Sender   ActorId         // 发送者
	Receiver ActorId         // 接收者
	Message  any             // 消息
}

func newRemoteMessage(opts *MessageOptions, receiver ActorRef, msg Message) RemoteMessage {
	var m = RemoteMessage{
		Opts:     opts,
		Receiver: receiver.GetId(),
		Message:  msg,
	}
	if opts.Sender != nil {
		m.Sender = opts.Sender.GetId()
	}
	return m
}

func ParseRemoteMessage(data []byte, callback func(err error)) (RemoteMessageEvent, error) {
	var m RemoteMessage
	err := gob.Decode(data, &m)
	if err != nil {
		return RemoteMessageEvent{}, err
	}

	return RemoteMessageEvent{
		Message:  m,
		Callback: callback,
	}, nil
}
