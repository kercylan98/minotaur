package server

import (
	"github.com/panjf2000/gnet"
	"minotaur/utils/synchronization"
	"time"
)

type gNet struct {
	*Server
	connections *synchronization.Map[string, *Conn]
}

func (slf *gNet) OnInitComplete(server gnet.Server) (action gnet.Action) {
	slf.connections = synchronization.NewMap[string, *Conn]()
	return
}

func (slf *gNet) OnShutdown(server gnet.Server) {
	for k := range slf.connections.Map() {
		slf.connections.Delete(k)
	}
	slf.connections = nil
}

func (slf *gNet) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	conn := newGNetConn(c)
	slf.connections.Set(c.RemoteAddr().String(), conn)
	slf.OnConnectionOpenedEvent(conn)
	return
}

func (slf *gNet) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	slf.OnConnectionClosedEvent(slf.connections.DeleteGet(c.RemoteAddr().String()))
	return
}

func (slf *gNet) PreWrite(c gnet.Conn) {
	return
}

func (slf *gNet) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (slf *gNet) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if conn, exist := slf.connections.GetExist(c.RemoteAddr().String()); exist {
		slf.Server.PushMessage(MessageTypePacket, conn, packet)
		return nil, gnet.None
	} else {
		return nil, gnet.Close
	}
}

func (slf *gNet) Tick() (delay time.Duration, action gnet.Action) {
	return
}
