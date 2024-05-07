package rpc

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/balancer"
	"sync"
)

// NewClient 创建一个新的 RPC 客户端
func NewClient(discovery Discovery) *Client {
	c := &Client{
		discovery: discovery,
		services:  make(map[string]balancer.Balancer[InstanceId, *ClientService]),
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	go c.run()
	return c
}

type Client struct {
	ctx       context.Context
	cancel    context.CancelFunc
	discovery Discovery                                                // 服务发现器
	services  map[string]balancer.Balancer[InstanceId, *ClientService] // serviceName -> instanceId -> service
	rw        sync.RWMutex
}

// GetService 根据负载均衡策略获取一个服务
func (c *Client) GetService(service ServiceName) (CallableService, error) {
	c.rw.RLock()
	instances, exist := c.services[service]
	if !exist {
		c.rw.RUnlock()
		return nil, ErrServiceNotFound
	}
	c.rw.RUnlock()

	instance, err := instances.Select()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (c *Client) UnaryCall(service ServiceName, route ...Route) UnaryCaller {
	return func(ctx context.Context, params any) (Reader, error) {
		target, err := c.GetService(service)
		if err != nil {
			return nil, err
		}
		return target.UnaryCall(route...)(ctx, params)
	}
}

func (c *Client) UnaryNotifyCall(service ServiceName, route ...Route) UnaryNotifyCaller {
	return func(ctx context.Context, params any) error {
		target, err := c.GetService(service)
		if err != nil {
			return err
		}
		return target.UnaryNotifyCall(route...)(ctx, params)
	}
}

func (c *Client) AsyncUnaryCall(service ServiceName, route ...Route) AsyncUnaryCaller {
	return func(ctx context.Context, params any, callback func(reader Reader, err error)) {
		target, err := c.GetService(service)
		if err != nil {
			callback(nil, err)
			return
		}
		target.AsyncUnaryCall(route...)(ctx, params, callback)
	}
}

func (c *Client) AsyncNotifyCall(service ServiceName, route ...Route) AsyncNotifyCaller {
	return func(params any) error {
		target, err := c.GetService(service)
		if err != nil {
			return err
		}
		return target.AsyncNotifyCall(route...)(params)
	}
}

// run 运行 RPC 客户端，开始监听服务
func (c *Client) run() {
	var registerCh = c.discovery.WatchRegister()
	var unregisterCh = c.discovery.WatchUnregister()
	for {
		select {
		case <-c.ctx.Done():
			return
		case service := <-registerCh:
			c.onRegister(service)
		case service := <-unregisterCh:
			c.onUnregister(service)
		}
	}
}

// Close 关闭 RPC 客户端
func (c *Client) Close() error {
	c.cancel()
	return c.discovery.Close()
}

func (c *Client) onRegister(service CallableService) {
	serviceInfo := service.GetServiceInfo()
	c.rw.Lock()
	instances, exist := c.services[serviceInfo.Name]
	if !exist {
		instances = balancer.NewRoundRobin[InstanceId, *ClientService]()
		c.services[serviceInfo.Name] = instances
	}
	c.rw.Unlock()
	instances.Add(&ClientService{service})
}

func (c *Client) onUnregister(service CallableService) {
	serviceInfo := service.GetServiceInfo()
	c.rw.Lock()
	instances, exist := c.services[serviceInfo.Name]
	if !exist {
		c.rw.Unlock()
		return
	}
	c.rw.Unlock()

	instances.Remove(&ClientService{service})
}
