package server

// connPacket 连接包
type connPacket struct {
	websocketMessageType int             // websocket 消息类型
	packet               []byte          // 数据包
	callback             func(err error) // 回调函数
}
