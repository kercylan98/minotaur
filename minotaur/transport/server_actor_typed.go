package transport

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

type ServerActorTyped interface {
	// Launch 通过指定的网络启动服务器
	Launch(network Network)

	// Shutdown 关闭服务器
	Shutdown()

	// SubscribeConnOpenedEvent 订阅连接打开事件
	SubscribeConnOpenedEvent(subscribeId vivid.SubscribeId, subscriber vivid.Subscriber, handler func(ServerConnOpenedEvent), options ...vivid.SubscribeOption)

	// SubscribeConnClosedEvent 订阅连接关闭事件
	SubscribeConnClosedEvent(subscribeId vivid.SubscribeId, subscriber vivid.Subscriber, handler func(ServerConnClosedEvent), options ...vivid.SubscribeOption)
}

type ServerActorTypedImpl struct {
	ref vivid.ActorRef
}

func (s *ServerActorTypedImpl) Launch(network Network) {
	s.ref.Tell(ServerLaunchMessage{Network: network})
}

func (s *ServerActorTypedImpl) Shutdown() {
	s.ref.Tell(ServerShutdownMessage{})
}

func (s *ServerActorTypedImpl) SubscribeConnOpenedEvent(subscribeId vivid.SubscribeId, subscriber vivid.Subscriber, handler func(ServerConnOpenedEvent), options ...vivid.SubscribeOption) {
	s.ref.GetSystem().Subscribe(subscribeId, subscriber, handler, options...)
}

func (s *ServerActorTypedImpl) SubscribeConnClosedEvent(subscribeId vivid.SubscribeId, subscriber vivid.Subscriber, handler func(ServerConnClosedEvent), options ...vivid.SubscribeOption) {
	s.ref.GetSystem().Subscribe(subscribeId, subscriber, handler, options...)
}
