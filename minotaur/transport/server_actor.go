package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"net"
	"reflect"
)

type (
	// ServerLaunchMessage 服务器启动消息，服务器在收到该消息后将开始运行，运行过程是异步的
	ServerLaunchMessage struct {
		Network Network // 网络接口
	}

	// ServerShutdownMessage 服务器关闭消息，服务器在收到该消息后将停止运行
	ServerShutdownMessage struct{}

	ServerConnOpenedMessage struct {
		conn   net.Conn
		writer ConnWriter
	}

	ServerConnClosedMessage struct {
		conn net.Conn
	}

	ServerSubscribeConnOpenedMessage struct {
		SubscribeId vivid.SubscribeId
		Handler     func(ctx vivid.MessageContext, event ServerConnOpenedEvent)
		Options     []vivid.SubscribeOption
	}

	ServerSubscribeConnClosedMessage struct {
		SubscribeId vivid.SubscribeId
		Handler     func(ctx vivid.MessageContext, event ServerConnClosedEvent)
		Options     []vivid.SubscribeOption
	}
)

type (
	ServerConnOpenedEvent struct {
		*ConnActor
	}

	ServerConnClosedEvent struct {
		*ConnActor
	}
)

func NewServerActor(system *vivid.ActorSystem, options ...*vivid.ActorOptions[*ServerActor]) vivid.TypedActorRef[ServerActorTyped] {
	ref := vivid.ActorOf[*ServerActor](system, options...)
	return vivid.Typed[ServerActorTyped](ref, &ServerActorTypedImpl{
		ref: ref,
	})
}

type ServerActor struct {
	ServerActorTyped
	network     Network
	core        *serverCore
	connections map[net.Conn]*ConnActor
}

func (s *ServerActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnBoot:
		s.onBoot(ctx)
	case vivid.OnTerminate:
		s.onServerShutdown(ctx, ServerShutdownMessage{})
	case ServerLaunchMessage:
		s.onServerLaunch(ctx, m)
	case ServerShutdownMessage:
		s.onServerShutdown(ctx, m)
	case ServerConnOpenedMessage:
		s.onServerConnOpened(ctx, m)
	case ServerConnClosedMessage:
		s.onServerConnClosed(ctx, m)
	case ServerSubscribeConnOpenedMessage:
		ctx.Become(vivid.BehaviorOf[ServerConnOpenedEvent](m.Handler))
		ctx.GetSystem().Subscribe(m.SubscribeId, ctx, ServerConnOpenedEvent{})
	case ServerSubscribeConnClosedMessage:
		ctx.Become(vivid.BehaviorOf[ServerConnClosedEvent](m.Handler))
		ctx.GetSystem().Subscribe(m.SubscribeId, ctx, ServerConnClosedEvent{})
	}
}

func (s *ServerActor) onBoot(ctx vivid.MessageContext) {
	s.connections = make(map[net.Conn]*ConnActor)
	s.core = new(serverCore).init(ctx.GetRef())
}

func (s *ServerActor) onServerLaunch(ctx vivid.MessageContext, m ServerLaunchMessage) {
	s.network = m.Network
	ip, err := network.IP()
	if err != nil {
		panic(err)
	}

	log.Info("Minotaur Server", log.String("", "============================================================================"))
	log.Info("Minotaur Server", log.String("", "RunningInfo"), log.String("network", reflect.TypeOf(s.network).String()), log.String("listen", fmt.Sprintf("%s://%s%s", s.network.Schema(), ip.String(), s.network.Address())))
	log.Info("Minotaur Server", log.String("", "============================================================================"))

	go func() {
		if err = s.network.Launch(ctx, s.core); err != nil {
			panic(err)
		}
	}()
}

func (s *ServerActor) onServerShutdown(ctx vivid.MessageContext, m ServerShutdownMessage) {
	log.Info("Minotaur Server", log.String("", "shutdown"))
	ctx.Reply(s.network.Shutdown())
}

func (s *ServerActor) onServerConnOpened(ctx vivid.MessageContext, m ServerConnOpenedMessage) {
	conn := newConn(ctx.GetContext(), m.conn, m.writer)
	vivid.ActorOfI(ctx, conn, func(options *vivid.ActorOptions[*ConnActor]) {
		options.WithName(fmt.Sprintf("conn-%s", m.conn.RemoteAddr().String()))
	})
	s.connections[m.conn] = conn
	ctx.Reply(ConnCore(conn))
	ctx.GetSystem().Publish(ctx, ServerConnOpenedEvent{ConnActor: conn})
}

func (s *ServerActor) onServerConnClosed(ctx vivid.MessageContext, m ServerConnClosedMessage) {
	conn, ok := s.connections[m.conn]
	if ok {
		delete(s.connections, m.conn)
		if conn.reader != nil {
			conn.reader.Tell(vivid.OnTerminate{})
		}
		if conn.writer != nil {
			conn.writer.Tell(vivid.OnTerminate{})
		}

		ctx.GetSystem().Publish(ctx, ServerConnClosedEvent{ConnActor: conn})
	}
}
