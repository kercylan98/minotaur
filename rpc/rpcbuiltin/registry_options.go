package rpcbuiltin

import (
	"github.com/kercylan98/minotaur/rpc"
	"time"
)

type NatsRegistryOptions struct {
	BucketName        string        // 用于存储服务信息的 Bucket 名称，默认为 NatsDefaultBucketName
	BucketDesc        string        // Bucket 描述信息，默认为 NatsDefaultBucketDesc
	KeyPrefix         string        // 用于存储服务信息的 Key 前缀，默认为 NatsDefaultBkvBucketPrefix
	TTL               time.Duration // 首次初始化时的服务注册信息存储桶的过期时间，默认为 NatsDefaultTTL
	KeepAliveInterval time.Duration // 服务注册信息的保活间隔，默认为 NatsDefaultKeepAlive
}

func NewNatsRegistryOptions() *NatsRegistryOptions {
	return &NatsRegistryOptions{
		BucketName:        rpc.NatsDefaultBucketName,
		BucketDesc:        rpc.NatsDefaultBucketDesc,
		KeyPrefix:         rpc.NatsDefaultBkvBucketPrefix,
		TTL:               rpc.NatsDefaultTTL,
		KeepAliveInterval: rpc.NatsDefaultKeepAlive,
	}
}

func (o *NatsRegistryOptions) WithBucketName(bucketName string) *NatsRegistryOptions {
	o.BucketName = bucketName
	return o
}

func (o *NatsRegistryOptions) WithBucketDesc(bucketDesc string) *NatsRegistryOptions {
	o.BucketDesc = bucketDesc
	return o
}

func (o *NatsRegistryOptions) WithKeyPrefix(keyPrefix string) *NatsRegistryOptions {
	o.KeyPrefix = keyPrefix
	return o
}

func (o *NatsRegistryOptions) WithTTL(ttl time.Duration) *NatsRegistryOptions {
	o.TTL = ttl
	return o
}

func (o *NatsRegistryOptions) WithKeepAliveInterval(interval time.Duration) *NatsRegistryOptions {
	o.KeepAliveInterval = interval
	return o
}

func (o *NatsRegistryOptions) Apply(opts ...*NatsRegistryOptions) *NatsRegistryOptions {
	for _, opt := range opts {
		o.BucketName = opt.BucketName
		o.BucketDesc = opt.BucketDesc
		o.KeyPrefix = opt.KeyPrefix
		o.TTL = opt.TTL
		o.KeepAliveInterval = opt.KeepAliveInterval
	}
	return o
}
