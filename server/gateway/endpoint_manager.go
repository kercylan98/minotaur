package gateway

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/random"
)

// NewEndpointManager 创建网关端点管理器
func NewEndpointManager() *EndpointManager {
	em := &EndpointManager{
		endpoints: concurrent.NewBalanceMap[string, []*Endpoint](),
		memory:    haxmap.New[string, *Endpoint](),
		selector: func(endpoints []*Endpoint) *Endpoint {
			return endpoints[random.Int(0, len(endpoints)-1)]
		},
	}
	return em
}

// EndpointManager 网关端点管理器
type EndpointManager struct {
	endpoints *concurrent.BalanceMap[string, []*Endpoint]
	memory    *haxmap.Map[string, *Endpoint]
	selector  func([]*Endpoint) *Endpoint
}

// GetEndpoint 获取端点
//   - name: 端点名称
//   - id: 使用端点的连接标识
func (slf *EndpointManager) GetEndpoint(name, id string) (*Endpoint, error) {
	endpoint, exist := slf.memory.Get(id)
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
		var available = make([]*Endpoint, 0, len(endpoints))
		for _, e := range endpoints {
			if !e.offline && e.state > 0 {
				available = append(available, e)
			}
		}
		if len(available) == 0 {
			return
		}
		endpoint = slf.selector(available)
	})
	if endpoint == nil {
		return nil, ErrEndpointNotExists
	}
	slf.memory.Set(id, endpoint)
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
