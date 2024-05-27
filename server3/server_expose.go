package server

import "github.com/kercylan98/minotaur/vivid"

type Server interface {
	Run()
}

type _Server struct {
	actor vivid.ActorRef
}

func (s *_Server) Run() {
	s.actor.Tell(onLaunchServerTellMessage{})
}
