package vivid

import "github.com/kercylan98/minotaur/rpc"

type ActorSystemOption func(*ActorSystemOptions)

type ActorSystemOptions struct {
	rpcSrv    rpc.Server    // RPC 服务
	discovery rpc.Discovery // 服务发现
}

func (o *ActorSystemOptions) apply(options ...ActorSystemOption) *ActorSystemOptions {
	for _, option := range options {
		option(o)
	}
	return o
}

func WithDiscovery(server rpc.Server, discovery rpc.Discovery) ActorSystemOption {
	return func(options *ActorSystemOptions) {
		options.rpcSrv = server
		options.discovery = discovery
	}
}
