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
		opts:      NewFIFOOptions().Apply(opts...),
		status:    fifoStateNone,
		cond:      sync.NewCond(&sync.Mutex{}),
		closed:    make(chan struct{}),
		handler:   handler,
		instantly: make(chan *instantlyMessage, 1),
	}
	f.buffer = buffer.NewRing[MessageContext](int(f.opts.BufferSize))
	return f
}

// FIFO 是一个先进先出的消息队列
type FIFO struct {
	opts      *FIFOOptions                 // 配置
	status    fifoState                    // 队列状态
	cond      *sync.Cond                   // 消息队列条件变量
	buffer    *buffer.Ring[MessageContext] // 消息缓冲区
	closed    chan struct{}                // 关闭信号
	handler   func(message MessageContext) // 消息处理函数
	instantly chan *instantlyMessage       // 立即处理的消息
}

func (m *FIFO) Start() {
	m.cond.L.Lock()
	if m.status != fifoStateNone {
		m.cond.L.Unlock()
		return
	}
	m.status = fifoStateRunning
	m.cond.L.Unlock()

	m.closed = make(chan struct{})
	go func(f *FIFO) {
		defer func(f *FIFO) {
			close(f.closed)
			f.buffer.Reset()
		}(f)

		for {
			f.processInstantly()

			f.cond.L.Lock()
			elements := f.buffer.ReadAll()
			if len(elements) == 0 {
				if f.status == fifoStateStopping {
					f.status = fifoStateStopped
					f.cond.L.Unlock()
					return
				}
				f.cond.Wait()
				elements = f.buffer.ReadAll() // 重新读取消息
			}
			f.cond.L.Unlock()

			for i := 0; i < len(elements); i++ {
				f.processInstantly()
				elem := elements[i]
				f.handler(elem)
			}
		}
	}(m)
}

func (m *FIFO) Stop() {
	m.cond.L.Lock()
	if m.status != fifoStateRunning {
		m.cond.L.Unlock()
		return
	}
	m.status = fifoStateStopping
	m.cond.L.Unlock()

	m.cond.Signal()

	<-m.closed
}

func (m *FIFO) Enqueue(message MessageContext, instantly bool) bool {
	m.cond.L.Lock()
	if m.status != fifoStateNone {
		if m.status != fifoStateRunning {
			if m.opts.StopMode != FIFOStopModeDrain {
				m.cond.L.Unlock()
				return false
			}
		}
	}

	if instantly {
		m.cond.L.Unlock()
		elem := &instantlyMessage{message: message}
		elem.mu.Lock()
		m.instantly <- elem
		m.cond.Broadcast()
		elem.mu.Lock()
		elem.mu.Unlock()
		return true
	}

	m.buffer.Write(message)
	m.cond.L.Unlock()
	m.cond.Broadcast()
	return true
}

func (m *FIFO) reset() {
	m.cond.L.Lock()
	if m.status < fifoStateStopping {
		m.cond.L.Unlock()
		m.Stop()
	} else {
		m.cond.L.Unlock()
	}

	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	m.buffer.Reset()
	m.status = fifoStateNone

	close(m.instantly)
	m.instantly = make(chan *instantlyMessage, 1)
}

func (m *FIFO) processInstantly() {
	select {
	case elem := <-m.instantly:
		defer elem.mu.Unlock()
		m.handler(elem.message)
	default:
	}
}
