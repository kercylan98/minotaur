package server_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	srv := server.New(server.NetworkWebsocket, server.WithPProf())
	srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn, err any) {
		fmt.Println("关闭", conn.GetID(), err, "Count", srv.GetOnlineCount())
	})

	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}

func TestNewClient(t *testing.T) {
	count := 500
	for i := 0; i < count; i++ {
		fmt.Println("启动", i+1)
		cli := client.NewWebsocket("ws://172.29.5.138:9999")
		cli.RegConnectionReceivePacketEvent(func(conn *client.Client, wst int, packet []byte) {
			fmt.Println(time.Now().Unix(), "收到", string(packet))
		})
		cli.RegConnectionClosedEvent(func(conn *client.Client, err any) {
			fmt.Println("关闭", err)
		})
		cli.RegConnectionOpenedEvent(func(conn *client.Client) {
			go func() {
				for i < count {
					time.Sleep(time.Second)
				}
				for {
					for i := 0; i < 10; i++ {
						cli.WriteWS(2, []byte("hello"))
					}
				}
			}()
		})
		if err := cli.Run(); err != nil {
			panic(err)
		}
	}

	time.Sleep(times.Week)
}
