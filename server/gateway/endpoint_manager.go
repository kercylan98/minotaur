package gateway

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/slice"
)

// NewEndpointManager 创建网关端点管理器
func NewEndpointManager() *EndpointManager {
	em := &EndpointManager{
		endpoints: concurrent.NewBalanceMap[string, []*Endpoint](),
		memory:    concurrent.NewBalanceMap[string, *Endpoint](),
	}
	return em
}

// EndpointManager 网关端点管理器
type EndpointManager struct {
	endpoints *concurrent.BalanceMap[string, []*Endpoint]
	memory    *concurrent.BalanceMap[string, *Endpoint]
}

// GetEndpoint 获取端点
func (slf *EndpointManager) GetEndpoint(name string, conn *server.Conn) (*Endpoint, error) {
	endpoint, exist := slf.memory.GetExist(conn.GetID())
	if exist {
		return endpoint, nil
	}
	slf.endpoints.Atom(func(m map[string][]*Endpoint) {
		endpoints, exist := m[name]
		if !exist {
			return
		}
		if len(endpoints) == 0 {
			return
		}
		// 随机获取
		endpoints = slice.Copy(endpoints)
		slice.Shuffle(endpoints)
		for _, e := range endpoints {
			if e.offline || e.state <= 0 {
				continue
			}
			endpoint = e
		}
	})
	if endpoint == nil {
		return nil, ErrEndpointNotExists
	}
	slf.memory.Set(conn.GetID(), endpoint)
	return endpoint, nil
}

// AddEndpoint 添加端点
func (slf *EndpointManager) AddEndpoint(endpoint *Endpoint) error {
	if endpoint.client.IsConnected() {
		return ErrCannotAddRunningEndpoint
	}
	for _, e := range slf.endpoints.Get(endpoint.name) {
		if e.address == endpoint.address {
			return ErrEndpointAlreadyExists
		}
	}
	go endpoint.Connect()
	slf.endpoints.Atom(func(m map[string][]*Endpoint) {
		m[endpoint.name] = append(m[endpoint.name], endpoint)
	})
	return nil
}

// RemoveEndpoint 移除端点
func (slf *EndpointManager) RemoveEndpoint(endpoint *Endpoint) error {
	slf.endpoints.Atom(func(m map[string][]*Endpoint) {
		var endpoints []*Endpoint
		endpoints, exist := m[endpoint.name]
		if !exist {
			return
		}
		for i, e := range endpoints {
			if e.address == endpoint.address {
				endpoints = append(endpoints[:i], endpoints[i+1:]...)
				break
			}
		}
		m[endpoint.name] = endpoints
	})
	return nil
}
