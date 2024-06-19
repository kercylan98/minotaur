package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/cluster"
	"github.com/kercylan98/minotaur/toolkit/log"
)

type ActorSystemOption func(options *ActorSystemOptions)

type ActorSystemOptions struct {
	options        []ActorSystemOption
	ClusterOptions []cluster.Option
	Logger         *log.Logger
}

// WithLogger 设置日志记录器
func (o *ActorSystemOptions) WithLogger(logger *log.Logger) *ActorSystemOptions {
	o.options = append(o.options, func(opts *ActorSystemOptions) {
		opts.Logger = logger
	})
	return o
}

// WithCluster 设置集群配置
func (o *ActorSystemOptions) WithCluster(options ...cluster.Option) *ActorSystemOptions {
	o.options = append(o.options, func(opts *ActorSystemOptions) {
		opts.ClusterOptions = append(opts.ClusterOptions, options...)
	})
	return o
}

func NewActorSystemOptions() *ActorSystemOptions {
	return &ActorSystemOptions{}
}

func (o *ActorSystemOptions) apply(options ...*ActorSystemOptions) *ActorSystemOptions {
	for _, opt := range options {
		for _, option := range opt.options {
			option(o)
		}
	}
	return o
}
