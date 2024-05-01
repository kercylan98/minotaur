package channels

import (
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"sync"
)

// NewUnboundedRing 创建一个并发安全的基于环形缓冲区实现的无界缓冲区
func NewUnboundedRing[T any](ctx context.Context) *UnboundedRing[T] {
	r := &UnboundedRing[T]{
		ctx:  ctx,
		ring: buffer.NewRing[T](1024),
		ch:   make(chan T, 1024),
	}
	r.cond = sync.NewCond(&r.rw)

	go r.run()
	return r
}

// UnboundedRing 是并发安全的，基于环形缓冲区实现的无界缓冲区
type UnboundedRing[T any] struct {
	ctx    context.Context
	ring   *buffer.Ring[T]
	rw     sync.RWMutex
	cond   *sync.Cond
	ch     chan T
	closed bool
}

// Put 将数据放入缓冲区
func (r *UnboundedRing[T]) Put(v ...T) error {
	if len(v) == 0 {
		return nil
	}
	r.rw.Lock()
	if r.closed {
		r.rw.Unlock()
		return errors.New("unbounded ring is closed")
	}
	for _, t := range v {
		r.ring.Write(t)
	}
	r.rw.Unlock()
	r.cond.Signal()
	return nil
}

// Get 获取可接收消息的读取通道
func (r *UnboundedRing[T]) Get() <-chan T {
	return r.ch
}

// Close 关闭缓冲区
func (r *UnboundedRing[T]) Close() {
	r.rw.RLock()
	r.closed = true
	r.rw.RUnlock()
	r.cond.Signal()
}

// IsClosed 是否已关闭
func (r *UnboundedRing[T]) IsClosed() bool {
	r.rw.RLock()
	defer r.rw.RUnlock()
	return r.closed
}

func (r *UnboundedRing[T]) run() {
	for {
		select {
		case <-r.ctx.Done():
			r.Close()
		default:
			r.rw.Lock()
			if r.ring.IsEmpty() {
				if r.closed { // 如果已关闭并且没有数据，则关闭通道
					close(r.ch)
					r.rw.Unlock()
					return
				}
				// 等待数据
				r.cond.Wait()
			}
			vs := r.ring.ReadAll()
			r.rw.Unlock()
			for _, v := range vs {
				r.ch <- v
			}
		}
	}
}
