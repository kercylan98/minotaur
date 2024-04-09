package brokers

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/loadbalancer"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"runtime"
	"sync"
	"sync/atomic"
)

const (
	sparseGoroutineStatusNone    = iota - 1 // 未运行
	sparseGoroutineStatusRunning            // 运行中
	sparseGoroutineStatusClosing            // 关闭中
	sparseGoroutineStatusClosed             // 已关闭
)

type (
	SparseGoroutineMessageHandler func(handler nexus.EventExecutor)
)

func NewSparseGoroutine[I, T comparable](queueFactory func(index int) nexus.Queue[I, T], handler SparseGoroutineMessageHandler) nexus.Broker[I, T] {
	s := &SparseGoroutine[I, T]{
		lb:           loadbalancer.NewRoundRobin[I, nexus.Queue[I, T]](),
		queues:       make(map[I]nexus.Queue[I, T]),
		state:        sparseGoroutineStatusNone,
		location:     make(map[T]I),
		handler:      handler,
		queueFactory: queueFactory,
	}
	defaultNum := runtime.NumCPU()
	s.queueRW.Lock()
	for i := 0; i < defaultNum; i++ {
		queue := s.queueFactory(i)
		s.lb.Add(queue) // 运行前添加到负载均衡器，未运行时允许接收消息
		queueId := queue.GetId()
		if _, exist := s.queues[queueId]; exist {
			panic(fmt.Errorf("duplicate queue id: %v", queueId))
		}
		s.queues[queueId] = queue
	}
	s.queueRW.Unlock()

	return s
}

type SparseGoroutine[I, T comparable] struct {
	state           int32                                          // 状态
	queueSize       int                                            // 队列管道大小
	queueBufferSize int                                            // 队列缓冲区大小
	queues          map[I]nexus.Queue[I, T]                        // 所有使用的队列
	queueRW         sync.RWMutex                                   // 队列读写锁
	location        map[T]I                                        // Topic 所在队列 Id 映射
	locationRW      sync.RWMutex                                   // 所在队列 ID 映射锁
	lb              *loadbalancer.RoundRobin[I, nexus.Queue[I, T]] // 负载均衡器
	wg              sync.WaitGroup                                 // 等待组
	handler         SparseGoroutineMessageHandler                  // 消息处理器

	queueFactory func(index int) nexus.Queue[I, T]
}

// Run 启动 Reactor，运行队列
func (s *SparseGoroutine[I, T]) Run() {
	if !atomic.CompareAndSwapInt32(&s.state, sparseGoroutineStatusNone, sparseGoroutineStatusRunning) {
		return
	}
	s.queueRW.Lock()
	queues := s.queues
	for _, queue := range queues {
		s.wg.Add(1)
		go queue.Run()

		go func(r *SparseGoroutine[I, T], queue nexus.Queue[I, T]) {
			defer r.wg.Done()
			for h := range queue.Consume() {
				h.Exec(
					// onProcess
					func(topic T, event nexus.EventExecutor) {
						s.handler(event)
					},
					// onFinish
					func(topic T, last bool) {
						if !last {
							return
						}
						s.locationRW.Lock()
						defer s.locationRW.Unlock()
						delete(s.location, topic)
					},
				)
			}
		}(s, queue)
	}
	s.queueRW.Unlock()
	s.wg.Wait()
}

func (s *SparseGoroutine[I, T]) Close() {
	if !atomic.CompareAndSwapInt32(&s.state, sparseGoroutineStatusRunning, sparseGoroutineStatusClosing) {
		return
	}
	s.queueRW.Lock()
	var wg sync.WaitGroup
	wg.Add(len(s.queues))
	for _, queue := range s.queues {
		go func(queue nexus.Queue[I, T]) {
			defer wg.Done()
			queue.Close()
		}(queue)
	}
	wg.Wait()
	atomic.StoreInt32(&s.state, sparseGoroutineStatusClosed)
	s.queueRW.Unlock()
}

// Publish 将消息分发到特定 topic，当 topic 首次使用时，将会根据负载均衡策略选择一个队列
//   - 设置 count 会增加消息的外部计数，当 SparseGoroutine 关闭时会等待外部计数归零
func (s *SparseGoroutine[I, T]) Publish(topic T, event nexus.Event[I, T]) error {
	s.queueRW.RLock()
	if atomic.LoadInt32(&s.state) > sparseGoroutineStatusClosing {
		s.queueRW.RUnlock()
		return fmt.Errorf("broker closing or closed")
	}

	var next nexus.Queue[I, T]
	s.locationRW.RLock()
	i, exist := s.location[topic]
	s.locationRW.RUnlock()
	if !exist {
		s.locationRW.Lock()
		if i, exist = s.location[topic]; !exist {
			next = s.lb.Next()
			s.location[topic] = next.GetId()
		} else {
			next = s.queues[i]
		}
		s.locationRW.Unlock()
	} else {
		next = s.queues[i]
	}
	s.queueRW.RUnlock()

	event.OnInitialize(context.Background(), s)
	return next.Publish(topic, event)
}
