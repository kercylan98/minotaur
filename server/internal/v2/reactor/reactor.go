package reactor

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/server/internal/v2/loadbalancer"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/super"
	"log/slog"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

const (
	StatusNone    = iota - 1 // 事件循环未运行
	StatusRunning            // 事件循环运行中
	StatusClosing            // 事件循环关闭中
	StatusClosed             // 事件循环已关闭
)

var sysIdent = &identifiable{ident: "system"}

// NewReactor 创建一个新的 Reactor 实例，初始化系统级别的队列和多个 Socket 对应的队列
func NewReactor[M any](systemQueueSize, queueSize int, handler MessageHandler[M], errorHandler ErrorHandler[M]) *Reactor[M] {
	r := &Reactor[M]{
		logger:       log.Default().Logger,
		systemQueue:  newQueue[M](-1, systemQueueSize, 1024),
		identifiers:  haxmap.New[string, *identifiable](),
		lb:           loadbalancer.NewRoundRobin[int, *queue[M]](),
		errorHandler: errorHandler,
		queueSize:    queueSize,
		state:        StatusNone,
	}

	defaultNum := runtime.NumCPU()
	if defaultNum < 1 {
		defaultNum = 1
	}

	r.queueRW.Lock()
	for i := 0; i < defaultNum; i++ {
		r.noneLockAddQueue()
	}
	r.queueRW.Unlock()

	r.handler = func(q *queue[M], ident *identifiable, msg M) {
		defer func(ident *identifiable, msg M) {
			if err := super.RecoverTransform(recover()); err != nil {
				defer func(msg M) {
					if err = super.RecoverTransform(recover()); err != nil {
						fmt.Println(err)
						debug.PrintStack()
					}
				}(msg)
				if r.errorHandler != nil {
					r.errorHandler(msg, err)
				} else {
					fmt.Println(err)
					debug.PrintStack()
				}
			}

			if atomic.AddInt64(&ident.n, -1) == 0 {
				r.queueRW.Lock()
				r.identifiers.Del(ident.ident)
				r.queueRW.Unlock()
			}

		}(ident, msg)
		if handler != nil {
			handler(msg)
		}
	}

	return r
}

// Reactor 是一个消息反应器，管理系统级别的队列和多个 Socket 对应的队列
type Reactor[M any] struct {
	logger       *slog.Logger                             // 日志记录器
	state        int32                                    // 状态
	systemQueue  *queue[M]                                // 系统级别的队列
	queueSize    int                                      // Socket 队列大小
	queues       []*queue[M]                              // Socket 使用的队列
	queueRW      sync.RWMutex                             // 队列读写锁
	identifiers  *haxmap.Map[string, *identifiable]       // 标识符到队列索引的映射及消息计数
	lb           *loadbalancer.RoundRobin[int, *queue[M]] // 负载均衡器
	wg           sync.WaitGroup                           // 等待组
	cwg          sync.WaitGroup                           // 关闭等待组
	handler      queueMessageHandler[M]                   // 消息处理器
	errorHandler ErrorHandler[M]                          // 错误处理器
	debug        bool                                     // 是否开启调试模式
}

// SetLogger 设置日志记录器
func (r *Reactor[M]) SetLogger(logger *slog.Logger) *Reactor[M] {
	r.logger = logger
	return r
}

// SetDebug 设置是否开启调试模式
func (r *Reactor[M]) SetDebug(debug bool) *Reactor[M] {
	r.debug = debug
	return r
}

// SystemDispatch 将消息分发到系统级别的队列
func (r *Reactor[M]) SystemDispatch(msg M) error {
	if atomic.LoadInt32(&r.state) > StatusRunning {
		r.queueRW.RUnlock()
		return fmt.Errorf("reactor closing or closed")
	}
	return r.systemQueue.push(sysIdent, msg)
}

// Dispatch 将消息分发到 ident 使用的队列，当 ident 首次使用时，将会根据负载均衡策略选择一个队列
func (r *Reactor[M]) Dispatch(ident string, msg M) error {
	r.queueRW.RLock()
	if atomic.LoadInt32(&r.state) > StatusRunning {
		r.queueRW.RUnlock()
		return fmt.Errorf("reactor closing or closed")
	}
	next := r.lb.Next()
	i, _ := r.identifiers.GetOrSet(ident, &identifiable{ident: ident})
	q := r.queues[next.Id()]
	atomic.AddInt64(&i.n, 1)
	r.queueRW.RUnlock()
	return q.push(i, msg)
}

// Run 启动 Reactor，运行系统级别的队列和多个 Socket 对应的队列
func (r *Reactor[M]) Run() {
	if !atomic.CompareAndSwapInt32(&r.state, StatusNone, StatusRunning) {
		return
	}
	r.queueRW.Lock()
	r.runQueue(r.systemQueue)
	for i := 0; i < len(r.queues); i++ {
		r.runQueue(r.queues[i])
	}
	r.queueRW.Unlock()
	r.wg.Wait()
}

func (r *Reactor[M]) noneLockAddQueue() {
	q := newQueue[M](len(r.queues), r.queueSize, 1024*8)
	r.lb.Add(q) // 运行前添加到负载均衡器，未运行时允许接收消息
	r.queues = append(r.queues, q)
}

func (r *Reactor[M]) noneLockDelQueue(q *queue[M]) {
	idx := q.Id()
	if idx < 0 || idx >= len(r.queues) || r.queues[idx] != q {
		return
	}
	r.queues = append(r.queues[:idx], r.queues[idx+1:]...)
	for i := idx; i < len(r.queues); i++ {
		r.queues[i].idx = i
	}
}

func (r *Reactor[M]) runQueue(q *queue[M]) {
	r.wg.Add(1)
	q.setClosedHandler(func(q *queue[M]) {
		// 关闭时正在等待关闭完成，外部已加锁，无需再次加锁
		r.noneLockDelQueue(q)
		r.cwg.Done()
	})
	go q.run()

	go func(r *Reactor[M], q *queue[M]) {
		defer r.wg.Done()
		for m := range q.read() {
			r.handler(q, m.ident, m.msg)
		}
	}(r, q)
}

func (r *Reactor[M]) Close() {
	if !atomic.CompareAndSwapInt32(&r.state, StatusRunning, StatusClosing) {
		return
	}
	r.queueRW.Lock()
	r.cwg.Add(len(r.queues) + 1)
	for _, q := range append(r.queues, r.systemQueue) {
		q.Close()
	}
	r.cwg.Wait()
	atomic.StoreInt32(&r.state, StatusClosed)
	r.queueRW.Unlock()
}
