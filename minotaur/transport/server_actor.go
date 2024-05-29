package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"net"
	"reflect"
)

type (
	// ServerLaunchMessage 服务器启动消息，服务器在收到该消息后将开始运行，运行过程是异步的
	ServerLaunchMessage struct {
		Network  Network      // 网络接口
		EventBus *pulse.Pulse // 事件总线
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
)

type (
	ServerConnOpenedEvent struct {
		*ConnActor
	}

	ServerConnClosedEvent struct {
		*ConnActor
	}
)

type ServerActor struct {
	network     Network
	pulse       *pulse.Pulse
	actor       vivid.ActorRef
	core        *serverCore
	connections map[net.Conn]*ConnActor
}

func (s *ServerActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		s.onPreStart(ctx)
	case vivid.OnDestroy:
		s.onServerShutdown(ctx, ServerShutdownMessage{})
	case ServerLaunchMessage:
		s.onServerLaunch(ctx, m)
	case ServerShutdownMessage:
		s.onServerShutdown(ctx, m)
	case ServerConnOpenedMessage:
		s.onServerConnOpened(ctx, m)
	case ServerConnClosedMessage:
		s.onServerConnClosed(ctx, m)
	}
}

func (s *ServerActor) onPreStart(ctx vivid.MessageContext) {
	s.connections = make(map[net.Conn]*ConnActor)
	s.actor = ctx.GetReceiver()
	s.core = new(serverCore).init(ctx.GetReceiver())
}

func (s *ServerActor) onServerLaunch(ctx vivid.MessageContext, m ServerLaunchMessage) {
	s.network = m.Network
	s.pulse = m.EventBus
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
	conn := newConn(s.pulse, ctx.GetContext(), m.conn, m.writer)
	s.connections[m.conn] = conn
	ctx.Reply(ConnCore(conn))
	s.pulse.Publish(s.actor, ServerConnOpenedEvent{ConnActor: conn})
}

func (s *ServerActor) onServerConnClosed(ctx vivid.MessageContext, m ServerConnClosedMessage) {
	conn, ok := s.connections[m.conn]
	if ok {
		delete(s.connections, m.conn)
		if conn.reader != nil {
			conn.reader.Tell(vivid.OnDestroy{})
		}
		if conn.writer != nil {
			conn.writer.Tell(vivid.OnDestroy{})
		}

		s.pulse.Publish(s.actor, ServerConnClosedEvent{ConnActor: conn})
	}
}
