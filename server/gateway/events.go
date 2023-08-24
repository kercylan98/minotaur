package gateway

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/slice"
)

type (
	ConnectionOpenedEventHandle             func(gateway *Gateway, conn *server.Conn)
	ConnectionClosedEventHandle             func(gateway *Gateway, conn *server.Conn)
	ConnectionReceivePacketEventHandle      func(gateway *Gateway, conn *server.Conn, packet []byte)
	EndpointConnectOpenedEventHandle        func(gateway *Gateway, endpoint *Endpoint)
	EndpointConnectClosedEventHandle        func(gateway *Gateway, endpoint *Endpoint)
	EndpointConnectReceivePacketEventHandle func(gateway *Gateway, endpoint *Endpoint, conn *server.Conn, packet []byte)
)

func newEvents() *events {
	return &events{
		connectionOpenedEventHandles:             slice.NewPriority[ConnectionOpenedEventHandle](),
		connectionClosedEventHandles:             slice.NewPriority[ConnectionClosedEventHandle](),
		connectionReceivePacketEventHandles:      slice.NewPriority[ConnectionReceivePacketEventHandle](),
		endpointConnectOpenedEventHandles:        slice.NewPriority[EndpointConnectOpenedEventHandle](),
		endpointConnectClosedEventHandles:        slice.NewPriority[EndpointConnectClosedEventHandle](),
		endpointConnectReceivePacketEventHandles: slice.NewPriority[EndpointConnectReceivePacketEventHandle](),
	}
}

type events struct {
	connectionOpenedEventHandles             *slice.Priority[ConnectionOpenedEventHandle]
	connectionClosedEventHandles             *slice.Priority[ConnectionClosedEventHandle]
	connectionReceivePacketEventHandles      *slice.Priority[ConnectionReceivePacketEventHandle]
	endpointConnectOpenedEventHandles        *slice.Priority[EndpointConnectOpenedEventHandle]
	endpointConnectClosedEventHandles        *slice.Priority[EndpointConnectClosedEventHandle]
	endpointConnectReceivePacketEventHandles *slice.Priority[EndpointConnectReceivePacketEventHandle]
}

// RegConnectionOpenedEventHandle 注册客户端连接打开事件处理函数
func (slf *events) RegConnectionOpenedEventHandle(handle ConnectionOpenedEventHandle, priority ...int) {
	slf.connectionOpenedEventHandles.Append(handle, slice.GetValue(priority, 0))
}

func (slf *events) OnConnectionOpenedEvent(gateway *Gateway, conn *server.Conn) {
	slf.connectionOpenedEventHandles.RangeValue(func(index int, value ConnectionOpenedEventHandle) bool {
		value(gateway, conn)
		return true
	})
}

// RegConnectionClosedEventHandle 注册客户端连接关闭事件处理函数
func (slf *events) RegConnectionClosedEventHandle(handle ConnectionClosedEventHandle, priority ...int) {
	slf.connectionClosedEventHandles.Append(handle, slice.GetValue(priority, 0))
}

func (slf *events) OnConnectionClosedEvent(gateway *Gateway, conn *server.Conn) {
	slf.connectionClosedEventHandles.RangeValue(func(index int, value ConnectionClosedEventHandle) bool {
		value(gateway, conn)
		return true
	})
}

// RegConnectionReceivePacketEventHandle 注册客户端连接接收数据包事件处理函数
func (slf *events) RegConnectionReceivePacketEventHandle(handle ConnectionReceivePacketEventHandle, priority ...int) {
	slf.connectionReceivePacketEventHandles.Append(handle, slice.GetValue(priority, 0))
}

func (slf *events) OnConnectionReceivePacketEvent(gateway *Gateway, conn *server.Conn, packet []byte) {
	slf.connectionReceivePacketEventHandles.RangeValue(func(index int, value ConnectionReceivePacketEventHandle) bool {
		value(gateway, conn, packet)
		return true
	})
}

// RegEndpointConnectOpenedEventHandle 注册端点连接打开事件处理函数
func (slf *events) RegEndpointConnectOpenedEventHandle(handle EndpointConnectOpenedEventHandle, priority ...int) {
	slf.endpointConnectOpenedEventHandles.Append(handle, slice.GetValue(priority, 0))
}

func (slf *events) OnEndpointConnectOpenedEvent(gateway *Gateway, endpoint *Endpoint) {
	slf.endpointConnectOpenedEventHandles.RangeValue(func(index int, value EndpointConnectOpenedEventHandle) bool {
		value(gateway, endpoint)
		return true
	})
}

// RegEndpointConnectClosedEventHandle 注册端点连接关闭事件处理函数
func (slf *events) RegEndpointConnectClosedEventHandle(handle EndpointConnectClosedEventHandle, priority ...int) {
	slf.endpointConnectClosedEventHandles.Append(handle, slice.GetValue(priority, 0))
}

func (slf *events) OnEndpointConnectClosedEvent(gateway *Gateway, endpoint *Endpoint) {
	slf.endpointConnectClosedEventHandles.RangeValue(func(index int, value EndpointConnectClosedEventHandle) bool {
		value(gateway, endpoint)
		return true
	})
}

// RegEndpointConnectReceivePacketEventHandle 注册端点连接接收数据包事件处理函数
func (slf *events) RegEndpointConnectReceivePacketEventHandle(handle EndpointConnectReceivePacketEventHandle, priority ...int) {
	slf.endpointConnectReceivePacketEventHandles.Append(handle, slice.GetValue(priority, 0))
}

func (slf *events) OnEndpointConnectReceivePacketEvent(gateway *Gateway, endpoint *Endpoint, conn *server.Conn, packet []byte) {
	slf.endpointConnectReceivePacketEventHandles.RangeValue(func(index int, value EndpointConnectReceivePacketEventHandle) bool {
		value(gateway, endpoint, conn, packet)
		return true
	})
}
