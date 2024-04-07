package server

import (
	"github.com/kercylan98/minotaur/utils/log/v2"
	"os"
	"sync"
	"time"
)

func NewOptions() *Options {
	return DefaultOptions()
}

func DefaultOptions() *Options {
	return &Options{
		serverMessageChannelSize:       1024 * 4,
		actorMessageChannelSize:        1024,
		serverMessageBufferInitialSize: 1024,
		actorMessageBufferInitialSize:  1024,
		messageErrorHandler:            nil,
		lifeCycleLimit:                 0,
		logger:                         log.NewLogger(log.NewHandler(os.Stdout, log.DefaultOptions().WithCallerSkip(-1).WithLevel(log.LevelInfo))),
	}
}

type Options struct {
	server                         *server
	rw                             sync.RWMutex
	serverMessageChannelSize       int                                          // 服务器 Actor 消息处理管道大小
	actorMessageChannelSize        int                                          // Actor 消息处理管道大小
	serverMessageBufferInitialSize int                                          // 服务器 Actor 消息写入缓冲区初始化大小
	actorMessageBufferInitialSize  int                                          // Actor 消息写入缓冲区初始化大小
	messageErrorHandler            func(srv Server, message Message, err error) // 消息错误处理器
	lifeCycleLimit                 time.Duration                                // 服务器生命周期上限，在服务器启动后达到生命周期上限将关闭服务器
	logger                         *log.Logger                                  // 日志记录器
	debug                          bool                                         // Debug 模式
	syncLowMessageDuration         time.Duration                                // 同步慢消息时间
	asyncLowMessageDuration        time.Duration                                // 异步慢消息时间
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

		opt.serverMessageChannelSize = option.serverMessageChannelSize
		opt.actorMessageChannelSize = option.actorMessageChannelSize
		opt.serverMessageBufferInitialSize = option.serverMessageBufferInitialSize
		opt.actorMessageBufferInitialSize = option.actorMessageBufferInitialSize
		opt.messageErrorHandler = option.messageErrorHandler
		opt.lifeCycleLimit = option.lifeCycleLimit
		opt.logger = option.logger
		opt.debug = option.debug
		opt.syncLowMessageDuration = option.syncLowMessageDuration
		opt.asyncLowMessageDuration = option.asyncLowMessageDuration

		option.rw.RUnlock()
	}
	if opt.server != nil && !opt.server.state.LaunchedAt.IsZero() {
		opt.active()
		if opt.server.reactor != nil {
			opt.server.reactor.SetLogger(opt.logger)
		}
	}
}

func (opt *Options) active() {
	opt.server.notify.lifeCycleTime <- opt.GetLifeCycleLimit()
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

// IsDebug 获取当前服务器是否是 Debug 模式
func (opt *Options) IsDebug() bool {
	return getOptionsValue(opt, func(opt *Options) bool {
		return opt.debug
	})
}

// WithLogger 设置服务器的日志记录器
//   - 该函数支持运行时设置
func (opt *Options) WithLogger(logger *log.Logger) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.logger = logger
		if opt.server != nil && opt.server.reactor != nil {
			opt.server.reactor.SetLogger(opt.logger)
		}
	})
}

// GetLogger 获取服务器的日志记录器
func (opt *Options) GetLogger() *log.Logger {
	return getOptionsValue(opt, func(opt *Options) *log.Logger {
		return opt.logger
	})
}

// WithServerMessageChannelSize 设置服务器 Actor 用于处理消息的管道大小，当管道由于逻辑阻塞而导致满载时，会导致新消息无法及时从缓冲区拿出，从而增加内存的消耗，但是并不会影响消息的写入
func (opt *Options) WithServerMessageChannelSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.serverMessageChannelSize = size
	})
}

func (opt *Options) GetServerMessageChannelSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.serverMessageChannelSize
	})
}

// WithActorMessageChannelSize 设置 Actor 用于处理消息的管道大小，当管道由于逻辑阻塞而导致满载时，会导致新消息无法及时从缓冲区拿出，从而增加内存的消耗，但是并不会影响消息的写入
func (opt *Options) WithActorMessageChannelSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.actorMessageChannelSize = size
	})
}

func (opt *Options) GetActorMessageChannelSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.actorMessageChannelSize
	})
}

// WithServerMessageBufferInitialSize 设置服务器 Actor 消息环形缓冲区 buffer.Ring 的初始大小，适当的值可以避免频繁扩容
//   - 由于扩容是按照当前大小的 2 倍进行扩容，过大的值也可能会导致内存消耗过高
func (opt *Options) WithServerMessageBufferInitialSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.serverMessageBufferInitialSize = size
	})
}

func (opt *Options) GetServerMessageBufferInitialSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.serverMessageBufferInitialSize
	})
}

// WithActorMessageBufferInitialSize 设置 Actor 消息环形缓冲区 buffer.Ring 的初始大小，适当的值可以避免频繁扩容
//   - 由于扩容是按照当前大小的 2 倍进行扩容，过大的值也可能会导致内存消耗过高
func (opt *Options) WithActorMessageBufferInitialSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.actorMessageBufferInitialSize = size
	})
}

func (opt *Options) GetActorMessageBufferInitialSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.actorMessageBufferInitialSize
	})
}

// WithMessageErrorHandler 设置消息错误处理器，当消息处理出现错误时，会调用该处理器进行处理
//   - 如果在运行时设置，后续消息错误将会使用新的 handler 进行处理
func (opt *Options) WithMessageErrorHandler(handler func(srv Server, message Message, err error)) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.messageErrorHandler = handler
	})
}

func (opt *Options) GetMessageErrorHandler() func(srv Server, message Message, err error) {
	return getOptionsValue(opt, func(opt *Options) func(srv Server, message Message, err error) {
		return opt.messageErrorHandler
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
