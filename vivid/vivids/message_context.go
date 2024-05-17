package vivids

type MessageContext interface {
	ActorContext

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
