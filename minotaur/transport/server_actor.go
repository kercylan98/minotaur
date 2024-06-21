package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"net"
	"reflect"
)

func NewServerActor(system *vivid.ActorSystem, options ...*vivid.ActorOptions[*ServerActor]) ServerActorTyped {
	return vivid.ActorOfT[*ServerActor, ServerActorTyped](system, options...)
}

type ServerActor struct {
	vivid.ActorRef
	typed       ServerActorTyped
	network     Network
	connections map[net.Conn]ConnActorTyped
	restart     bool
}

func (s *ServerActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnRestart:
		s.restart = true
	case vivid.OnBoot:
		s.ActorRef = ctx
		s.onBoot(ctx)
		if s.restart {
			s.onServerLaunch(ctx, ServerLaunchMessage{Network: s.network})
		}
	case vivid.OnActorTyped[ServerActorTyped]:
		s.typed = m.Typed
	case vivid.OnTerminate:
		s.onServerTerminate(ctx)
	case ServerLaunchMessage:
		s.onServerLaunch(ctx, m)
	case ServerConnOpenedMessage:
		s.onServerConnOpened(ctx, m)
	case ServerConnClosedMessage:
		s.onServerConnClosed(ctx, m)
	}
}

func (s *ServerActor) onBoot(ctx vivid.MessageContext) {
	s.connections = make(map[net.Conn]ConnActorTyped)
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
		if err = s.network.Launch(ctx, s.typed); err != nil {
			panic(err)
		}
	}()
}

func (s *ServerActor) onServerTerminate(ctx vivid.MessageContext) {
	log.Info("Minotaur Server", log.String("", "shutdown"))
	ctx.Reply(s.network.Shutdown())
}

func (s *ServerActor) onServerConnOpened(ctx vivid.MessageContext, m ServerConnOpenedMessage) {
	// 创建连接 Actor 并绑定到服务器
	connTyped := vivid.ActorOfFT[*ConnActor, ConnActorTyped](ctx, func(options *vivid.ActorOptions[*ConnActor]) {
		options.
			WithName("conn-" + m.conn.RemoteAddr().String()).
			WithStopOnParentRestart(true).
			WithSupervisor(func(message, reason vivid.Message) vivid.Directive {
				return vivid.DirectiveStop
			})
	})
	s.connections[m.conn] = connTyped
	connTyped.Init(m.conn, m.writer)
	ctx.Reply(connTyped)

	// 发布连接打开事件
	ctx.GetSystem().Publish(ctx, ServerConnectionOpenedEvent{Conn: connTyped})
}

func (s *ServerActor) onServerConnClosed(ctx vivid.MessageContext, m ServerConnClosedMessage) {
	conn, ok := s.connections[m.conn]
	if ok {
		delete(s.connections, m.conn)
		conn.Stop(true)
	}
}

func (s *ServerActor) Attach(conn net.Conn, writer ConnWriter) ConnActorTyped {
	typed, _ := s.Ask(ServerConnOpenedMessage{conn, writer}).(ConnActorTyped)
	return typed
}

func (s *ServerActor) Detach(conn net.Conn) {
	s.Tell(ServerConnClosedMessage{conn})
}
