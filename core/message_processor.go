package core

// MessageProcessor 消息处理器是针对不同类型的消息进行处理的接口
type MessageProcessor interface {
	ProcessUserMessage(msg Message, recoverMessage ...Message)
	ProcessSystemMessage(msg Message)
	ProcessRecover(reason Message)
}
