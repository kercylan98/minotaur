package rpcbuiltin

import (
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/internal/discovery"
	"github.com/nats-io/nats.go"
)

// NewDiscoveryWithNats 创建基于 Nats 的服务发现器
func NewDiscoveryWithNats(conn *nats.Conn, js nats.JetStreamContext, opts ...*NatsDiscoveryOptions) (rpc.Discovery, error) {
	var opt = NewNatsDiscoveryOptions().Apply(opts...)
	return discovery.NewNats(conn, js, opt.BucketName, opt.BucketDesc, opt.KeyPrefix, opt.TTL)
}
