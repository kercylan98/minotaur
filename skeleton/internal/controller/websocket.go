package controller

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
)

var _ application.Controller = (*WebSocketController)(nil)

func NewWebSocketController() *WebSocketController {
	return &WebSocketController{}
}

type WebSocketController struct {
	modules struct {
		actorSystem module.ActorSystemModule
		fiber       module.FiberModule
	}

	socketFactory socket.Factory
}

func (w *WebSocketController) OnInitialize(ctx *application.Context, loader *application.ServiceLoader) (err error) {
	w.modules.actorSystem = application.LoadModule[module.ActorSystemModule](ctx)
	w.modules.fiber = application.LoadModule[module.FiberModule](ctx)

	w.socketFactory = socket.NewFactory(w.modules.actorSystem.ActorSystem())

	w.modules.fiber.Fiber().Get("/websocket", fiber.UseFiberHandler[*application.Context](websocket.New(func(conn *websocket.Conn) {
		socket.ProduceFiberSocketV2(w.socketFactory, conn, newWebSocketActor())
	})))

	return
}
