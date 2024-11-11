package prc

import (
	"google.golang.org/grpc"
	"time"
)

type sharedStream interface {
	Process
	Send(*SharedMessage) error
	Recv() (*SharedMessage, error)
	Close()
	LoadArchives(messages [][]byte)
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
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.batches) > 0 {
		c.shared.streamArchiveLock.Lock()
		c.shared.streamArchives[c.address] = c.batches
		c.shared.streamArchiveTimeout[c.address] = time.AfterFunc(c.shared.config.disconnectionMessageRetentionTime, func() {
			c.shared.streamArchiveLock.Lock()
			delete(c.shared.streamArchives, c.address)
			delete(c.shared.streamArchiveTimeout, c.address)
			c.shared.streamArchiveLock.Unlock()
		})
		c.shared.streamArchiveLock.Unlock()
	}
}

func (c *clientStream) LoadArchives(messages [][]byte) {
	if len(messages) == 0 {
		return
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	c.batches = append(messages, c.batches...)
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

	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.batches) > 0 {
		s.shared.streamArchiveLock.Lock()
		s.shared.streamArchives[s.address] = s.batches
		s.shared.streamArchiveTimeout[s.address] = time.AfterFunc(s.shared.config.disconnectionMessageRetentionTime, func() {
			s.shared.streamArchiveLock.Lock()
			delete(s.shared.streamArchives, s.address)
			delete(s.shared.streamArchiveTimeout, s.address)
			s.shared.streamArchiveLock.Unlock()
		})
		s.shared.streamArchiveLock.Unlock()
	}
}

func (s *serverStream) LoadArchives(messages [][]byte) {
	if len(messages) == 0 {
		return
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.batches = append(messages, s.batches...)
}
