package server

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/collection/listings"
	"time"
)

type (
	LaunchedEventHandler                  func(srv Server, ip string, t time.Time)              // 服务器启动事件
	ShutdownEventHandler                  func(srv Server)                                      // 服务器关闭事件
	ConnectionOpenedEventHandler          func(srv Server, conn Conn)                           // 连接打开事件
	ConnectionOpenedAfterEventHandler     func(srv Server, conn Conn)                           // 连接打开后事件
	ConnectionClosedEventHandler          func(srv Server, conn Conn, err error)                // 连接关闭事件
	ConnectionReceivePacketEventHandler   func(srv Server, conn Conn, packet Packet)            // 连接接收数据包事件
	ConnectionAsyncWriteErrorEventHandler func(srv Server, conn Conn, packet Packet, err error) // 连接异步写入数据错误事件
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

	// RegisterConnectionOpenedAfterEvent 注册连接打开后事件，当新连接创建完毕后将会触发该事件
	//  - 该事件将在连接的 Actor 中运行，由于多个连接将使用同一个消息队列，不建议执行阻塞操作，当连接没有设置消息队列时，将在系统队列中执行
	RegisterConnectionOpenedAfterEvent(handler ConnectionOpenedAfterEventHandler, priority ...int)

	// RegisterConnectionClosedEvent 注册连接关闭事件，当连接关闭后将会触发该事件
	//  - 该事件将在系统级 Actor 中运行，不应执行阻塞操作
	RegisterConnectionClosedEvent(handler ConnectionClosedEventHandler, priority ...int)

	// RegisterConnectionReceivePacketEvent 注册连接接收数据包事件，当连接接收到数据包后将会触发该事件
	//  - 该事件将在连接的 Actor 中运行，不应执行阻塞操作
	RegisterConnectionReceivePacketEvent(handler ConnectionReceivePacketEventHandler, priority ...int)

	// RegisterConnectionAsyncWriteErrorEvent 注册连接异步写入数据错误事件，当连接异步写入数据失败时将会触发该事件
	//  - 该事件将在连接的 Actor 中运行，不应执行阻塞操作
	RegisterConnectionAsyncWriteErrorEvent(handler ConnectionAsyncWriteErrorEventHandler, priority ...int)
}

type events struct {
	*server

	launchedEventHandlers                  listings.SyncPrioritySlice[LaunchedEventHandler]
	shutdownEventHandlers                  listings.SyncPrioritySlice[ShutdownEventHandler]
	connectionOpenedEventHandlers          listings.SyncPrioritySlice[ConnectionOpenedEventHandler]
	connectionOpenedAfterEventHandlers     listings.SyncPrioritySlice[ConnectionOpenedAfterEventHandler]
	connectionClosedEventHandlers          listings.SyncPrioritySlice[ConnectionClosedEventHandler]
	connectionReceivePacketEventHandlers   listings.SyncPrioritySlice[ConnectionReceivePacketEventHandler]
	connectionAsyncWriteErrorEventHandlers listings.SyncPrioritySlice[ConnectionAsyncWriteErrorEventHandler]
}

func (s *events) init(srv *server) *events {
	s.server = srv
	return s
}

func (s *events) RegisterLaunchedEvent(handler LaunchedEventHandler, priority ...int) {
	s.launchedEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onLaunched() {
	s.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context) {
		s.launchedEventHandlers.RangeValue(func(index int, value LaunchedEventHandler) bool {
			value(s.server, s.server.state.Ip, s.server.state.LaunchedAt)
			return true
		})
	})
}

func (s *events) RegisterConnectionOpenedEvent(handler ConnectionOpenedEventHandler, priority ...int) {
	s.connectionOpenedEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionOpened(conn Conn) {
	s.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context) {
		s.connectionOpenedEventHandlers.RangeValue(func(index int, value ConnectionOpenedEventHandler) bool {
			value(s, conn)
			return true
		})
	})
}

func (s *events) RegisterConnectionOpenedAfterEvent(handler ConnectionOpenedAfterEventHandler, priority ...int) {
	s.connectionOpenedAfterEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionOpenedAfter(conn Conn) {
	s.PublishSyncMessage(conn.GetQueue(), func(ctx context.Context) {
		s.connectionOpenedAfterEventHandlers.RangeValue(func(index int, value ConnectionOpenedAfterEventHandler) bool {
			value(s, conn)
			return true
		})
	})
}

func (s *events) RegisterConnectionClosedEvent(handler ConnectionClosedEventHandler, priority ...int) {
	s.connectionClosedEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionClosed(conn Conn, err error) {
	s.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context) {
		s.connectionClosedEventHandlers.RangeValue(func(index int, value ConnectionClosedEventHandler) bool {
			value(s, conn, err)
			return true
		})
	})
}

func (s *events) RegisterConnectionReceivePacketEvent(handler ConnectionReceivePacketEventHandler, priority ...int) {
	s.connectionReceivePacketEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onConnectionReceivePacket(conn *conn, packet Packet) {
	s.PublishSyncMessage(conn.GetQueue(), func(ctx context.Context) {
		s.connectionReceivePacketEventHandlers.RangeValue(func(index int, value ConnectionReceivePacketEventHandler) bool {
			value(s, conn, packet)
			return true
		})
	})
}

func (s *events) RegisterConnectionAsyncWriteErrorEvent(handler ConnectionAsyncWriteErrorEventHandler, priority ...int) {
	s.connectionAsyncWriteErrorEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) OnConnectionAsyncWriteError(conn Conn, packet Packet, err error) {
	s.PublishSyncMessage(conn.GetQueue(), func(ctx context.Context) {
		s.connectionAsyncWriteErrorEventHandlers.RangeValue(func(index int, value ConnectionAsyncWriteErrorEventHandler) bool {
			value(s, conn, packet, err)
			return true
		})
	})
}

func (s *events) RegisterShutdownEvent(handler ShutdownEventHandler, priority ...int) {
	s.shutdownEventHandlers.AppendByOptionalPriority(handler, priority...)
}

func (s *events) onShutdown() {
	s.PublishSyncMessage(s.getSysQueue(), func(ctx context.Context) {
		s.shutdownEventHandlers.RangeValue(func(index int, value ShutdownEventHandler) bool {
			value(s)
			return true
		})
	})
}
