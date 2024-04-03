package server

import (
	"net"
)

// ConnWriter 用于兼容不同 Network 的连接数据写入器
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

	// WriteContext 写入数据
	WriteContext(data []byte, context interface{}) error

	PushSyncMessage(handler func(srv Server, conn Conn))

	PushAsyncMessage(handler func(srv Server, conn Conn) error, callbacks ...func(srv Server, conn Conn, err error))
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
	conn   net.Conn   // 连接
	writer ConnWriter // 写入器
	actor  string     // Actor 名称
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

func (c *conn) WriteContext(data []byte, context interface{}) error {
	return c.writer(NewPacket(data).SetContext(context))
}

func (c *conn) PushSyncMessage(handler func(srv Server, conn Conn)) {
	if err := c.server.reactor.Dispatch(c.actor, SyncMessage(c.server, func(srv *server) {
		handler(srv, c)
	})); err != nil {
		panic(err)
	}
}

func (c *conn) PushAsyncMessage(handler func(srv Server, conn Conn) error, callbacks ...func(srv Server, conn Conn, err error)) {
	if err := c.server.reactor.Dispatch(c.actor, AsyncMessage(c.server, c.actor, func(srv *server) error {
		return handler(srv, c)
	}, func(srv *server, err error) {
		for _, callback := range callbacks {
			callback(srv, c, err)
		}
	})); err != nil {
		panic(err)
	}
}
