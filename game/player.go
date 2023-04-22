package game

// Player 玩家
type Player[ID comparable] interface {
	// GetID 用户玩家ID
	GetID() ID
	// Send 发送数据包
	Send(packet []byte) error
	// Close 关闭玩家并且释放其资源
	Close()
}
