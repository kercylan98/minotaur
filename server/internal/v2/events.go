package server

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"reflect"
	"time"

	"github.com/kercylan98/minotaur/utils/collection/listings"
)

type (
	LaunchedEventHandler                func(srv Server, ip string, t time.Time)
	ShutdownEventHandler                func(srv Server)
	ConnectionOpenedEventHandler        func(srv Server, conn Conn)
	ConnectionClosedEventHandler        func(srv Server, conn Conn, err error)
	ConnectionReceivePacketEventHandler func(srv Server, conn Conn, packet Packet)
)

type Events interface {
	// RegisterLaunchedEvent 注册服务器启动事件，当服务器启动后将会触发该事件
	//  - 该事件将在系统级 Actor 中运行，该事件中阻塞会导致服务器启动延迟
	RegisterLaunchedEvent(handler LaunchedEventHandler, priority ...int)

	// RegisterShutdownEvent 注册服务器关闭事件，当服务器关闭时将会触发该事件，当该事件处理完毕后服务器将关闭
	//  - 该事件将在系统级 Actor 中运行，该事件中阻塞会导致服务器关闭延迟
	//  - 该事件未执行完毕前，服务器的一切均正常运行
	RegisterShutdownEvent(handler ShutdownEventHandler, priority ...int)

	// RegisterConnectionOpenedEvent 注册连接打开事件，当新连接创建完毕时将会触发该事件
	//  - 该事件将在系统级 Actor 中运行，不应执行阻塞操作
	RegisterConnectionOpenedEvent(handler ConnectionOpenedEventHandler, priority ...int)

	// RegisterConnectionClosedEvent 注册连接关闭事件，当连接关闭后将会触发该事件
	//  - 该事件将在系统级 Actor 中运行，不应执行阻塞操作
	RegisterConnectionClosedEvent(handler ConnectionClosedEventHandler, priority ...int)

	// RegisterConnectionReceivePacketEvent 注册连接接收数据包事件，当连接接收到数据包后将会触发该事件
	//  - 该事件将在连接的 Actor 中运行，不应执行阻塞操作
	RegisterConnectionReceivePacketEvent(handler ConnectionReceivePacketEventHandler, priority ...int)
}

type events struct {
	*server

	launchedEventHandlers                listings.SyncPrioritySlice[LaunchedEventHandler]
	shutdownEventHandlers                listings.SyncPrioritySlice[ShutdownEventHandler]
	connectionOpenedEventHandlers        listings.SyncPrioritySlice[ConnectionOpenedEventHandler]
	connectionClosedEventHandlers        listings.SyncPrioritySlice[ConnectionClosedEventHandler]
	connectionReceivePacketEventHandlers listings.SyncPrioritySlice[ConnectionReceivePacketEventHandler]
}

func (s *events) init(srv *server) *events {
	s.server = srv
	return s
}

func (s *events) RegisterLaunchedEvent(handler LaunchedEventHandler, priority ...int) {
	s.launchedEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onLaunched() {
	s.Options.getManyOptions(func(opt *Options) {
		opt.logger.Info("Minotaur Server", log.String("", "============================================================================"))
		opt.logger.Info("Minotaur Server", log.String("", "RunningInfo"), log.String("network", reflect.TypeOf(s.network).String()), log.String("listen", fmt.Sprintf("%s://%s%s", s.network.Schema(), s.server.state.Ip, s.network.Address())))
		opt.logger.Info("Minotaur Server", log.String("", "============================================================================"))
	})

	s.PushMessage(GenerateSystemSyncMessage(func(srv Server) {
		s.launchedEventHandlers.RangeValue(func(index int, value LaunchedEventHandler) bool {
			value(s.server, s.server.state.Ip, s.server.state.LaunchedAt)
			return true
		})
	}))
}

func (s *events) RegisterConnectionOpenedEvent(handler ConnectionOpenedEventHandler, priority ...int) {
	s.connectionOpenedEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionOpened(conn Conn) {
	s.PushMessage(GenerateSystemSyncMessage(func(srv Server) {
		s.connectionOpenedEventHandlers.RangeValue(func(index int, value ConnectionOpenedEventHandler) bool {
			value(s.server, conn)
			return true
		})
	}))
}

func (s *events) RegisterConnectionClosedEvent(handler ConnectionClosedEventHandler, priority ...int) {
	s.connectionClosedEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionClosed(conn Conn, err error) {
	s.PushMessage(GenerateSystemSyncMessage(func(srv Server) {
		s.connectionClosedEventHandlers.RangeValue(func(index int, value ConnectionClosedEventHandler) bool {
			value(s.server, conn, err)
			return true
		})
	}))
}

func (s *events) RegisterConnectionReceivePacketEvent(handler ConnectionReceivePacketEventHandler, priority ...int) {
	s.connectionReceivePacketEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionReceivePacket(conn *conn, packet Packet) {
	conn.getDispatchHandler()(GenerateConnSyncMessage(conn, func(srv Server, conn Conn) {
		s.connectionReceivePacketEventHandlers.RangeValue(func(index int, value ConnectionReceivePacketEventHandler) bool {
			value(s.server, conn, packet)
			return true
		})
	}))
}

func (s *events) RegisterShutdownEvent(handler ShutdownEventHandler, priority ...int) {
	s.shutdownEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onShutdown() {
	s.PushSystemMessage(GenerateSystemSyncMessage(func(srv Server) {
		s.shutdownEventHandlers.RangeValue(func(index int, value ShutdownEventHandler) bool {
			value(s.server)
			return true
		})
	}))
}
