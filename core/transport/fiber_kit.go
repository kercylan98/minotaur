package transport

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/core/vivid"
	"sync/atomic"
)

type FiberKit struct {
	ownerRef    vivid.ActorRef
	app         *fiber.App
	actorSystem *vivid.ActorSystem
	fws         *FiberWebSocket
}

func (k *FiberKit) App() *fiber.App {
	return k.app
}

func (k *FiberKit) System() *vivid.ActorSystem {
	return k.actorSystem
}

func (k *FiberKit) WebSocket(path string, rulePath ...string) *FiberWebSocket {
	k.fws = &FiberWebSocket{kit: k}
	k.fws.init()
	k.app.Use(path, func(c *fiber.Ctx) (err error) {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}
		if err = k.fws.upgradeHook(k, c); err != nil {
			return err
		}
		return c.Next()
	})

	var rp = path
	if len(rulePath) > 0 {
		rp = rulePath[0]
	}

	k.app.Get(rp, websocket.New(func(c *websocket.Conn) {
		var err error
		var fiberCtx = &FiberContext{conn: c}
		var rootActor = k.System().Context()
		var result vivid.Message
		var status = new(atomic.Uint32)

		if result, err = rootActor.FutureAsk(k.ownerRef, (*fiberConnectionOpenedMessage)(newFiberConnActor(k.ownerRef, status, k, fiberCtx, c))).Result(); err != nil {
			return
		}

		ref := result.(vivid.ActorRef)

		var (
			mt  int
			msg []byte
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				if status.CompareAndSwap(fiberConnStatusOnline, fiberConnStatusClosed) {
					rootActor.Tell(ref, err)
				}
				break
			}
			rootActor.Tell(ref, fiberReceivePacketMessage{packet: NewPacket(msg).SetContext(mt)})
		}
	}))

	return k.fws
}
