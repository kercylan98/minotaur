package vivid

import "reflect"

var messageContextType = reflect.TypeOf((*MessageContext)(nil)).Elem()

type MessageContext interface {

	// GetSeq 用于获取消息序号
	GetSeq() uint64

	// Reply 用于回复消息
	Reply(msg Message) error

	// GetSenderId 用于获取消息发送者
	GetSenderId() ActorId

	// GetReceiverId 用于获取消息接收者
	GetReceiverId() ActorId

	// GetMessage 用于获取消息
	GetMessage() Message
}

func newMessageContext(system *ActorSystem, senderId, receiverId ActorId, message Message, options *MessageOptions) MessageContext {
	return &messageContext{
		system:     system,
		Options:    options,
		Message:    message,
		SenderId:   senderId,
		ReceiverId: receiverId,
		Seq:        system.seq.Add(1),
	}
}

type messageContext struct {
	system *ActorSystem // 生产该消息的 ActorSystem，不参与序列化

	Options      *MessageOptions
	SenderId     ActorId // 隐式发送者
	ReceiverId   ActorId // 接收者
	Message      Message // 消息
	ReplyMessage Message // 回复的消息
	Seq          uint64  // 消息序号（响应消息 Seq 同请求消息 Seq）
}

func (c *messageContext) Reply(msg Message) error {
	var clone = &messageContext{
		Options:      c.Options,
		SenderId:     c.ReceiverId,
		ReceiverId:   c.SenderId,
		Message:      nil,
		ReplyMessage: msg,
		Seq:          c.Seq,
	}

	return c.system.sendCtx(c.ReceiverId, clone)
}

func (c *messageContext) GetSeq() uint64 {
	return c.Seq
}

func (c *messageContext) GetSenderId() ActorId {
	return c.Options.SenderId
}

func (c *messageContext) GetReceiverId() ActorId {
	return c.ReceiverId
}

func (c *messageContext) GetMessage() Message {
	return c.Message
}

func parseMessageContext(system *ActorSystem, data []byte) (MessageContext, error) {
	ctx := new(messageContext)
	if err := gob.Decode(data, ctx); err != nil {
		return nil, err
	}
	ctx.system = system
	return ctx, nil
}
