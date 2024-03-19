package server

import "github.com/panjf2000/gnet/v2"

func NewServer(trafficker Trafficker) *Server {
	srv := &Server{
		trafficker: trafficker,
	}
	return srv
}

type Server struct {
	trafficker Trafficker
}

func (s *Server) Run(protoAddr string) (err error) {
	var handler *eventHandler
	handler, err = newEventHandler(new(Options), s.trafficker)
	return gnet.Run(handler, protoAddr)
}
