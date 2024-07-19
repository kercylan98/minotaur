package prc

import (
	"context"
	"errors"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc/codec"
	"github.com/puzpuzpuz/xsync/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net"
)

// NewShared 创建一个资源控制器的共享
func NewShared(rc *ResourceController) *Shared {
	s := &Shared{
		rc:      rc,
		streams: xsync.NewMapOf[PhysicalAddress, sharedStream](),
		codec:   codec.NewProtobuf(),
	}
	s.streamServer = newSharedServer(s)
	return s
}

// Shared 资源控制器的共享
type Shared struct {
	*UnimplementedSharedServer
	codec        codec.Codec
	streamServer *sharedServer
	rc           *ResourceController // 共享的资源控制器
	grpc         *grpc.Server
	streams      *xsync.MapOf[PhysicalAddress, sharedStream]
}

// open 打开一个资源控制器
func (s *Shared) open(address PhysicalAddress) (sharedStream, error) {
	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := NewSharedClient(cc)
	server, err := client.StreamHandler(context.Background())
	if err != nil {
		return nil, err
	}

	if err = server.Send(&SharedMessage{
		MessageType: &SharedMessage_Handshake{Handshake: &Handshake{Address: s.rc.GetPhysicalAddress()}},
	}); err != nil {
		return nil, err
	}

	stream := newServerStream(address, s, server, cc)
	go func() {
		if err = s.streaming(address, stream); err != nil {
			panic(err)
		}
	}()

	return stream, nil
}

func (s *Shared) streaming(address PhysicalAddress, stream sharedStream) (err error) {
	s.attachStream(address, stream)
	defer func() {
		s.detachStream(address)
	}()

	var message *SharedMessage
	for {
		message, err = stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		switch m := message.MessageType.(type) {
		case *SharedMessage_DeliveryMessage:
			s.onDeliveryMessage(stream, address, m.DeliveryMessage)
		}
	}
}

// Share 共享资源控制器
func (s *Shared) Share() error {
	listener, err := net.Listen("tcp", s.rc.GetPhysicalAddress())
	if err != nil {
		return err
	}

	s.grpc = grpc.NewServer()
	s.grpc.RegisterService(&Shared_ServiceDesc, s.streamServer)

	go func() {
		if err := s.grpc.Serve(listener); err != nil {
			panic(err)
		}
	}()

	s.rc.RegisterResolver(FunctionalPhysicalAddressResolver(func(id *ProcessId) Process {
		process, exist := s.streams.Load(id.PhysicalAddress)
		if exist {
			return process.(sharedStream)
		}

		var err error
		process, err = s.open(id.PhysicalAddress)
		if err != nil {
			panic(err)
		}
		return process
	}))

	return nil
}

func (s *Shared) attachStream(address PhysicalAddress, stream sharedStream) {
	s.streams.Store(address, stream)
}

func (s *Shared) detachStream(address PhysicalAddress) {
	stream, loaded := s.streams.LoadAndDelete(address)
	if loaded {
		stream.Close()
	}
}

func (s *Shared) onDeliveryMessage(stream sharedStream, address PhysicalAddress, m *DeliveryMessage) {
	message, err := s.codec.Decode(m.MessageType, m.MessageData)
	if err != nil {
		panic(err)
	}

	sender, receiver := NewProcessRef(m.Sender), NewProcessRef(m.Receiver)
	receiverProcess := s.rc.GetProcess(receiver)

	if m.System {
		receiverProcess.DeliverySystemMessage(receiver, sender, nil, message)
	} else {
		receiverProcess.DeliveryUserMessage(receiver, sender, nil, message)
	}
}
