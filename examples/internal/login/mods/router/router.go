package router

import (
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net/http"
)

type Router interface {
	vivid.Mod

	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type Mod struct {
	Handler *network.HttpServe
}

func (m *Mod) OnLifeCycle(ctx vivid.ActorContext, lifeCycle vivid.ModLifeCycle) {
	switch lifeCycle {
	default:
	}
}

func (m *Mod) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.Handler.HandleFunc(pattern, handler)
}
