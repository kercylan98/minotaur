package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"sync/atomic"
	"unsafe"
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

type defaultMailbox struct {
	queue       *lfQueue
	systemQueue *lfQueue
	processor   core.MessageProcessor
	dispatcher  Dispatcher
	status      uint32
	num         int32
	suspended   uint32
}

func newDefaultMailbox() *defaultMailbox {
	return &defaultMailbox{
		queue:       newLfQueue(),
		systemQueue: newLfQueue(),
	}
}

func (m *defaultMailbox) OnInit(processor core.MessageProcessor, dispatcher Dispatcher) {
	m.processor = processor
	m.dispatcher = dispatcher
}

func (m *defaultMailbox) DeliveryUserMessage(message Message) {
	m.queue.Enqueue(unsafe.Pointer(&message))
	atomic.AddInt32(&m.num, 1)
	m.dispatch()
}

func (m *defaultMailbox) DeliverySystemMessage(message Message) {
	m.systemQueue.Enqueue(unsafe.Pointer(&message))
	atomic.AddInt32(&m.num, 1)
	m.dispatch()
}

func (m *defaultMailbox) dispatch() {
	if atomic.CompareAndSwapUint32(&m.status, defaultMailboxStatusIdle, defaultMailboxStatusRunning) {
		m.dispatcher.Dispatch(m.process)
	}
}

func (m *defaultMailbox) process() {
	m.processHandle()
	for atomic.LoadInt32(&m.num) > 0 && (atomic.LoadUint32(&m.suspended) == 1) {
		m.processHandle()
	}

	atomic.StoreUint32(&m.status, defaultMailboxStatusIdle)
}

func (m *defaultMailbox) processHandle() {
	defer func() {
		if reason := recover(); reason != nil {
			m.processor.ProcessRecover(reason)
		}
	}()

	var msg Message
	for {
		if ptr := m.systemQueue.Dequeue(); ptr != nil {
			msg = *(*Message)(ptr)
			atomic.AddInt32(&m.num, -1)
			switch msg.(type) {
			case _OnSuspendMailbox:
				atomic.StoreUint32(&m.suspended, 1)
			case _OnResumeMailbox:
				atomic.StoreUint32(&m.suspended, 0)
			default:
				m.processor.ProcessSystemMessage(msg)
			}
			continue
		}

		if atomic.LoadUint32(&m.suspended) == 1 {
			return
		}

		if ptr := m.queue.Dequeue(); ptr != nil {
			msg = *(*Message)(ptr)
			atomic.AddInt32(&m.num, -1)
			m.processor.ProcessUserMessage(msg)
			continue
		}
		break
	}
}
