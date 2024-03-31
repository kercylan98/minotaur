package server

import (
	"net"
)

type ConnWriter func(packet Packet) error

type Conn interface {
	// SetActor 设置连接使用的 Actor 名称
	SetActor(actor string)

	// GetActor 获取连接使用的 Actor 名称
	GetActor() string

	// WritePacket 写入一个 Packet
	WritePacket(packet Packet) error

	// Write 写入数据
	Write(data []byte) (n int, err error)

	// WriteBytes 写入数据
	WriteBytes(data []byte) error
}

func newConn(c net.Conn, connWriter ConnWriter) *conn {
	return &conn{
		conn:   c,
		writer: connWriter,
	}
}

type conn struct {
	conn   net.Conn
	writer ConnWriter
	actor  string
}

func (c *conn) SetActor(actor string) {
	c.actor = actor
}

func (c *conn) GetActor() string {
	return c.actor
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
