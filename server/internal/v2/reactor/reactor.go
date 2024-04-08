package reactor

import (
	"fmt"
	"github.com/kercylan98/minotaur/server/internal/v2/loadbalancer"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"github.com/kercylan98/minotaur/utils/super"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

const (
	statusNone    = iota - 1 // 事件循环未运行
	statusRunning            // 事件循环运行中
	statusClosing            // 事件循环关闭中
	statusClosed             // 事件循环已关闭
)

// NewReactor 创建一个新的 Reactor 实例，初始化系统级别的队列和多个 Socket 对应的队列
func NewReactor[Queue comparable, M queue.Message[Queue]](queueSize, queueBufferSize int, handler MessageHandler[Queue, M], errorHandler ErrorHandler[Queue, M]) *Reactor[Queue, M] {
	if handler == nil {

	}
	r := &Reactor[Queue, M]{
		lb:              loadbalancer.NewRoundRobin[int, *queue.Queue[int, Queue, M]](),
		errorHandler:    errorHandler,
		queueSize:       queueSize,
		queueBufferSize: queueBufferSize,
		state:           statusNone,
		handler:         handler,
		location:        make(map[Queue]int),
	}
	r.logger.Store(log.GetLogger())

	defaultNum := runtime.NumCPU()
	r.queueRW.Lock()
	for i := 0; i < defaultNum; i++ {
		r.noneLockAddQueue()
	}
	r.queueRW.Unlock()

	return r
}

// Reactor 是一个消息反应器，管理系统级别的队列和多个 Socket 对应的队列
type Reactor[Queue comparable, M queue.Message[Queue]] struct {
	logger          atomic.Pointer[log.Logger]                                 // 日志记录器
	state           int32                                                      // 状态
	queueSize       int                                                        // 队列管道大小
	queueBufferSize int                                                        // 队列缓冲区大小
	queues          []*queue.Queue[int, Queue, M]                              // 所有使用的队列
	queueRW         sync.RWMutex                                               // 队列读写锁
	location        map[Queue]int                                              // 所在队列 ID 映射
	locationRW      sync.RWMutex                                               // 所在队列 ID 映射锁
	lb              *loadbalancer.RoundRobin[int, *queue.Queue[int, Queue, M]] // 负载均衡器
	wg              sync.WaitGroup                                             // 等待组
	cwg             sync.WaitGroup                                             // 关闭等待组
	handler         MessageHandler[Queue, M]                                   // 消息处理器
	errorHandler    ErrorHandler[Queue, M]                                     // 错误处理器
}

// SetLogger 设置日志记录器
func (r *Reactor[Queue, M]) SetLogger(logger *log.Logger) {
	r.logger.Store(logger)
}

// GetLogger 获取日志记录器
func (r *Reactor[Queue, M]) GetLogger() *log.Logger {
	return r.logger.Load()
}

// process 消息处理
func (r *Reactor[Queue, M]) process(msg M) {
	defer func(msg M) {
		if err := super.RecoverTransform(recover()); err != nil {
			if r.errorHandler != nil {
				r.errorHandler(msg, err)
			} else {
				r.GetLogger().Error("Reactor", log.String("action", "process"), log.Any("queue", msg.GetQueue()), log.Err(err))
				debug.PrintStack()
			}
		}
	}(msg)

	r.handler(msg)
}

// Dispatch 将消息分发到 ident 使用的队列，当 ident 首次使用时，将会根据负载均衡策略选择一个队列
//   - 设置 count 会增加消息的外部计数，当 Reactor 关闭时会等待外部计数归零
//   - 当 ident 为空字符串时候，将发送到
func (r *Reactor[Queue, M]) Dispatch(ident Queue, msg M) error {
	r.queueRW.RLock()
	if atomic.LoadInt32(&r.state) > statusClosing {
		r.queueRW.RUnlock()
		return fmt.Errorf("reactor closing or closed")
	}

	var next *queue.Queue[int, Queue, M]
	r.locationRW.RLock()
	i, exist := r.location[ident]
	r.locationRW.RUnlock()
	if !exist {
		r.locationRW.Lock()
		if i, exist = r.location[ident]; !exist {
			next = r.lb.Next()
			r.location[ident] = next.Id()
			r.logger.Load().Debug("Reactor", log.String("action", "bind"), log.Any("ident", ident), log.Any("queue", next.Id()))
		} else {
			next = r.queues[i]
		}
		r.locationRW.Unlock()
	} else {
		next = r.queues[i]
	}
	r.queueRW.RUnlock()
	return next.Push(ident, msg)
}

// Run 启动 Reactor，运行系统级别的队列和多个 Socket 对应的队列
func (r *Reactor[Queue, M]) Run(callbacks ...func(queues []*queue.Queue[int, Queue, M])) {
	if !atomic.CompareAndSwapInt32(&r.state, statusNone, statusRunning) {
		return
	}
	r.queueRW.Lock()
	queues := r.queues
	for _, q := range queues {
		r.runQueue(q)
	}
	r.queueRW.Unlock()
	for _, callback := range callbacks {
		callback(queues)
	}
	r.wg.Wait()
}

func (r *Reactor[Queue, M]) noneLockAddQueue() {
	q := queue.New[int, Queue, M](len(r.queues), r.queueSize, r.queueBufferSize)
	r.lb.Add(q) // 运行前添加到负载均衡器，未运行时允许接收消息
	r.queues = append(r.queues, q)
}

func (r *Reactor[Queue, M]) noneLockDelQueue(q *queue.Queue[int, Queue, M]) {
	idx := q.Id()
	if idx < 0 || idx >= len(r.queues) || r.queues[idx] != q {
		return
	}
	r.queues = append(r.queues[:idx], r.queues[idx+1:]...)
	for i := idx; i < len(r.queues); i++ {
		r.queues[i].SetId(i)
	}
}

func (r *Reactor[Queue, M]) runQueue(q *queue.Queue[int, Queue, M]) {
	r.wg.Add(1)
	q.SetClosedHandler(func(q *queue.Queue[int, Queue, M]) {
		// 关闭时正在等待关闭完成，外部已加锁，无需再次加锁
		r.noneLockDelQueue(q)
		r.cwg.Done()
		r.logger.Load().Debug("Reactor", log.String("action", "close"), log.Any("queue", q.Id()))
	})
	go q.Run()

	go func(r *Reactor[Queue, M], q *queue.Queue[int, Queue, M]) {
		defer r.wg.Done()
		for m := range q.Read() {
			m(r.process, r.processFinish)
		}
	}(r, q)
}

func (r *Reactor[Queue, M]) Close() {
	if !atomic.CompareAndSwapInt32(&r.state, statusRunning, statusClosing) {
		return
	}
	r.queueRW.Lock()
	r.cwg.Add(len(r.queues) + 1)
	for _, q := range r.queues {
		q.Close()
	}
	r.cwg.Wait()
	atomic.StoreInt32(&r.state, statusClosed)
	r.queueRW.Unlock()
}

func (r *Reactor[Queue, M]) processFinish(m M, last bool) {
	if !last {
		return
	}
	queueName := m.GetQueue()

	r.locationRW.RLock()
	mq, exist := r.location[queueName]
	r.locationRW.RUnlock()
	if exist {
		r.locationRW.Lock()
		defer r.locationRW.Unlock()
		mq, exist = r.location[queueName]
		if exist {
			delete(r.location, queueName)
			r.queueRW.RLock()
			mq := r.queues[mq]
			r.queueRW.RUnlock()
			r.logger.Load().Debug("Reactor", log.String("action", "unbind"), log.Any("queueName", queueName), log.Any("queue", mq.Id()))
		}
	}
}
