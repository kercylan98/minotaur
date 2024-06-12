package minotaur

import (
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/toolkit/log"
)

type Option func(*Options)

type Options struct {
	Logger            *log.Logger       // 日志记录器
	ActorSystemName   string            // Actor 系统名称
	EventBusActorName string            // 事件总线 Actor 名称
	Network           transport.Network // 网络

	LaunchedHooks []func(app *Application) // 启动钩子
}

// defaultApply 设置缺省值
func (o *Options) defaultApply() *Options {
	if o.ActorSystemName == "" {
		o.ActorSystemName = "minotaur"
	}
	if o.EventBusActorName == "" {
		o.EventBusActorName = "event_bus"
	}
	if o.Logger == nil {
		o.Logger = log.GetDefault()
	}
	return o
}

func (o *Options) apply(options ...Option) *Options {
	for _, option := range options {
		option(o)
	}
	return o.defaultApply()
}

func WithLaunchedHook(hooks ...func(app *Application)) Option {
	return func(o *Options) {
		o.LaunchedHooks = append(o.LaunchedHooks, hooks...)
	}
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

func WithActorSystemName(name string) Option {
	return func(o *Options) {
		o.ActorSystemName = name
	}
}

func WithEventBusActorName(name string) Option {
	return func(o *Options) {
		o.EventBusActorName = name
	}
}
