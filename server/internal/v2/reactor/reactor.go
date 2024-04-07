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
func NewReactor[M queue.Message](systemQueueSize, queueSize, systemBufferSize, queueBufferSize int, handler MessageHandler[M], errorHandler ErrorHandler[M]) *Reactor[M] {
	if handler == nil {

	}
	r := &Reactor[M]{
		systemQueue:     queue.New[int, string, M](-1, systemQueueSize, systemBufferSize),
		lb:              loadbalancer.NewRoundRobin[int, *queue.Queue[int, string, M]](),
		errorHandler:    errorHandler,
		queueSize:       queueSize,
		queueBufferSize: queueBufferSize,
		state:           statusNone,
		handler:         handler,
		location:        make(map[string]int),
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
type Reactor[M queue.Message] struct {
	logger          atomic.Pointer[log.Logger]                                  // 日志记录器
	state           int32                                                       // 状态
	systemQueue     *queue.Queue[int, string, M]                                // 系统级别的队列
	queueSize       int                                                         // 队列管道大小
	queueBufferSize int                                                         // 队列缓冲区大小
	queues          []*queue.Queue[int, string, M]                              // 所有使用的队列
	queueRW         sync.RWMutex                                                // 队列读写锁
	location        map[string]int                                              // 所在队列 ID 映射
	locationRW      sync.RWMutex                                                // 所在队列 ID 映射锁
	lb              *loadbalancer.RoundRobin[int, *queue.Queue[int, string, M]] // 负载均衡器
	wg              sync.WaitGroup                                              // 等待组
	cwg             sync.WaitGroup                                              // 关闭等待组
	handler         MessageHandler[M]                                           // 消息处理器
	errorHandler    ErrorHandler[M]                                             // 错误处理器
}

// SetLogger 设置日志记录器
func (r *Reactor[M]) SetLogger(logger *log.Logger) {
	r.logger.Store(logger)
}

// GetLogger 获取日志记录器
func (r *Reactor[M]) GetLogger() *log.Logger {
	return r.logger.Load()
}

// process 消息处理
func (r *Reactor[M]) process(msg queue.MessageWrapper[int, string, M]) {
	defer func(msg queue.MessageWrapper[int, string, M]) {
		if err := super.RecoverTransform(recover()); err != nil {
			if r.errorHandler != nil {
				r.errorHandler(msg, err)
			} else {
				r.GetLogger().Error("Reactor", log.String("action", "process"), log.Any("ident", msg.Ident()), log.Int("queue", msg.Queue().Id()), log.Err(err))
				debug.PrintStack()
			}
		}
	}(msg)

	r.handler(msg)
}

// SystemDispatch 将消息分发到系统级别的队列
func (r *Reactor[M]) SystemDispatch(msg M) error {
	if atomic.LoadInt32(&r.state) > statusClosing {
		r.queueRW.RUnlock()
		return fmt.Errorf("reactor closing or closed")
	}
	return r.systemQueue.Push(false, "", msg)
}

// IdentDispatch 将消息分发到 ident 使用的队列，当 ident 首次使用时，将会根据负载均衡策略选择一个队列
//   - 设置 count 会增加消息的外部计数，当 Reactor 关闭时会等待外部计数归零
func (r *Reactor[M]) IdentDispatch(ident string, msg M) error {
	r.queueRW.RLock()
	if atomic.LoadInt32(&r.state) > statusClosing {
		r.queueRW.RUnlock()
		return fmt.Errorf("reactor closing or closed")
	}

	var next *queue.Queue[int, string, M]
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
	return next.Push(true, ident, msg)
}

// Run 启动 Reactor，运行系统级别的队列和多个 Socket 对应的队列
func (r *Reactor[M]) Run(callbacks ...func(queues []*queue.Queue[int, string, M])) {
	if !atomic.CompareAndSwapInt32(&r.state, statusNone, statusRunning) {
		return
	}
	r.queueRW.Lock()
	queues := append([]*queue.Queue[int, string, M]{r.systemQueue}, r.queues...)
	for _, q := range queues {
		r.runQueue(q)
	}
	r.queueRW.Unlock()
	for _, callback := range callbacks {
		callback(queues)
	}
	r.wg.Wait()
}

func (r *Reactor[M]) noneLockAddQueue() {
	q := queue.New[int, string, M](len(r.queues), r.queueSize, r.queueBufferSize)
	r.lb.Add(q) // 运行前添加到负载均衡器，未运行时允许接收消息
	r.queues = append(r.queues, q)
}

func (r *Reactor[M]) noneLockDelQueue(q *queue.Queue[int, string, M]) {
	idx := q.Id()
	if idx < 0 || idx >= len(r.queues) || r.queues[idx] != q {
		return
	}
	r.queues = append(r.queues[:idx], r.queues[idx+1:]...)
	for i := idx; i < len(r.queues); i++ {
		r.queues[i].SetId(i)
	}
}

func (r *Reactor[M]) runQueue(q *queue.Queue[int, string, M]) {
	r.wg.Add(1)
	q.SetClosedHandler(func(q *queue.Queue[int, string, M]) {
		// 关闭时正在等待关闭完成，外部已加锁，无需再次加锁
		r.noneLockDelQueue(q)
		r.cwg.Done()
		r.logger.Load().Debug("Reactor", log.String("action", "close"), log.Any("queue", q.Id()))
	})
	go q.Run()

	go func(r *Reactor[M], q *queue.Queue[int, string, M]) {
		defer r.wg.Done()
		for m := range q.Read() {
			m(r.process, func(m queue.MessageWrapper[int, string, M], last bool) {
				if last {
					r.locationRW.RLock()
					mq, exist := r.location[m.Ident()]
					r.locationRW.RUnlock()
					if exist {
						r.locationRW.Lock()
						defer r.locationRW.Unlock()
						mq, exist = r.location[m.Ident()]
						if exist {
							delete(r.location, m.Ident())
							r.queueRW.RLock()
							mq := r.queues[mq]
							r.queueRW.RUnlock()
							r.logger.Load().Debug("Reactor", log.String("action", "unbind"), log.Any("ident", m.Ident()), log.Any("queue", mq.Id()))
						}
					}
				}
			})
		}
	}(r, q)
}

func (r *Reactor[M]) Close() {
	if !atomic.CompareAndSwapInt32(&r.state, statusRunning, statusClosing) {
		return
	}
	r.queueRW.Lock()
	r.cwg.Add(len(r.queues) + 1)
	for _, q := range append(r.queues, r.systemQueue) {
		q.Close()
	}
	r.cwg.Wait()
	atomic.StoreInt32(&r.state, statusClosed)
	r.queueRW.Unlock()
}
