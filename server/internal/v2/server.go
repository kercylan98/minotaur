package server

import (
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"golang.org/x/net/context"
)

type Server interface {
	Events

	Run() error
	Shutdown() error
}

type server struct {
	*controller
	*events
	ctx     context.Context
	cancel  context.CancelFunc
	network Network
	reactor *reactor.Reactor[Message]
}

func NewServer(network Network) Server {
	srv := &server{
		network: network,
	}
	srv.ctx, srv.cancel = context.WithCancel(context.Background())
	srv.controller = new(controller).init(srv)
	srv.events = new(events).init(srv)
	srv.reactor = reactor.NewReactor[Message](1024*8, 1024, func(msg Message) {
		msg.Execute()
	}, nil)
	return srv
}

func (s *server) Run() (err error) {
	go s.reactor.Run()
	if err = s.network.OnSetup(s.ctx, s); err != nil {
		return
	}

	if err = s.network.OnRun(); err != nil {
		panic(err)
	}
	return
}

func (s *server) Shutdown() (err error) {
	defer s.cancel()
	err = s.network.OnShutdown()
	s.reactor.Close()
	return
}
