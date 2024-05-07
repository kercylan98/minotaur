package server_test

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/network"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	srv := server.NewServer(network.WebSocket(":9999"),
		server.NewOptions().
			WithLifeCycleLimit(chrono.Day*3).
			WithLogger(log.GetLogger()).
			WithEventOptions(nexus.NewEventOptions().WithLowHandlerTrace(true, func(cost time.Duration, stack []byte) {
				t.Log("low handler trace", cost.String())
				fmt.Println(string(stack))
			})),
	)

	srv.RegisterConnectionOpenedEvent(func(srv server.Server, conn server.Conn) {
		if err := conn.WritePacket(server.NewPacket([]byte("hello")).SetContext(ws.OpText)); err != nil {
			t.Error(err)
		}

		conn.WriteWebSocketText([]byte("hello text"))

		srv.PublishAsyncMessage("123", func(ctx context.Context) error {
			time.Sleep(time.Second)
			return nil
		}, func(ctx context.Context, err error) {
			time.Sleep(time.Second)
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
