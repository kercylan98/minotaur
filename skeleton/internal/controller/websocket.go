package controller

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
)

var _ application.Controller = (*WebSocketController)(nil)

func NewWebSocketController(actorSystem *vivid.ActorSystem) *WebSocketController {
	return &WebSocketController{
		actorSystem:   actorSystem,
		socketFactory: socket.NewFactory(actorSystem),
	}
}

type WebSocketController struct {
	modules struct {
		fiber module.FiberModule
	}
	actorSystem   *vivid.ActorSystem
	socketFactory socket.Factory
}

func (w *WebSocketController) OnInitialize(ctx *application.Context, loader *application.ServiceLoader) (err error) {
	w.modules.fiber = application.LoadModule[module.FiberModule](ctx)

	w.modules.fiber.Fiber().Get("/websocket", fiber.UseFiberHandler[*application.Context](websocket.New(func(conn *websocket.Conn) {
		socket.ProduceFiberSocketV2(w.socketFactory, conn, newWebSocketActor())
	})))

	return
}
