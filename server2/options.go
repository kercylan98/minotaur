package server

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"os"
	"runtime"
	"sync"
	"time"
)

func NewOptions() *Options {
	return DefaultOptions()
}

func DefaultOptions() *Options {
	opt := &Options{
		messageChannelSize:       1024 * 4,
		messageBufferInitialSize: 1024,
		lifeCycleLimit:           0,
		logger:                   log.NewLogger(log.NewHandler(os.Stdout, log.DefaultOptions().WithCallerSkip(-1).WithLevel(log.LevelInfo))),
		sparseGoroutineNum:       runtime.NumCPU(),
		eventOptions:             nexus.NewEventOptions().WithLowHandlerThreshold(0, nil),
	}
	return opt
}

type Options struct {
	server                   *server
	rw                       sync.RWMutex
	messageChannelSize       int                 // 服务器消息处理管道大小
	messageBufferInitialSize int                 // 服务器消息写入缓冲区初始化大小
	lifeCycleLimit           time.Duration       // 服务器生命周期上限，在服务器启动后达到生命周期上限将关闭服务器
	logger                   *log.Logger         // 日志记录器
	debug                    bool                // Debug 模式
	syncLowMessageDuration   time.Duration       // 同步慢消息时间
	asyncLowMessageDuration  time.Duration       // 异步慢消息时间
	sparseGoroutineNum       int                 // 稀疏 goroutine 数量
	eventOptions             *nexus.EventOptions // 事件选项
	zombieConnectionDeadline time.Duration       // 僵尸连接超时时间
}

func (opt *Options) init(srv *server) *Options {
	opt.server = srv
	return opt
}

func (opt *Options) Apply(options ...*Options) {
	opt.rw.Lock()
	defer opt.rw.Unlock()
	for _, option := range options {
		option.rw.RLock()

		opt.messageChannelSize = option.messageChannelSize
		opt.messageBufferInitialSize = option.messageBufferInitialSize
		opt.lifeCycleLimit = option.lifeCycleLimit
		opt.logger = option.logger
		opt.debug = option.debug
		opt.syncLowMessageDuration = option.syncLowMessageDuration
		opt.asyncLowMessageDuration = option.asyncLowMessageDuration
		opt.sparseGoroutineNum = option.sparseGoroutineNum
		opt.eventOptions = option.eventOptions
		opt.zombieConnectionDeadline = option.zombieConnectionDeadline

		option.rw.RUnlock()
	}
	if opt.server != nil && !opt.server.state.LaunchedAt.IsZero() {
		opt.active()
	}
}

func (opt *Options) active() {
	opt.server.notify.lifeCycleTime <- opt.GetLifeCycleLimit()
}

// WithZombieConnectionDeadline 设置僵尸连接超时时间，当连接超过该时间未收到任何消息时，将关闭连接
func (opt *Options) WithZombieConnectionDeadline(deadline time.Duration) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.zombieConnectionDeadline = deadline
	})
}

func (opt *Options) GetZombieConnectionDeadline() time.Duration {
	return getOptionsValue(opt, func(opt *Options) time.Duration {
		return opt.zombieConnectionDeadline
	})
}

// WithEventOptions 设置服务器事件选项
//   - 该函数支持运行时设置
func (opt *Options) WithEventOptions(eventOptions *nexus.EventOptions) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.eventOptions = eventOptions
	})
}

func (opt *Options) GetEventOptions() *nexus.EventOptions {
	return getOptionsValue(opt, func(opt *Options) *nexus.EventOptions {
		return opt.eventOptions
	})
}

// WithSparseGoroutineNum 设置服务器稀疏 goroutine 数量
//   - 该函数在运行时设置无效
func (opt *Options) WithSparseGoroutineNum(num int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.sparseGoroutineNum = num
	})
}

func (opt *Options) GetSparseGoroutineNum() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.sparseGoroutineNum
	})
}

// WithSyncLowMessageMonitor 设置同步消息的慢消息监测时间
func (opt *Options) WithSyncLowMessageMonitor(duration time.Duration) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.syncLowMessageDuration = duration
	})
}

func (opt *Options) GetSyncLowMessageDuration() time.Duration {
	return getOptionsValue(opt, func(opt *Options) time.Duration {
		return opt.syncLowMessageDuration
	})
}

// WithAsyncLowMessageMonitor 设置异步消息的慢消息监测时间
func (opt *Options) WithAsyncLowMessageMonitor(duration time.Duration) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.asyncLowMessageDuration = duration
	})
}

func (opt *Options) GetAsyncLowMessageDuration() time.Duration {
	return getOptionsValue(opt, func(opt *Options) time.Duration {
		return opt.asyncLowMessageDuration
	})
}

// WithDebug 设置 Debug 模式是否开启
//   - 该函数支持运行时设置
func (opt *Options) WithDebug(debug bool) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.debug = true
	})
}

// WithLogger 设置服务器的日志记录器
//   - 该函数支持运行时设置
func (opt *Options) WithLogger(logger *log.Logger) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.logger = logger
	})
}

// GetLogger 获取服务器的日志记录器
func (opt *Options) GetLogger() *log.Logger {
	return getOptionsValue(opt, func(opt *Options) *log.Logger {
		return opt.logger
	})
}

// WithMessageChannelSize 设置服务器 Actor 用于处理消息的管道大小，当管道由于逻辑阻塞而导致满载时，会导致新消息无法及时从缓冲区拿出，从而增加内存的消耗，但是并不会影响消息的写入
func (opt *Options) WithMessageChannelSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.messageChannelSize = size
	})
}

func (opt *Options) GetServerMessageChannelSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.messageChannelSize
	})
}

// WithMessageBufferInitialSize 设置服务器 Actor 消息环形缓冲区 buffer.Ring 的初始大小，适当的值可以避免频繁扩容
//   - 由于扩容是按照当前大小的 2 倍进行扩容，过大的值也可能会导致内存消耗过高
func (opt *Options) WithMessageBufferInitialSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.messageBufferInitialSize = size
	})
}

func (opt *Options) GetServerMessageBufferInitialSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.messageBufferInitialSize
	})
}

// WithLifeCycleLimit 设置服务器生命周期上限，在服务器启动后达到生命周期上限将关闭服务器
//   - 如果设置为 <= 0 的值，将不限制服务器生命周期
//   - 该函数支持运行时设置
func (opt *Options) WithLifeCycleLimit(limit time.Duration) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.lifeCycleLimit = limit
	})
}

// WithLifeCycleEnd 设置服务器生命周期终点，在服务器达到该时间后将关闭服务器
//   - 如果设置 end 为零值或小于当前时间的值，将不限制服务器生命周期
//   - 该函数支持运行时设置
func (opt *Options) WithLifeCycleEnd(end time.Time) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		now := time.Now()
		if end.Before(now) {
			opt.lifeCycleLimit = 0
			return
		}
		opt.lifeCycleLimit = end.Sub(now)
	})
}

func (opt *Options) GetLifeCycleLimit() time.Duration {
	return getOptionsValue(opt, func(opt *Options) time.Duration {
		return opt.lifeCycleLimit
	})
}

func (opt *Options) modifyOptionsValue(handler func(opt *Options)) *Options {
	opt.rw.Lock()
	handler(opt)
	opt.rw.Unlock()
	return opt
}

func (opt *Options) getManyOptions(handler func(opt *Options)) {
	opt.rw.RLock()
	defer opt.rw.RUnlock()
	handler(opt)
}

func getOptionsValue[V any](opt *Options, handler func(opt *Options) V) V {
	opt.rw.RLock()
	defer opt.rw.RUnlock()
	return handler(opt)
}
