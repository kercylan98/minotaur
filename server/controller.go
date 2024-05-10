package server

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/panjf2000/ants/v2"
	"net"
)

const (
	zombieConnTimeoutTask = "zct|t" // 僵尸连接超时任务
)

// Controller 控制器是暴露 Server 对用户非公开的接口信息，适用于功能性的拓展
type Controller interface {
	// GetServer 获取服务器
	GetServer() Server
	// RegisterConnection 注册连接
	RegisterConnection(conn net.Conn, writer ConnWriter, callback func(conn Conn, descriptor *ConnDescriptor))
	// EliminateConnection 消除连接
	EliminateConnection(conn net.Conn, err error)
	// ReactPacket 反应连接数据包
	ReactPacket(conn net.Conn, packet Packet)
	// GetAnts 获取服务器异步池
	GetAnts() *ants.Pool
	// OnConnectionAsyncWriteError 注册连接异步写入数据错误事件
	OnConnectionAsyncWriteError(conn Conn, packet Packet, err error)
	// GetServerLogger 获取服务器日志记录器
	GetServerLogger() *log.Logger
	// GetScheduler 获取调度器
	GetScheduler() *chrono.Scheduler
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

func (s *controller) RegisterConnection(conn net.Conn, writer ConnWriter, callback func(conn Conn, descriptor *ConnDescriptor)) {
	s.GetScheduler().RegisterAfterTask(zombieConnTimeoutTask+conn.RemoteAddr().String(), s.GetZombieConnectionDeadline(), func() {
		_ = conn.Close()
	})

	s.server.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context) {
		c := newConn(s.server, conn, writer)
		s.server.connections[conn] = c
		if callback != nil {
			callback(c, &c.descriptor)
		}
		s.events.onConnectionOpened(c)
	})
}

func (s *controller) EliminateConnection(conn net.Conn, err error) {
	s.GetScheduler().UnregisterTask(zombieConnTimeoutTask + conn.RemoteAddr().String())

	s.server.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context) {
		c, exist := s.server.connections[conn]
		if !exist {
			return
		}
		delete(s.server.connections, conn)
		s.server.events.onConnectionClosed(c, err)
		c.DelQueue()
		c.Close()
	})
}

func (s *controller) ReactPacket(conn net.Conn, packet Packet) {
	s.GetScheduler().RegisterAfterTask(zombieConnTimeoutTask+conn.RemoteAddr().String(), s.GetZombieConnectionDeadline(), func() {
		_ = conn.Close()
	})

	s.server.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context) {
		c, exist := s.server.connections[conn]
		if !exist {
			return
		}
		s.PublishSyncMessage(c.GetQueue(), func(ctx context.Context) {
			s.events.onConnectionReceivePacket(c, packet)
		})
	})
}

func (s *controller) GetServerLogger() *log.Logger {
	return s.server.Options.GetLogger()
}

func (s *controller) GetScheduler() *chrono.Scheduler {
	return s.server.scheduler
}
