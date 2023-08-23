package client_test

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"testing"
	"time"
)

func TestUnixDomainSocket_Write(t *testing.T) {
	var closed = make(chan struct{})
	srv := server.New(server.NetworkUnix)
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		t.Log(string(packet))
		conn.Write(packet)
	})
	srv.RegStartFinishEvent(func(srv *server.Server) {
		time.Sleep(time.Second)
		cli := client.NewUnixDomainSocket("./test.sock")
		cli.RegConnectionOpenedEvent(func(conn *client.Client) {
			conn.Write([]byte("Hello~"))
		})
		cli.RegConnectionReceivePacketEvent(func(conn *client.Client, wst int, packet []byte) {
			t.Log(packet)
			closed <- struct{}{}
		})
		if err := cli.Run(); err != nil {
			panic(err)
		}
	})
	go func() {
		if err := srv.Run("./test.sock"); err != nil {
			panic(err)
		}
	}()

	<-closed
	srv.Shutdown()
}
