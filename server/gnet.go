package server

import (
	"bytes"
	"github.com/panjf2000/gnet"
	"time"
)

type gNet struct {
	*Server
}

func (slf *gNet) OnInitComplete(server gnet.Server) (action gnet.Action) {
	return
}

func (slf *gNet) OnShutdown(server gnet.Server) {
	return
}

func (slf *gNet) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	conn := newGNetConn(slf.Server, c)
	c.SetContext(conn)
	slf.OnConnectionOpenedEvent(conn)
	return
}

func (slf *gNet) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	conn := c.Context().(*Conn)
	conn.Close(err)
	return
}

func (slf *gNet) PreWrite(c gnet.Conn) {
	return
}

func (slf *gNet) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (slf *gNet) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	slf.Server.PushPacketMessage(c.Context().(*Conn), 0, bytes.Clone(packet))
	return nil, gnet.None
}

func (slf *gNet) Tick() (delay time.Duration, action gnet.Action) {
	return
}
