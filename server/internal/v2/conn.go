package server

import (
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	messageEvents "github.com/kercylan98/minotaur/toolkit/nexus/events"
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

	// WriteContext 通过指定上下文写入数据，该函数会创建一个新的 Packet 并设置上下文
	WriteContext(data []byte, context any) error

	// WriteWebSocketText 写入 WebSocket 文本数据
	WriteWebSocketText(data []byte) error

	// WriteWebSocketBinary 写入 WebSocket 二进制数据
	WriteWebSocketBinary(data []byte) error

	// PublishMessage 发布消息到指定 topic，该函数将被允许使用自定义的消息事件
	PublishMessage(topic string, event nexus.Event[int, string])

	// PublishSystemMessage 发布消息到系统队列，该函数将被允许使用自定义的消息事件
	PublishSystemMessage(event nexus.Event[int, string])

	// PublishQueueMessage 发布消息到当前连接队列，该函数将被允许使用自定义的消息事件
	PublishQueueMessage(event nexus.Event[int, string])

	// PublishSyncMessage 发布同步消息到指定 topic
	PublishSyncMessage(topic string, handler messageEvents.SynchronousHandler)

	// PublishAsyncMessage 发布异步消息到指定 topic，当包含多个 callback 时，仅首个生效
	PublishAsyncMessage(topic string, handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler)

	// PublishSystemSyncMessage 发布同步消息到系统队列
	PublishSystemSyncMessage(handler messageEvents.SynchronousHandler)

	// PublishSystemAsyncMessage 发布异步消息到系统队列
	PublishSystemAsyncMessage(handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler)

	// PublishQueueSyncMessage 发布同步消息到当前连接队列
	PublishQueueSyncMessage(handler messageEvents.SynchronousHandler)

	// PublishQueueAsyncMessage 发布异步消息到当前连接队列
	PublishQueueAsyncMessage(handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler)
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
	if c.conn == nil {
		return len(data), c.writer(NewPacket(data))
	}
	return c.conn.Write(data)
}

func (c *conn) WriteBytes(data []byte) error {
	if c.conn == nil {
		return c.writer(NewPacket(data))
	}
	_, err := c.conn.Write(data)
	return err
}

func (c *conn) WriteContext(data []byte, context any) error {
	return c.writer(NewPacket(data).SetContext(context))
}

func (c *conn) WriteWebSocketText(data []byte) error {
	return c.WriteContext(data, ws.OpText)
}

func (c *conn) WriteWebSocketBinary(data []byte) error {
	return c.WriteContext(data, ws.OpBinary)
}

func (c *conn) PublishMessage(topic string, event nexus.Event[int, string]) {
	c.server.PublishMessage(topic, event)
}

func (c *conn) PublishSystemMessage(event nexus.Event[int, string]) {
	c.server.PublishSystemMessage(event)
}

func (c *conn) PublishQueueMessage(event nexus.Event[int, string]) {
	c.server.PublishMessage(c.GetQueue(), event)
}

func (c *conn) PublishSyncMessage(topic string, handler messageEvents.SynchronousHandler) {
	c.server.PublishSyncMessage(topic, handler)
}

func (c *conn) PublishAsyncMessage(topic string, handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler) {
	c.server.PublishAsyncMessage(topic, handler, callback...)
}

func (c *conn) PublishSystemSyncMessage(handler messageEvents.SynchronousHandler) {
	c.server.PublishSystemSyncMessage(handler)
}

func (c *conn) PublishSystemAsyncMessage(handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler) {
	c.server.PublishSystemAsyncMessage(handler, callback...)
}

func (c *conn) PublishQueueSyncMessage(handler messageEvents.SynchronousHandler) {
	c.server.PublishSyncMessage(c.GetQueue(), handler)
}

func (c *conn) PublishQueueAsyncMessage(handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler) {
	c.server.PublishAsyncMessage(c.GetQueue(), handler, callback...)
}
