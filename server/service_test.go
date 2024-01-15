package server_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"testing"
	"time"
)

type TestService struct{}

func (ts *TestService) OnInit(srv *server.Server) {
	srv.RegStartFinishEvent(func(srv *server.Server) {
		fmt.Println("server start finish")
	})

	srv.RegStopEvent(func(srv *server.Server) {
		fmt.Println("server stop")
	})
}

func TestBindService(t *testing.T) {
	var cases = []struct {
		name string
	}{
		{name: "TestBindService"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			srv := server.New(server.NetworkNone, server.WithLimitLife(time.Millisecond))
			server.BindService(srv, new(TestService))
			if err := srv.RunNone(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
