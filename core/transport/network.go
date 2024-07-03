package transport

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/grpc"
	"math"
	"net"
)

var _ vivid.TransportModule = &Network{}
var _ vivid.PriorityModule = &Network{}

// NewNetwork 创建一个网络模块，该模块用于给 ActorSystem 赋予网络通信的能力，支持跨网络的 Actor 通信
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

type messageWrapper struct {
	sender   *core.ProcessRef
	receiver core.Address
	message  core.Message
	system   bool
}

type Network struct {
	support *vivid.ModuleSupport
	address core.Address     // 指定 ActorSystem 地址
	server  *server          // 处理流消息的服务器
	em      *endpointManager // 远程端点管理器
	codec   core.Codec       // 消息编解码器
	grpc    *grpc.Server     // grpc 服务
}

func (n *Network) Priority() int {
	return math.MinInt
}

func (n *Network) OnLoad(support *vivid.ModuleSupport, hasTransportModule bool) {
	n.server = newServer(n)
	n.em = newEndpointManager(n)
	n.codec = core.NewProtobufExpandCodec()
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

	n.grpc = grpc.NewServer()
	n.grpc.RegisterService(&ActorSystemCommunication_ServiceDesc, n.server)

	go func() {
		if err := n.grpc.Serve(listener); err != nil {
			log.Error("Network", "onLaunch", "failed to serve", err)
		}
	}()
}

func (n *Network) send(sender *core.ProcessRef, receiver core.Address, message core.Message, system bool) {
	ref := n.em.getEndpoint(receiver)
	n.support.System().Context().Tell(ref, messageWrapper{
		sender:   sender,
		receiver: receiver,
		message:  message,
		system:   system,
	})
}

func (n *Network) OnShutdown() {
	n.em.close()
	n.grpc.GracefulStop()
}
