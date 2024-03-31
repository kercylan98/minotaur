package server

import "golang.org/x/net/context"

type Server interface {
	Run() error
	Shutdown() error
}

type server struct {
	*controller
	ctx     context.Context
	cancel  context.CancelFunc
	network Network
}

func NewServer(network Network) Server {
	srv := &server{
		network: network,
	}
	srv.ctx, srv.cancel = context.WithCancel(context.Background())
	srv.controller = new(controller).init(srv)
	return srv
}

func (s *server) Run() (err error) {
	if err = s.network.OnSetup(s.ctx, s); err != nil {
		return
	}

	if err = s.network.OnRun(); err != nil {
		panic(err)
	}
	return
}

func (s *server) Shutdown() (err error) {
	defer s.server.cancel()
	err = s.network.OnShutdown()
	return
}
