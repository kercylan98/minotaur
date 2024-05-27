package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type ServerCore interface {
	// Attach 用于将一个连接绑定到服务器上，当绑定失败时，会获得一个 nil 的返回值
	Attach(conn net.Conn, writer ConnWriter) ConnCore

	// Detach 用于将一个连接从服务器上解绑
	Detach(conn net.Conn)
}

type serverCore struct {
	serverActor vivid.ActorRef
}

func (s *serverCore) init(ref vivid.ActorRef) *serverCore {
	s.serverActor = ref
	return s
}

func (s *serverCore) Attach(conn net.Conn, writer ConnWriter) ConnCore {
	core, _ := s.serverActor.Ask(ConnOpenedMessage{conn, writer}).(ConnCore)
	return core
}

func (s *serverCore) Detach(conn net.Conn) {
	s.serverActor.Tell(ConnClosedMessage{conn})
}
