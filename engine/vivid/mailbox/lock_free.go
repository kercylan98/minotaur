package mailbox

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid/dispatcher"
	"github.com/kercylan98/minotaur/toolkit/queues"
	"sync/atomic"
	"unsafe"
)

const (
	lockFreeStatusIdle uint32 = iota
	lockFreeStatusRunning
)

var _ Mailbox = &LockFree{}

// NewLockFree 创建一个基于无锁队列实现的邮箱，该邮箱基于 queues.LFQueue 实现
//   - 默认邮箱在 userMessageBatchLimit 大于 1 时需要注意，一批用户消息将会被处理，而不会被系统消息抢先执行
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
	atomic.StoreUint32(&m.suspended, 1)
}

func (m *LockFree) Resume() {
	atomic.StoreUint32(&m.suspended, 0)
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
	if atomic.CompareAndSwapUint32(&m.status, lockFreeStatusIdle, lockFreeStatusRunning) {
		m.dispatcher.Dispatch(m.process)
	}
}

func (m *LockFree) process() {
	for {
		m.processHandle()
		atomic.StoreUint32(&m.status, lockFreeStatusIdle)
		notEmpty := atomic.LoadInt32(&m.sysNum) > 0 || (atomic.LoadUint32(&m.suspended) == 0 && atomic.LoadInt32(&m.userNum) > 0)
		if !notEmpty {
			break
		} else if !atomic.CompareAndSwapUint32(&m.status, lockFreeStatusIdle, lockFreeStatusRunning) {
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
