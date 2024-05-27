package transport_test

import (
	"github.com/kercylan98/minotaur/minotaur/server3/network"
	"github.com/kercylan98/minotaur/vivid"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestNewServer(t *testing.T) {
	system := vivid.NewActorSystem("test-srv")
	srv := server.NewServer(&system, network.WebSocket(":8899"))
	srv.Run()

	var s = make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-s

	system.Shutdown()
	t.Log("Server shutdown")
}
