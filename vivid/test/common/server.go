package common

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/network"
)

func NewServer(address string) *Server {
	return &Server{
		c:    make(chan []byte, 1024),
		addr: address,
	}
}

type Server struct {
	c    chan []byte
	addr string
}

func (s *Server) Run() error {
	srv := server.NewServer(network.Tcp(s.addr))
	srv.RegisterConnectionReceivePacketEvent(func(srv server.Server, conn server.Conn, packet server.Packet) {
		s.c <- packet.GetBytes()
	})
	return srv.Run()
}

func (s *Server) Shutdown() error {
	return nil
}

func (s *Server) C() <-chan []byte {
	return s.c
}
