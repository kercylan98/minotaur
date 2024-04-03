package queue

// MessageHandler 消息处理器支持传入两个函数对消息进行处理
//   - 在 handler 内可以执行对消息的逻辑
//   - 在 finisher 函数中可以接收到该消息是否是最后一条消息
type MessageHandler[Id, Ident comparable, M Message] func(
	handler func(m MessageWrapper[Id, Ident, M]),
	finisher func(m MessageWrapper[Id, Ident, M], last bool),
)
