package server_test

import (
	"context"
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/server/internal/v2"
	"github.com/kercylan98/minotaur/server/internal/v2/network"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"github.com/kercylan98/minotaur/utils/random"
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

	srv := server.NewServer(network.WebSocket(":9999"), server.NewOptions().WithLifeCycleLimit(times.Day*3).WithLogger(log.GetLogger()))

	var tm = make(map[string]bool)

	srv.RegisterConnectionOpenedEvent(func(srv server.Server, conn server.Conn) {
		if err := conn.WritePacket(server.NewPacket([]byte("hello")).SetContext(ws.OpText)); err != nil {
			t.Error(err)
		}

		conn.WriteWebSocketText([]byte("hello text"))

		srv.PublishAsyncMessage("123", func(ctx context.Context) error {
			return nil
		}, func(ctx context.Context, err error) {
			for i := 0; i < 10000000; i++ {
				_ = tm["1"]
				tm["1"] = random.Bool()
				conn.WriteWebSocketText([]byte("hello text"))
			}
		})
	})

	srv.RegisterConnectionReceivePacketEvent(func(srv server.Server, conn server.Conn, packet server.Packet) {
		if err := conn.WritePacket(packet); err != nil {
			panic(err)
		}
		//srv.PushAsyncMessage(func(srv server.Server) error {
		//	for i := 0; i < 3; i++ {
		//		time.Sleep(time.Second)
		//	}
		//	return nil
		//}, func(srv server.Server, err error) {
		//	t.Log("callback")
		//})
	})

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
