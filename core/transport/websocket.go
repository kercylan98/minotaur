package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
)

func NewWebSocket(addr string, pattern ...string) vivid.Module {
	n := &userNetwork{}
	n.producer = func() vivid.Actor {
		return newGnetEngine(n.support, schemaWebSocket, addr, pattern...)
	}
	return n
}

var _ vivid.Module = &userNetwork{}

type userNetwork struct {
	support  *vivid.ModuleSupport
	producer vivid.ActorProducer
}

func (w *userNetwork) OnLoad(support *vivid.ModuleSupport) {
	w.support = support
	w.support.System().ActorOf(w.producer)
}
