package application

import (
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
)

func newConn() *Conn {
	return &Conn{}
}

type Conn struct {
}

func (c *Conn) OnOpened(ctx vivid.ActorContext, socket socket.Socket) {
}

func (c *Conn) OnClose(ctx vivid.ActorContext, socket socket.Socket, err error) {
}

func (c *Conn) OnPacket(ctx vivid.ActorContext, socket socket.Socket, packet *socket.Packet) {
	ctx.System().Logger().Info("receive packet", log.String("data", string(packet.GetData())))
	socket.WritePacket(packet)
}

func (c *Conn) OnReceive(ctx vivid.ActorContext) {

}
