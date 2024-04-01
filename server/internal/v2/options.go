package server

import (
	"sync"
	"time"
)

func NewOptions() *Options {
	return DefaultOptions()
}

func DefaultOptions() *Options {
	return &Options{
		ServerMessageChannelSize:       1024 * 4,
		ActorMessageChannelSize:        1024,
		ServerMessageBufferInitialSize: 1024,
		ActorMessageBufferInitialSize:  1024,
		MessageErrorHandler:            nil,
		LifeCycleLimit:                 0,
	}
}

type Options struct {
	server                         *server
	rw                             sync.RWMutex
	ServerMessageChannelSize       int                                          // 服务器 Actor 消息处理管道大小
	ActorMessageChannelSize        int                                          // Actor 消息处理管道大小
	ServerMessageBufferInitialSize int                                          // 服务器 Actor 消息写入缓冲区初始化大小
	ActorMessageBufferInitialSize  int                                          // Actor 消息写入缓冲区初始化大小
	MessageErrorHandler            func(srv Server, message Message, err error) // 消息错误处理器
	LifeCycleLimit                 time.Duration                                // 服务器生命周期上限，在服务器启动后达到生命周期上限将关闭服务器
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

		opt.ServerMessageChannelSize = option.ServerMessageChannelSize
		opt.ActorMessageChannelSize = option.ActorMessageChannelSize
		opt.ServerMessageBufferInitialSize = option.ServerMessageBufferInitialSize
		opt.ActorMessageBufferInitialSize = option.ActorMessageBufferInitialSize
		opt.MessageErrorHandler = option.MessageErrorHandler
		opt.LifeCycleLimit = option.LifeCycleLimit

		option.rw.RUnlock()
	}
	if opt.server != nil && !opt.server.state.LaunchedAt.IsZero() {
		opt.active()
	}
}

func (opt *Options) active() {
	opt.server.notify.lifeCycleTime <- opt.getLifeCycleLimit()
}

// WithServerMessageChannelSize 设置服务器 Actor 用于处理消息的管道大小，当管道由于逻辑阻塞而导致满载时，会导致新消息无法及时从缓冲区拿出，从而增加内存的消耗，但是并不会影响消息的写入
func (opt *Options) WithServerMessageChannelSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.ServerMessageChannelSize = size
	})
}

func (opt *Options) getServerMessageChannelSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.ServerMessageChannelSize
	})
}

// WithActorMessageChannelSize 设置 Actor 用于处理消息的管道大小，当管道由于逻辑阻塞而导致满载时，会导致新消息无法及时从缓冲区拿出，从而增加内存的消耗，但是并不会影响消息的写入
func (opt *Options) WithActorMessageChannelSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.ActorMessageChannelSize = size
	})
}

func (opt *Options) getActorMessageChannelSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.ActorMessageChannelSize
	})
}

// WithServerMessageBufferInitialSize 设置服务器 Actor 消息环形缓冲区 buffer.Ring 的初始大小，适当的值可以避免频繁扩容
//   - 由于扩容是按照当前大小的 2 倍进行扩容，过大的值也可能会导致内存消耗过高
func (opt *Options) WithServerMessageBufferInitialSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.ServerMessageBufferInitialSize = size
	})
}

func (opt *Options) getServerMessageBufferInitialSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.ServerMessageBufferInitialSize
	})
}

// WithActorMessageBufferInitialSize 设置 Actor 消息环形缓冲区 buffer.Ring 的初始大小，适当的值可以避免频繁扩容
//   - 由于扩容是按照当前大小的 2 倍进行扩容，过大的值也可能会导致内存消耗过高
func (opt *Options) WithActorMessageBufferInitialSize(size int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.ActorMessageBufferInitialSize = size
	})
}

func (opt *Options) getActorMessageBufferInitialSize() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return opt.ActorMessageBufferInitialSize
	})
}

// WithMessageErrorHandler 设置消息错误处理器，当消息处理出现错误时，会调用该处理器进行处理
//   - 如果在运行时设置，后续消息错误将会使用新的 handler 进行处理
func (opt *Options) WithMessageErrorHandler(handler func(srv Server, message Message, err error)) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.MessageErrorHandler = handler
	})
}

func (opt *Options) getMessageErrorHandler() func(srv Server, message Message, err error) {
	return getOptionsValue(opt, func(opt *Options) func(srv Server, message Message, err error) {
		return opt.MessageErrorHandler
	})
}

// WithLifeCycleLimit 设置服务器生命周期上限，在服务器启动后达到生命周期上限将关闭服务器
//   - 如果设置为 <= 0 的值，将不限制服务器生命周期
//   - 该函数支持运行时设置
func (opt *Options) WithLifeCycleLimit(limit time.Duration) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.LifeCycleLimit = limit
	})
}

// WithLifeCycleEnd 设置服务器生命周期终点，在服务器达到该时间后将关闭服务器
//   - 如果设置 end 为零值或小于当前时间的值，将不限制服务器生命周期
//   - 该函数支持运行时设置
func (opt *Options) WithLifeCycleEnd(end time.Time) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		now := time.Now()
		if end.Before(now) {
			opt.LifeCycleLimit = 0
			return
		}
		opt.LifeCycleLimit = end.Sub(now)
	})
}

// getLifeCycleLimit 获取服务器生命周期上限
func (opt *Options) getLifeCycleLimit() time.Duration {
	return getOptionsValue(opt, func(opt *Options) time.Duration {
		return opt.LifeCycleLimit
	})
}

func (opt *Options) modifyOptionsValue(handler func(opt *Options)) *Options {
	opt.rw.Lock()
	handler(opt)
	opt.rw.Unlock()
	return opt
}

func getOptionsValue[V any](opt *Options, handler func(opt *Options) V) V {
	opt.rw.RLock()
	defer opt.rw.RUnlock()
	return handler(opt)
}
