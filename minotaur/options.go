package minotaur

import (
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"time"
)

// NetworkShutdownMode 网络关闭模式
type NetworkShutdownMode uint8

const (
	// NetworkShutdownModeForce 强制关闭，直接关闭网络
	NetworkShutdownModeForce NetworkShutdownMode = iota

	// NetworkShutdownModeGraceful 优雅关闭，将会等待所有所有连接断开且持续一段时间没有新连接后关闭网络。
	NetworkShutdownModeGraceful
)

type Option func(*Options)

type Options struct {
	Logger                                *log.Logger                 // 日志记录器
	ActorSystemName                       string                      // Actor 系统名称
	EventBusActorName                     string                      // 事件总线 Actor 名称
	Network                               transport.Network           // 网络
	ActorSystemOptions                    []*vivid.ActorSystemOptions // Actor 系统配置
	NetworkShutdownMode                   NetworkShutdownMode         // 网络关闭模式
	NetworkGracefulShutdownTimeout        time.Duration               // 优雅关闭超时时间
	NetworkGracefulShutdownWaitTime       time.Duration               // 优雅关闭额外等待时间
	NetworkGracefulShutdownPromptHandler  func(attrs ...log.Attr)     // 优雅关闭提示处理器
	NetworkGracefulShutdownPromptInterval time.Duration               // 优雅关闭提示间隔
}

func defaultOptions() *Options {
	return &Options{
		Logger:                                log.GetDefault(),
		ActorSystemName:                       "app",
		EventBusActorName:                     "event_bus",
		NetworkShutdownMode:                   NetworkShutdownModeForce,
		NetworkGracefulShutdownWaitTime:       time.Second * 30,
		NetworkGracefulShutdownPromptInterval: time.Second,
	}
}

func (o *Options) apply(options ...Option) *Options {
	for _, option := range options {
		option(o)
	}
	return o
}

func WithLogger(logger *log.Logger) Option {
	return func(o *Options) {
		o.Logger = logger
	}
}

func WithNetwork(network transport.Network) Option {
	return func(o *Options) {
		o.Network = network
	}
}

func WithNetworkShutdownMode(mode NetworkShutdownMode) Option {
	return func(o *Options) {
		o.NetworkShutdownMode = mode
	}
}

func WithNetworkGracefulShutdownTimeout(timeout time.Duration) Option {
	if timeout < 0 {
		timeout = 0
	}
	return func(o *Options) {
		o.NetworkGracefulShutdownTimeout = timeout
	}
}

func WithNetworkGracefulShutdownWaitTime(wait time.Duration) Option {
	if wait < 0 {
		wait = 0
	}
	return func(o *Options) {
		o.NetworkGracefulShutdownWaitTime = wait
	}
}

func WithNetworkGracefulShutdownPromptHandler(interval time.Duration, handler func(attrs ...log.Attr)) Option {
	if interval < 0 {
		interval = time.Second
	}
	return func(o *Options) {
		o.NetworkGracefulShutdownPromptHandler = handler
		o.NetworkGracefulShutdownPromptInterval = interval
	}
}

func WithActorSystemName(name string) Option {
	return func(o *Options) {
		o.ActorSystemName = name
	}
}

func WithActorSystemOptions(options ...*vivid.ActorSystemOptions) Option {
	return func(o *Options) {
		o.ActorSystemOptions = append(o.ActorSystemOptions, options...)
	}
}

func WithEventBusActorName(name string) Option {
	return func(o *Options) {
		o.EventBusActorName = name
	}
}
