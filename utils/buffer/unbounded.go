package buffer

import (
	"sync"
)

// NewUnbounded 创建一个无界缓冲区
//   - generateNil: 生成空值的函数，该函数仅需始终返回 nil 即可
//
// 该缓冲区来源于 gRPC 的实现，用于在不使用额外 goroutine 的情况下实现无界缓冲区
//   - 该缓冲区的所有方法都是线程安全的，除了用于同步的互斥锁外，不会阻塞任何东西
func NewUnbounded[V any](generateNil func() V) *Unbounded[V] {
	return &Unbounded[V]{c: make(chan V, 1), nil: generateNil()}
}

// NewUnboundedN 与 NewUnbounded 相同，只是省略了 generateNil 参数
func NewUnboundedN[V any]() *Unbounded[V] {
	return NewUnbounded[V](func() (v V) {
		return v
	})
}

// Unbounded 是无界缓冲区的实现
type Unbounded[V any] struct {
	c       chan V
	closed  bool
	mu      sync.Mutex
	backlog []V
	nil     V
}

// Put 将数据放入缓冲区
func (slf *Unbounded[V]) Put(t V) {
	slf.mu.Lock()
	defer slf.mu.Unlock()
	if slf.closed {
		return
	}
	if len(slf.backlog) == 0 {
		select {
		case slf.c <- t:
			return
		default:
		}
	}
	slf.backlog = append(slf.backlog, t)
}

// Load 将缓冲区中的数据发送到读取通道中，如果缓冲区中没有数据，则不会发送
//   - 在每次 Get 后都应该执行该函数
func (slf *Unbounded[V]) Load() {
	slf.mu.Lock()
	defer slf.mu.Unlock()
	if slf.closed {
		return
	}
	if len(slf.backlog) > 0 {
		select {
		case slf.c <- slf.backlog[0]:
			slf.backlog[0] = slf.nil
			slf.backlog = slf.backlog[1:]
		default:
		}
	}
}

// Get 获取读取通道
func (slf *Unbounded[V]) Get() <-chan V {
	return slf.c
}

// Close 关闭
func (slf *Unbounded[V]) Close() {
	slf.mu.Lock()
	defer slf.mu.Unlock()
	if slf.closed {
		return
	}
	slf.closed = true
	close(slf.c)
}

// IsClosed 是否已关闭
func (slf *Unbounded[V]) IsClosed() bool {
	slf.mu.Lock()
	defer slf.mu.Unlock()
	return slf.closed
}
