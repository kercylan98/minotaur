package minotaur

import (
	"github.com/kercylan98/minotaur/minotaur/transport"
)

type Option func(*Options)

type Options struct {
	ActorSystemName   string            // Actor 系统名称
	EventBusActorName string            // 事件总线 Actor 名称
	Network           transport.Network // 网络
}

// defaultApply 设置缺省值
func (o *Options) defaultApply() *Options {
	if o.ActorSystemName == "" {
		o.ActorSystemName = "minotaur"
	}
	if o.EventBusActorName == "" {
		o.EventBusActorName = "event_bus"
	}
	return o
}

func (o *Options) apply(options ...Option) *Options {
	for _, option := range options {
		option(o)
	}
	return o.defaultApply()
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
