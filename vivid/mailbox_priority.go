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
		opts:    NewPriorityOptions().Apply(opts...),
		status:  priorityStateNone,
		cond:    sync.NewCond(&sync.Mutex{}),
		closed:  make(chan struct{}),
		handler: handler,
	}
	p.buffer = make(priorityHeap, 0, p.opts.BufferSize)
	return p
}

type Priority struct {
	opts    *PriorityOptions             // 配置
	status  priorityState                // 队列状态
	cond    *sync.Cond                   // 消息队列条件变量
	buffer  priorityHeap                 // 消息缓冲区
	closed  chan struct{}                // 关闭信号
	handler func(message MessageContext) // 消息处理函数
}

func (p *Priority) Start() {
	p.cond.L.Lock()
	if p.status != priorityStateNone {
		p.cond.L.Unlock()
		return
	}
	p.status = priorityStateRunning
	p.cond.L.Unlock()

	p.closed = make(chan struct{})
	go func(p *Priority) {
		defer func(p *Priority) {
			close(p.closed)
			if err := recover(); err != nil {
				panic(err)
			}
		}(p)

		for {
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
	}(p)
}

func (p *Priority) Stop() {
	p.cond.L.Lock()
	if p.status != priorityStateRunning {
		p.cond.L.Unlock()
		return
	}
	p.status = priorityStateStopping
	p.cond.L.Unlock()
	p.cond.Signal()
}

func (p *Priority) Enqueue(message MessageContext) bool {
	p.cond.L.Lock()
	if p.status != priorityStateRunning {
		p.cond.L.Unlock()
		return false
	}
	heap.Push(&p.buffer, message)
	p.cond.L.Unlock()
	p.cond.Signal()
	return true
}

func (p *Priority) reset() {
	p.cond.L.Lock()
	p.buffer = p.buffer[:0]
	p.cond.L.Unlock()
}
