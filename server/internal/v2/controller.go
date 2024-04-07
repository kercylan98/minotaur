package server

import (
	"github.com/kercylan98/minotaur/utils/log"
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
	// PushSystemMessage 推送系统消息
	PushSystemMessage(message Message, errorHandlers ...func(err error))
	// PushIdentMessage 推送标识消息
	PushIdentMessage(ident string, message Message, errorHandlers ...func(err error))
	// MessageErrProcess 消息错误处理
	MessageErrProcess(message Message, err error)
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

func (s *controller) MessageErrProcess(message Message, err error) {
	if err == nil {
		return
	}
	if s.server.messageErrorHandler != nil {
		s.server.messageErrorHandler(s.server, message, err)
	} else {
		s.server.GetLogger().Error("Server", log.Err(err))
	}
}

func (s *controller) GetAnts() *ants.Pool {
	return s.server.ants
}

func (s *controller) PushSystemMessage(message Message, errorHandlers ...func(err error)) {
	if err := s.server.reactor.SystemDispatch(message); err != nil {
		for _, f := range errorHandlers {
			f(err)
		}
		s.MessageErrProcess(message, err)
	}
}

func (s *controller) PushIdentMessage(ident string, message Message, errorHandlers ...func(err error)) {
	if err := s.server.reactor.IdentDispatch(ident, message); err != nil {
		for _, f := range errorHandlers {
			f(err)
		}
		s.MessageErrProcess(message, err)
	}
}

func (s *controller) RegisterConnection(conn net.Conn, writer ConnWriter) {
	s.PushSystemMessage(GenerateSystemSyncMessage(func(srv Server) {
		c := newConn(s.server, conn, writer)
		s.server.connections[conn] = c
		s.events.onConnectionOpened(c)
	}))
}

func (s *controller) EliminateConnection(conn net.Conn, err error) {
	s.PushSystemMessage(GenerateSystemSyncMessage(func(srv Server) {
		c, exist := s.server.connections[conn]
		if !exist {
			return
		}
		delete(s.server.connections, conn)
		s.server.events.onConnectionClosed(c, err)
	}))
}

func (s *controller) ReactPacket(conn net.Conn, packet Packet) {
	s.PushSystemMessage(GenerateSystemSyncMessage(func(srv Server) {
		c, exist := s.server.connections[conn]
		if !exist {
			return
		}
		ident, exist := c.GetActor()
		if !exist {
			s.PushSystemMessage(GenerateSystemSyncMessage(func(srv Server) {
				s.events.onConnectionReceivePacket(c, packet)
			}))
		} else {
			s.PushIdentMessage(ident, GenerateSystemSyncMessage(func(srv Server) {
				s.events.onConnectionReceivePacket(c, packet)
			}))
		}
	}))
}
