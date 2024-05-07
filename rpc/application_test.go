package rpc_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/rpcbuiltin"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	var opts = []nats.Option{
		nats.ReconnectWait(time.Second),
		nats.MaxReconnects(-1),
	}

	conn, err := nats.Connect("127.0.0.1:4222", opts...)
	if err != nil {
		panic(err)
	}
	js, err := conn.JetStream()
	if err != nil {
		panic(err)
	}

	discovery, err := rpcbuiltin.NewDiscoveryWithNats(conn, js)
	if err != nil {
		panic(err)
	}
	var cli = rpc.NewClient(discovery)

	for {
		if err := func() error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			resp, err := cli.UnaryCall("test-app", "account", "login")(ctx, struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}{
				Username: "test-un",
				Password: "test-pw",
			})
			if err != nil {
				return err
			}

			var result string
			resp.ReadTo(&result)

			fmt.Println(result)

			return nil
		}(); err != nil && !errors.Is(err, rpc.ErrServiceNotFound) {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}

func TestApplication(t *testing.T) {
	var opts = []nats.Option{
		nats.ReconnectWait(time.Second),
		nats.MaxReconnects(-1),
	}

	conn, err := nats.Connect("127.0.0.1:4222", opts...)
	if err != nil {
		panic(err)
	}
	js, err := conn.JetStream()
	if err != nil {
		panic(err)
	}

	registry, err := rpcbuiltin.NewRegistryWithNats(conn, js)
	if err != nil {
		panic(err)
	}

	var app = rpc.NewApplication(rpc.Service{
		Name:       "test-app",
		InstanceId: "test-app-1",
	}, registry)
	app.Register("account", "login").Unary(func(reader rpc.Reader) any {
		var params struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		reader.ReadTo(&params)

		fmt.Println("account login", params)

		return "ok"
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}
