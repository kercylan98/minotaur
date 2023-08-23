package client

type (
	ConnectionClosedEventHandle        func(conn *Client, err any)
	ConnectionOpenedEventHandle        func(conn *Client)
	ConnectionReceivePacketEventHandle func(conn *Client, wst int, packet []byte)
)

type events struct {
	ConnectionClosedEventHandles        []ConnectionClosedEventHandle
	ConnectionOpenedEventHandles        []ConnectionOpenedEventHandle
	ConnectionReceivePacketEventHandles []ConnectionReceivePacketEventHandle
}

// RegConnectionClosedEvent 注册连接关闭事件
func (slf *events) RegConnectionClosedEvent(handle ConnectionClosedEventHandle) {
	slf.ConnectionClosedEventHandles = append(slf.ConnectionClosedEventHandles, handle)
}

func (slf *events) OnConnectionClosedEvent(conn *Client, err any) {
	for _, handle := range slf.ConnectionClosedEventHandles {
		handle(conn, err)
	}
}

// RegConnectionOpenedEvent 注册连接打开事件
func (slf *events) RegConnectionOpenedEvent(handle ConnectionOpenedEventHandle) {
	slf.ConnectionOpenedEventHandles = append(slf.ConnectionOpenedEventHandles, handle)
}

func (slf *events) OnConnectionOpenedEvent(conn *Client) {
	for _, handle := range slf.ConnectionOpenedEventHandles {
		handle(conn)
	}
}

// RegConnectionReceivePacketEvent 注册连接接收数据包事件
func (slf *events) RegConnectionReceivePacketEvent(handle ConnectionReceivePacketEventHandle) {
	slf.ConnectionReceivePacketEventHandles = append(slf.ConnectionReceivePacketEventHandles, handle)
}

func (slf *events) OnConnectionReceivePacketEvent(conn *Client, wst int, packet []byte) {
	for _, handle := range slf.ConnectionReceivePacketEventHandles {
		handle(conn, wst, packet)
	}
}
