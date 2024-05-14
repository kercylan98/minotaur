package vivid

// RemoteMessageEvent 远程消息事件
type RemoteMessageEvent struct {
	Message  RemoteMessage
	Callback func(err error) // 回调函数，即便是 err 为 nil 也会触发回调
}

func (e RemoteMessageEvent) reply(system *ActorSystem, msg Message) error {

	senderRef, err := system.GetActor(e.Message.Sender)
	if err != nil {
		return err
	}

	data, err := gob.Encode(newRemoteMessage(system, senderRef, msg, new(MessageOptions).apply()))
	if err != nil {
		return err
	}

	cli, err := system.opts.ClientFactory(system.opts.Network, e.Message.Receiver.Host(), e.Message.Receiver.Port())
	if err != nil {
		return err
	}

	return cli.Tell(data)
}

// RemoteMessage 远程消息是内部转换的可序列化的消息，被用于 ActorSystem 之间的通信
type RemoteMessage struct {
	Opts     *MessageOptions // 消息选项
	Sender   ActorId         // 发送者
	Receiver ActorId         // 接收者
	Message  any             // 消息
	Seq      uint64          // 序号
}

func newRemoteMessage(system *ActorSystem, receiver ActorRef, msg Message, opts *MessageOptions) RemoteMessage {
	var m = RemoteMessage{
		Opts:     opts,
		Receiver: receiver.GetId(),
		Message:  msg,
		Seq:      system.seq.Add(1),
	}
	if opts != nil && opts.sender != nil {
		m.Sender = opts.sender.GetId()
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
