package transport

import (
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

func NewServer(system *vivid.ActorSystem, options ...*vivid.ActorOptions[*ServerActor]) Server {
	ref := vivid.ActorOf[*ServerActor](system, append(options, vivid.NewActorOptions[*ServerActor]())...)
	srv := Server{
		srvActor: ref,
	}
	return srv
}

type Server struct {
	srvActor vivid.ActorRef
	eventBus *pulse.Pulse
}

func (s *Server) Launch(pulse *pulse.Pulse, network Network) {
	s.srvActor.Tell(ServerLaunchMessage{
		Network:  network,
		EventBus: pulse,
	})
}
