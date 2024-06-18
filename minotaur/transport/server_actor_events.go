package transport

type (
	// ServerConnectionOpenedEvent 当服务器接收到新连接并成功建立连接后，将触发该事件
	ServerConnectionOpenedEvent struct {
		Conn ConnActorTyped // 连接
	}
)
