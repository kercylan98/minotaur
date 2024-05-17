package vivids

// ActorRef 是 Actor 的引用
type ActorRef interface {
	// GetId 用于获取 Actor 的 ActorId
	GetId() ActorId

	// Tell 用于向 Actor 发送消息
	Tell(msg Message, opts ...MessageOption) error

	// Ask 用于向 Actor 发送消息并等待返回结果
	Ask(msg Message, opts ...MessageOption) (Message, error)
}
