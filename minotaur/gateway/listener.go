package gateway

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/balancer"
)

type ListenerActorBindAddressMessage struct {
	Address Address
}

type ListenerActorBindBalancerMessage struct {
	Balancer balancer.Balancer[EndpointId, *Endpoint]
}

type ListenerActorBindEndpointMessage struct {
	Endpoint *Endpoint
}

type Listener interface {
	// Start 启动监听器开始监听端口活动
	Start(ctx context.Context) error

	// Stop 停止监听器
	Stop() error

	// SetAddress 设置监听器的地址
	SetAddress(addr Address)

	// Address 获取监听器的地址
	Address() Address

	// AddEndpoint 添加一个端点
	AddEndpoint(endpoint *Endpoint)

	// RemoveEndpoint 移除一个端点
	RemoveEndpoint(id EndpointId)
}

type listenerBindEvent struct {
	listener Listener
	callback []func(err error)
}
