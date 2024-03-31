package server

import "net"

type Controller interface {
	RegisterConnection(conn net.Conn, writer ConnWriter)
	EliminateConnection(conn net.Conn, err error)
	ReactPacket(conn net.Conn, packet Packet)
}

type controller struct {
	*server
	connections map[net.Conn]*conn
}

func (s *controller) init(srv *server) *controller {
	s.server = srv
	s.connections = make(map[net.Conn]*conn)
	return s
}

func (s *controller) RegisterConnection(conn net.Conn, writer ConnWriter) {
	if err := s.server.reactor.SystemDispatch(HandlerMessage(s.server, func(srv *server) {
		srv.connections[conn] = newConn(conn, writer)
	})); err != nil {
		panic(err)
	}
}

func (s *controller) EliminateConnection(conn net.Conn, err error) {
	if err := s.server.reactor.SystemDispatch(HandlerMessage(s.server, func(srv *server) {
		delete(srv.connections, conn)
	})); err != nil {
		panic(err)
	}
}

func (s *controller) ReactPacket(conn net.Conn, packet Packet) {
	if err := s.server.reactor.SystemDispatch(HandlerMessage(s.server, func(srv *server) {
		c, exist := srv.connections[conn]
		if !exist {
			return
		}

		if err := srv.reactor.Dispatch(c.GetActor(), HandlerMessage(srv, func(srv *server) {
			srv.events.onConnectionReceivePacket(c, packet)
		})); err != nil {
			panic(err)
		}
	})); err != nil {
		panic(err)
	}
}
