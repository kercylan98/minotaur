package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/core/vivid/supervisor"
	"reflect"
)

func NewWebSocket(addr string, pattern ...string) *GNET {
	return newGNET(schemaWebSocket, addr, pattern...)
}

func NewTCP(addr string) *GNET {
	return newGNET(schemaTcp, addr)
}

func NewTCP4(addr string) *GNET {
	return newGNET(schemaTcp4, addr)
}

func NewTCP6(addr string) *GNET {
	return newGNET(schemaTcp6, addr)
}

func NewUDP(addr string) *GNET {
	return newGNET(schemaUdp, addr)
}

func NewUDP4(addr string) *GNET {
	return newGNET(schemaUdp4, addr)
}

func NewUDP6(addr string) *GNET {
	return newGNET(schemaUdp6, addr)
}

func NewUnix(addr string) *GNET {
	return newGNET(schemaUnix, addr)
}

func newGNET(schema, addr string, pattern ...string) *GNET {
	return &GNET{
		addr:    addr,
		schema:  schema,
		pattern: pattern,
	}
}

type GNET struct {
	support  *vivid.ModuleSupport
	addr     string
	schema   string
	services []GNETService
	pattern  []string
}

func (n *GNET) BindService(services ...GNETService) *GNET {
	n.services = append(n.services, services...)
	return n
}

func (n *GNET) OnLoad(support *vivid.ModuleSupport, hasTransportModule bool) {
	// field load
	n.support = support

	// init kit
	kit := &GNETKit{
		actorSystem: support.System(),
	}

	// init services
	for _, service := range n.services {
		service.OnInit(kit)
	}

	// init actor
	actorType := reflect.TypeOf((*gnetActor)(nil)).Elem().Name()
	kit.ownerRef = n.support.System().ActorOf(func() vivid.Actor {
		return newGNETActor(n, kit, n.schema, n.addr, n.pattern...)
	}, func(options *vivid.ActorOptions) {
		options.WithNamePrefix(actorType)
		options.WithSupervisorStrategy(supervisor.Stop())
	})
}
