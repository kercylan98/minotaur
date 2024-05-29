package transport

import (
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

type (
	ConnReceivePacketMessage struct {
		Packet
	}
)

func newConn(eventBus *pulse.Pulse, server vivid.ActorContext, c net.Conn, writer ConnWriter) *ConnActor {
	conn := &ConnActor{
		eventBus:   eventBus,
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
	eventBus   *pulse.Pulse
	server     vivid.ActorContext
	conn       net.Conn
	connWriter ConnWriter
	reader     vivid.ActorRef
	writer     vivid.ActorRef
}

func (c *ConnActor) OnReceive(ctx vivid.MessageContext) {
	switch ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		c.reader = ctx.GetReceiver()
		c.writer = vivid.ActorOf[*ConnWriteActor](c.server, vivid.NewActorOptions[*ConnWriteActor]().WithConstruct(func() *ConnWriteActor {
			return &ConnWriteActor{
				conn:   c.conn,
				writer: c.connWriter,
			}
		}()))
	}
}

func (c *ConnActor) Write(packet Packet) {
	if c.writer == nil {
		return
	}
	c.writer.Tell(connWriteMessage{packet})
}

func (c *ConnActor) React(packet Packet) {
	if c.reader == nil {
		return
	}
	c.reader.Tell(ConnReceivePacketMessage{packet})
}
