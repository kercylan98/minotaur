package builtin

import "minotaur/server"

func NewPlayer[ID comparable](id ID, conn *server.Conn) *Player[ID] {
	return &Player[ID]{
		id:   id,
		conn: conn,
	}
}

// Player 游戏玩家
type Player[ID comparable] struct {
	id   ID
	conn *server.Conn
}

func (slf *Player[ID]) GetID() ID {
	return slf.id
}

// Send 向该玩家发送数据
func (slf *Player[ID]) Send(packet []byte) error {
	return slf.conn.Write(packet)
}

// Close 关闭玩家
func (slf *Player[ID]) Close() {
	slf.conn.Close()
}
