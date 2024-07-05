package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"net"
)

func NewConn(c net.Conn, system *vivid.ActorSystem, reader, writer vivid.ActorRef) *Conn {
	return &Conn{
		actorSystem: system,
		conn:        c,
		reader:      reader,
		writer:      writer,
	}
}

type Conn struct {
	actorSystem *vivid.ActorSystem
	conn        net.Conn
	reader      vivid.ActorRef
	writer      vivid.ActorRef
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) Close() {
	c.actorSystem.TerminateGracefully(c.reader)
}

func (c *Conn) WritePacket(packet Packet) {
	c.actorSystem.Context().Tell(c.writer, packet)
}

func (c *Conn) Write(bytes []byte) {
	c.WritePacket(NewPacket(bytes))
}

func (c *Conn) WriteString(str string) {
	c.Write([]byte(str))
}
