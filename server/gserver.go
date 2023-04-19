package server

import (
	"github.com/panjf2000/gnet"
	"minotaur/utils/synchronization"
	"time"
)

type gServer struct {
	*Server
	connections *synchronization.Map[string, *Conn]
}

func (slf *gServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	slf.connections = synchronization.NewMap[string, *Conn]()
	return
}

func (slf *gServer) OnShutdown(server gnet.Server) {
	for k := range slf.connections.Map() {
		slf.connections.Delete(k)
	}
	slf.connections = nil
	return
}

func (slf *gServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	conn := newGNetConn(c)
	slf.connections.Set(c.RemoteAddr().String(), conn)
	slf.OnConnectionOpenedEvent(conn)
	return
}

func (slf *gServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	slf.OnConnectionClosedEvent(slf.connections.DeleteGet(c.RemoteAddr().String()))
	return
}

func (slf *gServer) PreWrite(c gnet.Conn) {
	return
}

func (slf *gServer) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (slf *gServer) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if conn, exist := slf.connections.GetExist(c.RemoteAddr().String()); exist {
		slf.Server.PushMessage(MessageTypePacket, conn, packet)
		return nil, gnet.None
	} else {
		return nil, gnet.Close
	}
}

func (slf *gServer) Tick() (delay time.Duration, action gnet.Action) {
	return
}
