package server

import (
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/vivids"
	"net"
)

type ConnectionWriter func(packet Packet) error

func newConn(c net.Conn, connWriter ConnectionWriter) *Conn {
	return &Conn{
		conn:   c,
		writer: connWriter,
	}
}

type (
	Conn struct {
		vivid.BasicActor
		conn   net.Conn         // 连接
		writer ConnectionWriter // 写入器
	}

	ConnTellMessageClose struct{}
)

func (c *Conn) OnPreStart(ctx vivids.ActorContext) (err error) {
	return
}

func (c *Conn) OnReceived(ctx vivids.MessageContext) (err error) {
	switch e := ctx.GetMessage().(type) {
	case NetworkConnectionClosedMessage:
		ctx.NotifyTerminated() // 已关闭，销毁 Actor 即可
	case NetworkConnectionReceivedMessage:
		return c.writer(e.Packet)
	}

	return
}

func (c *Conn) OnDestroy(ctx vivids.ActorContext) (err error) {
	return c.conn.Close() // 关闭连接
}
