package client

import "github.com/kercylan98/minotaur/server"

type (
	WebsocketConnectionClosedEventHandle        func(conn *Websocket, err any)
	WebsocketConnectionOpenedEventHandle        func(conn *Websocket)
	WebsocketConnectionReceivePacketEventHandle func(conn *Websocket, packet server.Packet)
)

type websocketEvents struct {
	websocketConnectionClosedEventHandles        []WebsocketConnectionClosedEventHandle
	websocketConnectionOpenedEventHandles        []WebsocketConnectionOpenedEventHandle
	websocketConnectionReceivePacketEventHandles []WebsocketConnectionReceivePacketEventHandle
}

// RegWebsocketConnectionClosedEvent 注册连接关闭事件
func (slf *websocketEvents) RegWebsocketConnectionClosedEvent(handle WebsocketConnectionClosedEventHandle) {
	slf.websocketConnectionClosedEventHandles = append(slf.websocketConnectionClosedEventHandles, handle)
}

func (slf *websocketEvents) OnWebsocketConnectionClosedEvent(conn *Websocket, err any) {
	for _, handle := range slf.websocketConnectionClosedEventHandles {
		handle(conn, err)
	}
}

// RegWebsocketConnectionOpenedEvent 注册连接打开事件
func (slf *websocketEvents) RegWebsocketConnectionOpenedEvent(handle WebsocketConnectionOpenedEventHandle) {
	slf.websocketConnectionOpenedEventHandles = append(slf.websocketConnectionOpenedEventHandles, handle)
}

func (slf *websocketEvents) OnWebsocketConnectionOpenedEvent(conn *Websocket) {
	for _, handle := range slf.websocketConnectionOpenedEventHandles {
		handle(conn)
	}
}

// RegWebsocketConnectionReceivePacketEvent 注册连接接收数据包事件
func (slf *websocketEvents) RegWebsocketConnectionReceivePacketEvent(handle WebsocketConnectionReceivePacketEventHandle) {
	slf.websocketConnectionReceivePacketEventHandles = append(slf.websocketConnectionReceivePacketEventHandles, handle)
}

func (slf *websocketEvents) OnWebsocketConnectionReceivePacketEvent(conn *Websocket, packet server.Packet) {
	for _, handle := range slf.websocketConnectionReceivePacketEventHandles {
		handle(conn, packet)
	}
}
