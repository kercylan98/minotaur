package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/panjf2000/gnet/v2"
	"net"
)

func NewConn(c gnet.Conn, ctx vivid.ActorContext, writer vivid.ActorRef) *Conn {
	return &Conn{
		ActorContext: ctx,
		conn:         c,
		writer:       writer,
	}
}

type Conn struct {
	vivid.ActorContext
	conn   gnet.Conn
	writer vivid.ActorRef
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) Close() {
	c.ActorContext.Terminate(c.ActorContext.Ref())
}

func (c *Conn) WritePacket(packet Packet) {
	c.ActorContext.Tell(c.writer, packet)
}

func (c *Conn) Write(bytes []byte) {
	c.WritePacket(NewPacket(bytes))
}

func (c *Conn) WriteString(str string) {
	c.Write([]byte(str))
}
