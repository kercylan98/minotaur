package transport

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/grpc"
	"net"
)

var _ vivid.TransportModule = &Network{}

func NewNetwork(address string) *Network {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		panic(err)
	}
	n := &Network{
		address: core.NewRootAddress("", "", host, convert.StringToUint16(port)),
	}

	return n
}

type Network struct {
	support *vivid.ModuleSupport
	address core.Address     // 指定 ActorSystem 地址
	server  *server          // 处理流消息的服务器
	em      *endpointManager // 远程端点管理器
}

func (n *Network) OnLoad(support *vivid.ModuleSupport) {
	n.server = newServer(n)
	n.em = newEndpointManager(n)
	n.support = support
	n.support.RegAddressResolver(func(address core.Address) core.Process {
		return newRemoteActor(n, address)
	})

	n.launch()
}

func (n *Network) ActorSystemAddress() core.Address {
	return n.address
}

func (n *Network) launch() {
	listener, err := net.Listen("tcp", n.address.Address())
	if err != nil {
		panic(err)
	}

	grpcSrv := grpc.NewServer()
	grpcSrv.RegisterService(&ActorSystemCommunication_ServiceDesc, n.server)

	go func() {
		if err := grpcSrv.Serve(listener); err != nil {
			log.Error("Network", "onLaunch", "failed to serve", err)
		}
	}()
}

func (n *Network) send(sender *core.ProcessRef, receiver core.Address, message core.Message) {
	e := n.em.getEndpoint(receiver)
	e.send(sender, receiver, message)
}
