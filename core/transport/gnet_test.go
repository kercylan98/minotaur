package transport_test

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"testing"
	"time"
)

type EchoGNETService struct {
	kit *transport.GNETKit
}

func (e *EchoGNETService) OnInit(kit *transport.GNETKit) {
	e.kit = kit

	e.kit.ConnectionOpenedHook(func(kit *transport.GNETKit, conn *transport.Conn) error {
		fmt.Println("connection opened")
		conn.WritePacket(transport.NewPacket([]byte("hello")).SetContext(ws.OpText))
		//conn.Close()
		return nil
	}).
		ConnectionClosedHook(func(kit *transport.GNETKit, conn *transport.Conn, err error) {
			fmt.Println("connection closed", err)
		}).
		ConnectionPacketHook(func(kit *transport.GNETKit, conn *transport.Conn, packet transport.Packet) error {
			conn.WritePacket(packet)
			return nil
		})

}

func TestNewGNet(t *testing.T) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
		options.WithModule(transport.NewWebSocket(":8080", "/ws").BindService(new(EchoGNETService)))
	})

	time.Sleep(time.Second * 9999999)
}
