package server

import (
	"bytes"
	"github.com/panjf2000/gnet"
	"time"
)

type gNet struct {
	*Server
	state chan<- error
}

func (g *gNet) OnInitComplete(server gnet.Server) (action gnet.Action) {
	if g.state != nil {
		g.state <- nil
		g.state = nil
	}
	return
}

func (g *gNet) OnShutdown(server gnet.Server) {
	return
}

func (g *gNet) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	conn := newGNetConn(g.Server, c)
	c.SetContext(conn)
	g.OnConnectionOpenedEvent(conn)
	return
}

func (g *gNet) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	conn := c.Context().(*Conn)
	conn.Close(err)
	return
}

func (g *gNet) PreWrite(c gnet.Conn) {
	return
}

func (g *gNet) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (g *gNet) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	g.Server.PushPacketMessage(c.Context().(*Conn), 0, bytes.Clone(packet))
	return nil, gnet.None
}

func (g *gNet) Tick() (delay time.Duration, action gnet.Action) {
	return
}
