package transport

import (
	"context"
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"io"
	"net"
	"sync"
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
	UnimplementedServiceServer
	address       core.Address
	support       *vivid.ModuleSupport
	connectionsRw sync.RWMutex
	connections   map[Service_OnConnectionOpenServer]struct{}
}

func (n *Network) OnLoad(support *vivid.ModuleSupport) {
	support.SetAddressResolverOnlyAddress()
	n.support = support
	n.support.System().ActorOf(func() vivid.Actor {
		return n
	})
	n.support.RegAddressResolver(n.onCreateEndpoint)
}

func (n *Network) ActorSystemAddress() core.Address {
	return n.address
}

func (n *Network) OnConnectionOpen(stream Service_OnConnectionOpenServer) error {
	for {
		msg, err := stream.Recv()
		switch {
		case errors.Is(err, io.EOF):
			return nil
		case err != nil:
			return err
		}

		switch m := msg.MessageType.(type) {
		case *DistributedMessage_ConnectionOpened:
			n.connectionsRw.Lock()
			if n.connections == nil {
				n.connections = make(map[Service_OnConnectionOpenServer]struct{})
			}
			n.connections[stream] = struct{}{}
			n.connectionsRw.Unlock()
			log.Debug("connection opened", log.String("address", core.Address(m.ConnectionOpened.Address).String()))
		case *DistributedMessage_ConnectionClosed:
			n.connectionsRw.Lock()
			delete(n.connections, stream)
			n.connectionsRw.Unlock()
			log.Debug("connection closed")
		case *DistributedMessage_ConnectionMessage:
			typ, findMessageErr := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(m.ConnectionMessage.Message.TypeName))
			if findMessageErr != nil {
				log.Error("find message error", log.Err(err))
				break
			}

			message := typ.New().Interface()

			if unmarshalErr := proto.Unmarshal(m.ConnectionMessage.Message.MessageData, message); unmarshalErr != nil {
				log.Error("unmarshal message error", log.Err(unmarshalErr))
				break
			}

			sender := core.NewProcessRef(core.Address(m.ConnectionMessage.SenderAddress))
			receiver := n.support.GetProcess(core.Address(m.ConnectionMessage.ReceiverAddress))
			receiver.SendUserMessage(sender, message)
		}
	}
}

func (n *Network) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		n.onLaunch(ctx)
	case vivid.FutureForwardMessage:
		n.onFutureForwardMessage(ctx, m)
	}
}

func (n *Network) onLaunch(ctx vivid.ActorContext) {
	listener, err := net.Listen("tcp", n.address.Address())
	if err != nil {
		panic(err)
	}

	grpcSrv := grpc.NewServer()
	grpcSrv.RegisterService(&Service_ServiceDesc, n)

	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		log.Debug("network start listen", log.String("address", n.address.Address()))
		return grpcSrv.Serve(listener)
	})
}

func (n *Network) onFutureForwardMessage(ctx vivid.ActorContext, m vivid.FutureForwardMessage) {
	if m.Error != nil {
		fmt.Println("network shutdown with error", m.Error)
	}
	ctx.Terminate(ctx.Ref())
}

func (n *Network) onCreateEndpoint(address core.Address) core.Process {
	return &Network{address: address}
}

func (n *Network) GetAddress() core.Address {
	return n.address
}

func (n *Network) SendUserMessage(sender *core.ProcessRef, message core.Message) {
	conn, err := grpc.NewClient(n.address.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	stream, err := conn.NewStream(context.Background(), &Service_ServiceDesc.Streams[0], "/remote.Service/OnConnectionOpen")
	if err != nil {
		panic(err)
	}
	if err = stream.SendMsg(&DistributedMessage{
		MessageType: &DistributedMessage_ConnectionOpened{
			ConnectionOpened: &ConnectionOpened{
				Address: string(n.address),
			},
		},
	}); err != nil {
		panic(err)
	}
}

func (n *Network) SendSystemMessage(sender *core.ProcessRef, message core.Message) {

}

func (n *Network) Terminate(ref *core.ProcessRef) {

}
