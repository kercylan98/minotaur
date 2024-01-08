package buffer

import (
	"sync"
)

// NewRingUnbounded 创建一个并发安全的基于环形缓冲区实现的无界缓冲区
func NewRingUnbounded[T any](bufferSize int) *RingUnbounded[T] {
	ru := &RingUnbounded[T]{
		ring:         NewRing[T](1024),
		rc:           make(chan T, bufferSize),
		closedSignal: make(chan struct{}),
	}
	ru.cond = sync.NewCond(&ru.rrm)

	ru.process()
	return ru
}

// RingUnbounded 基于环形缓冲区实现的无界缓冲区
type RingUnbounded[T any] struct {
	ring         *Ring[T]
	rrm          sync.Mutex
	cond         *sync.Cond
	rc           chan T
	closed       bool
	closedMutex  sync.RWMutex
	closedSignal chan struct{}
}

// Write 写入数据
func (b *RingUnbounded[T]) Write(v T) {
	b.closedMutex.RLock()
	defer b.closedMutex.RUnlock()
	if b.closed {
		return
	}

	b.rrm.Lock()
	b.ring.Write(v)
	b.cond.Signal()
	b.rrm.Unlock()
}

// Read 读取数据
func (b *RingUnbounded[T]) Read() <-chan T {
	return b.rc
}

// Closed 判断缓冲区是否已关闭
func (b *RingUnbounded[T]) Closed() bool {
	b.closedMutex.RLock()
	defer b.closedMutex.RUnlock()
	return b.closed
}

// Close 关闭缓冲区，关闭后将不再接收新数据，但是已有数据仍然可以读取
func (b *RingUnbounded[T]) Close() <-chan struct{} {
	b.closedMutex.Lock()
	defer b.closedMutex.Unlock()
	if b.closed {
		return b.closedSignal
	}
	b.closed = true

	b.rrm.Lock()
	b.cond.Signal()
	b.rrm.Unlock()
	return b.closedSignal
}

func (b *RingUnbounded[T]) process() {
	go func(b *RingUnbounded[T]) {
		for {
			b.closedMutex.RLock()
			b.rrm.Lock()
			vs := b.ring.ReadAll()
			if len(vs) == 0 && !b.closed {
				b.closedMutex.RUnlock()
				b.cond.Wait()
			} else {
				b.closedMutex.RUnlock()
			}
			b.rrm.Unlock()

			b.closedMutex.RLock()
			if b.closed && len(vs) == 0 {
				close(b.rc)
				close(b.closedSignal)
				b.closedMutex.RUnlock()
				break
			}

			for _, v := range vs {
				b.rc <- v
			}
			b.closedMutex.RUnlock()
		}
	}(b)
}
