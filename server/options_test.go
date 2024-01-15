package server_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
	"time"
)

func TestWithLowMessageDuration(t *testing.T) {
	var cases = []struct {
		name     string
		duration time.Duration
	}{
		{name: "TestWithLowMessageDuration", duration: server.DefaultLowMessageDuration},
		{name: "TestWithLowMessageDuration_Zero", duration: 0},
		{name: "TestWithLowMessageDuration_Negative", duration: -server.DefaultAsyncLowMessageDuration},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			networks := server.GetNetworks()
			for i := 0; i < len(networks); i++ {
				low := false
				network := networks[i]
				srv := server.New(network,
					server.WithLowMessageDuration(c.duration),
				)
				srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
					low = true
					srv.Shutdown()
				})
				srv.RegStartFinishEvent(func(srv *server.Server) {
					if c.duration <= 0 {
						srv.Shutdown()
						return
					}
					time.Sleep(server.DefaultLowMessageDuration)
				})
				var lis string
				switch network {
				case server.NetworkNone, server.NetworkUnix:
					lis = "addr"
				default:
					lis = fmt.Sprintf(":%d", random.UsablePort())
				}
				if err := srv.Run(lis); err != nil {
					t.Fatalf("%s run error: %s", network, err)
				}
				if !low && c.duration > 0 {
					t.Fatalf("%s low message not exec", network)
				}
			}
		})
	}
}

func TestWithAsyncLowMessageDuration(t *testing.T) {
	var cases = []struct {
		name     string
		duration time.Duration
	}{
		{name: "TestWithAsyncLowMessageDuration", duration: time.Millisecond * 100},
		{name: "TestWithAsyncLowMessageDuration_Zero", duration: 0},
		{name: "TestWithAsyncLowMessageDuration_Negative", duration: -server.DefaultAsyncLowMessageDuration},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			networks := server.GetNetworks()
			for i := 0; i < len(networks); i++ {
				low := false
				network := networks[i]
				srv := server.New(network,
					server.WithAsyncLowMessageDuration(c.duration),
				)
				srv.RegMessageLowExecEvent(func(srv *server.Server, message *server.Message, cost time.Duration) {
					low = true
					srv.Shutdown()
				})
				srv.RegStartFinishEvent(func(srv *server.Server) {
					if c.duration <= 0 {
						srv.Shutdown()
						return
					}
					srv.PushAsyncMessage(func() error {
						time.Sleep(c.duration)
						return nil
					}, nil)
				})
				var lis string
				switch network {
				case server.NetworkNone, server.NetworkUnix:
					lis = fmt.Sprintf("%s%d", "addr", random.Int(0, 9999))
				default:
					lis = fmt.Sprintf(":%d", random.UsablePort())
				}
				if err := srv.Run(lis); err != nil {
					t.Fatalf("%s run error: %s", network, err)
				}
				if !low && c.duration > 0 {
					t.Fatalf("%s low message not exec", network)
				}
			}
		})
	}
}
