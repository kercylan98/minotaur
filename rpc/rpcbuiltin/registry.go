package rpcbuiltin

import (
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/registry"
	"github.com/nats-io/nats.go"
)

// NewRegistryWithNatsE 创建基于 Nats 注册中心的注册器
func NewRegistryWithNatsE(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsRegistryOptions) (rpc.Registry, error) {
	var opt = NewNatsRegistryOptions().Apply(opts...)
	return registry.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL, opt.KeepAliveInterval)
}

// NewRegistryWithNats 创建基于 Nats 注册中心的注册器，如果出现错误则 panic
func NewRegistryWithNats(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsRegistryOptions) rpc.Registry {
	var opt = NewNatsRegistryOptions().Apply(opts...)
	r, err := registry.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL, opt.KeepAliveInterval)
	if err != nil {
		panic(err)
	}
	return r
}
