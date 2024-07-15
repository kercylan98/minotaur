package transport

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/collection"
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
	closed     bool
}

func (em *endpointManager) getEndpoint(address core.Address) vivid.ActorRef {
	em.endpointRw.Lock()
	if em.closed {
		em.endpointRw.Unlock()
		return core.NewProcessRef(em.network.support.GetDeadLetter().GetAddress())
	}
	ref, exist := em.endpoints[address.PhysicalAddress()]
	if !exist {
		ref = em.network.support.System().ActorOf(func() vivid.Actor {
			return newEndpoint(em.network, address)
		}, func(options *vivid.ActorOptions) {
			options.WithNamePrefix("endpoint")
			options.WithName(address.PhysicalAddress())
			options.WithMailbox(func() vivid.Mailbox {
				return vivid.NewDefaultMailbox(128)
			})
		})
		em.endpoints[address.PhysicalAddress()] = ref
	}
	em.endpointRw.Unlock()
	return ref
}

func (em *endpointManager) delEndpoint(address core.Address) {
	em.endpointRw.Lock()
	delete(em.endpoints, address.PhysicalAddress())
	em.endpointRw.Unlock()
}

func (em *endpointManager) close() {
	em.endpointRw.Lock()
	em.closed = true
	endpoints := collection.CloneMap(em.endpoints)
	em.endpointRw.Unlock()

	for _, ref := range endpoints {
		em.network.support.System().Context().Terminate(ref)
	}
}
