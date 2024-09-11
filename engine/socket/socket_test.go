package socket_test

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
)

func TestSocket(t *testing.T) {
	system := vivid.NewActorSystem()
	socketFactory := socket.NewFactory(system)
	fiberApp := fiber.New()
	fiberApp.Get("ws", websocket.New(func(conn *websocket.Conn) {
		c := socketFactory.Produce(new(Conn), func(packet []byte, ctx any) error {
			return conn.WriteMessage(ctx.(int), packet)
		}, func() error {
			if err := conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
				return err
			}
			return conn.Close()
		})

		defer c.Close()
		for {
			typ, bytes, err := conn.ReadMessage()
			if err != nil {
				c.Close(err)
				return
			}

			c.React(bytes, typ)
		}
	}))

	if err := fiberApp.Listen(":8080"); err != nil {
		panic(err)
	}
}

type Conn struct {
}

func (c *Conn) OnOpened(ctx vivid.ActorContext, socket socket.Socket) {
	fmt.Println("opened")
	socket.Write([]byte("opened"), websocket.TextMessage)
}

func (c *Conn) OnClose(ctx vivid.ActorContext, socket socket.Socket, err error) {
	fmt.Println("close:", err)
	socket.Write([]byte("close"), websocket.TextMessage)
}

func (c *Conn) OnPacket(ctx vivid.ActorContext, socket socket.Socket, packet *socket.Packet) {
	fmt.Println("received:", string(packet.GetData()))
	if string(packet.GetData()) == "close" {
		socket.Close(fmt.Errorf("close"))
	}
	socket.WritePacket(packet)
}

func (c *Conn) OnReceive(ctx vivid.ActorContext) {

}
