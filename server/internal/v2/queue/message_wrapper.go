package queue

// MessageWrapper 提供了对外部消息的包装，用于方便的获取消息信息
type MessageWrapper[Id, Ident comparable, M Message] struct {
	queue *Queue[Id, Ident, M] // 处理消息的队列
	ident Ident                // 消息所有人
	msg   M                    // 消息信息
}

// Queue 返回处理该消息的队列
func (m MessageWrapper[Id, Ident, M]) Queue() *Queue[Id, Ident, M] {
	return m.queue
}

// Ident 返回消息的所有人
func (m MessageWrapper[Id, Ident, M]) Ident() Ident {
	return m.ident
}

// Message 返回消息的具体实例
func (m MessageWrapper[Id, Ident, M]) Message() M {
	return m.msg
}
