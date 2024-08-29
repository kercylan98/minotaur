package mailbox

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/toolkit/queues"
	"sync/atomic"
	"unsafe"
)

var _ Mailbox = &Sync{}

func NewSync(recipient Recipient) *Sync {
	return &Sync{
		queue:       queues.NewLFQueue(),
		systemQueue: queues.NewLFQueue(),
		recipient:   recipient,
	}
}

type Sync struct {
	queue       *queues.LFQueue
	systemQueue *queues.LFQueue
	recipient   Recipient
	status      uint32
	sysNum      int32
	userNum     int32
	suspended   uint32
}

func (m *Sync) Suspend() {
	atomic.StoreUint32(&m.suspended, 1)
}

func (m *Sync) Resume() {
	atomic.StoreUint32(&m.suspended, 0)
	m.dispatch()
}

func (m *Sync) DeliveryUserMessage(message prc.Message) {
	m.queue.Push(unsafe.Pointer(&message))
	atomic.AddInt32(&m.userNum, 1)
	m.dispatch()
}

func (m *Sync) DeliverySystemMessage(message prc.Message) {
	m.systemQueue.Push(unsafe.Pointer(&message))
	atomic.AddInt32(&m.sysNum, 1)
	m.dispatch()
}

func (m *Sync) dispatch() {
	if atomic.CompareAndSwapUint32(&m.status, mailboxStatusIdle, mailboxStatusRunning) {
		m.process()
	}
}

func (m *Sync) process() {
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

func (m *Sync) processHandle() {
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
