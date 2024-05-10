package server_test

import (
	"fmt"
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
			WithZombieConnectionDeadline(time.Second*5).
			WithLifeCycleLimit(chrono.Day*3).
			WithLogger(log.GetLogger()).
			WithEventOptions(nexus.NewEventOptions().WithDeadLockThreshold(time.Second*5, func(stack []byte) {
				t.Log("dead lock")
				fmt.Println(string(stack))
			})).
			WithIndependentGoroutineBroker(),
	)

	srv.RegisterConnectionOpenedEvent(func(srv server.Server, conn server.Conn) {

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
