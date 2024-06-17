package cluster

import "github.com/kercylan98/minotaur/toolkit/log"

type Option func(*Options)

type Options struct {
	ClusterName          string   // 集群名称
	Region               string   // 地域
	Zone                 string   // 可用区
	ShardId              uint16   // 分片ID
	Weight               uint64   // 权重
	BindAddr             string   // 绑定地址
	BindPort             uint16   // 绑定端口
	AdvertiseAddr        string   // 广播地址
	AdvertisePort        uint16   // 广播端口
	DefaultJoinAddresses []string // 默认加入地址
	Logger               *log.Logger
}

func WithLogger(logger *log.Logger) Option {
	return func(o *Options) {
		o.Logger = logger
	}
}

func WithDefaultJoinAddresses(defaultJoinAddresses ...string) Option {
	return func(o *Options) {
		o.DefaultJoinAddresses = defaultJoinAddresses
	}
}

func WithClusterName(clusterName string) Option {
	return func(o *Options) {
		o.ClusterName = clusterName
	}
}

func WithRegion(region string) Option {
	return func(o *Options) {
		o.Region = region
	}
}

func WithZone(zone string) Option {
	return func(o *Options) {
		o.Zone = zone
	}
}

func WithShardId(shardId uint16) Option {
	return func(o *Options) {
		o.ShardId = shardId
	}
}

func WithWeight(weight uint64) Option {
	return func(o *Options) {
		o.Weight = weight
	}
}

func WithBindAddr(bindAddr string) Option {
	return func(o *Options) {
		o.BindAddr = bindAddr
	}
}

func WithBindPort(bindPort uint16) Option {
	return func(o *Options) {
		o.BindPort = bindPort
	}
}

func WithAdvertiseAddr(advertiseAddr string) Option {
	return func(o *Options) {
		o.AdvertiseAddr = advertiseAddr
	}
}

func WithAdvertisePort(advertisePort uint16) Option {
	return func(o *Options) {
		o.AdvertisePort = advertisePort
	}
}

func (o *Options) apply(opts ...Option) *Options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}
