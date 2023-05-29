package game

import "github.com/kercylan98/minotaur/server"

// Player 玩家
type Player[ID comparable] interface {
	// UseConn 指定连接
	UseConn(conn *server.Conn)
	// Close 关闭玩家并且释放其资源
	Close()
}
