package stream

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/panjf2000/gnet/v2"
)

func newGNETConnWrapper(system *vivid.ActorSystem, conn gnet.Conn) *gnetConnWrapper {
	return &gnetConnWrapper{
		system: system,
		conn:   conn,
	}
}

type gnetConnWrapper struct {
	system *vivid.ActorSystem
	conn   gnet.Conn
}

func (g *gnetConnWrapper) Write(packet *Packet) error {
	return g.conn.AsyncWrite(packet.Data(), func(c gnet.Conn, err error) error {
		ref := g.conn.Context().(vivid.ActorRef)
		g.system.Tell(ref, err)
		return nil
	})
}

func (g *gnetConnWrapper) Close() error {
	return g.conn.Close()
}
