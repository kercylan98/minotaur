package server

import "net"

type Controller interface {
	Run() error
	Shutdown() error
}

type controller struct {
	*server
}

func (s *controller) init(srv *server) *controller {
	s.server = srv
	return s
}

func (s *controller) RegisterConn(conn net.Conn, writer ConnWriter) {

}

func (s *controller) UnRegisterConn() {

}
