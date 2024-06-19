package gateway

import (
	"errors"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/balancer"
	"github.com/kercylan98/minotaur/toolkit/log"
	"net"
)

type L4TCPListenerActor struct {
	address   Address
	srv       vivid.TypedActorRef[transport.ServerActorTyped]
	endpoints balancer.Balancer[EndpointId, *Endpoint]
}

func (l *L4TCPListenerActor) OnReceive(ctx vivid.MessageContext) {
	switch ctx.GetMessage().(type) {
	case vivid.OnBoot:
		ctx.Become(vivid.BehaviorOf(l.onBindAddress))
		ctx.Become(vivid.BehaviorOf(l.onBindBalancer))
		ctx.Become(vivid.BehaviorOf(l.onBindEndpoint))
		ctx.Become(vivid.BehaviorOf(l.onConnectionOpened))
	}
}

func (l *L4TCPListenerActor) onBindAddress(ctx vivid.MessageContext, message ListenerActorBindAddressMessage) {
	if l.address == message.Address {
		return
	}
	l.address = message.Address

	if l.srv != nil {
		l.srv.Stop()
		l.srv = nil
	}

	l.srv = transport.NewServerActor(ctx.GetSystem(), vivid.NewActorOptions[*transport.ServerActor]().
		WithSupervisor(func(message, reason vivid.Message) vivid.Directive {
			return vivid.DirectiveRestart
		}),
	)

	l.srv.Tell(transport.ServerLaunchMessage{
		Network: network.Tcp(message.Address),
	})

	l.srv.Api().SubscribeConnOpenedEvent(ctx)
}

func (l *L4TCPListenerActor) onConnectionOpened(ctx vivid.MessageContext, message transport.ServerConnectionOpenedEvent) {
	// 选择端点
	endpoint, err := l.endpoints.Select()
	if err != nil {
		ctx.GetSystem().GetLogger().Error("SelectEndpoint", log.Err(err))
		message.Conn.Stop()
		return
	}

	// 连接到端点
	tcpAddr, err := net.ResolveTCPAddr("tcp", endpoint.GetAddress())
	if err != nil {
		ctx.GetSystem().GetLogger().Error("ResolveTcpAddr", log.Err(err))
		message.Conn.Stop()
		return
	}

	endpointConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		ctx.GetSystem().GetLogger().Error("DialTcp", log.Err(err))
		message.Conn.Stop()
		return
	}

	message.Conn.Api().SetTerminateHandler(func(ctx vivid.MessageContext, conn transport.Conn, message vivid.OnTerminate) {
		endpointConn.Close()
	})
	message.Conn.Api().SetPacketHandler(func(ctx vivid.MessageContext, conn transport.Conn, packet transport.ConnectionReactPacketMessage) {
		endpointConn.Write(packet.Packet.GetBytes())
	})
}

func (l *L4TCPListenerActor) onBindBalancer(ctx vivid.MessageContext, message ListenerActorBindBalancerMessage) {
	l.endpoints = message.Balancer
}

func (l *L4TCPListenerActor) onBindEndpoint(ctx vivid.MessageContext, message ListenerActorBindEndpointMessage) {
	if l.endpoints == nil {
		ctx.GetSystem().GetLogger().Error("BindEndpoint", log.Err(errors.New("endpoints balancer is nil")))
		return
	}

	l.endpoints.Add(message.Endpoint)
}
