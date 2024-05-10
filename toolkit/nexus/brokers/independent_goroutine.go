package brokers

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"sync"
	"sync/atomic"
)

const (
	independentGoroutineStatusNone    = iota // 未运行
	independentGoroutineStatusRunning        // 运行中
	independentGoroutineStatusClosing        // 关闭中
	independentGoroutineStatusClosed         // 已关闭
)

type (
	IndependentGoroutineBinder[T comparable]   func(topic T)
	IndependentGoroutineUnBinder[T comparable] func(topic T)
	IndependentGoroutineMessageHandler         func(handler nexus.EventExecutor)
)

// NewIndependentGoroutine 创建一个 IndependentGoroutine，该 IndependentGoroutine 支持通过 queueFactory 指定队列工厂函数
//   - 为避免队列高频创建和销毁，创建该队列将返回一个绑定函数和解绑函数，当不再使用时，调用解绑函数，解绑函数会在队列为空时关闭队列
func NewIndependentGoroutine[I constraints.Ordered, T comparable](queueFactory func(index int) nexus.Queue[I, T], handler IndependentGoroutineMessageHandler, options ...*IndependentGoroutineOptions[I, T]) (nexus.Broker[I, T], IndependentGoroutineBinder[T], IndependentGoroutineUnBinder[T]) {
	i := &IndependentGoroutine[I, T]{
		opts:         NewIndependentGoroutineOptions[I, T]().Apply(options...),
		queues:       make(map[T]nexus.Queue[I, T]),
		queueFactory: queueFactory,
		bindCounter:  make(map[T]int),
		handler:      handler,
		closed:       make(chan struct{}),
	}

	return i, i.bind, i.unBind
}

// IndependentGoroutine 每个队列都使用独立的 goroutine 来处理消息
type IndependentGoroutine[I constraints.Ordered, T comparable] struct {
	state  int32         // 状态
	closed chan struct{} // 关闭信号
	opts   *IndependentGoroutineOptions[I, T]

	guid         int                                // 队列唯一标识
	queues       map[T]nexus.Queue[I, T]            // 所有使用的队列
	queueRW      sync.RWMutex                       // 队列读写锁
	queueFactory func(index int) nexus.Queue[I, T]  // 队列工厂函数
	bindCounter  map[T]int                          // 绑定计数器
	handler      IndependentGoroutineMessageHandler // 消息处理器
}

func (i *IndependentGoroutine[I, T]) bind(topic T) {
	i.queueRW.Lock()
	curr := i.bindCounter[topic] + 1
	i.bindCounter[topic] = curr
	i.queueRW.Unlock()

	if i.opts.queueBindCounterChangedHook != nil {
		i.opts.queueBindCounterChangedHook(topic, curr)
	}
}

func (i *IndependentGoroutine[I, T]) unBind(topic T) {
	i.queueRW.Lock()
	curr := i.bindCounter[topic] - 1
	if curr <= 0 {
		queue, exist := i.queues[topic]
		delete(i.bindCounter, topic)
		delete(i.queues, topic)
		if exist {
			// 尝试关闭
			queue.Close()
			if i.opts.queueClosedHook != nil {
				i.opts.queueClosedHook(topic, queue, len(i.queues))
			}
		}

	} else {
		i.bindCounter[topic] = curr
	}
	i.queueRW.Unlock()

	if i.opts.queueBindCounterChangedHook != nil {
		i.opts.queueBindCounterChangedHook(topic, curr)
	}
}

func (i *IndependentGoroutine[I, T]) runQueue(topic T, queue nexus.Queue[I, T], num int) {
	go queue.Run()
	if i.opts.queueCreatedHook != nil {
		i.opts.queueCreatedHook(topic, queue, num)
	}
	for h := range queue.Consume() {
		h.Exec(
			// onProcess
			func(topic T, event nexus.EventExecutor) {
				i.handler(event)
			},
			// onFinish
			func(topic T, last bool) {
				if !last {
					return
				}
				// 为最后一条消息时，检查队列绑定计数器，如果为 0 则关闭队列
				i.queueRW.Lock()
				defer i.queueRW.Unlock()
				if i.bindCounter[topic] <= 0 && i.queues[topic] != nil { // 防止重复关闭
					queue.Close()
					delete(i.queues, topic)
					if i.opts.queueClosedHook != nil {
						i.opts.queueClosedHook(topic, queue, len(i.queues))
					}
				}
			},
		)
	}
}

func (i *IndependentGoroutine[I, T]) Run() {
	if !atomic.CompareAndSwapInt32(&i.state, independentGoroutineStatusNone, independentGoroutineStatusRunning) {
		return
	}

	<-i.closed
}

func (i *IndependentGoroutine[I, T]) Close() {
	if !atomic.CompareAndSwapInt32(&i.state, independentGoroutineStatusRunning, independentGoroutineStatusClosing) {
		return
	}

	// 由于关闭中也可以产生新队列，关闭时需多轮检查
	defer close(i.closed)
	for {
		var wg = new(sync.WaitGroup)
		i.queueRW.RLock()
		if len(i.queues) == 0 {
			atomic.StoreInt32(&i.state, independentGoroutineStatusClosed)
			i.queueRW.RUnlock()
			break
		}

		// 向所有队列发送关闭信号，异步等待关闭，释放锁，避免新消息无法进入
		for key, queue := range i.queues {
			count := i.bindCounter[key]
			if count > 0 {
				continue
			}

			wg.Add(1)
			go func(wg *sync.WaitGroup, queue nexus.Queue[I, T]) {
				defer wg.Done()
				queue.Close()

				i.queueRW.Lock()
				defer i.queueRW.Unlock()

				delete(i.queues, key)
				if i.opts.queueClosedHook != nil {
					i.opts.queueClosedHook(key, queue, len(i.queues))
				}
			}(wg, queue)
		}

		i.queueRW.RUnlock()

		// 等待本轮关闭进入下一轮
		wg.Wait()
	}
}

func (i *IndependentGoroutine[I, T]) Publish(topic T, event nexus.Event[I, T]) error {
	if atomic.LoadInt32(&i.state) > independentGoroutineStatusClosing {
		return fmt.Errorf("broker closed")
	}

	i.queueRW.RLock()
	queue, exist := i.queues[topic]
	i.queueRW.RUnlock()
	if !exist {
		// 双重检查确保锁粒度最小
		i.queueRW.Lock()
		if queue, exist = i.queues[topic]; !exist {
			// 创建新队列
			i.guid++
			queue = i.queueFactory(i.guid)
			go i.runQueue(topic, queue, len(i.queues)+1)
			i.queues[topic] = queue
		}
		i.queueRW.Unlock()
	}

	return queue.Publish(topic, event)
}
