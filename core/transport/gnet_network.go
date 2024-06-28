package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
)

func NewWebSocket(addr string, pattern ...string) vivid.Module {
	return newGnetNetwork(schemaWebSocket, addr, pattern...)
}

func NewTCP(addr string) vivid.Module {
	return newGnetNetwork(schemaTcp, addr)
}

func NewTCP4(addr string) vivid.Module {
	return newGnetNetwork(schemaTcp4, addr)
}

func NewTCP6(addr string) vivid.Module {
	return newGnetNetwork(schemaTcp6, addr)
}

func NewUDP(addr string) vivid.Module {
	return newGnetNetwork(schemaUdp, addr)
}

func NewUDP4(addr string) vivid.Module {
	return newGnetNetwork(schemaUdp4, addr)
}

func NewUDP6(addr string) vivid.Module {
	return newGnetNetwork(schemaUdp6, addr)
}

func NewUnix(addr string) vivid.Module {
	return newGnetNetwork(schemaUnix, addr)
}

func newGnetNetwork(schema, addr string, pattern ...string) vivid.Module {
	n := &gnetNetwork{}
	n.producer = func() vivid.Actor {
		return newGnetEngine(n.support, schema, addr, pattern...)
	}
	return n
}

var _ vivid.Module = &gnetNetwork{}

type gnetNetwork struct {
	support  *vivid.ModuleSupport
	producer vivid.ActorProducer
}

func (w *gnetNetwork) OnLoad(support *vivid.ModuleSupport) {
	w.support = support
	w.support.System().ActorOf(w.producer, func(options *vivid.ActorOptions) {
		options.WithName("un")
	})
}
