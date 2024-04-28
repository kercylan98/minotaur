package server

import (
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

// NewBot 创建一个机器人，目前仅支持 Socket 服务器
func NewBot(srv *Server, options ...BotOption) *Bot {
	if !srv.IsSocket() {
		panic(fmt.Errorf("server type[%s] is not socket", srv.network))
	}
	bot := &Bot{
		conn: newBotConn(srv),
	}
	for _, option := range options {
		option(bot)
	}
	return bot
}

type Bot struct {
	conn   *Conn
	joined atomic.Bool
}

// JoinServer 加入服务器
func (slf *Bot) JoinServer() {
	if slf.joined.Swap(true) {
		slf.conn.server.OnConnectionClosedEvent(slf.conn, nil)
	}
	slf.conn.server.OnConnectionOpenedEvent(slf.conn)
}

// LeaveServer 离开服务器
func (slf *Bot) LeaveServer() {
	if slf.joined.Swap(false) {
		slf.conn.server.OnConnectionClosedEvent(slf.conn, nil)
	}
}

// SetNetworkDelay 设置网络延迟和波动范围
//   - delay 延迟
//   - fluctuation 波动范围
func (slf *Bot) SetNetworkDelay(delay, fluctuation time.Duration) {
	slf.conn.delay = delay
	slf.conn.fluctuation = fluctuation
}

// SetWriter 设置写入器
func (slf *Bot) SetWriter(writer io.Writer) {
	slf.conn.botWriter.Store(&writer)
}

// SendPacket 发送数据包到服务器
func (slf *Bot) SendPacket(packet []byte) {
	if slf.conn.server.IsOnline(slf.conn.GetID()) {
		slf.conn.server.PushPacketMessage(slf.conn, 0, packet)
	}
}

// SendWSPacket 发送 WebSocket 数据包到服务器
func (slf *Bot) SendWSPacket(wst int, packet []byte) {
	if slf.conn.server.IsOnline(slf.conn.GetID()) {
		slf.conn.server.PushPacketMessage(slf.conn, wst, packet)
	}
}
