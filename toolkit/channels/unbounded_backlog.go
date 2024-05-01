package channels

import (
	"sync"
)

// NewUnboundedBacklog 创建一个并发安全的，基于 channel 和缓冲队列实现的无界缓冲区
//
// 该缓冲区来源于 gRPC 的实现，用于在不使用额外 goroutine 的情况下实现无界缓冲区
//   - 该缓冲区的所有方法都是线程安全的，除了用于同步的互斥锁外，不会阻塞任何东西
func NewUnboundedBacklog[V any]() *UnboundedBacklog[V] {
	return &UnboundedBacklog[V]{c: make(chan V, 1)}
}

// UnboundedBacklog 是并发安全的无界缓冲区的实现
type UnboundedBacklog[V any] struct {
	c       chan V
	closed  bool
	mu      sync.Mutex
	backlog []V
}

// Put 将数据放入缓冲区
func (ub *UnboundedBacklog[V]) Put(t V) {
	ub.mu.Lock()
	defer ub.mu.Unlock()
	if ub.closed {
		return
	}
	if len(ub.backlog) == 0 {
		select {
		case ub.c <- t:
			return
		default:
		}
	}
	ub.backlog = append(ub.backlog, t)
}

// Load 将缓冲区中的数据发送到读取通道中，如果缓冲区中没有数据，则不会发送
//   - 在每次 Get 后都应该执行该函数
func (ub *UnboundedBacklog[V]) Load() {
	ub.mu.Lock()
	defer ub.mu.Unlock()
	if ub.closed {
		return
	}
	if len(ub.backlog) > 0 {
		select {
		case ub.c <- ub.backlog[0]:
			ub.backlog = ub.backlog[1:]
		default:
		}
	}
}

// Get 获取可接收消息的读取通道，需要注意的是，每次读取成功都应该通过 Load 函数将缓冲区中的数据加载到读取通道中
func (ub *UnboundedBacklog[V]) Get() <-chan V {
	return ub.c
}

// Close 关闭
func (ub *UnboundedBacklog[V]) Close() {
	ub.mu.Lock()
	defer ub.mu.Unlock()
	if ub.closed {
		return
	}
	ub.closed = true
	close(ub.c)
}

// IsClosed 是否已关闭
func (ub *UnboundedBacklog[V]) IsClosed() bool {
	ub.mu.Lock()
	defer ub.mu.Unlock()
	return ub.closed
}
