package core

// MessageProcessor 消息处理器是针对不同类型的消息进行处理的接口
type MessageProcessor interface {
	// ProcessUserMessage 处理用户消息，recoverMessage 是实现方可选的行为，通常用于在处理完毕时将当前正在处理的消息恢复为原始状态
	ProcessUserMessage(msg Message, recoverMessage ...Message)

	// ProcessSystemMessage 处理系统消息
	ProcessSystemMessage(msg Message)

	// ProcessRecover 当处理消息的过程中引发了 panic，recover 会将 panic 的原因传递给 ProcessRecover
	ProcessRecover(reason Message)
}
