package transport_test

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"testing"
	"time"
)

type EchoFiberService struct {
	kit *transport.FiberKit
}

func (e *EchoFiberService) OnInit(kit *transport.FiberKit) {
	e.kit = kit

	e.kit.WebSocket("/ws").
		ConnectionOpenedHook(func(kit *transport.FiberKit, ctx *transport.FiberContext, conn *transport.Conn) error {
			fmt.Println("connection opened")
			conn.WritePacket(transport.NewPacket([]byte("hello")).SetContext(websocket.TextMessage))
			//conn.Close()
			return nil
		}).
		ConnectionClosedHook(func(kit *transport.FiberKit, ctx *transport.FiberContext, conn *transport.Conn, err error) {
			fmt.Println("connection closed", err)
		}).
		ConnectionPacketHook(func(kit *transport.FiberKit, ctx *transport.FiberContext, conn *transport.Conn, packet transport.Packet) error {
			conn.WritePacket(packet)
			return nil
		})

}

func TestNewFiber(t *testing.T) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
		options.WithModule(transport.NewFiber(":8080").BindService(new(EchoFiberService)))
	})

	time.Sleep(time.Second * 9999999)
}
