package transport

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"sync"
)

func newEndpoint(network *Network, address core.Address) *endpoint {
	e := &endpoint{
		once:    new(sync.Once),
		network: network,
		address: address,
	}
	return e
}

type endpoint struct {
	once    *sync.Once
	mutex   sync.Mutex
	network *Network
	address core.Address
	stream  ActorSystemCommunication_StreamHandlerClient
}

func (e *endpoint) initConn() {
	e.once.Do(func() {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		addr := e.address.Address()
		cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			e.once = new(sync.Once)
			log.Error("Network", log.String("connection_endpoint", addr), log.Err(err))
			return
		}
		log.Debug("Network", log.String("connection_endpoint", addr))

		client := NewActorSystemCommunicationClient(cc)
		stream, err := client.StreamHandler(context.Background())
		if err != nil {
			e.once = new(sync.Once)
			log.Error("Network", log.String("connection_endpoint", addr), log.Err(err))
			return
		}

		err = stream.Send(&DistributedMessage{
			MessageType: &DistributedMessage_ConnectionOpen{
				ConnectionOpen: &ConnectionOpen{
					Address: string(e.address.Address()),
				},
			},
		})
		if err != nil {
			log.Error("Network", log.String("connection_endpoint", addr), log.Err(err))
			stream.CloseSend()
			e.once = new(sync.Once)
			return
		}

		e.stream = stream
		wait := make(chan struct{})
		var opened sync.Once
		defer opened.Do(func() { close(wait) })

		go func() {
			defer func() {
				e.mutex.Lock()
				defer e.mutex.Unlock()
				e.once = new(sync.Once)
			}()
			for {
				in, err := stream.Recv()
				if err != nil {
					return
				}

				switch m := in.MessageType.(type) {
				case *DistributedMessage_ConnectionOpened:
					opened.Do(func() { close(wait) })
				case *DistributedMessage_ConnectionMessage:
					e.network.server.onConnectionMessage(m)
				}
			}
		}()

		<-wait
	})
}

func (e *endpoint) send(sender *core.ProcessRef, receiver core.Address, message core.Message) {
	e.initConn()

	var (
		protoMessage   proto.Message
		isProtoMessage bool
		typeName       string
		bytes          []byte
		err            error
	)

	var processMessage = func(message core.Message) (typeName string, bytes []byte) {
		if protoMessage, isProtoMessage = message.(proto.Message); isProtoMessage {
			typeName = string(proto.MessageName(protoMessage))
			if bytes, err = proto.Marshal(protoMessage); err != nil {
				panic(err)
				return
			}
		} else {
			switch v := message.(type) {
			case string:
				typeName = "string"
				bytes = []byte(v)
			}
		}
		return
	}

	if rm, ok := message.(vivid.RegulatoryMessage); ok {
		sender = rm.Sender
		typeName, bytes = processMessage(rm.Message)
		message = &RegulatoryMessage{
			SenderAddress: []byte(rm.Sender.Address()),
			Message: &Message{
				TypeName:    typeName,
				MessageData: bytes,
			},
		}
	}

	typeName, bytes = processMessage(message)

	err = e.stream.Send(&DistributedMessage{
		MessageType: &DistributedMessage_ConnectionMessage{
			ConnectionMessage: &ConnectionMessage{
				SenderAddress:   []byte(sender.Address()),
				ReceiverAddress: []byte(receiver),
				Message: &Message{
					TypeName:    typeName,
					MessageData: bytes,
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}
}
