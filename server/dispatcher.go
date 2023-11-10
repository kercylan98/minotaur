package server

import "github.com/kercylan98/minotaur/utils/buffer"

// generateDispatcher 生成消息分发器
func generateDispatcher(handler func(message *Message)) *dispatcher {
	return &dispatcher{
		buffer:  buffer.NewUnboundedN[*Message](),
		handler: handler,
	}
}

// dispatcher 消息分发器
type dispatcher struct {
	buffer  *buffer.Unbounded[*Message]
	handler func(message *Message)
}

func (slf *dispatcher) start() {
	for {
		select {
		case message, ok := <-slf.buffer.Get():
			if !ok {
				return
			}
			slf.buffer.Load()
			slf.handler(message)
		}
	}
}

func (slf *dispatcher) put(message *Message) {
	slf.buffer.Put(message)
}

func (slf *dispatcher) close() {
	slf.buffer.Close()
}
