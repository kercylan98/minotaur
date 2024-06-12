package transport

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type (
	ConnReceivePacketMessage struct {
		Packet
	}
)

func newConn(server vivid.ActorContext, c net.Conn, writer ConnWriter) *ConnActor {
	conn := &ConnActor{
		server:     server,
		conn:       c,
		connWriter: writer,
	}
	return conn
}

type Conn interface {
	// Write 向连接内写入数据包
	Write(packet Packet)
}

type ConnCore interface {
	Conn

	// React 处理数据包
	React(packet Packet)
}

type ConnActor struct {
	server     vivid.ActorContext
	conn       net.Conn
	connWriter ConnWriter
	reader     vivid.ActorRef
	writer     vivid.ActorRef
}

func (c *ConnActor) OnReceive(ctx vivid.MessageContext) {
	switch ctx.GetMessage().(type) {
	case vivid.OnBoot:
		c.reader = ctx.GetRef()
		c.writer = vivid.ActorOf[*ConnWriteActor](c.server, vivid.NewActorOptions[*ConnWriteActor]().
			WithName(fmt.Sprintf("conn-write-%s", c.conn.RemoteAddr().String())).
			WithConstruct(func() *ConnWriteActor {
				return &ConnWriteActor{
					conn:   c.conn,
					writer: c.connWriter,
				}
			}()))
	case ConnReceivePacketMessage:

	}
}

func (c *ConnActor) Write(packet Packet) {
	if c.writer == nil {
		panic("should not happen, writer is nil")
	}
	c.writer.Tell(connWriteMessage{packet})
}

func (c *ConnActor) React(packet Packet) {
	if c.reader == nil {
		panic("should not happen, reader is nil")
	}
	c.reader.Tell(ConnReceivePacketMessage{packet})
}

func (c *ConnActor) Close() {
	_ = c.conn.Close()
	c.reader.Tell(vivid.OnTerminate{})
	c.writer.Tell(vivid.OnTerminate{})
}
