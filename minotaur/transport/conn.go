package transport

import (
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net"
)

func newConn(eventBus *pulse.Pulse, server vivid.ActorContext, c net.Conn, writer ConnWriter) *conn {
	conn := &conn{}
	conn.reader = vivid.ActorOf[*ConnReadActor](server, vivid.NewActorOptions[*ConnReadActor]().WithConstruct(func() *ConnReadActor {
		return &ConnReadActor{
			conn:     c,
			eventBus: eventBus,
		}
	}()))
	conn.writer = vivid.ActorOf[*ConnWriteActor](server, vivid.NewActorOptions[*ConnWriteActor]().WithConstruct(func() *ConnWriteActor {
		return &ConnWriteActor{
			conn:   c,
			writer: writer,
		}
	}()))
	return conn
}

type Conn interface {
	Write(packet Packet)
}

type ConnCore interface {
	Conn

	// React 处理数据包
	React(packet Packet)
}

type conn struct {
	reader vivid.ActorRef
	writer vivid.ActorRef
}

func (c *conn) Write(packet Packet) {
	c.writer.Tell(connWriteMessage{packet})
}

func (c *conn) React(packet Packet) {
	c.reader.Tell(connReceivePacketMessage{packet})
}
