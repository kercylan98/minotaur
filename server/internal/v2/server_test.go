package server_test

import (
	"github.com/kercylan98/minotaur/server/internal/v2"
	"github.com/kercylan98/minotaur/server/internal/v2/network"
	"testing"
)

func TestNewServer(t *testing.T) {
	srv := server.server.NewServer(network.WebSocket(":9999"))
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
