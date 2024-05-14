package queues

import (
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"sync"
	"sync/atomic"
)

const (
	fifoStateNone     = fifoState(iota) // 未启动状态
	fifoStateRunning                    // 运行中状态
	fifoStateStopping                   // 停止中状态
	fifoStateStopped                    // 已停止状态
)

const (
	FIFOStopModeInstantly = FifoStopMode(iota) // 立刻停止消息队列，新消息将不再接收，缓冲区内未处理的消息将被丢弃
	FIFOStopModeGraceful                       // 优雅停止消息队列，新消息将不再接收，等待未处理的消息处理完毕后再停止
	FIFOStopModeDrain                          // 新消息将继续被接收，等待消息队列处理完毕且没有新消息后再停止
)

type (
	fifoState    = int32 // 状态
	FifoStopMode = int8  // FIFO 消息队列的停止模式，目前支持 FIFOStopModeInstantly、 FIFOStopModeGraceful、 FIFOStopModeDrain
)

func NewFIFO[T any](opts ...*FIFOOptions) *FIFO[T] {
	f := &FIFO[T]{
		opts:   NewFIFOOptions().Apply(opts...),
		status: fifoStateNone,
		cond:   sync.NewCond(&sync.Mutex{}),
		closed: make(chan struct{}),
	}
	f.buffer = buffer.NewRing[T](int(f.opts.BufferSize))
	f.dequeue = make(chan T, f.opts.DequeueBufferSize)
	return f
}

// FIFO 是一个先进先出的消息队列
type FIFO[T any] struct {
	opts    *FIFOOptions    // 配置
	status  fifoState       // 队列状态
	cond    *sync.Cond      // 消息队列条件变量
	buffer  *buffer.Ring[T] // 消息缓t冲区
	closed  chan struct{}   // 关闭信号
	dequeue chan T          // 出列通道
}

func (f *FIFO[T]) Start() {
	if !atomic.CompareAndSwapInt32(&f.status, fifoStateNone, fifoStateRunning) {
		return
	}
	defer func(f *FIFO[T]) {
		close(f.closed)
		close(f.dequeue)
		f.buffer.Reset()
	}(f)

	for {
		f.cond.L.Lock()
		elements := f.buffer.ReadAll()
		if len(elements) == 0 || (f.opts.StopMode != FIFOStopModeDrain && atomic.LoadInt32(&f.status) == fifoStateStopping) {
			// 此刻队列没有消息且处于停止中状态，将会关闭 Actor
			if atomic.CompareAndSwapInt32(&f.status, fifoStateStopping, fifoStateStopped) {
				f.cond.L.Unlock()
				return
			}
			f.cond.Wait()

			// 重新读取消息
			elements = f.buffer.ReadAll()
		}
		f.cond.L.Unlock()

		for i := 0; i < len(elements); i++ {
			elem := elements[i]
			if f.opts.DequeueFullBlock {
				f.dequeue <- elem
			} else {
				select {
				case f.dequeue <- elem:
				default:
					if f.opts.PickUpDiscarded {
						f.Enqueue(elem)
					}
				}
			}
		}
	}
}

func (f *FIFO[T]) Stop() {
	f.cond.L.Lock()
	if !atomic.CompareAndSwapInt32(&f.status, fifoStateRunning, fifoStateStopping) {
		f.cond.L.Unlock()
		return
	}
	// 避免循环位于等待状态且一直没有新消息进入无法退出循环
	f.cond.Signal()
	f.cond.L.Unlock()

	<-f.closed
}

func (f *FIFO[T]) Enqueue(elem T) {
	if atomic.LoadInt32(&f.status) == fifoStateStopped && f.opts.StopMode != FIFOStopModeDrain {
		return
	}

	f.cond.L.Lock()
	f.buffer.Write(elem)
	f.cond.Signal()
	f.cond.L.Unlock()
}

func (f *FIFO[T]) Dequeue() <-chan T {
	return f.dequeue
}
