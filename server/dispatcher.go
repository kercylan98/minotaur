package server

import (
	"github.com/alphadose/haxmap"
)

var dispatcherUnique = struct{}{}

// generateDispatcher 生成消息分发器
func generateDispatcher(size int, name string, handler func(dispatcher *dispatcher, message *Message)) *dispatcher {
	return &dispatcher{
		name:    name,
		buffer:  make(chan *Message, size),
		handler: handler,
		uniques: haxmap.New[string, struct{}](),
	}
}

// dispatcher 消息分发器
type dispatcher struct {
	name    string
	buffer  chan *Message
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
		case message, ok := <-slf.buffer:
			if !ok {
				return
			}
			slf.handler(slf, message)
		}
	}
}

func (slf *dispatcher) put(message *Message) {
	slf.buffer <- message
}

func (slf *dispatcher) close() {
	close(slf.buffer)
}

func (slf *dispatcher) transfer(target *dispatcher) {
	if target == nil {
		return
	}
	for {
		select {
		case message, ok := <-slf.buffer:
			if !ok {
				return
			}
			target.buffer <- message
		}
	}
}
