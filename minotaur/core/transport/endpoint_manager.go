package transport

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"sync"
)

func newEndpointManager(network *Network) *endpointManager {
	em := &endpointManager{
		network:   network,
		endpoints: make(map[string]*endpoint),
	}
	return em
}

type endpointManager struct {
	network    *Network
	endpoints  map[string]*endpoint
	endpointRw sync.RWMutex
}

func (em *endpointManager) getEndpoint(address core.Address) *endpoint {
	em.endpointRw.Lock()
	e, exist := em.endpoints[address.Address()]
	if !exist {
		e = newEndpoint(em.network, address)
		em.endpoints[address.Address()] = e
	}
	em.endpointRw.Unlock()
	return e
}
