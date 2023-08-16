package client

import "github.com/kercylan98/minotaur/server"

type (
	ConnectionClosedEventHandle        func(conn *Websocket, err any)
	ConnectionOpenedEventHandle        func(conn *Websocket)
	ConnectionReceivePacketEventHandle func(conn *Websocket, packet server.Packet)
)

type websocketEvents struct {
	connectionClosedEventHandles        []ConnectionClosedEventHandle
	connectionOpenedEventHandles        []ConnectionOpenedEventHandle
	connectionReceivePacketEventHandles []ConnectionReceivePacketEventHandle
}

// RegConnectionClosedEvent 注册连接关闭事件
func (slf *websocketEvents) RegConnectionClosedEvent(handle ConnectionClosedEventHandle) {
	slf.connectionClosedEventHandles = append(slf.connectionClosedEventHandles, handle)
}

func (slf *websocketEvents) OnConnectionClosedEvent(conn *Websocket, err any) {
	for _, handle := range slf.connectionClosedEventHandles {
		handle(conn, err)
	}
}

// RegConnectionOpenedEvent 注册连接打开事件
func (slf *websocketEvents) RegConnectionOpenedEvent(handle ConnectionOpenedEventHandle) {
	slf.connectionOpenedEventHandles = append(slf.connectionOpenedEventHandles, handle)
}

func (slf *websocketEvents) OnConnectionOpenedEvent(conn *Websocket) {
	for _, handle := range slf.connectionOpenedEventHandles {
		handle(conn)
	}
}

// RegConnectionReceivePacketEvent 注册连接接收数据包事件
func (slf *websocketEvents) RegConnectionReceivePacketEvent(handle ConnectionReceivePacketEventHandle) {
	slf.connectionReceivePacketEventHandles = append(slf.connectionReceivePacketEventHandles, handle)
}

func (slf *websocketEvents) OnConnectionReceivePacketEvent(conn *Websocket, packet server.Packet) {
	for _, handle := range slf.connectionReceivePacketEventHandles {
		handle(conn, packet)
	}
}
