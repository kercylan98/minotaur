package server_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/utils/times"
	"golang.org/x/time/rate"
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
	for i := 0; i < 1000; i++ {
		id := i
		fmt.Println("启动", i+1)
		cli := client.NewWebsocket("ws://127.0.0.1:9999")
		cli.RegConnectionReceivePacketEvent(func(conn *client.Client, wst int, packet []byte) {
			fmt.Println("收到", id+1, string(packet))
		})
		cli.RegConnectionOpenedEvent(func(conn *client.Client) {
			go func() {
				for i < 1000 {
					time.Sleep(time.Second)
				}
				for {
					time.Sleep(time.Millisecond * 100)
					cli.WriteWS(2, []byte("hello"))
				}
			}()
		})
		if err := cli.Run(); err != nil {
			panic(err)
		}
	}

	time.Sleep(times.Week)
}
