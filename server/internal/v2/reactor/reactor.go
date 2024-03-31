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
	"time"
)

// NewReactor 创建一个新的 Reactor 实例，初始化系统级别的队列和多个 Socket 对应的队列
func NewReactor[M any](systemQueueSize, socketQueueSize int, handler MessageHandler[M], errorHandler ErrorHandler[M]) *Reactor[M] {
	r := &Reactor[M]{
		logger:          log.Default().Logger,
		systemQueue:     newQueue[M](-1, systemQueueSize, 1024*16),
		identifiers:     haxmap.New[string, int](),
		lb:              loadbalancer.NewRoundRobin[int, *queue[M]](),
		errorHandler:    errorHandler,
		socketQueueSize: socketQueueSize,
	}

	r.handler = func(q *queue[M], msg M) {
		defer func(msg M) {
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
		}(msg)
		var startedAt = time.Now()
		if handler != nil {
			handler(msg)
		}
		r.log(log.String("action", "handle"), log.Int("queue", q.Id()), log.Int64("cost/ns", time.Since(startedAt).Nanoseconds()))
	}

	return r
}

// Reactor 是一个消息反应器，管理系统级别的队列和多个 Socket 对应的队列
type Reactor[M any] struct {
	logger          *slog.Logger                             // 日志记录器
	systemQueue     *queue[M]                                // 系统级别的队列
	socketQueueSize int                                      // Socket 队列大小
	queues          []*queue[M]                              // Socket 使用的队列
	identifiers     *haxmap.Map[string, int]                 // 标识符到队列索引的映射
	lb              *loadbalancer.RoundRobin[int, *queue[M]] // 负载均衡器
	wg              sync.WaitGroup                           // 等待组
	handler         queueMessageHandler[M]                   // 消息处理器
	errorHandler    ErrorHandler[M]                          // 错误处理器
	debug           bool                                     // 是否开启调试模式
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
	return r.systemQueue.push(msg)
}

// Dispatch 将消息分发到 identifier 使用的队列，当 identifier 首次使用时，将会根据负载均衡策略选择一个队列
func (r *Reactor[M]) Dispatch(identifier string, msg M) error {
	next := r.lb.Next()
	if next == nil {
		return r.Dispatch(identifier, msg)
	}
	idx, _ := r.identifiers.GetOrSet(identifier, next.Id())
	q := r.queues[idx]
	r.log(log.String("action", "dispatch"), log.String("identifier", identifier), log.Int("queue", q.Id()))
	return q.push(msg)
}

// Run 启动 Reactor，运行系统级别的队列和多个 Socket 对应的队列
func (r *Reactor[M]) Run() {
	r.initQueue(r.systemQueue)
	for i := 0; i < runtime.NumCPU(); i++ {
		r.addQueue()
	}
	r.wg.Wait()
}

func (r *Reactor[M]) addQueue() {
	r.log(log.String("action", "add queue"), log.Int("queue", len(r.queues)))
	r.wg.Add(1)
	q := newQueue[M](len(r.queues), r.socketQueueSize, 1024*8)
	r.initQueue(q)
	r.queues = append(r.queues, q)
}

func (r *Reactor[M]) removeQueue(q *queue[M]) {
	idx := q.Id()
	if idx < 0 || idx >= len(r.queues) || r.queues[idx] != q {
		return
	}
	r.queues = append(r.queues[:idx], r.queues[idx+1:]...)
	for i := idx; i < len(r.queues); i++ {
		r.queues[i].idx = i
	}
	r.log(log.String("action", "remove queue"), log.Int("queue", len(r.queues)))
}

func (r *Reactor[M]) initQueue(q *queue[M]) {
	r.wg.Add(1)
	go func(r *Reactor[M], q *queue[M]) {
		defer r.wg.Done()
		go q.run()
		if q.idx >= 0 {
			r.lb.Add(q)
		}
		for m := range q.read() {
			r.handler(q, m)
		}
	}(r, q)
	r.log(log.String("action", "run queue"), log.Int("queue", q.Id()))
}

func (r *Reactor[M]) Close() {
	queues := append(r.queues, r.systemQueue)
	for _, q := range queues {
		q.Close()
	}
}

func (r *Reactor[M]) log(args ...any) {
	if !r.debug {
		return
	}
	r.logger.Debug("Reactor", args...)
}
