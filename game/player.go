package game

// Player 玩家
type Player[ID comparable] interface {
	// GetID 用户玩家ID
	GetID() ID
	// Send 发送数据包
	//   - messageType: websocket模式中指定消息类型
	Send(packet []byte, messageType ...int)
	// SyncSend 同步发送数据包
	SyncSend(packet []byte, messageType ...int) error
	// Close 关闭玩家并且释放其资源
	Close()
}
