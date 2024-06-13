package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type ServerActorExpandTyped interface {
	// Attach 用于将一个连接绑定到服务器上，当绑定失败时，会获得一个 nil 的返回值
	Attach(conn net.Conn, writer ConnWriter) vivid.TypedActorRef[ConnActorExpandTyped]

	// Detach 用于将一个连接从服务器上解绑
	Detach(conn net.Conn)
}

type ServerActorExpandTypedImpl struct {
	ServerActorRef vivid.ActorRef
}

func (s *ServerActorExpandTypedImpl) Attach(conn net.Conn, writer ConnWriter) vivid.TypedActorRef[ConnActorExpandTyped] {
	connActorExpandTyped, _ := s.ServerActorRef.Ask(ServerConnOpenedMessage{conn, writer}).(vivid.TypedActorRef[ConnActorExpandTyped])
	return connActorExpandTyped
}

func (s *ServerActorExpandTypedImpl) Detach(conn net.Conn) {
	s.ServerActorRef.Tell(ServerConnClosedMessage{conn})
}
