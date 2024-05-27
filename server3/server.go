package server

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"github.com/kercylan98/minotaur/vivid"
	"net"
	"reflect"
)

func NewServer(system *vivid.ActorSystem, network Network, options ...*vivid.ActorOptions[*server]) Server {
	ref := vivid.ActorOf[*server](system, append(options, vivid.NewActorOptions[*server]().
		WithConstruct(func() *server {
			srv := &server{
				network:     network,
				connections: make(map[net.Conn]*conn),
			}
			srv.core = new(_Core).init(srv)
			return srv
		}()),
	)...)
	return &_Server{
		actor: ref,
	}
}

type server struct {
	core        *_Core
	actor       vivid.ActorRef
	network     Network
	connections map[net.Conn]*conn
}

func (s *server) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		s.actor = ctx.GetReceiver()
	case vivid.OnDestroy:
		s.onShutdown(ctx, onShutdownServerAskMessage{})
	case onLaunchServerTellMessage:
		s.onLaunch(ctx, m)
	case onShutdownServerAskMessage:
		s.onShutdown(ctx, m)
	case onConnectionOpenedAskMessage:
		s.onConnectionOpened(ctx, m)
	case onConnectionClosedTellMessage:
		s.onConnectionClosed(ctx, m)
	}
}

func (s *server) onLaunch(ctx vivid.MessageContext, m onLaunchServerTellMessage) {
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

func (s *server) onShutdown(ctx vivid.MessageContext, m onShutdownServerAskMessage) {
	log.Info("Minotaur Server", log.String("", "shutdown"))
	ctx.Reply(s.network.Shutdown())
}

func (s *server) onConnectionOpened(ctx vivid.MessageContext, m onConnectionOpenedAskMessage) {
	ref := vivid.ActorOf[*conn](ctx, vivid.NewActorOptions[*conn]().WithConstruct(m.conn))
	s.connections[m.conn.conn] = m.conn
	ctx.Reply(ref)
}

func (s *server) onConnectionClosed(ctx vivid.MessageContext, m onConnectionClosedTellMessage) {
	connection, ok := s.connections[m.conn]
	if ok {
		delete(s.connections, m.conn)
		connection.actor.Tell(vivid.OnDestroy{})
	}
}
