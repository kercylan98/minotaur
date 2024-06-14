package vivid

import (
	"container/heap"
	"sync"
)

const (
	priorityStateNone     = priorityState(iota) // 未启动状态
	priorityStateRunning                        // 运行中状态
	priorityStateStopping                       // 停止中状态
	priorityStateStopped                        // 已停止状态
)

const (
	PriorityStopModeInstantly = PriorityStopMode(iota) // 立刻停止消息队列，新消息将不再接收，缓冲区内未处理的消息将被丢弃
	PriorityStopModeGraceful                           // 优雅停止消息队列，新消息将不再接收，等待未处理的消息处理完毕后再停止
	PriorityStopModeDrain                              // 新消息将继续被接收，等待消息队列处理完毕且没有新消息后再停止
)

type (
	priorityState    = int32 // 状态
	PriorityStopMode = int8  // Priority 消息队列的停止模式，目前支持 PriorityStopModeInstantly、 PriorityStopModeGraceful、 PriorityStopModeDrain
)

func NewPriority(handler func(message MessageContext), opts ...*PriorityOptions) *Priority {
	p := &Priority{
		opts:      NewPriorityOptions().Apply(opts...),
		status:    priorityStateNone,
		cond:      sync.NewCond(&sync.Mutex{}),
		closed:    make(chan struct{}),
		handler:   handler,
		instantly: make(chan *instantlyMessage, 1),
	}
	p.buffer = make(priorityHeap, 0, p.opts.BufferSize)
	return p
}

type Priority struct {
	opts      *PriorityOptions             // 配置
	status    priorityState                // 队列状态
	cond      *sync.Cond                   // 消息队列条件变量
	buffer    priorityHeap                 // 消息缓冲区
	closed    chan struct{}                // 关闭信号
	handler   func(message MessageContext) // 消息处理函数
	instantly chan *instantlyMessage       // 立即处理的消息
}

func (m *Priority) Start() {
	m.cond.L.Lock()
	if m.status != priorityStateNone {
		m.cond.L.Unlock()
		return
	}
	m.status = priorityStateRunning
	m.cond.L.Unlock()

	m.closed = make(chan struct{})
	go func(p *Priority) {
		defer func(p *Priority) {
			close(p.closed)
			if err := recover(); err != nil {
				panic(err)
			}
		}(p)

		for {
			p.processInstantly()
			p.cond.L.Lock()
			if p.buffer.Len() == 0 {
				if p.status == priorityStateStopping {
					p.status = priorityStateStopped
					p.cond.L.Unlock()
					break
				}
				p.cond.Wait()
				if p.buffer.Len() == 0 {
					p.cond.L.Unlock()
					continue
				}
			}
			msg := heap.Pop(&p.buffer).(MessageContext)
			p.cond.L.Unlock()

			p.handler(msg)
		}
	}(m)
}

func (m *Priority) Stop() {
	m.cond.L.Lock()
	if m.status != priorityStateRunning {
		m.cond.L.Unlock()
		return
	}
	m.status = priorityStateStopping
	m.cond.L.Unlock()
	m.cond.Signal()
}

func (m *Priority) Enqueue(message MessageContext, instantly bool) bool {
	m.cond.L.Lock()
	if m.status != priorityStateRunning {
		m.cond.L.Unlock()
		return false
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

	heap.Push(&m.buffer, message)
	m.cond.L.Unlock()
	m.cond.Signal()
	return true
}

func (m *Priority) reset() {
	m.cond.L.Lock()
	if m.status < fifoStateStopping {
		m.cond.L.Unlock()
		m.Stop()
	} else {
		m.cond.L.Unlock()
	}

	m.cond.L.Lock()
	defer m.cond.L.Unlock()

	m.status = priorityStateNone
	m.buffer = m.buffer[:0]
}

func (m *Priority) processInstantly() {
	select {
	case elem := <-m.instantly:
		defer elem.mu.Unlock()
		m.handler(elem.message)
	default:
	}
}
