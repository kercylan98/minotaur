package server

import (
	"context"
	"github.com/kercylan98/minotaur/utils/super"
	"sync"
)

type Server interface {
	Run() error
	Shutdown() error
}

type server struct {
	ctx         *super.CancelContext
	networks    []Network
	connections *connections
}

func NewServer(networks ...Network) Server {
	srv := &server{
		ctx:      super.WithCancelContext(context.Background()),
		networks: networks,
	}
	srv.connections = new(connections).init(srv.ctx)
	return srv
}

func (s *server) Run() (err error) {
	for _, network := range s.networks {
		if err = network.OnSetup(s.ctx, s.connections); err != nil {
			return
		}
	}

	var group = new(sync.WaitGroup)
	for _, network := range s.networks {
		group.Add(1)
		go func(ctx *super.CancelContext, group *sync.WaitGroup, network Network) {
			defer group.Done()
			if err = network.OnRun(ctx); err != nil {
				panic(err)
			}
		}(s.ctx, group, network)
	}
	group.Wait()

	return
}

func (s *server) Shutdown() (err error) {
	defer s.ctx.Cancel()
	for _, network := range s.networks {
		if err = network.OnShutdown(); err != nil {
			return
		}
	}
	return
}
