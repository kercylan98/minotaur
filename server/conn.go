package server

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"net"
	"strings"
	"sync"
)

// newKcpConn 创建一个处理KCP的连接
func newKcpConn(server *Server, session *kcp.UDPSession) *Conn {
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
			packets:    make(chan *connPacket, 1024*10),
			mutex:      new(sync.Mutex),
			server:     server,
			remoteAddr: session.RemoteAddr(),
			ip:         session.RemoteAddr().String(),
			kcp:        session,
			data:       map[any]any{},
		},
	}
	if index := strings.LastIndex(c.ip, ":"); index != -1 {
		c.ip = c.ip[0:index]
	}
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go c.writeLoop(wait)
	wait.Wait()
	return c
}

// newKcpConn 创建一个处理GNet的连接
func newGNetConn(server *Server, conn gnet.Conn) *Conn {
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
			packets:    make(chan *connPacket, 1024*10),
			mutex:      new(sync.Mutex),
			server:     server,
			remoteAddr: conn.RemoteAddr(),
			ip:         conn.RemoteAddr().String(),
			gn:         conn,
			data:       map[any]any{},
		},
	}
	if index := strings.LastIndex(c.ip, ":"); index != -1 {
		c.ip = c.ip[0:index]
	}
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go c.writeLoop(wait)
	wait.Wait()
	return c
}

// newKcpConn 创建一个处理WebSocket的连接
func newWebsocketConn(server *Server, ws *websocket.Conn, ip string) *Conn {
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
			packets:    make(chan *connPacket, 1024*10),
			mutex:      new(sync.Mutex),
			server:     server,
			remoteAddr: ws.RemoteAddr(),
			ip:         ip,
			ws:         ws,
			data:       map[any]any{},
		},
	}
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go c.writeLoop(wait)
	wait.Wait()
	return c
}

// newGatewayConn 创建一个处理网关消息的连接
func newGatewayConn(conn *Conn, connId string) *Conn {
	c := &Conn{
		//ctx: server.ctx,
		connection: &connection{
			packets: make(chan *connPacket, 1024*10),
			mutex:   new(sync.Mutex),
			server:  conn.server,
			data:    map[any]any{},
		},
	}
	c.gw = func(packet []byte) {
		conn.Write(packet)
	}
	return c
}

// NewEmptyConn 创建一个适用于测试的空连接
func NewEmptyConn(server *Server) *Conn {
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
			packets:    make(chan *connPacket, 1024*10),
			mutex:      new(sync.Mutex),
			server:     server,
			remoteAddr: &net.TCPAddr{},
			ip:         "0.0.0.0:0",
			data:       map[any]any{},
		},
	}
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go c.writeLoop(wait)
	wait.Wait()
	return c
}

// Conn 服务器连接单次会话的包装
type Conn struct {
	*connection
	ctx context.Context
}

// connection 长久保持的连接
type connection struct {
	server     *Server
	mutex      *sync.Mutex
	close      sync.Once
	closed     bool
	remoteAddr net.Addr
	ip         string
	ws         *websocket.Conn
	gn         gnet.Conn
	kcp        *kcp.UDPSession
	gw         func(packet []byte)
	data       map[any]any
	packetPool *concurrent.Pool[*connPacket]
	packets    chan *connPacket
}

// IsEmpty 是否是空连接
func (slf *Conn) IsEmpty() bool {
	return slf.ws == nil && slf.gn == nil && slf.kcp == nil && slf.gw == nil
}

// RemoteAddr 获取远程地址
func (slf *Conn) RemoteAddr() net.Addr {
	return slf.remoteAddr
}

// GetID 获取连接ID
//   - 为远程地址的字符串形式
func (slf *Conn) GetID() string {
	return slf.remoteAddr.String()
}

// GetIP 获取连接IP
func (slf *Conn) GetIP() string {
	return slf.ip
}

// IsClosed 是否已经关闭
func (slf *Conn) IsClosed() bool {
	return slf.closed
}

// Close 关闭连接
func (slf *Conn) Close(err ...error) {
	slf.close.Do(func() {
		slf.mutex.Lock()
		defer slf.mutex.Unlock()
		slf.closed = true
		if slf.ws != nil {
			_ = slf.ws.Close()
		} else if slf.gn != nil {
			_ = slf.gn.Close()
		} else if slf.kcp != nil {
			_ = slf.kcp.Close()
		}
		if slf.packetPool != nil {
			slf.packetPool.Close()
		}
		slf.packetPool = nil
		if slf.packets != nil {
			close(slf.packets)
			slf.packets = nil
		}
		if len(err) > 0 {
			slf.server.OnConnectionClosedEvent(slf, err[0])
			return
		}
		slf.server.OnConnectionClosedEvent(slf, nil)
	})
}

// SetData 设置连接数据，该数据将在连接关闭前始终存在
func (slf *Conn) SetData(key, value any) *Conn {
	slf.data[key] = value
	return slf
}

// GetData 获取连接数据
func (slf *Conn) GetData(key any) any {
	return slf.data[key]
}

// SetMessageData 设置消息数据，该数据将在消息处理完成后释放
func (slf *Conn) SetMessageData(key, value any) *Conn {
	slf.ctx = context.WithValue(slf.ctx, key, value)
	return slf
}

// GetMessageData 获取消息数据
func (slf *Conn) GetMessageData(key any) any {
	return slf.ctx.Value(key)
}

// ReleaseData 释放数据
func (slf *Conn) ReleaseData() *Conn {
	for k := range slf.data {
		delete(slf.data, k)
	}
	return slf
}

// IsWebsocket 是否是websocket连接
func (slf *Conn) IsWebsocket() bool {
	return slf.server.network == NetworkWebsocket
}

// GetWST 获取websocket消息类型
func (slf *Conn) GetWST() int {
	wst, _ := slf.ctx.Value(contextKeyWST).(int)
	return wst
}

// SetWST 设置websocket消息类型
func (slf *Conn) SetWST(wst int) *Conn {
	slf.ctx = context.WithValue(slf.ctx, contextKeyWST, wst)
	return slf
}

// Write 向连接中写入数据
//   - messageType: websocket模式中指定消息类型
func (slf *Conn) Write(packet []byte, callback ...func(err error)) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	if slf.gw != nil {
		slf.gw(packet)
		return
	}
	packet = slf.server.OnConnectionWritePacketBeforeEvent(slf, packet)
	if slf.packetPool == nil || slf.packets == nil {
		return
	}
	cp := slf.packetPool.Get()
	cp.wst = slf.GetWST()
	cp.packet = packet
	if len(callback) > 0 {
		cp.callback = callback[0]
	}
	slf.packets <- cp
}

// writeLoop 写循环
func (slf *Conn) writeLoop(wait *sync.WaitGroup) {
	slf.packetPool = concurrent.NewPool[*connPacket](10*1024,
		func() *connPacket {
			return &connPacket{}
		}, func(data *connPacket) {
			data.wst = 0
			data.packet = nil
			data.callback = nil
		},
	)
	defer func() {
		if err := recover(); err != nil {
			slf.Close()
		}
	}()
	wait.Done()
	for packet := range slf.packets {

		data := packet
		var err error
		if slf.IsWebsocket() {
			err = slf.ws.WriteMessage(data.wst, data.packet)
		} else {
			if slf.gn != nil {
				switch slf.server.network {
				case NetworkUdp, NetworkUdp4, NetworkUdp6:
					err = slf.gn.SendTo(data.packet)
				default:
					err = slf.gn.AsyncWrite(data.packet)
				}
			} else if slf.kcp != nil {
				_, err = slf.kcp.Write(data.packet)
			}
		}
		callback := data.callback
		slf.packetPool.Release(data)

		if callback != nil {
			callback(err)
		}
		if err != nil {
			panic(err)
		}
	}
}
