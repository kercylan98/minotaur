package rpcbuiltin

import (
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/discovery"
	"github.com/nats-io/nats.go"
)

// NewDiscoveryWithNatsE 创建基于 Nats 的服务发现器
func NewDiscoveryWithNatsE(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsDiscoveryOptions) (rpc.Discovery, error) {
	var opt = NewNatsDiscoveryOptions().Apply(opts...)
	return discovery.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL)
}

// NewDiscoveryWithNats 创建基于 Nats 的服务注册器，当创建失败时会 panic
func NewDiscoveryWithNats(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsDiscoveryOptions) rpc.Discovery {
	var opt = NewNatsDiscoveryOptions().Apply(opts...)
	d, err := discovery.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL)
	if err != nil {
		panic(err)
	}
	return d
}
