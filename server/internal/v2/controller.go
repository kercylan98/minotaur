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
	if err := s.server.reactor.DispatchWithSystem(SyncMessage(s.server, func(srv *server) {
		c := newConn(s.server, conn, writer)
		srv.connections[conn] = c
		s.events.onConnectionOpened(c)
	})); err != nil {
		panic(err)
	}
}

func (s *controller) EliminateConnection(conn net.Conn, err error) {
	if err := s.server.reactor.DispatchWithSystem(SyncMessage(s.server, func(srv *server) {
		c, exist := srv.connections[conn]
		if !exist {
			return
		}
		delete(srv.connections, conn)
		srv.events.onConnectionClosed(c, err)
	})); err != nil {
		panic(err)
	}
}

func (s *controller) ReactPacket(conn net.Conn, packet Packet) {
	if err := s.server.reactor.DispatchWithSystem(SyncMessage(s.server, func(srv *server) {
		c, exist := srv.connections[conn]
		if !exist {
			return
		}
		srv.events.onConnectionReceivePacket(c, packet)
	})); err != nil {
		panic(err)
	}
}
