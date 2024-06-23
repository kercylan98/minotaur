package vivid

import (
	"sync/atomic"
	"unsafe"
)

// Lock-free queue implementation
type lfQueue struct {
	head, tail unsafe.Pointer
}

type lfNode struct {
	value unsafe.Pointer
	next  unsafe.Pointer
}

func newLfQueue() *lfQueue {
	node := unsafe.Pointer(&lfNode{})
	return &lfQueue{head: node, tail: node}
}

func (q *lfQueue) Enqueue(value unsafe.Pointer) {
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

func (q *lfQueue) Dequeue() unsafe.Pointer {
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
