package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"net"
)

func NewConn(c net.Conn, system *vivid.ActorSystem, ref vivid.ActorRef) *Conn {
	return &Conn{
		actorSystem: system,
		conn:        c,
		ref:         ref,
	}
}

type Conn struct {
	actorSystem *vivid.ActorSystem
	conn        net.Conn
	ref         vivid.ActorRef
}

func (c *Conn) Ref() vivid.ActorRef {
	return c.ref
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) Close() {
	c.actorSystem.TerminateGracefully(c.ref)
}

func (c *Conn) WritePacket(packet Packet) {
	c.actorSystem.Context().Tell(c.ref, packet)
}

func (c *Conn) Write(bytes []byte) {
	c.WritePacket(NewPacket(bytes))
}

func (c *Conn) WriteString(str string) {
	c.Write([]byte(str))
}
