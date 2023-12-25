package server_test

import (
	"github.com/kercylan98/minotaur/server"
	"testing"
	"time"
)

type TestService struct {
}

func (ts *TestService) OnInit(srv *server.Server) {
	srv.RegStartFinishEvent(func(srv *server.Server) {
		println("Server started")
	})

	srv.RegStopEvent(func(srv *server.Server) {
		println("Server stopped")
	})
}

func TestBindService(t *testing.T) {
	srv := server.New(server.NetworkNone, server.WithLimitLife(time.Second))

	server.BindService(srv, new(TestService))

	if err := srv.RunNone(); err != nil {
		panic(err)
	}
}
