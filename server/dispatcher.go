package server

import (
	"context"
	"github.com/alphadose/haxmap"
	"sync"
)

var dispatcherUnique = struct{}{}

// generateDispatcher 生成消息分发器
func generateDispatcher(size int, name string, handler func(dispatcher *dispatcher, message *Message)) *dispatcher {
	d := &dispatcher{
		name:       name,
		buffer:     make(chan *Message, size),
		handler:    handler,
		uniques:    haxmap.New[string, struct{}](),
		queueMutex: new(sync.Mutex),
	}
	d.ctx, d.cancel = context.WithCancel(context.Background())
	d.queueCond = sync.NewCond(d.queueMutex)
	return d
}

// dispatcher 消息分发器
type dispatcher struct {
	name       string
	buffer     chan *Message
	uniques    *haxmap.Map[string, struct{}]
	handler    func(dispatcher *dispatcher, message *Message)
	ctx        context.Context
	cancel     context.CancelFunc
	queue      []*Message
	queueMutex *sync.Mutex
	queueCond  *sync.Cond
}

func (d *dispatcher) unique(name string) bool {
	_, loaded := d.uniques.GetOrSet(name, dispatcherUnique)
	return loaded
}

func (d *dispatcher) antiUnique(name string) {
	d.uniques.Del(name)
}

func (d *dispatcher) start() {
	d.process()
	for {
		select {
		case message, ok := <-d.buffer:
			if !ok {
				return
			}
			d.handler(d, message)
		}
	}
}

func (d *dispatcher) process() {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				d.queueMutex.Lock()
				if len(d.queue) == 0 {
					d.queueCond.Wait()
				}
				messages := make([]*Message, len(d.queue))
				copy(messages, d.queue)
				d.queue = d.queue[:0]
				d.queueMutex.Unlock()
				for _, message := range messages {
					select {
					case d.buffer <- message:
					}
				}
			}
		}
	}(d.ctx)
}

func (d *dispatcher) put(message *Message) {
	d.queueMutex.Lock()
	d.queue = append(d.queue, message)
	d.queueCond.Signal()
	defer d.queueMutex.Unlock()
}

func (d *dispatcher) close() {
	close(d.buffer)
	d.cancel()
}

func (d *dispatcher) transfer(target *dispatcher) {
	if target == nil {
		return
	}
	for {
		select {
		case message, ok := <-d.buffer:
			if !ok {
				return
			}
			target.buffer <- message
		}
	}
}
