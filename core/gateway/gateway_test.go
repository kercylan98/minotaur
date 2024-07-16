package gateway_test

import (
	"github.com/kercylan98/minotaur/core/gateway"
	"github.com/kercylan98/minotaur/core/vivid"
	"os"
	"testing"
)

func TestGatewayServer(t *testing.T) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(gateway.NewGateway(":8888"))
	}).Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.ShutdownGracefully()
	})
}

func TestGatewayClient(t *testing.T) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(gateway.NewService(":8888", &gateway.ServiceInfo{
			Name:        "test",
			Description: "test service",
			Address:     ":8080",
			NetworkType: gateway.ServiceNetworkType_SNT_TCP,
			Weight:      100,
		}))
	}).Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.ShutdownGracefully()
	})
}
