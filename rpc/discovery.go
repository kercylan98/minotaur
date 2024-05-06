package rpc

import (
	"github.com/kercylan98/minotaur/rpc/internal/discovery"
	"github.com/nats-io/nats.go"
)

// Discovery 表示一个 RPC 服务发现器
type Discovery interface {
	// WatchRegister 监听服务注册事件
	WatchRegister() <-chan CallableService

	// WatchUnregister 监听服务注销事件
	WatchUnregister() <-chan CallableService

	// Close 关闭服务发现器
	Close() error
}

// NewDiscoveryWithNats 创建基于 Nats 的服务发现器
func NewDiscoveryWithNats(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsDiscoveryOptions) (Discovery, error) {
	var opt = NewNatsDiscoveryOptions().Apply(opts...)
	return discovery.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL)
}
