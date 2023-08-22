package client_test

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"testing"
	"time"
)

func TestUnixDomainSocket_Write(t *testing.T) {
	var close = make(chan struct{})
	srv := server.New(server.NetworkUnix)
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet server.Packet) {
		t.Log(packet)
		conn.Write(packet)
	})
	srv.RegStartFinishEvent(func(srv *server.Server) {
		time.Sleep(time.Second)
		cli := client.NewUnixDomainSocket("./test.sock")
		cli.RegUDSConnectionOpenedEvent(func(conn *client.UnixDomainSocket) {
			conn.Write(server.NewPacketString("Hello~"))
		})
		cli.RegUDSConnectionReceivePacketEvent(func(conn *client.UnixDomainSocket, packet server.Packet) {
			t.Log(packet)
			close <- struct{}{}
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

	<-close
	srv.Shutdown()
}
