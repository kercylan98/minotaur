package unsafevivid

import vivid "github.com/kercylan98/minotaur/vivid/vivids"

func NewMessageContext(system *ActorSystem, senderId, receiverId vivid.ActorId, message vivid.Message, options *vivid.MessageOptions) *MessageContext {
	return &MessageContext{
		system:     system,
		Options:    options,
		Message:    message,
		SenderId:   senderId,
		ReceiverId: receiverId,
		Seq:        system.seq.Add(1),
	}
}

func NewMessageContextWithBytes(system *ActorSystem, data []byte) (*MessageContext, error) {
	ctx := new(MessageContext)
	if err := gob.Decode(data, ctx); err != nil {
		return nil, err
	}
	ctx.system = system
	return ctx, nil
}

type MessageContext struct {
	vivid.ActorContext              // 消息接收者的 ActorContext，不参与序列化
	system             *ActorSystem // 生产该消息的 ActorSystem，不参与序列化

	Options       *vivid.MessageOptions // 消息选项
	SenderId      vivid.ActorId         // 隐式发送者
	ReceiverId    vivid.ActorId         // 接收者
	Message       vivid.Message         // 消息
	ReplyMessage  vivid.Message         // 回复的消息
	Seq           uint64                // 消息序号（响应消息 Seq 同请求消息 Seq）
	RemoteNetwork string                // 远程网络地址
	RemoteHost    string                // 远程主机地址
	RemotePort    uint16                // 远程端口
}

func (c *MessageContext) Reply(msg vivid.Message) error {
	var clone = &MessageContext{
		Options:      c.Options,
		SenderId:     c.ReceiverId,
		ReceiverId:   c.SenderId,
		Message:      nil,
		ReplyMessage: msg,
		Seq:          c.Seq,
	}
	if clone.ReceiverId == "" && c.RemoteNetwork != "" && c.RemoteHost != "" {
		// 匿名远程 Actor，通过网络发送消息
		client, err := c.system.opts.ClientFactory(c.system, c.RemoteNetwork, c.RemoteHost, c.RemotePort)
		if err != nil {
			return err
		}
		data, err := gob.Encode(clone)
		if err != nil {
			return err
		}
		return client.Exec(data)
	}

	// 本地回复
	c.system.replyWaitersLock.Lock()
	waiter, exist := c.system.replyWaiters[clone.Seq]
	c.system.replyWaitersLock.Unlock()
	if exist {
		waiter <- clone.ReplyMessage
	}
	return nil
}

func (c *MessageContext) GetSeq() uint64 {
	return c.Seq
}

func (c *MessageContext) GetSenderId() vivid.ActorId {
	return c.Options.SenderId
}

func (c *MessageContext) GetReceiverId() vivid.ActorId {
	return c.ReceiverId
}

func (c *MessageContext) GetMessage() vivid.Message {
	return c.Message
}
