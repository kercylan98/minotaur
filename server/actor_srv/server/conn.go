package server

import (
	"errors"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/vivids"
	"net"
)

type ConnectionWriter func(packet Packet) error

type ConnProps struct {
	Conn        net.Conn
	Writer      ConnectionWriter
	MessageChan chan Packet
}

type Conn struct {
	vivid.BasicActor
	ConnProps
}

func (c *Conn) OnPreStart(ctx vivids.ActorContext) (err error) {
	var ok bool
	c.ConnProps, ok = ctx.GetProps().(ConnProps)
	if !ok {
		return errors.New("invalid ConnProps")
	}

	return
}

func (c *Conn) OnReceived(ctx vivids.MessageContext) (err error) {
	switch e := ctx.GetMessage().(type) {
	case error:
		ctx.NotifyTerminated() // 已关闭，销毁 Actor 即可
	case NetworkConnectionReceivedMessage:
		return c.ConnProps.Writer(e.Packet)
	}

	return
}

func (c *Conn) OnDestroy(ctx vivids.ActorContext) (err error) {
	return c.ConnProps.Conn.Close() // 关闭连接
}
