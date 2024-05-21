package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"sync"
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

func NewFIFO(handler func(message MessageContext), opts ...*FIFOOptions) *FIFO {
	f := &FIFO{
		opts:    NewFIFOOptions().Apply(opts...),
		status:  fifoStateNone,
		cond:    sync.NewCond(&sync.Mutex{}),
		closed:  make(chan struct{}),
		handler: handler,
	}
	f.buffer = buffer.NewRing[MessageContext](int(f.opts.BufferSize))
	return f
}

// FIFO 是一个先进先出的消息队列
type FIFO struct {
	opts    *FIFOOptions                 // 配置
	status  fifoState                    // 队列状态
	cond    *sync.Cond                   // 消息队列条件变量
	buffer  *buffer.Ring[MessageContext] // 消息缓冲区
	closed  chan struct{}                // 关闭信号
	handler func(message MessageContext) // 消息处理函数
}

func (f *FIFO) Start() {
	f.cond.L.Lock()
	if f.status != fifoStateNone {
		f.cond.L.Unlock()
		return
	}
	f.status = fifoStateRunning
	f.cond.L.Unlock()

	f.closed = make(chan struct{})
	defer func(f *FIFO) {
		close(f.closed)
		f.buffer.Reset()
	}(f)

	for {
		f.cond.L.Lock()
		elements := f.buffer.ReadAll()
		if len(elements) == 0 || (f.opts.StopMode == FIFOStopModeInstantly && f.status == fifoStateStopping) {
			// 此刻队列没有消息且处于停止中状态，将会关闭 Actor
			if f.status == fifoStateStopping {
				f.status = fifoStateStopped
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
			f.handler(elem)
		}
	}
}

func (f *FIFO) Stop() {
	f.cond.L.Lock()
	if f.status != fifoStateRunning {
		f.cond.L.Unlock()
		return
	}
	f.status = fifoStateStopping
	f.cond.L.Unlock()

	f.cond.Signal()

	<-f.closed
}

func (f *FIFO) Enqueue(message MessageContext) {
	f.cond.L.Lock()
	if f.status != fifoStateNone {
		if f.status != fifoStateRunning {
			if f.opts.StopMode != FIFOStopModeDrain {
				f.cond.L.Unlock()
				return
			}
		}
	}

	f.buffer.Write(message)
	f.cond.L.Unlock()
	f.cond.Broadcast()
}

func (f *FIFO) reset() {
	f.cond.L.Lock()
	if f.status < fifoStateStopping {
		f.cond.L.Unlock()
		f.Stop()
	} else {
		f.cond.L.Unlock()
		return
	}

	f.cond.L.Lock()
	defer f.cond.L.Unlock()

	f.buffer.Reset()
	f.status = fifoStateNone
}
