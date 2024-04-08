package queue

// MessageHandler 消息处理器支持传入两个函数对消息进行处理
//   - 在 handler 内可以执行对消息的逻辑
//   - 在 finisher 函数中可以接收到该消息是否是最后一条消息
type MessageHandler[Id, Q comparable, M Message[Q]] func(
	handler func(m M),
	finisher func(m M, last bool),
)
