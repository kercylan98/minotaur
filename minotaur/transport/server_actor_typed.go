package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

type ServerActorTyped interface {
	// SubscribeConnOpenedEvent 订阅连接打开事件
	SubscribeConnOpenedEvent(subscriber vivid.Subscriber, options ...vivid.SubscribeOption)
}

type ServerActorTypedImpl struct {
	ServerActorRef vivid.ActorRef
}

func (s *ServerActorTypedImpl) SubscribeConnOpenedEvent(subscriber vivid.Subscriber, options ...vivid.SubscribeOption) {
	s.ServerActorRef.GetSystem().Subscribe(subscriber, ServerConnectionOpenedEvent{}, options...)
}
