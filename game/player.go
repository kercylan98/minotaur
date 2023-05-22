package game

import "github.com/kercylan98/minotaur/server"

// Player 玩家
type Player[ID comparable] interface {
	// GetID 用户玩家ID
	GetID() ID
	// UseConn 指定连接
	UseConn(conn *server.Conn)
	// Send 发送数据包
	//   - messageType: websocket模式中指定消息类型
	Send(packet []byte, messageType ...int)
	// Close 关闭玩家并且释放其资源
	Close()
}
