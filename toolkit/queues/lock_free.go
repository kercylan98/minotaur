package queues

import (
	"sync/atomic"
	"unsafe"
)

// LFQueue 无锁（Lock-free）队列实现
type LFQueue struct {
	head, tail unsafe.Pointer
}

type lfNode struct {
	value unsafe.Pointer
	next  unsafe.Pointer
}

func NewLFQueue() *LFQueue {
	node := unsafe.Pointer(&lfNode{})
	return &LFQueue{head: node, tail: node}
}

func (q *LFQueue) Push(value unsafe.Pointer) {
	node := unsafe.Pointer(&lfNode{value: value})
	for {
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*lfNode)(tail).next)
		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				if atomic.CompareAndSwapPointer(&(*lfNode)(tail).next, next, node) {
					atomic.CompareAndSwapPointer(&q.tail, tail, node)
					return
				}
			} else {
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			}
		}
	}
}

func (q *LFQueue) Pop() unsafe.Pointer {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*lfNode)(head).next)
		if head == atomic.LoadPointer(&q.head) {
			if head == tail {
				if next == nil {
					return nil
				}
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			} else {
				value := (*lfNode)(next).value
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					return value
				}
			}
		}
	}
}
