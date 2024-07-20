package prc

import (
	"google.golang.org/grpc"
)

type sharedStream interface {
	Process
	Send(*SharedMessage) error
	Recv() (*SharedMessage, error)
	Close()
}

func newClientStream(address PhysicalAddress, shared *Shared, stream Shared_StreamHandlerServer) *clientStream {
	s := &clientStream{
		stream: stream,
	}
	s.sharedStreamProcess = newSharedStreamProcess(address, s, shared)
	return s
}

func newServerStream(address PhysicalAddress, shared *Shared, stream Shared_StreamHandlerClient, conn *grpc.ClientConn) *serverStream {
	s := &serverStream{
		cc:     conn,
		stream: stream,
	}
	s.sharedStreamProcess = newSharedStreamProcess(address, s, shared)
	return s
}

type clientStream struct {
	*sharedStreamProcess
	stream Shared_StreamHandlerServer
}

func (c *clientStream) Send(message *SharedMessage) error {
	return c.stream.Send(message)
}

func (c *clientStream) Recv() (*SharedMessage, error) {
	return c.stream.Recv()
}

func (c *clientStream) Close() {
	return
}

type serverStream struct {
	*sharedStreamProcess
	cc     *grpc.ClientConn
	stream Shared_StreamHandlerClient
}

func (s *serverStream) Send(message *SharedMessage) error {
	return s.stream.Send(message)
}

func (s *serverStream) Recv() (*SharedMessage, error) {
	return s.stream.Recv()
}

func (s *serverStream) Close() {
	_ = s.stream.CloseSend()
	_ = s.cc.Close()
}
