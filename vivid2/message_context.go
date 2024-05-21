package vivid

type MessageContext interface {
	ActorContext

	// GetReceiver 获取消息的接收者
	GetReceiver() ActorRef

	// GetMessage 获取消息内容
	GetMessage() Message

	// Reply 回复消息
	Reply(Message)
}

// _LocalMessageContext 本地消息上下文
type _LocalMessageContext struct {
	ActorContext // 接收者的上下文
	sender       ActorRef
	message      Message
	seq          uint64
	network      string
	host         string
	port         uint16
}

func (c *_LocalMessageContext) GetReceiver() ActorRef {
	return c.ActorContext.(*_ActorCore)
}

func (c *_LocalMessageContext) GetMessage() Message {
	return c.message
}

func (c *_LocalMessageContext) Reply(message Message) {
	if c.sender == nil {
		system := c.ActorContext.GetSystem()
		if c.network == system.network && c.host == system.host && c.port == system.port {
			system.askWaitsLock.RLock()
			wait, exist := system.askWaits[c.seq]
			system.askWaitsLock.RUnlock()
			if !exist {
				return
			}
			wait <- message
			return
		}
	}
	c.ActorContext.(*_ActorCore).system.sendMessage(c.sender, message, func(options *MessageOptions) {
		options.replySeq = c.seq
	})
}

// _RemoteMessageContext 远程消息上下文
type _RemoteMessageContext struct {
	system     *ActorSystem // Actor 系统
	ReceiverId ActorId      // 接收者 ID
	SenderId   ActorId      // 发送者 ID
	Message    Message      // 消息内容
	Seq        uint64       // 消息序号
	ReplySeq   uint64       // 回复消息序号
	Network    string       // 网络地址
	Host       string       // 主机地址
	Port       uint16       // 端口
}

func (c *_RemoteMessageContext) GetReceiver() ActorRef {
	return &_RemoteActorRef{
		system:  c.system,
		actorId: c.ReceiverId,
	}
}

func (c *_RemoteMessageContext) GetMessage() Message {
	return c.Message
}

func (c *_RemoteMessageContext) Reply(message Message) {
	panic("not implemented")
}
