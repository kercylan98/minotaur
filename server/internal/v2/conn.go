package server

import (
	"go.uber.org/atomic"
	"net"
)

// ConnWriter 用于兼容不同 Network 的连接数据写入器
type ConnWriter func(packet Packet) error

type Conn interface {
	// SetQueue 设置连接使用的消息队列名称
	SetQueue(queue string)

	// DelQueue 删除连接使用的消息队列，删除后连接将在系统队列执行消息
	DelQueue()

	// GetQueue 获取连接使用的消息队列名称
	GetQueue() string

	// WritePacket 写入一个 Packet
	WritePacket(packet Packet) error

	// Write 写入数据
	Write(data []byte) (n int, err error)

	// WriteBytes 写入数据
	WriteBytes(data []byte) error

	// WriteContext 写入数据
	WriteContext(data []byte, context interface{}) error
}

func newConn(srv *server, c net.Conn, connWriter ConnWriter) *conn {
	return &conn{
		server: srv,
		conn:   c,
		writer: connWriter,
	}
}

type conn struct {
	server *server
	conn   net.Conn               // 连接
	writer ConnWriter             // 写入器
	queue  atomic.Pointer[string] // Actor 名称
}

func (c *conn) SetQueue(queue string) {
	c.queue.Store(&queue)
}

func (c *conn) DelQueue() {
	c.queue.Store(nil)
}

func (c *conn) GetQueue() string {
	ident := c.queue.Load()
	if ident == nil {
		return ""
	}
	return *ident
}

func (c *conn) WritePacket(packet Packet) error {
	return c.writer(packet)
}

func (c *conn) Write(data []byte) (n int, err error) {
	return c.conn.Write(data)
}

func (c *conn) WriteBytes(data []byte) error {
	_, err := c.conn.Write(data)
	return err
}

func (c *conn) WriteContext(data []byte, context interface{}) error {
	return c.writer(NewPacket(data).SetContext(context))
}
