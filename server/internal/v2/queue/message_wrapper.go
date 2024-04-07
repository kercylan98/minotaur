package queue

func messageWrapper[Id, Ident comparable, M Message](queue *Queue[Id, Ident, M], hasIdent bool, ident Ident, msg M) MessageWrapper[Id, Ident, M] {
	return MessageWrapper[Id, Ident, M]{
		queue:    queue,
		hasIdent: hasIdent,
		ident:    ident,
		msg:      msg,
	}
}

// MessageWrapper 提供了对外部消息的包装，用于方便的获取消息信息
type MessageWrapper[Id, Ident comparable, M Message] struct {
	queue    *Queue[Id, Ident, M] // 处理消息的队列
	ident    Ident                // 消息所有人
	msg      M                    // 消息信息
	hasIdent bool                 // 是否拥有所有人
}

// Queue 返回处理该消息的队列
func (m MessageWrapper[Id, Ident, M]) Queue() *Queue[Id, Ident, M] {
	return m.queue
}

// Ident 返回消息的所有人
func (m MessageWrapper[Id, Ident, M]) Ident() Ident {
	return m.ident
}

// HasIdent 返回消息是否拥有所有人
func (m MessageWrapper[Id, Ident, M]) HasIdent() bool {
	return m.hasIdent
}

// Message 返回消息的具体实例
func (m MessageWrapper[Id, Ident, M]) Message() M {
	return m.msg
}
