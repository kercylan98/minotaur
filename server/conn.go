package server

import (
	"context"
	"errors"
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	messageEvents "github.com/kercylan98/minotaur/toolkit/nexus/events"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ConnStatusNormal ConnStatus = iota // 正常状态
	ConnStatusClosed                   // 关闭状态
)

// ConnStatus 连接状态
type ConnStatus = int32

// ConnWriter 用于兼容不同 Network 的连接数据写入器
type ConnWriter func(packet Packet) error

type Conn interface {
	// SetQueue 设置连接使用的消息队列名称
	SetQueue(queue string)

	// DelQueue 删除连接使用的消息队列，删除后连接将在系统队列执行消息
	DelQueue()

	// GetQueue 获取连接使用的消息队列名称
	GetQueue() string

	// AsyncWrite 异步的写入数据，通过该函数可以将读写进行分离，以便写入数据时不会对业务逻辑产生阻塞
	AsyncWrite(data []byte)

	// AsyncWritePacket 异步的写入数据，通过该函数可以将读写进行分离，以便写入数据时不会对业务逻辑产生阻塞
	AsyncWritePacket(packet Packet)

	// AsyncWriteContext 异步的通过指定上下文写入数据，该函数会创建一个新的 Packet 并设置上下文
	AsyncWriteContext(data []byte, context any) error

	// AsyncWriteWebSocketText 异步的写入 WebSocket 文本数据
	AsyncWriteWebSocketText(data []byte) error

	// AsyncWriteWebSocketBinary 异步的写入 WebSocket 二进制数据
	AsyncWriteWebSocketBinary(data []byte) error

	// Write 同步的写入数据
	Write(data []byte) (n int, err error)

	// WriteBytes 同步的写入数据
	WriteBytes(data []byte) error

	// WriteContext 同步的通过指定上下文写入数据，该函数会创建一个新的 Packet 并设置上下文
	WriteContext(data []byte, context any) error

	// WriteWebSocketText 同步的写入 WebSocket 文本数据
	WriteWebSocketText(data []byte) error

	// WriteWebSocketBinary 同步的写入 WebSocket 二进制数据
	WriteWebSocketBinary(data []byte) error

	// WritePacket 写入一个 Packet
	WritePacket(packet Packet) error

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

	// Close 关闭连接
	Close()

	// IsClosed 是否已关闭
	IsClosed() bool
}

func newConn(srv *server, c net.Conn, connWriter ConnWriter) *conn {
	ic := &conn{
		server:      srv,
		conn:        c,
		writer:      connWriter,
		state:       ConnStatusNormal,
		writeBuffer: buffer.NewRing[Packet](1024),
		writeCond:   sync.NewCond(new(sync.Mutex)),
	}
	ic.ctx, ic.cancel = context.WithCancel(srv.ctx)
	go ic.initWriteQueue()
	return ic
}

type conn struct {
	server *server                // 连接所属服务器
	conn   net.Conn               // 连接
	writer ConnWriter             // 写入器
	queue  atomic.Pointer[string] // Actor 名称
	ctx    context.Context        // 连接上下文
	cancel context.CancelFunc     // 连接上下文取消函数

	// 写分离队列相关
	state         int32                // 连接状态
	writeBuffer   *buffer.Ring[Packet] // 写入队列
	writeCond     *sync.Cond           // 写入条件
	leakDetection context.Context      // 泄漏检测
	leakCancel    context.CancelFunc   // 泄漏检测取消函数
}

// initWriteQueue 初始化写入队列
func (c *conn) initWriteQueue() {
	for {
		select {
		case <-c.ctx.Done():
			// 通知条件变化，退出交给 default 分支处理，避免直接 return 导致条件无法被通知，从而导致死锁
			c.writeCond.Signal()
		default:
			c.writeCond.L.Lock()
			for c.writeBuffer.Len() == 0 {
				if atomic.LoadInt32(&c.state) == ConnStatusClosed {
					// 连接已关闭，退出写入队列
					c.writeCond.L.Unlock()
					c.leakCancel()
					return
				}
				// 等待写入队列有数据
				c.writeCond.Wait()
			}
			packets := c.writeBuffer.ReadAll()
			c.writeCond.L.Unlock()

			for _, p := range packets {
				if err := c.writer(p); err != nil {
					if atomic.LoadInt32(&c.state) == ConnStatusClosed {
						break // 连接已关闭，退出写入循环，交给下一次循环处理关闭
					}
					c.server.events.OnConnectionAsyncWriteError(c, p, err)
				}
			}
		}
	}
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

func (c *conn) AsyncWrite(data []byte) {
	c.AsyncWritePacket(NewPacket(data))
}

func (c *conn) AsyncWritePacket(packet Packet) {
	if atomic.LoadInt32(&c.state) == ConnStatusClosed {
		return
	}
	c.writeCond.L.Lock()
	c.writeBuffer.Write(packet)
	c.writeCond.L.Unlock()
	c.writeCond.Signal()
}

func (c *conn) AsyncWriteContext(data []byte, context any) error {
	return c.WritePacket(NewPacket(data).SetContext(context))
}

func (c *conn) AsyncWriteWebSocketText(data []byte) error {
	return c.AsyncWriteContext(data, ws.OpText)
}

func (c *conn) AsyncWriteWebSocketBinary(data []byte) error {
	return c.AsyncWriteContext(data, ws.OpBinary)
}

func (c *conn) WritePacket(packet Packet) error {
	return c.writer(packet)
}

func (c *conn) Write(data []byte) (n int, err error) {
	return len(data), c.writer(NewPacket(data))
}

func (c *conn) WriteBytes(data []byte) error {
	return c.writer(NewPacket(data))
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

func (c *conn) Close() {
	if !atomic.CompareAndSwapInt32(&c.state, ConnStatusNormal, ConnStatusClosed) {
		return // 已经关闭
	}

	defer func(c *conn) {
		// 写入队列泄漏检测
		c.leakDetection, c.leakCancel = context.WithTimeout(context.Background(), time.Second*3)
		go func(ctx context.Context, c *conn) {
			<-ctx.Done()
			switch {
			case errors.Is(ctx.Err(), context.DeadlineExceeded):
				c.server.GetLogger().Warn("Close", log.String("addr", c.conn.RemoteAddr().String()), log.String("info", "connection write queue leak"))
			}
		}(c.leakDetection, c)

		// 在退出的时候，可能写入队列处于 default 分支，等待条件变化，此时需要通知条件变化，避免死锁
		c.cancel()
		c.writeCond.Signal()
	}(c)

	// 清理资源
	_ = c.conn.Close()
}

func (c *conn) IsClosed() bool {
	return atomic.LoadInt32(&c.state) == ConnStatusClosed
}
