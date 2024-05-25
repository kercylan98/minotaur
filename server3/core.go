package server

import "net"

type Core interface {
	BindConn(conn net.Conn, writer ConnWriter)
}

type _Core struct {
	srv *server
}

func (c *_Core) init(srv *server) *_Core {
	c.srv = srv
	return c
}

func (c *_Core) BindConn(netConn net.Conn, writer ConnWriter) {
	connection := &conn{
		conn:   netConn,
		writer: writer,
	}
	c.srv.actor.Tell(onConnectionOpenedMessage{connection})
}
