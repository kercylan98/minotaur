package minotaur

import "github.com/kercylan98/minotaur/minotaur/transport"

type Option func(*Options)

type Options struct {
	Network transport.Network // 网络
}

func (o *Options) apply(options ...Option) *Options {
	for _, option := range options {
		option(o)
	}
	return o
}

func WithNetwork(network transport.Network) Option {
	return func(o *Options) {
		o.Network = network
	}
}
