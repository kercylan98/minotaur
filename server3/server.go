package server

import (
	"github.com/kercylan98/minotaur/vivid"
)

func NewServer(system *vivid.ActorSystem, network Network, options ...*vivid.ActorOptions[*server]) Server {
	ref := vivid.ActorOf[*server](system, append(options, vivid.NewActorOptions[*server]().
		WithProps(func() *server {
			srv := &server{
				network: network,
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
	connections map[vivid.ActorId]vivid.ActorRef
}

func (s *server) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		s.actor = ctx.GetReceiver()
	case onLaunchServerAskMessage:
		s.onLaunch(ctx, m)
	case onShutdownServerAskMessage:
		s.onShutdown(ctx, m)
	case onConnectionOpenedMessage:
		s.onConnectionOpened(ctx, m)
	}
}

func (s *server) onLaunch(ctx vivid.MessageContext, m onLaunchServerAskMessage) {
	ctx.Reply(s.network.Launch(ctx, s.core))
}

func (s *server) onShutdown(ctx vivid.MessageContext, m onShutdownServerAskMessage) {
	ctx.Reply(s.network.Shutdown())
}

func (s *server) onConnectionOpened(ctx vivid.MessageContext, m onConnectionOpenedMessage) {
	connActor := vivid.ActorOf[*conn](ctx, vivid.NewActorOptions[*conn]().WithProps(m.conn))
	s.connections[vivid.GetActorIdByActorRef(connActor)] = connActor
}
