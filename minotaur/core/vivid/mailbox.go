package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"sync/atomic"
)

type Mailbox interface {
	// OnInit 在邮箱被初始化时将会被调用，processor 是邮箱的消息处理器，dispatcher 是邮箱的消息分发器
	OnInit(processor core.MessageProcessor, dispatcher Dispatcher)

	// DeliveryUserMessage 投递用户消息到邮箱
	DeliveryUserMessage(message Message)

	// DeliverySystemMessage 投递系统消息到邮箱
	DeliverySystemMessage(message Message)
}

const (
	defaultMailboxStatusIdle uint32 = iota
	defaultMailboxStatusRunning
)

var _ Mailbox = &defaultMailbox{}

func newDefaultMailbox() *defaultMailbox {
	return &defaultMailbox{
		queue:       unboundedQueuePool.Get(),
		systemQueue: unboundedQueuePool.Get(),
	}
}

type defaultMailbox struct {
	queue       core.Queue
	systemQueue core.Queue
	processor   core.MessageProcessor
	dispatcher  Dispatcher
	status      uint32
	num         int32
}

func (m *defaultMailbox) OnInit(processor core.MessageProcessor, dispatcher Dispatcher) {
	m.processor = processor
	m.dispatcher = dispatcher
}

func (m *defaultMailbox) DeliveryUserMessage(message Message) {
	m.queue.Enqueue(message)
	atomic.AddInt32(&m.num, 1)
	m.dispatch()
}

func (m *defaultMailbox) DeliverySystemMessage(message Message) {
	m.systemQueue.Enqueue(message)
	atomic.AddInt32(&m.num, 1)
	m.dispatch()
}

func (m *defaultMailbox) dispatch() {
	if atomic.CompareAndSwapUint32(&m.status, defaultMailboxStatusIdle, defaultMailboxStatusRunning) {
		m.dispatcher.Dispatch(m.process)
	}
}

func (m *defaultMailbox) process() {
	var msg Message
	for {
		if msg = m.systemQueue.Dequeue(); msg != nil {
			atomic.AddInt32(&m.num, -1)
			m.processor.ProcessSystemMessage(msg)
			continue
		}

		if msg = m.queue.Dequeue(); msg == nil {
			break
		}

		atomic.AddInt32(&m.num, -1)
		m.processor.ProcessUserMessage(msg)
	}

	atomic.StoreUint32(&m.status, defaultMailboxStatusIdle)

	if atomic.LoadInt32(&m.num) > 0 {
		m.dispatch()
	}
}
