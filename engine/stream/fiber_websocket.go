package stream

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/prc"
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
		wrappers := make(map[prc.LogicalAddress]*fiberWebSocketWrapper)
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *websocket.Conn:
				wrapper := &fiberWebSocketWrapper{system: system, Conn: m, closed: make(chan struct{})}
				wrapper.streamRef = ctx.ActorOfF(func() vivid.Actor {
					return NewStream(wrapper, configurator...)
				}, func(descriptor *vivid.ActorDescriptor) {
					descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
						return supervision.StopStrategy()
					}))
				})
				wrappers[wrapper.streamRef.LogicalAddress] = wrapper
				ctx.Reply(wrapper)
			case *vivid.OnTerminated:
				key := m.TerminatedActor.LogicalAddress
				wrapper, exist := wrappers[key]
				if exist {
					close(wrapper.closed)
					delete(wrappers, key)
				}
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
		message, err := system.FutureAsk(server, conn, -1).Result()
		if err != nil {
			return
		}
		wrapper, ok := message.(*fiberWebSocketWrapper)
		if !ok {
			return
		}

		var (
			mt  int
			msg []byte
		)
		for {
			if mt, msg, err = conn.ReadMessage(); err != nil {
				system.Tell(wrapper.streamRef, err)
				break
			}
			system.Tell(wrapper.streamRef, NewPacketDC(msg, mt))
		}
		<-wrapper.closed
	})

	// 返回 Upgrader 函数
	return func(ctx *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(ctx) {
			return fiber.ErrUpgradeRequired
		}

		return handler(ctx)
	}
}
