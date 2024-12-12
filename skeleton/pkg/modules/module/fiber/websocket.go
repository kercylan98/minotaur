package fiber

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/modules"
)

func NewWebSocket() modules.FiberWebSocket {
	return &webSocket{}
}

type webSocket struct {
	fiber modules.Fiber
}

func (w *webSocket) ImportDependencies(getter func(name string) application.Module) {
	w.fiber = getter("fiber")
}

func (w *webSocket) Name() string {
	return "websocket"
}

func (w *webSocket) Setup(ctx *application.Context) error {
	return nil
}

func (w *webSocket) OnReceive(ctx vivid.ActorContext) {

}
