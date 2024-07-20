package prc

import (
	"errors"
)

func newSharedServer(s *Shared) *sharedServer {
	return &sharedServer{shared: s}
}

type sharedServer struct {
	*UnimplementedSharedServer
	shared *Shared
}

// StreamHandler 监听来自客户端的流
func (s *sharedServer) StreamHandler(client Shared_StreamHandlerServer) error {
	handshakeMessage, err := client.Recv()
	if err != nil {
		return err
	}

	handshake, ok := handshakeMessage.MessageType.(*SharedMessage_Handshake)
	if !ok {
		return errors.New("handshake message is expected")
	}

	if handshake.Handshake.Address == s.shared.rc.GetPhysicalAddress() {
		return errors.New("loop-back connection is not allowed")
	}

	return s.shared.streaming(handshake.Handshake.Address, newClientStream(handshake.Handshake.Address, s.shared, client))
}
