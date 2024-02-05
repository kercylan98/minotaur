package server

import (
	"github.com/kercylan98/minotaur/server"
	"time"
)

type Service struct {
	srv *server.Server
}

func (s *Service) OnInit() {
	s.srv = server.New(server.NetworkNone, server.WithLimitLife(time.Second*3))
}

func (s *Service) OnPreload() {

}

func (s *Service) OnMount() {

}

func (s *Service) OnBlock() {
	if err := s.srv.RunNone(); err != nil {
		panic(err)
	}
}
