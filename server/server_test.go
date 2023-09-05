package server_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/utils/times"
	"golang.org/x/time/rate"
	"sync/atomic"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	limiter := rate.NewLimiter(rate.Every(time.Second), 100)
	srv := server.New(server.NetworkWebsocket, server.WithMessageBufferSize(1024*1024), server.WithPProf())
	srv.RegMessageExecBeforeEvent(func(srv *server.Server, message *server.Message) bool {
		t, c := srv.TimeoutContext(time.Second * 5)
		defer c()
		if err := limiter.Wait(t); err != nil {
			return false
		}
		return true
	})
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}

func TestNewClient(t *testing.T) {
	var total atomic.Int64
	for i := 0; i < 1000; i++ {
		cli := client.NewWebsocket("ws://127.0.0.1:9999")
		cli.RegConnectionReceivePacketEvent(func(conn *client.Client, wst int, packet []byte) {
			fmt.Println(string(packet))
		})
		cli.RegConnectionOpenedEvent(func(conn *client.Client) {
			go func() {
				for {
					cli.WriteWS(2, []byte("hello"))
					total.Add(1)
				}
			}()
		})
		if err := cli.Run(); err != nil {
			panic(err)
		}
	}

	time.Sleep(times.Week)
}
