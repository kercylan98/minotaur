package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type ServerActorTyped interface {
	vivid.ActorTyped

	// Attach 用于将一个连接绑定到服务器上，当绑定失败时，会获得一个 nil 的返回值
	Attach(conn net.Conn, writer ConnWriter) ConnActorTyped

	// Detach 用于将一个连接从服务器上解绑
	Detach(conn net.Conn)

	// SubscribeConnOpenedEvent 订阅连接打开事件
	SubscribeConnOpenedEvent(subscriber vivid.Subscriber, options ...vivid.SubscribeOption)
}

func (s *ServerActor) SubscribeConnOpenedEvent(subscriber vivid.Subscriber, options ...vivid.SubscribeOption) {
	s.GetSystem().Subscribe(subscriber, ServerConnectionOpenedEvent{}, options...)
}
