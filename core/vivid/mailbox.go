package vivid

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/queues"
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

var _ Mailbox = &LockFreeMailbox{}

type LockFreeMailbox struct {
	queue                 *queues.LFQueue
	systemQueue           *queues.LFQueue
	processor             core.MessageProcessor
	dispatcher            Dispatcher
	status                uint32
	sysNum                int32
	userNum               int32
	suspended             uint32
	userMessageBatchLimit int
}

// NewDefaultMailbox 创建一个默认的邮箱，该邮箱基于 queues.LFQueue 实现
//   - 默认邮箱在 userMessageBatchLimit 大于 1 时需要注意，一批用户消息将会被处理，而不会被系统消息抢先执行
func NewDefaultMailbox(userMessageBatchLimit int) *LockFreeMailbox {
	if userMessageBatchLimit <= 1 {
		userMessageBatchLimit = 1
	}
	return &LockFreeMailbox{
		queue:                 queues.NewLFQueue(),
		systemQueue:           queues.NewLFQueue(),
		userMessageBatchLimit: userMessageBatchLimit,
	}
}

func (m *LockFreeMailbox) OnInit(processor core.MessageProcessor, dispatcher Dispatcher) {
	m.processor = processor
	m.dispatcher = dispatcher
}

func (m *LockFreeMailbox) DeliveryUserMessage(message Message) {
	m.queue.Push(unsafe.Pointer(&message))
	atomic.AddInt32(&m.userNum, 1)
	m.dispatch()
}

func (m *LockFreeMailbox) DeliverySystemMessage(message Message) {
	m.systemQueue.Push(unsafe.Pointer(&message))
	atomic.AddInt32(&m.sysNum, 1)
	m.dispatch()
}

func (m *LockFreeMailbox) dispatch() {
	if atomic.CompareAndSwapUint32(&m.status, defaultMailboxStatusIdle, defaultMailboxStatusRunning) {
		m.dispatcher.Dispatch(m.process)
	}
}

func (m *LockFreeMailbox) process() {
	for atomic.LoadInt32(&m.sysNum) > 0 || (atomic.LoadUint32(&m.suspended) == 0 && atomic.LoadInt32(&m.userNum) > 0) {
		m.processHandle()

		// 处理完成后可能会有新的消息到达，所以需要再次尝试
		for atomic.LoadInt32(&m.sysNum) > 0 || (atomic.LoadUint32(&m.suspended) == 0 && atomic.LoadInt32(&m.userNum) > 0) {
			m.processHandle()
		}

		atomic.CompareAndSwapUint32(&m.status, defaultMailboxStatusRunning, defaultMailboxStatusIdle)
	}
}

func (m *LockFreeMailbox) processHandle() {
	defer func() {
		if reason := recover(); reason != nil {
			m.processor.ProcessRecover(reason)
		}
	}()

	var msg Message
	var messages []Message
	for {
		if ptr := m.systemQueue.Pop(); ptr != nil {
			msg = *(*Message)(ptr)
			atomic.AddInt32(&m.sysNum, -1)
			switch msg.(type) {
			case OnSuspendMailbox:
				atomic.StoreUint32(&m.suspended, 1)
			case OnResumeMailbox:
				atomic.StoreUint32(&m.suspended, 0)
			default:
				m.processor.ProcessSystemMessage(msg)
			}
			continue
		}

		if atomic.LoadUint32(&m.suspended) == 1 {
			return
		}

		if m.userMessageBatchLimit > 1 {
			if ptrList := m.queue.BatchPop(m.userMessageBatchLimit); len(ptrList) != 0 {
				if len(messages) < len(ptrList) {
					messages = make([]Message, len(ptrList))
				}
				for i := 0; i < len(ptrList); i++ {
					messages[i] = *(*Message)(ptrList[i])
				}
				atomic.AddInt32(&m.userNum, int32(-len(messages)))
				m.processor.ProcessUserMessage(messages)
				messages = messages[:0]
				continue
			}
			break
		}

		if ptr := m.queue.Pop(); ptr != nil {
			msg = *(*Message)(ptr)
			atomic.AddInt32(&m.userNum, -1)
			m.processor.ProcessUserMessage(msg)
			continue
		}
		break
	}
}
