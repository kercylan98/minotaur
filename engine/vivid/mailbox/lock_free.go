package mailbox

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid/dispatcher"
	"github.com/kercylan98/minotaur/toolkit/queues"
	"github.com/puzpuzpuz/xsync/v3"
	"sync/atomic"
	"unsafe"
)

var _ Mailbox = &LockFree{}

// NewLockFree 创建一个基于无锁队列实现的邮箱，该邮箱基于 queues.LFQueue 实现
func NewLockFree(dispatcher dispatcher.Dispatcher, recipient Recipient) *LockFree {
	return &LockFree{
		queue:       queues.NewLFQueue(),
		systemQueue: queues.NewLFQueue(),
		dispatcher:  dispatcher,
		recipient:   recipient,
	}
}

type LockFree struct {
	queue       *queues.LFQueue
	systemQueue *queues.LFQueue
	dispatcher  dispatcher.Dispatcher
	recipient   Recipient
	status      uint32
	sysNum      int32
	userNum     int32
	suspended   uint32
}

func (m *LockFree) Suspend() {
	xsync.NewCounter()
	atomic.StoreUint32(&m.suspended, 1)
}

func (m *LockFree) Resume() {
	atomic.StoreUint32(&m.suspended, 0)
	m.dispatch()
}

func (m *LockFree) DeliveryUserMessage(message prc.Message) {
	m.queue.Push(unsafe.Pointer(&message))
	atomic.AddInt32(&m.userNum, 1)
	m.dispatch()
}

func (m *LockFree) DeliverySystemMessage(message prc.Message) {
	m.systemQueue.Push(unsafe.Pointer(&message))
	atomic.AddInt32(&m.sysNum, 1)
	m.dispatch()
}

func (m *LockFree) dispatch() {
	if atomic.CompareAndSwapUint32(&m.status, mailboxStatusIdle, mailboxStatusRunning) {
		m.dispatcher.Dispatch(m.process)
	}
}

func (m *LockFree) process() {
	for {
		m.processHandle()
		atomic.StoreUint32(&m.status, mailboxStatusIdle)
		notEmpty := atomic.LoadInt32(&m.sysNum) > 0 || (atomic.LoadUint32(&m.suspended) == 0 && atomic.LoadInt32(&m.userNum) > 0)
		if !notEmpty {
			break
		} else if !atomic.CompareAndSwapUint32(&m.status, mailboxStatusIdle, mailboxStatusRunning) {
			break
		}
	}
}

func (m *LockFree) processHandle() {
	defer func() {
		if reason := recover(); reason != nil {
			m.recipient.ProcessAccident(reason)
		}
	}()

	var msg prc.Message
	for {
		if ptr := m.systemQueue.Pop(); ptr != nil {
			msg = *(*prc.Message)(ptr)
			atomic.AddInt32(&m.sysNum, -1)
			m.recipient.ProcessSystemMessage(msg)
			continue
		}

		if atomic.LoadUint32(&m.suspended) == 1 {
			return
		}

		if ptr := m.queue.Pop(); ptr != nil {
			msg = *(*prc.Message)(ptr)
			atomic.AddInt32(&m.userNum, -1)
			m.recipient.ProcessUserMessage(msg)
			continue
		}
		break
	}
}
