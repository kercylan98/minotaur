package client_test

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"sync"
	"testing"
)

func TestClient_WriteWS(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	srv := server.New(server.NetworkWebsocket)
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		srv.Shutdown()
	})
	srv.RegStopEvent(func(srv *server.Server) {
		wait.Done()
	})
	srv.RegMessageReadyEvent(func(srv *server.Server) {
		cli := client.NewWebsocket("ws://127.0.0.1:9999")
		cli.RegConnectionOpenedEvent(func(conn *client.Client) {
			conn.WriteWS(2, []byte("Hello"))
		})
		if err := cli.Run(); err != nil {
			panic(err)
		}
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}

	wait.Wait()
}
