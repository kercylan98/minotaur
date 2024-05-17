package server_test

import (
	"github.com/kercylan98/minotaur/server/actor_srv/server"
	"github.com/kercylan98/minotaur/server/actor_srv/server/network"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/vivids"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	system := vivid.NewActorSystem("TestServerSystem")
	if err := system.Run(); err != nil {
		t.Error(err)
	}

	srv := server.NewServer(network.WebSocket(":9966"))
	if ref, err := system.ActorOf(srv, vivids.NewActorOptions().WithName("server")); err != nil {
		t.Error(err)
	} else {
		if err = ref.Tell("start"); err != nil {
			t.Error(err)
		}
	}

	time.Sleep(chrono.Day)
}
