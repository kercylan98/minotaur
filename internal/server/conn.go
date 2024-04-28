package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/server/writeloop"
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/hub"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"io"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"time"
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
			openTime:   time.Now(),
		},
	}
	if index := strings.LastIndex(c.ip, ":"); index != -1 {
		c.ip = c.ip[0:index]
	}
	c.init()
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
			openTime:   time.Now(),
		},
	}
	if index := strings.LastIndex(c.ip, ":"); index != -1 {
		c.ip = c.ip[0:index]
	}
	c.init()
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
			openTime:   time.Now(),
		},
	}
	c.init()
	return c
}

// newBotConn 创建一个适用于测试等情况的机器人连接
func newBotConn(server *Server) *Conn {
	ip, port := random.NetIP(), random.Port()
	var writer io.Writer = os.Stdout
	c := &Conn{
		ctx: server.ctx,
		connection: &connection{
			server: server,
			remoteAddr: &net.TCPAddr{
				IP:   ip,
				Port: port,
				Zone: "",
			},
			ip:       fmt.Sprintf("BOT:%s:%d", ip.String(), port),
			data:     map[any]any{},
			openTime: time.Now(),
		},
	}
	c.botWriter.Store(&writer)
	c.init()
	return c
}

// Conn 服务器连接单次消息的包装
type Conn struct {
	*connection
	wst int
	ctx context.Context
}

// connection 长久保持的连接
type connection struct {
	server      *Server
	ticker      *timer.Ticker
	remoteAddr  net.Addr
	ip          string
	ws          *websocket.Conn
	gn          gnet.Conn
	kcp         *kcp.UDPSession
	gw          func(packet []byte)
	data        map[any]any
	closed      bool
	pool        *hub.ObjectPool[*connPacket]
	loop        writeloop.WriteLoop[*connPacket]
	mu          sync.Mutex
	openTime    time.Time
	delay       time.Duration
	fluctuation time.Duration
	botWriter   atomic.Pointer[io.Writer]
}

// Ticker 获取定时器
func (slf *Conn) Ticker() *timer.Ticker {
	return slf.ticker
}

// GetServer 获取服务器
func (slf *Conn) GetServer() *Server {
	return slf.server
}

// GetOpenTime 获取连接打开时间
func (slf *Conn) GetOpenTime() time.Time {
	return slf.openTime
}

// GetOnlineTime 获取连接在线时长
func (slf *Conn) GetOnlineTime() time.Duration {
	return time.Now().Sub(slf.openTime)
}

// GetWebsocketRequest 获取websocket请求
func (slf *Conn) GetWebsocketRequest() *http.Request {
	return slf.GetData(wsRequestKey).(*http.Request)
}

// IsBot 是否是机器人连接
func (slf *Conn) IsBot() bool {
	return slf != nil && slf.ws == nil && slf.gn == nil && slf.kcp == nil && slf.gw == nil
}

// RemoteAddr 获取远程地址
func (slf *Conn) RemoteAddr() net.Addr {
	return slf.remoteAddr
}

// GetID 获取连接ID
//   - 为远程地址的字符串形式
func (slf *Conn) GetID() string {
	if slf.IsBot() {
		return slf.ip
	}
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
	return collection.CloneMap(slf.data)
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

// GetWST 获取本次 websocket 消息类型
//   - 默认将与发送类型相同
func (slf *Conn) GetWST() int {
	return slf.wst
}

// SetWST 设置本次 websocket 消息类型
func (slf *Conn) SetWST(wst int) *Conn {
	slf.wst = wst
	return slf
}

// PushAsyncMessage 推送异步消息，该消息将通过 Server.PushShuntAsyncMessage 函数推送
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Conn) PushAsyncMessage(caller func() error, callback func(err error), mark ...log.Field) {
	slf.server.PushShuntAsyncMessage(slf, caller, callback, mark...)
}

// PushUniqueAsyncMessage 推送唯一异步消息，该消息将通过 Server.PushUniqueShuntAsyncMessage 函数推送
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
//   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息
func (slf *Conn) PushUniqueAsyncMessage(name string, caller func() error, callback func(err error), mark ...log.Field) {
	slf.server.PushUniqueShuntAsyncMessage(slf, name, caller, callback, mark...)
}

// Write 向连接中写入数据
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

func (slf *Conn) init() {
	if slf.server.ticker != nil && slf.server.connTickerSize > 0 {
		if slf.server.tickerAutonomy {
			slf.ticker = slf.server.tickerPool.GetTicker(slf.server.connTickerSize)
		} else {
			slf.ticker = slf.server.tickerPool.GetTicker(slf.server.connTickerSize, timer.WithCaller(func(name string, caller func()) {
				slf.server.PushShuntTickerMessage(slf, name, caller)
			}))
		}
	}
	slf.pool = hub.NewObjectPool[connPacket](
		func() *connPacket {
			return &connPacket{}
		}, func(data *connPacket) {
			data.wst = 0
			data.packet = nil
			data.callback = nil
		},
	)
	slf.loop = writeloop.NewChannel[*connPacket](slf.pool, slf.server.connWriteBufferSize, func(data *connPacket) error {
		if slf.server.runtime.packetWarnSize > 0 && len(data.packet) > slf.server.runtime.packetWarnSize {
			log.Warn("Conn.Put", log.String("State", "PacketWarn"), log.String("Reason", "PacketSize"), log.String("ID", slf.GetID()), log.Int("PacketSize", len(data.packet)))
		}
		var err error
		if slf.delay > 0 || slf.fluctuation > 0 {
			time.Sleep(random.Duration(int64(slf.delay-slf.fluctuation), int64(slf.delay+slf.fluctuation)))
			_, err = (*slf.botWriter.Load()).Write(data.packet)
			if data.callback != nil {
				data.callback(err)
			}
			return err
		}
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
		if slf.ticker != nil {
			slf.ticker.Release()
		}
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
	if slf.ticker != nil {
		slf.ticker.Release()
	}
	slf.loop.Close()
	slf.mu.Unlock()
	if len(err) > 0 {
		slf.server.OnConnectionClosedEvent(slf, err[0])
		return
	}
	slf.server.OnConnectionClosedEvent(slf, nil)
}
