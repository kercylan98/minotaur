package server

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/utils/buffer"
	"sync/atomic"
	"time"
)

var dispatcherUnique = struct{}{}

// generateDispatcher 生成消息分发器
func generateDispatcher(name string, handler func(dispatcher *dispatcher, message *Message)) *dispatcher {
	d := &dispatcher{
		name:    name,
		buf:     buffer.NewUnbounded[*Message](),
		handler: handler,
		uniques: haxmap.New[string, struct{}](),
	}
	return d
}

// dispatcher 消息分发器
type dispatcher struct {
	name     string
	buf      *buffer.Unbounded[*Message]
	uniques  *haxmap.Map[string, struct{}]
	handler  func(dispatcher *dispatcher, message *Message)
	closed   uint32
	msgCount int64
}

func (d *dispatcher) unique(name string) bool {
	_, loaded := d.uniques.GetOrSet(name, dispatcherUnique)
	return loaded
}

func (d *dispatcher) antiUnique(name string) {
	d.uniques.Del(name)
}

func (d *dispatcher) start() {
	defer d.buf.Close()
	for {
		select {
		case message, ok := <-d.buf.Get():
			if !ok {
				return
			}
			d.buf.Load()
			d.handler(d, message)

			if atomic.AddInt64(&d.msgCount, -1) <= 0 && atomic.LoadUint32(&d.closed) == 1 {
				return
			}
		}
	}
}

func (d *dispatcher) put(message *Message) {
	if atomic.CompareAndSwapUint32(&d.closed, 1, 1) {
		return
	}
	atomic.AddInt64(&d.msgCount, 1)
	d.buf.Put(message)
}

func (d *dispatcher) close() {
	atomic.CompareAndSwapUint32(&d.closed, 0, 1)

	go func() {
		for {
			if d.buf.IsClosed() {
				return
			}
			if atomic.LoadInt64(&d.msgCount) <= 0 {
				d.buf.Close()
				return
			}
			time.Sleep(time.Second)
		}
	}()
}
