package dispatcher

import (
	"context"
	"errors"
	"github.com/kercylan98/minotaur/utils/buffer"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/super"
	"sync"
	"sync/atomic"
)

// NewDispatcher 创建一个消息调度器
func NewDispatcher[M any](bufferSize int, handler func(M)) *Dispatcher[M] {
	d := &Dispatcher[M]{
		buf:     buffer.NewRing[M](bufferSize),
		bufCond: sync.NewCond(&sync.Mutex{}),
		ctx:     context.Background(),
	}

	d.BindErrorHandler(func(err error) {
		log.Error("dispatcher", log.Err(err))
	})

	d.handler = func(m M) {
		defer func() {
			if err := super.RecoverTransform(recover()); err != nil {
				d.errorHandler.Load().(func(error))(err)
			}
		}()
		handler(m)
	}

	return d
}

// Dispatcher 并发安全且不阻塞的消息调度器
type Dispatcher[M any] struct {
	buf          *buffer.Ring[M]
	bufCond      *sync.Cond
	handler      func(M)
	closed       bool
	ctx          context.Context
	errorHandler atomic.Value
}

// BindErrorHandler 绑定一个错误处理器到调度器中
func (d *Dispatcher[M]) BindErrorHandler(handler func(error)) {
	d.errorHandler.Store(handler)
}

// BindContext 绑定一个上下文到调度器中
func (d *Dispatcher[M]) BindContext(ctx context.Context) {
	d.bufCond.L.Lock()
	d.ctx = ctx
	if _, canceled := d.ctx.Deadline(); canceled {
		d.closed = true
		d.bufCond.Signal()
		d.bufCond.L.Unlock()
		return
	}
	d.bufCond.L.Unlock()
}

// Send 发送消息到调度器中等待处理
func (d *Dispatcher[M]) Send(m M) error {
	d.bufCond.L.Lock()
	if d.closed {
		d.bufCond.L.Unlock()
		return errors.New("dispatcher closed")
	}
	d.buf.Write(m)
	d.bufCond.Signal()
	d.bufCond.L.Unlock()
	return nil
}

// Start 阻塞式启动调度器，调用后将开始处理消息
func (d *Dispatcher[M]) Start() {
	for {
		select {
		case <-d.ctx.Done():
			d.Stop()
		default:
			d.bufCond.L.Lock()
			if d.buf.Len() == 0 {
				if d.closed {
					d.bufCond.L.Unlock()
					return
				}
				d.bufCond.Wait()
			}
			messages := d.buf.ReadAll()
			d.bufCond.L.Unlock()
			for _, msg := range messages {
				d.handler(msg)
			}
		}
	}
}

// Stop 停止调度器，调用后将不再接受新消息，但会处理完已有消息
func (d *Dispatcher[M]) Stop() {
	d.bufCond.L.Lock()
	d.closed = true
	d.bufCond.Signal()
	d.bufCond.L.Unlock()
}
