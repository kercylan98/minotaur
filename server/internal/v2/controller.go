package server

import (
	"context"
	"github.com/panjf2000/ants/v2"
	"net"
)

// Controller 控制器是暴露 Server 对用户非公开的接口信息，适用于功能性的拓展
type Controller interface {
	// GetServer 获取服务器
	GetServer() Server
	// RegisterConnection 注册连接
	RegisterConnection(conn net.Conn, writer ConnWriter)
	// EliminateConnection 消除连接
	EliminateConnection(conn net.Conn, err error)
	// ReactPacket 反应连接数据包
	ReactPacket(conn net.Conn, packet Packet)
	// GetAnts 获取服务器异步池
	GetAnts() *ants.Pool
}

type controller struct {
	*server
	connections map[net.Conn]*conn
}

func (s *controller) init(srv *server) *controller {
	s.server = srv
	s.connections = make(map[net.Conn]*conn)
	return s
}

func (s *controller) GetServer() Server {
	return s.server
}

func (s *controller) GetAnts() *ants.Pool {
	return s.server.ants
}

func (s *controller) RegisterConnection(conn net.Conn, writer ConnWriter) {
	s.server.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context, srv Server) {
		c := newConn(s.server, conn, writer)
		s.server.connections[conn] = c
		s.events.onConnectionOpened(c)
	})
}

func (s *controller) EliminateConnection(conn net.Conn, err error) {
	s.server.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context, srv Server) {
		c, exist := s.server.connections[conn]
		if !exist {
			return
		}
		delete(s.server.connections, conn)
		s.server.events.onConnectionClosed(c, err)
	})
}

func (s *controller) ReactPacket(conn net.Conn, packet Packet) {
	s.server.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context, srv Server) {
		c, exist := s.server.connections[conn]
		if !exist {
			return
		}
		s.PublishSyncMessage(c.GetActor(), func(ctx context.Context, srv Server) {
			s.events.onConnectionReceivePacket(c, packet)
		})
	})
}
