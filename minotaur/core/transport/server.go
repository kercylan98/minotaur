package transport

import (
	"errors"
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
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
		case *DistributedMessage_ConnectionMessage:
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
	}); err != nil {
		log.Debug("connectionClosed", log.String("address", conn.openedInfo.Address), log.Err(err))
	} else {
		log.Debug("connectionClosed", log.String("address", conn.openedInfo.Address))
	}
}

func (s *server) onConnectionMessage(m *DistributedMessage_ConnectionMessage) {
	var processMessage = func(message *Message) core.Message {
		typ, findMessageErr := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(message.TypeName))
		var userDefine any
		if findMessageErr != nil {
			switch message.TypeName {
			case "string":
				userDefine = string(message.MessageData)
			default:
				log.Error("find message error", log.Err(findMessageErr))
				return nil
			}
		}

		var m any
		if userDefine == nil {
			protoMessage := typ.New().Interface()
			if unmarshalErr := proto.Unmarshal(message.MessageData, protoMessage); unmarshalErr != nil {
				log.Error("unmarshal message error", log.Err(unmarshalErr))
				return nil
			}
			m = protoMessage
		} else {
			m = userDefine
		}

		return m
	}

	message := processMessage(m.ConnectionMessage.Message)

	switch v := message.(type) {
	case *RegulatoryMessage:
		message = vivid.RegulatoryMessage{
			Sender:  core.NewProcessRef(core.Address(v.SenderAddress)),
			Message: processMessage(v.Message),
		}
	}

	sender := core.NewProcessRef(core.Address(m.ConnectionMessage.SenderAddress))
	receiver := s.network.support.GetProcess(core.Address(m.ConnectionMessage.ReceiverAddress))
	receiver.SendUserMessage(sender, message)
}
