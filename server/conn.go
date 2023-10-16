package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/server/writeloop"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
)

var wsRequestKey = fmt.Sprintf("WS:REQ:%s", strings.ToUpper(random.HostName()))

// newKcpConn 创建一个处理KCP的连接
func newKcpConn(server *Server, session *kcp.UDPSession) *Conn {
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
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
	c.writeLoop()
	return c
}

// newKcpConn 创建一个处理GNet的连接
func newGNetConn(server *Server, conn gnet.Conn) *Conn {
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
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
	c.writeLoop()
	return c
}

// newKcpConn 创建一个处理WebSocket的连接
func newWebsocketConn(server *Server, ws *websocket.Conn, ip string) *Conn {
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
			server:     server,
			remoteAddr: ws.RemoteAddr(),
			ip:         ip,
			ws:         ws,
			data:       map[any]any{},
		},
	}
	c.writeLoop()
	return c
}

// newGatewayConn 创建一个处理网关消息的连接
func newGatewayConn(conn *Conn, connId string) *Conn {
	c := &Conn{
		//ctx: server.ctx,
		connection: &connection{
			server: conn.server,
			data:   map[any]any{},
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
			server:     server,
			remoteAddr: &net.TCPAddr{},
			ip:         "0.0.0.0:0",
			data:       map[any]any{},
		},
	}
	c.writeLoop()
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
	remoteAddr net.Addr
	ip         string
	ws         *websocket.Conn
	gn         gnet.Conn
	kcp        *kcp.UDPSession
	gw         func(packet []byte)
	data       map[any]any
	closed     bool
	pool       *concurrent.Pool[*connPacket]
	loop       *writeloop.WriteLoop[*connPacket]
	mu         sync.Mutex
}

// GetWebsocketRequest 获取websocket请求
func (slf *Conn) GetWebsocketRequest() *http.Request {
	return slf.GetData(wsRequestKey).(*http.Request)
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
	slf.mu.Lock()
	defer slf.mu.Unlock()
	return slf.closed
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

// ViewData 查看只读的连接数据
func (slf *Conn) ViewData() map[any]any {
	return hash.Copy(slf.data)
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
	if slf.gw != nil {
		slf.gw(packet)
		return
	}
	packet = slf.server.OnConnectionWritePacketBeforeEvent(slf, packet)
	slf.mu.Lock()
	defer slf.mu.Unlock()
	if slf.closed {
		return
	}
	cp := slf.pool.Get()
	cp.wst = slf.GetWST()
	cp.packet = packet
	if len(callback) > 0 {
		cp.callback = callback[0]
	}
	slf.loop.Put(cp)
}

// writeLoop 写循环
func (slf *Conn) writeLoop() {
	slf.pool = concurrent.NewPool[*connPacket](10*1024,
		func() *connPacket {
			return &connPacket{}
		}, func(data *connPacket) {
			data.wst = 0
			data.packet = nil
			data.callback = nil
		},
	)
	slf.loop = writeloop.NewWriteLoop[*connPacket](slf.pool, func(data *connPacket) error {
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
		if data.callback != nil {
			data.callback(err)
		}
		return err
	}, func(err any) {
		slf.Close(errors.New(fmt.Sprint(err)))
	})
}

// Close 关闭连接
func (slf *Conn) Close(err ...error) {
	slf.mu.Lock()
	if slf.closed {
		slf.mu.Unlock()
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Error("Conn.Close", log.String("State", "Panic"), log.Any("Error", err))
			debug.PrintStack()
			slf.mu.Unlock()
		}
	}()
	slf.closed = true
	if slf.ws != nil {
		_ = slf.ws.Close()
	} else if slf.gn != nil {
		_ = slf.gn.Close()
	} else if slf.kcp != nil {
		_ = slf.kcp.Close()
	}
	slf.pool.Close()
	slf.loop.Close()
	slf.mu.Unlock()
	if len(err) > 0 {
		slf.server.OnConnectionClosedEvent(slf, err[0])
		return
	}
	slf.server.OnConnectionClosedEvent(slf, nil)
}
