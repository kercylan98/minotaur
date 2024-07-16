package gateway

import (
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
)

func NewGateway(bindAddr string) *Gateway {
	return &Gateway{
		bindAddr: bindAddr,
	}
}

type Gateway struct {
	bindAddr string
	system   *vivid.ActorSystem
}

func (g *Gateway) OnLoad(support *vivid.ModuleSupport, hasTransportModule bool) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithName("gateway")
		options.WithModule(transport.NewTCP(g.bindAddr).BindService(new(bindServer)))
	})
}
