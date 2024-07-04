package transport_test

import (
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"testing"
	"time"
)

func TestNewWebSocket(t *testing.T) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewWebSocket(":8080", "/ws").
			SetPacketHandler(func(conn *transport.Conn, packet transport.Packet) {
				conn.WritePacket(packet)
			}).
			SetConnOpenedHandler(func(conn *transport.Conn) {
				t.Log("conn opened", conn.RemoteAddr().String())
			}).
			SetConnClosedHandler(func(conn *transport.Conn, err error) {
				t.Log("conn closed", conn.RemoteAddr().String(), err)
			}),
		)
	}).Shutdown()

	time.Sleep(time.Second * 1000)
}
