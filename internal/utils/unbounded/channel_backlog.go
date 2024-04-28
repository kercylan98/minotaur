package unbounded

import (
	"sync"
)

// NewChannelBacklog 创建一个并发安全的，基于 channel 和缓冲队列实现的无界缓冲区
//
// 该缓冲区来源于 gRPC 的实现，用于在不使用额外 goroutine 的情况下实现无界缓冲区
//   - 该缓冲区的所有方法都是线程安全的，除了用于同步的互斥锁外，不会阻塞任何东西
func NewChannelBacklog[V any]() *ChannelBacklog[V] {
	return &ChannelBacklog[V]{c: make(chan V, 1)}
}

// ChannelBacklog 是并发安全的无界缓冲区的实现
type ChannelBacklog[V any] struct {
	c       chan V
	closed  bool
	mu      sync.Mutex
	backlog []V
}

// Put 将数据放入缓冲区
func (cb *ChannelBacklog[V]) Put(t V) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	if cb.closed {
		return
	}
	if len(cb.backlog) == 0 {
		select {
		case cb.c <- t:
			return
		default:
		}
	}
	cb.backlog = append(cb.backlog, t)
}

// Load 将缓冲区中的数据发送到读取通道中，如果缓冲区中没有数据，则不会发送
//   - 在每次 Get 后都应该执行该函数
func (cb *ChannelBacklog[V]) Load() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	if cb.closed {
		return
	}
	if len(cb.backlog) > 0 {
		select {
		case cb.c <- cb.backlog[0]:
			cb.backlog = cb.backlog[1:]
		default:
		}
	}
}

// Get 获取可接收消息的读取通道，需要注意的是，每次读取成功都应该通过 Load 函数将缓冲区中的数据加载到读取通道中
func (cb *ChannelBacklog[V]) Get() <-chan V {
	return cb.c
}

// Close 关闭
func (cb *ChannelBacklog[V]) Close() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	if cb.closed {
		return
	}
	cb.closed = true
	close(cb.c)
}

// IsClosed 是否已关闭
func (cb *ChannelBacklog[V]) IsClosed() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.closed
}
