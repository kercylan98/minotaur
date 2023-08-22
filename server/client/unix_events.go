package client

import "github.com/kercylan98/minotaur/server"

type (
	UDSConnectionClosedEventHandle        func(conn *UnixDomainSocket, err any)
	UDSConnectionOpenedEventHandle        func(conn *UnixDomainSocket)
	UDSConnectionReceivePacketEventHandle func(conn *UnixDomainSocket, packet server.Packet)
)

type udsEvents struct {
	UDSConnectionClosedEventHandles        []UDSConnectionClosedEventHandle
	UDSConnectionOpenedEventHandles        []UDSConnectionOpenedEventHandle
	UDSConnectionReceivePacketEventHandles []UDSConnectionReceivePacketEventHandle
}

// RegUDSConnectionClosedEvent 注册连接关闭事件
func (slf *udsEvents) RegUDSConnectionClosedEvent(handle UDSConnectionClosedEventHandle) {
	slf.UDSConnectionClosedEventHandles = append(slf.UDSConnectionClosedEventHandles, handle)
}

func (slf *udsEvents) OnUDSConnectionClosedEvent(conn *UnixDomainSocket, err any) {
	for _, handle := range slf.UDSConnectionClosedEventHandles {
		handle(conn, err)
	}
}

// RegUDSConnectionOpenedEvent 注册连接打开事件
func (slf *udsEvents) RegUDSConnectionOpenedEvent(handle UDSConnectionOpenedEventHandle) {
	slf.UDSConnectionOpenedEventHandles = append(slf.UDSConnectionOpenedEventHandles, handle)
}

func (slf *udsEvents) OnUDSConnectionOpenedEvent(conn *UnixDomainSocket) {
	for _, handle := range slf.UDSConnectionOpenedEventHandles {
		handle(conn)
	}
}

// RegUDSConnectionReceivePacketEvent 注册连接接收数据包事件
func (slf *udsEvents) RegUDSConnectionReceivePacketEvent(handle UDSConnectionReceivePacketEventHandle) {
	slf.UDSConnectionReceivePacketEventHandles = append(slf.UDSConnectionReceivePacketEventHandles, handle)
}

func (slf *udsEvents) OnUDSConnectionReceivePacketEvent(conn *UnixDomainSocket, packet server.Packet) {
	for _, handle := range slf.UDSConnectionReceivePacketEventHandles {
		handle(conn, packet)
	}
}
