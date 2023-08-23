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
	slf.OnConnectionClosedEvent(c.Context().(*Conn), err)
	return
}

func (slf *gNet) PreWrite(c gnet.Conn) {
	return
}

func (slf *gNet) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (slf *gNet) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	PushPacketMessage(slf.Server, c.Context().(*Conn), 0, append(bytes.Clone(packet), 0))
	return nil, gnet.None
}

func (slf *gNet) Tick() (delay time.Duration, action gnet.Action) {
	return
}
