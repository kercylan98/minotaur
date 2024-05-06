package rpc

import (
	"github.com/kercylan98/minotaur/rpc/internal/registry"
	"github.com/nats-io/nats.go"
)

// Registry 表示一个 RPC 服务注册器，用于注册服务到注册中心
type Registry interface {
	// OnRegister 服务注册时触发
	OnRegister(service Service) error

	// OnUnregister 服务注销时触发
	OnUnregister() error

	// WatchCall 监听服务调用事件
	WatchCall() <-chan Caller
}

// NewRegistryWithNats 创建基于 Nats 注册中心的注册器
func NewRegistryWithNats(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsRegistryOptions) (Registry, error) {
	var opt = NewNatsRegistryOptions().Apply(opts...)
	return registry.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL, opt.KeepAliveInterval)
}
