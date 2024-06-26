package transport

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/core/vivid"
	"sync"
)

func newEndpointManager(network *Network) *endpointManager {
	em := &endpointManager{
		network:   network,
		endpoints: make(map[string]vivid.ActorRef),
	}
	return em
}

type endpointManager struct {
	network    *Network
	endpoints  map[string]vivid.ActorRef
	endpointRw sync.RWMutex
}

func (em *endpointManager) getEndpoint(address core.Address) vivid.ActorRef {
	em.endpointRw.Lock()
	ref, exist := em.endpoints[address.Address()]
	if !exist {
		ref = em.network.support.System().ActorOf(func() vivid.Actor {
			return newEndpoint(em.network, address)
		}, func(options *vivid.ActorOptions) {
			options.WithName("endpoint/" + address.Address())
			options.WithMailbox(func() vivid.Mailbox {
				return vivid.NewDefaultMailbox(128)
			})
		})
		em.endpoints[address.Address()] = ref
	}
	em.endpointRw.Unlock()
	return ref
}

func (em *endpointManager) delEndpoint(address core.Address) {
	em.endpointRw.Lock()
	delete(em.endpoints, address.Address())
	em.endpointRw.Unlock()
}
