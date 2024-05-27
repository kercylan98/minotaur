package server

import (
	"github.com/kercylan98/minotaur/vivid"
	"net"
)

// Core 是对于可自行实现功能提供的对外接口
type Core interface {
	// BindConn 用于绑定一个连接到服务器
	BindConn(conn net.Conn, writer ConnWriter) vivid.ActorRef
	// UnbindConn 用于解绑一个连接
	UnbindConn(conn net.Conn)
	// ProcessPacket 用于处理一个数据包
	ProcessPacket(conn vivid.ActorRef, packet Packet)
}

type _Core struct {
	srv *server
}

func (c *_Core) init(srv *server) *_Core {
	c.srv = srv
	return c
}

func (c *_Core) BindConn(netConn net.Conn, writer ConnWriter) vivid.ActorRef {
	connection := &conn{
		conn:   netConn,
		writer: writer,
	}
	reply := c.srv.actor.Ask(onConnectionOpenedAskMessage{connection})
	switch r := reply.(type) {
	case vivid.ActorRef:
		return r
	}
	return nil
}

func (c *_Core) UnbindConn(netConn net.Conn) {
	c.srv.actor.Tell(onConnectionClosedTellMessage{netConn})
}

func (c *_Core) ProcessPacket(conn vivid.ActorRef, packet Packet) {
	conn.Tell(onConnectionReceivedMessage{Packet: packet})
}
