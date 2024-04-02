package server_test

import (
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/server/internal/v2"
	"github.com/kercylan98/minotaur/server/internal/v2/network"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	server.EnableHttpPProf(":9998", "/debug/pprof", func(err error) {
		panic(err)
	})

	go func() {
		time.Sleep(time.Second * 5)
		server.DisableHttpPProf()
		time.Sleep(time.Second * 5)
		server.EnableHttpPProf(":9998", "/debug/pprof", func(err error) {
			panic(err)
		})
	}()

	srv := server.NewServer(network.WebSocket(":9999"), server.NewOptions().WithLifeCycleLimit(times.Second*3))

	srv.RegisterConnectionOpenedEvent(func(srv server.Server, conn server.Conn) {
		if err := conn.WritePacket(server.NewPacket([]byte("hello")).SetContext(ws.OpText)); err != nil {
			t.Error(err)
		}
	})

	srv.RegisterConnectionReceivePacketEvent(func(srv server.Server, conn server.Conn, packet server.Packet) {
		if err := conn.WritePacket(packet); err != nil {
			panic(err)
		}
	})

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
