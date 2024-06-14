package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"net"
	"reflect"
)

func NewServerActor(system *vivid.ActorSystem, options ...*vivid.ActorOptions[*ServerActor]) vivid.TypedActorRef[ServerActorTyped] {
	ref := vivid.ActorOf[*ServerActor](system, options...)
	return vivid.Typed[ServerActorTyped](ref, &ServerActorTypedImpl{
		ServerActorRef: ref,
	})
}

type ServerActor struct {
	ServerActorTyped
	network     Network
	connections map[net.Conn]vivid.TypedActorRef[ConnActorTyped]
	restart     bool
}

func (s *ServerActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnRestart:
		s.restart = true
	case vivid.OnBoot:
		s.onBoot(ctx)
		if s.restart {
			s.onServerLaunch(ctx, ServerLaunchMessage{Network: s.network})
		}
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
	}
}

func (s *ServerActor) onBoot(ctx vivid.MessageContext) {
	s.connections = make(map[net.Conn]vivid.TypedActorRef[ConnActorTyped])
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

	serverExpandTyped := vivid.Typed[ServerActorExpandTyped](ctx, &ServerActorExpandTypedImpl{
		ServerActorRef: ctx,
	})
	go func() {
		if err = s.network.Launch(ctx, serverExpandTyped); err != nil {
			panic(err)
		}
	}()
}

func (s *ServerActor) onServerShutdown(ctx vivid.MessageContext, m ServerShutdownMessage) {
	log.Info("Minotaur Server", log.String("", "shutdown"))
	ctx.Reply(s.network.Shutdown())
}

func (s *ServerActor) onServerConnOpened(ctx vivid.MessageContext, m ServerConnOpenedMessage) {
	// 创建连接 Actor 并绑定到服务器
	connRef := ctx.ActorOf(vivid.OfO(func(options *vivid.ActorOptions[*ConnActor]) {
		options.
			WithName("conn-" + m.conn.RemoteAddr().String()).
			WithStopOnParentRestart(true).
			WithSupervisor(func(message, reason vivid.Message) vivid.Directive {
				return vivid.DirectiveStop
			})
	}))
	var connActor *ConnActor
	connRef.Tell(ConnectionInitMessage{Conn: m.conn, Writer: m.writer, ActorHook: func(actor *ConnActor) {
		connActor = actor
	}}, vivid.WithInstantly(true)) // 由于是立即执行，可以直接获取写入器引用
	connTyped := connActor.Typed
	connExpandTyped := vivid.Typed[ConnActorExpandTyped](connRef, &ConnActorExpandTypedImpl{
		ConnActorRef: connRef,
	})

	s.connections[m.conn] = connTyped
	ctx.Reply(connExpandTyped)

	// 发布连接打开事件
	ctx.GetSystem().Publish(ctx, ServerConnectionOpenedEvent{Conn: connTyped})
}

func (s *ServerActor) onServerConnClosed(ctx vivid.MessageContext, m ServerConnClosedMessage) {
	conn, ok := s.connections[m.conn]
	if ok {
		delete(s.connections, m.conn)
		conn.Stop()
	}
}
