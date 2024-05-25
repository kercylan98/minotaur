package server

type (
	// LauncherServerEvent 启动服务器事件
	LauncherServerEvent struct {
	}

	// ShutdownServerEvent 关闭服务器事件
	ShutdownServerEvent struct {
	}

	// ConnectionOpenedEvent 连接打开事件
	ConnectionOpenedEvent struct {
	}

	// ConnectionClosedEvent 连接关闭事件
	ConnectionClosedEvent struct {
	}

	// ConnectionReceivedPacketEvent 接收到连接数据包事件
	ConnectionReceivedPacketEvent struct {
	}
)
