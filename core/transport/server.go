package transport

import (
	"errors"
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"sync"
)

var _ ActorSystemCommunicationServer = &server{}

func newServer(network *Network) *server {
	return &server{
		network: network,
	}
}

type server struct {
	UnimplementedActorSystemCommunicationServer
	network      *Network
	connections  map[ActorSystemCommunication_StreamHandlerServer]*serverConn
	connectionRw sync.RWMutex
}

func (s *server) StreamHandler(stream ActorSystemCommunication_StreamHandlerServer) error {
	defer s.onConnectionClosed(stream)

	for {
		msg, err := stream.Recv()
		switch {
		case errors.Is(err, io.EOF):
			return nil
		case err != nil:
			return err
		}

		switch m := msg.MessageType.(type) {
		case *DistributedMessage_ConnectionOpen:
			if err = s.onConnectionOpened(stream, m.ConnectionOpen); err != nil {
				return err
			}
		case *DistributedMessage_ConnectionClosed:
			return nil
		case *DistributedMessage_ConnectionMessageBatch:
			s.onConnectionMessage(m)
		}
	}
}

func (s *server) mustEmbedUnimplementedActorSystemCommunicationServer() {}

func (s *server) onConnectionOpened(stream ActorSystemCommunication_StreamHandlerServer, m *ConnectionOpen) error {
	s.connectionRw.Lock()
	if s.connections == nil {
		s.connections = make(map[ActorSystemCommunication_StreamHandlerServer]*serverConn)
	}
	s.connections[stream] = &serverConn{
		openedInfo: m,
	}
	s.connectionRw.Unlock()

	if err := stream.Send(&DistributedMessage{
		MessageType: &DistributedMessage_ConnectionOpened{ConnectionOpened: &ConnectionOpened{}},
	}); err != nil {
		log.Debug("connectionOpened", log.String("address", m.Address), log.Err(err))
		return err
	} else {
		log.Debug("connectionOpened", log.String("address", m.Address))
		return nil
	}
}

func (s *server) onConnectionClosed(stream ActorSystemCommunication_StreamHandlerServer) {
	s.connectionRw.Lock()
	conn, exist := s.connections[stream]
	if !exist {
		s.connectionRw.Unlock()
		return
	}
	delete(s.connections, stream)
	s.connectionRw.Unlock()
	if err := stream.Send(&DistributedMessage{
		MessageType: &DistributedMessage_ConnectionClosed{
			ConnectionClosed: &ConnectionClosed{},
		},
	}); err != nil && status.Code(err) != codes.Unavailable {
		log.Debug("connectionClosed", log.String("address", conn.openedInfo.Address), log.Err(err))
	} else {
		log.Debug("connectionClosed", log.String("address", conn.openedInfo.Address))
	}
}

func (s *server) onConnectionMessage(batch *DistributedMessage_ConnectionMessageBatch) {
	var b = batch.ConnectionMessageBatch
	for i := 0; i < len(b.MessageData); i++ {
		if b.Bad[i] {
			continue
		}

		message, err := s.network.codec.Decode(b.TypeName[i], b.MessageData[i])
		if err != nil {
			log.Error("Endpoint", log.String("type", "decode"), log.Err(err))
			continue
		}

		sender := core.Address(b.SenderAddress[i])
		receiver := core.Address(b.ReceiverAddress[i])
		regulatoryMessageSender := core.Address(b.RegulatoryMessageSenderAddress[i])
		if len(regulatoryMessageSender) > 0 {
			message = vivid.RegulatoryMessage{
				Sender:  core.NewProcessRef(regulatoryMessageSender),
				Message: message,
			}
		}
		system := b.System[i]
		if system {
			s.network.support.GetProcess(receiver).SendSystemMessage(core.NewProcessRef(sender), message)
		} else {
			s.network.support.GetProcess(receiver).SendUserMessage(core.NewProcessRef(sender), message)
		}
	}
}
