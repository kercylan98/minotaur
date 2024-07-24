package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/examples/internal/chat-room/internal/types"
	"strings"
)

func main() {
	fiberApp := fiber.New()
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)
	roomManager := system.ActorOfF(func() vivid.Actor {
		return types.NewRoomManager()
	})

	fiberApp.Get("/ws", stream.NewFiberWebSocketHandler(fiberApp, system, stream.FunctionalConfigurator(func(c *stream.Configuration) {
		var writer stream.Writer
		var conn *websocket.Conn
		var currRoom vivid.ActorRef
		c.WithPerformance(vivid.ActorFunctionalPerformance(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case stream.Writer:
				writer = m
			case *websocket.Conn:
				conn = m
				if conn.Query("userId") == "" {
					ctx.Terminate(ctx.Ref(), true)
					return
				}
			case *stream.Packet:
				str := string(m.Data())
				if strings.HasPrefix(str, "/") {
					var args = strings.Split(str, " ")
					var cmd = args[0]
					args = args[1:]
					switch cmd {
					case "/join":
						if len(args) > 0 {
							room, err := ctx.FutureAsk(roomManager, types.GetRoomMessage{RoomId: types.RoomId(args[0])}).Result()
							if err != nil {
								ctx.Tell(writer, m.Derivation([]byte("join room failed!")))
							}
							currRoom = room.(vivid.ActorRef)

							ctx.Tell(currRoom, types.JoinRoomMessage{
								UserId: types.UserId(conn.Query("userId")),
								User:   ctx.Ref(),
							})
						}
					}
				} else if strings.HasPrefix(str, "[") {
					ctx.Tell(writer, m)
				} else {
					if currRoom == nil {
						ctx.Tell(writer, m.Derivation([]byte("player send /join $roomId join room")))
						return
					}
					ctx.Tell(currRoom, &types.Chat{
						UserId: types.UserId(conn.Query("userId")),
						Packet: m,
					})
				}
			}
		}))
	})))

	if err := fiberApp.Listen(":8080"); err != nil {
		panic(err)
	}
}
