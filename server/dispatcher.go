package server

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/utils/buffer"
)

var dispatcherUnique = struct{}{}

// generateDispatcher 生成消息分发器
func generateDispatcher(name string, handler func(dispatcher *dispatcher, message *Message)) *dispatcher {
	return &dispatcher{
		name:    name,
		buffer:  buffer.NewUnboundedN[*Message](),
		handler: handler,
		uniques: haxmap.New[string, struct{}](),
	}
}

// dispatcher 消息分发器
type dispatcher struct {
	name    string
	buffer  *buffer.Unbounded[*Message]
	uniques *haxmap.Map[string, struct{}]
	handler func(dispatcher *dispatcher, message *Message)
}

func (slf *dispatcher) unique(name string) bool {
	_, loaded := slf.uniques.GetOrSet(name, dispatcherUnique)
	return loaded
}

func (slf *dispatcher) antiUnique(name string) {
	slf.uniques.Del(name)
}

func (slf *dispatcher) start() {
	for {
		select {
		case message, ok := <-slf.buffer.Get():
			if !ok {
				return
			}
			slf.buffer.Load()
			slf.handler(slf, message)
		}
	}
}

func (slf *dispatcher) put(message *Message) {
	slf.buffer.Put(message)
}

func (slf *dispatcher) close() {
	slf.buffer.Close()
}
