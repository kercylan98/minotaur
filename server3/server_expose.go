package server

import "github.com/kercylan98/minotaur/vivid"

type Server interface {
	Run() error
}

type _Server struct {
	actor vivid.ActorRef
}

func (s *_Server) Run() error {
	err, _ := s.actor.Ask(onLaunchServerAskMessage{}).(error)
	return err
}
