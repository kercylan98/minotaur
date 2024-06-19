package gateway_test

import (
	"github.com/kercylan98/minotaur/minotaur/gateway"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/balancer"
	"testing"
)

func TestGateway_BindListener(t *testing.T) {
	system := vivid.NewActorSystem("gateway")
	ref := system.ActorOf(vivid.OfO[*gateway.L4TCPListenerActor]())
	ref.Tell(gateway.ListenerActorBindAddressMessage{Address: ":8080"})
	ref.Tell(gateway.ListenerActorBindBalancerMessage{Balancer: balancer.NewConsistentHashWeight[gateway.EndpointId, *gateway.Endpoint](1)})
	ref.Tell(gateway.ListenerActorBindEndpointMessage{Endpoint: &gateway.Endpoint{
		Id:     1,
		Host:   "192.168.2.112",
		Port:   10000,
		Weight: 10,
	}})

	system.AwaitShutdown()
}
