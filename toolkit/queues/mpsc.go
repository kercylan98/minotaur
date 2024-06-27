package queues

import (
	"sync/atomic"
	"unsafe"
)

type mpscNode struct {
	next *mpscNode
	val  interface{}
}

type MPSC struct {
	head, tail *mpscNode
}

func NewMPSC() *MPSC {
	q := &MPSC{}
	stub := &mpscNode{}
	q.head = stub
	q.tail = stub
	return q
}

// Push 将 m 添加到队列末尾。
//   - 可以从多个 goroutine 安全地调用 Push
func (q *MPSC) Push(m any) {
	n := new(mpscNode)
	n.val = m
	prev := (*mpscNode)(atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), unsafe.Pointer(n)))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&prev.next)), unsafe.Pointer(n))
}

// Pop 从队列前面删除项目，如果队列为空则为 nil
//   - 必须从单个消费者 goroutine 调用
func (q *MPSC) Pop() any {
	tail := q.tail
	next := (*mpscNode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next))))
	if next != nil {
		q.tail = next
		v := next.val
		next.val = nil
		return v
	}
	return nil
}

// Empty 如果队列为空将返回 true
//   - 必须从单个消费者 goroutine 调用
func (q *MPSC) Empty() bool {
	tail := q.tail
	next := (*mpscNode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next))))
	return next == nil
}
