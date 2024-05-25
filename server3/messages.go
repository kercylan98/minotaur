package server

type (
	onLaunchServerAskMessage struct {
		Network Network // 网络接口
	}

	// onShutdownServerAskMessage 用于关闭服务器的消息，该消息回复一个 error 类型的值
	onShutdownServerAskMessage struct {
	}

	onConnectionOpenedMessage struct {
		conn *conn
	}
)
