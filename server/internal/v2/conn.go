package server

import (
	"go.uber.org/atomic"
	"net"
)

// ConnWriter 用于兼容不同 Network 的连接数据写入器
type ConnWriter func(packet Packet) error

type Conn interface {
	// SetActor 设置连接使用的 Actor 名称
	SetActor(actor string)

	// DelActor 删除连接使用的 Actor
	DelActor()

	// GetActor 获取连接使用的 Actor 名称及是否拥有 Actor 名称的状态
	GetActor() (string, bool)

	// WritePacket 写入一个 Packet
	WritePacket(packet Packet) error

	// Write 写入数据
	Write(data []byte) (n int, err error)

	// WriteBytes 写入数据
	WriteBytes(data []byte) error

	// WriteContext 写入数据
	WriteContext(data []byte, context interface{}) error

	// PushMessage 通过连接推送特定消息到队列中进行处理
	PushMessage(message Message)

	// PushSyncMessage 是 PushMessage 中对于 GenerateConnSyncMessage 的快捷方式
	PushSyncMessage(handler func(srv Server, conn Conn))

	// PushAsyncMessage 是 PushMessage 中对于 GenerateConnAsyncMessage 的快捷方式，当 callback 传入多个时，将仅有首个生效
	PushAsyncMessage(handler func(srv Server, conn Conn) error, callback ...func(srv Server, conn Conn, err error))
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
	actor  atomic.Pointer[string] // Actor 名称
}

func (c *conn) SetActor(actor string) {
	c.actor.Store(&actor)
}

func (c *conn) DelActor() {
	c.actor.Store(nil)
}

func (c *conn) GetActor() (string, bool) {
	ident := c.actor.Load()
	if ident == nil {
		return "", false
	}
	return *ident, true
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

func (c *conn) PushMessage(message Message) {
	c.getDispatchHandler()(message)
}

func (c *conn) PushSyncMessage(handler func(srv Server, conn Conn)) {
	c.PushMessage(GenerateConnSyncMessage(c, handler))
}

func (c *conn) PushAsyncMessage(handler func(srv Server, conn Conn) error, callback ...func(srv Server, conn Conn, err error)) {
	var cb func(srv Server, conn Conn, err error)
	if len(callback) > 0 {
		cb = callback[0]
	}
	c.PushMessage(GenerateConnAsyncMessage(c, handler, cb))
}

func (c *conn) getDispatchHandler() func(message Message) {
	var ident, exist = c.GetActor()
	return func(message Message) {
		if !exist {
			c.server.PushSystemMessage(message)
		} else {
			c.server.PushIdentMessage(ident, message)
		}
	}
}
