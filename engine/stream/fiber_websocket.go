package stream

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
)

// NewFiberWebSocketHandler 创建一个 fiber.Handler，它将会对请求进行 WebSocket 升级，并采用 Actor 对其进行维护。
//   - 这看来似乎是必须的，如果你要对消息进行处理，那么请在 Configurator 内指定 Configuration.WithPerformance 可选项，这可以让你得到 Stream 的 vivid.ActorContext，并对其进行控制。
//
// 在 Stream 的 vivid.ActorContext 中会传入三个特殊消息，你可以选择使用它们：
//   - stream.Writer：在 Stream Actor 启动后，将会收到一个写入器，这个写入器接收 *stream.Packet 类型的消息，它将会将消息写入到连接中。
//   - *websocket.Conn：这是 WebSocket 的特殊消息，你可以获取它来进行额外的操作。
//   - *stream.Packet：当收到该消息，也就意味着存在需要处理的数据包。
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
				ctx.Tell(ref, m)
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
