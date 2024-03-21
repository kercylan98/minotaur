package server

import (
	"context"
	"github.com/kercylan98/minotaur/utils/super"
)

type Server interface {
	Run() error
	Shutdown() error
}

type server struct {
	*networkCore
	ctx     *super.CancelContext
	network Network
}

func NewServer(network Network) Server {
	srv := &server{
		ctx:     super.WithCancelContext(context.Background()),
		network: network,
	}
	srv.networkCore = new(networkCore).init(srv)
	return srv
}

func (s *server) Run() (err error) {
	if err = s.network.OnSetup(s.ctx, s); err != nil {
		return
	}

	if err = s.network.OnRun(s.ctx); err != nil {
		panic(err)
	}
	return
}

func (s *server) Shutdown() (err error) {
	defer s.ctx.Cancel()
	err = s.network.OnShutdown()
	return
}
