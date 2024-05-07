package rpcbuiltin

import (
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/internal/registry"
	"github.com/nats-io/nats.go"
)

// NewRegistryWithNats 创建基于 Nats 注册中心的注册器
func NewRegistryWithNats(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsRegistryOptions) (rpc.Registry, error) {
	var opt = NewNatsRegistryOptions().Apply(opts...)
	return registry.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL, opt.KeepAliveInterval)
}
