package transport

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/core/vivid"
)

type FiberKit struct {
	fiberActorRef vivid.ActorRef
	fiberApp      *fiber.App
	actorSystem   *vivid.ActorSystem
}

func (k *FiberKit) Fiber() *fiber.App {
	return k.fiberApp
}

func (k *FiberKit) System() *vivid.ActorSystem {
	return k.actorSystem
}

func (k *FiberKit) WebSocket(path string, rulePath ...string) *FiberWebSocket {
	fws := &FiberWebSocket{kit: k}
	fws.init()
	k.fiberApp.Use(path, func(c *fiber.Ctx) (err error) {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}
		if err = fws.upgradeHook(k, c); err != nil {
			return err
		}
		return c.Next()
	})

	var rp = path
	if len(rulePath) > 0 {
		rp = rulePath[0]
	}

	k.fiberApp.Get(rp, websocket.New(func(c *websocket.Conn) {
		var err error
		var fiberCtx = &FiberContext{conn: c}
		var conn *Conn
		var rootActor = k.System().Context()
		var result vivid.Message

		result, err = rootActor.FutureAsk(k.fiberActorRef, actorOfMessage{
			ActorProducer: func() vivid.Actor {
				return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
					switch m := ctx.Message().(type) {
					case vivid.OnLaunch:
						conn = NewConn(&fiberConnWrapper{c}, ctx, ctx.Ref())
						if err = fws.connectionOpenedHook(k, fiberCtx, conn); err != nil {
							ctx.Terminate(ctx.Ref())
						}
					case Packet:
						if err = c.WriteMessage(m.GetContext().(int), m.GetBytes()); err != nil {
							ctx.Terminate(ctx.Ref())
						}
					case vivid.OnTerminate:
						_ = c.Close()
						fws.connectionClosedHook(k, fiberCtx, conn, err)
					}
				})
			},
			ActorOptionDefiner: func(options *vivid.ActorOptions) {
				options.WithName("conn-" + c.RemoteAddr().String())
			},
		}).Result()
		if err != nil {
			return
		}
		ref := result.(vivid.ActorRef)

		defer func() {
			rootActor.Terminate(ref)
		}()

		var (
			mt  int
			msg []byte
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				break
			}
			if err = fws.connectionPacketHook(k, fiberCtx, conn, NewPacket(msg).SetContext(mt)); err != nil {
				break
			}
		}

	}))

	return fws
}
