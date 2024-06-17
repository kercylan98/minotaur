package vivid

import "github.com/kercylan98/minotaur/minotaur/cluster"

type ActorSystemOption func(options *ActorSystemOptions)

type ActorSystemOptions struct {
	options        []ActorSystemOption
	ClusterOptions []cluster.Option
}

// WithCluster 用于设置集群配置
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
