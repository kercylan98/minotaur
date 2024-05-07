package rpcbuiltin

import (
	"github.com/kercylan98/minotaur/rpc"
	"time"
)

type NatsDiscoveryOptions struct {
	BucketName string        // 用于存储服务信息的 Bucket 名称，默认为 NatsDefaultBucketName
	BucketDesc string        // Bucket 描述信息，默认为 NatsDefaultBucketDesc
	KeyPrefix  string        // 用于存储服务信息的 Key 前缀，默认为 NatsDefaultBkvBucketPrefix
	TTL        time.Duration // 首次初始化时的服务注册信息存储桶的过期时间，默认为 NatsDefaultTTL
}

func NewNatsDiscoveryOptions() *NatsDiscoveryOptions {
	return &NatsDiscoveryOptions{
		BucketName: rpc.NatsDefaultBucketName,
		BucketDesc: rpc.NatsDefaultBucketDesc,
		KeyPrefix:  rpc.NatsDefaultBkvBucketPrefix,
		TTL:        rpc.NatsDefaultTTL,
	}
}

func (o *NatsDiscoveryOptions) WithBucketName(bucketName string) *NatsDiscoveryOptions {
	o.BucketName = bucketName
	return o
}

func (o *NatsDiscoveryOptions) WithBucketDesc(bucketDesc string) *NatsDiscoveryOptions {
	o.BucketDesc = bucketDesc
	return o
}

func (o *NatsDiscoveryOptions) WithKeyPrefix(keyPrefix string) *NatsDiscoveryOptions {
	o.KeyPrefix = keyPrefix
	return o
}

func (o *NatsDiscoveryOptions) WithTTL(ttl int) *NatsDiscoveryOptions {
	o.TTL = time.Duration(ttl)
	return o
}

func (o *NatsDiscoveryOptions) Apply(opts ...*NatsDiscoveryOptions) *NatsDiscoveryOptions {
	for _, opt := range opts {
		o.BucketName = opt.BucketName
		o.BucketDesc = opt.BucketDesc
		o.KeyPrefix = opt.KeyPrefix
		o.TTL = opt.TTL
	}
	return o
}
