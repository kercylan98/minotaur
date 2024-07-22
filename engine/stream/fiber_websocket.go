package stream

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
)

func NewFiberWebSocketHandler(app *fiber.App, system *vivid.ActorSystem, configurator ...Configurator) fiber.Handler {
	// 创建服务器 Actor
	server := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *websocket.Conn:
				wrapper := &fiberWebSocketWrapper{m}
				ref := ctx.ActorOfF(func() vivid.Actor {
					return NewStream(wrapper, configurator...)
				}, func(descriptor *vivid.ActorDescriptor) {
					descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
						return supervision.StopStrategy()
					}))
				})
				ctx.Reply(ref)
			}
		})
	})

	// 注册 fiber 停机 hook
	app.Hooks().OnShutdown(func() error {
		system.Terminate(server, true)
		return nil
	})

	// 定义连接 Reader Actor
	handler := websocket.New(func(conn *websocket.Conn) {
		message, err := system.FutureAsk(server, conn).Result()
		if err != nil {
			return
		}
		stream, ok := message.(vivid.ActorRef)
		if !ok {
			return
		}
		defer system.Terminate(stream, true)

		var (
			mt  int
			msg []byte
		)
		for {
			if mt, msg, err = conn.ReadMessage(); err != nil {
				system.Tell(stream, err)
				break
			}
			system.Tell(stream, NewPacketDC(msg, mt))
		}
	})

	// 返回 Upgrader 函数
	return func(ctx *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(ctx) {
			return fiber.ErrUpgradeRequired
		}

		return handler(ctx)
	}
}
