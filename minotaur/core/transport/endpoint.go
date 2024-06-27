package transport

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newEndpoint(network *Network, address core.Address) *endpoint {
	e := &endpoint{
		network: network,
		address: address.ParseToRoot(),
	}
	return e
}

type endpoint struct {
	network *Network
	address core.Address
	stream  ActorSystemCommunication_StreamHandlerClient
	conn    *grpc.ClientConn
}

func (e *endpoint) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		e.onLaunch(ctx, m)
	case vivid.OnTerminate:
		e.onTerminate(ctx, m)
	case []any:
		e.onSend(ctx, m)
	}
}

func (e *endpoint) onLaunch(ctx vivid.ActorContext, m vivid.OnLaunch) {
	addr := e.address.Address()
	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Endpoint", log.String("type", "connect"), log.Err(err))
		ctx.Terminate(ctx.Ref())
		return
	}
	e.conn = cc
	log.Debug("Endpoint", log.String("type", "connected"), log.String("addr", addr))

	client := NewActorSystemCommunicationClient(cc)
	stream, err := client.StreamHandler(context.Background())
	if err != nil {
		log.Error("Endpoint", log.String("type", "handshake"), log.Err(err))
		ctx.Terminate(ctx.Ref())
		return
	}

	err = stream.Send(&DistributedMessage{
		MessageType: &DistributedMessage_ConnectionOpen{
			ConnectionOpen: &ConnectionOpen{
				Address: e.network.address.Address(),
			},
		},
	})
	if err != nil {
		log.Error("Endpoint", log.String("type", "handshake"), log.Err(err))
		ctx.Terminate(ctx.Ref())
		return
	}

	in, err := stream.Recv()
	if err != nil {
		log.Error("Endpoint", log.String("type", "handshake"), log.Err(err))
		ctx.Terminate(ctx.Ref())
		return
	}
	if _, ok := in.MessageType.(*DistributedMessage_ConnectionOpened); !ok {
		log.Error("Endpoint", log.String("type", "handshake"), log.Err(err))
		ctx.Terminate(ctx.Ref())
		return
	}

	e.stream = stream
	go func() {
		defer func() {
			ctx.Terminate(ctx.Ref())
			log.Debug("Endpoint", log.String("type", "closed"))
		}()

		for {
			in, err := stream.Recv()
			if err != nil {
				return
			}

			switch m := in.MessageType.(type) {
			case *DistributedMessage_ConnectionMessageBatch:
				e.network.server.onConnectionMessage(m)
			default:
				log.Warn("Endpoint", log.String("type", "unknown message"))
			}
		}

	}()
}

func (e *endpoint) onTerminate(ctx vivid.ActorContext, m vivid.OnTerminate) {
	e.network.em.delEndpoint(e.address)
	if e.stream != nil {
		if err := e.stream.CloseSend(); err != nil {
			log.Error("Endpoint", log.String("type", "close"), log.String("instance", "stream"), log.Err(err))
		}
	}
	if e.conn != nil {
		if err := e.conn.Close(); err != nil {
			log.Error("Endpoint", log.String("type", "close"), log.String("instance", "conn"), log.Err(err))
		}
	}
}

func (e *endpoint) onSend(ctx vivid.ActorContext, m []any) {
	var batch = &ConnectionMessageBatch{
		SenderAddress:                  make([][]byte, len(m)),
		ReceiverAddress:                make([][]byte, len(m)),
		RegulatoryMessageSenderAddress: make([][]byte, len(m)),
		TypeName:                       make([]string, len(m)),
		MessageData:                    make([][]byte, len(m)),
		Bad:                            make([]bool, len(m)),
	}

	var bad = 0
	var ok bool
	for i, source := range m {
		var v messageWrapper
		if v, ok = source.(messageWrapper); !ok {
			batch.Bad[i] = true
			bad++
			log.Error("Endpoint", log.String("type", "convert"))
			continue
		}
		rm, ok := v.message.(vivid.RegulatoryMessage)
		if ok {
			batch.RegulatoryMessageSenderAddress[i] = []byte(rm.Sender.Address())
			v.message = rm.Message
		}
		var err error
		batch.TypeName[i], batch.MessageData[i], err = e.network.codec.Encode(v.message)
		if err != nil {
			batch.Bad[i] = true
			bad++
			log.Error("Endpoint", log.String("type", "encode"), log.Err(err))
			continue
		}
		batch.SenderAddress[i], batch.ReceiverAddress[i] = []byte(v.sender.Address()), []byte(v.receiver)
	}

	if bad == len(m) {
		return
	}

	err := e.stream.Send(&DistributedMessage{
		MessageType: &DistributedMessage_ConnectionMessageBatch{ConnectionMessageBatch: batch},
	})

	if err != nil {
		log.Error("Endpoint", log.String("type", "send"), log.Err(err))
		ctx.Terminate(ctx.Ref())
	}
}
