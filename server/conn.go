package server

import (
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"net"
	"strings"
	"sync"
	"time"
)

// newKcpConn 创建一个处理KCP的连接
func newKcpConn(server *Server, session *kcp.UDPSession) *Conn {
	c := &Conn{
		server:     server,
		remoteAddr: session.RemoteAddr(),
		ip:         session.RemoteAddr().String(),
		kcp:        session,
		write: func(data []byte) error {
			_, err := session.Write(data)
			return err
		},
		data: map[any]any{},
	}
	if index := strings.LastIndex(c.ip, ":"); index != -1 {
		c.ip = c.ip[0:index]
	}
	go c.writeLoop()
	return c
}

// newKcpConn 创建一个处理GNet的连接
func newGNetConn(server *Server, conn gnet.Conn) *Conn {
	c := &Conn{
		server:     server,
		remoteAddr: conn.RemoteAddr(),
		ip:         conn.RemoteAddr().String(),
		gn:         conn,
		write: func(data []byte) error {
			return conn.AsyncWrite(data)
		},
		data: map[any]any{},
	}
	if index := strings.LastIndex(c.ip, ":"); index != -1 {
		c.ip = c.ip[0:index]
	}
	go c.writeLoop()
	return c
}

// newKcpConn 创建一个处理WebSocket的连接
func newWebsocketConn(server *Server, ws *websocket.Conn, ip string) *Conn {
	c := &Conn{
		server:     server,
		remoteAddr: ws.RemoteAddr(),
		ip:         ip,
		ws:         ws,
		write: func(data []byte) error {
			return ws.WriteMessage(websocket.BinaryMessage, data)
		},
		data: map[any]any{},
	}
	go c.writeLoop()
	return c
}

// Conn 服务器连接
type Conn struct {
	server     *Server
	remoteAddr net.Addr
	ip         string
	ws         *websocket.Conn
	gn         gnet.Conn
	kcp        *kcp.UDPSession
	write      func(data []byte) error
	data       map[any]any
	mutex      sync.Mutex
	packetPool *synchronization.Pool[*connPacket]
	packets    []*connPacket
}

func (slf *Conn) RemoteAddr() net.Addr {
	return slf.remoteAddr
}

func (slf *Conn) GetID() string {
	return slf.remoteAddr.String()
}

func (slf *Conn) GetIP() string {
	return slf.ip
}

// Close 关闭连接
func (slf *Conn) Close() {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	if slf.ws != nil {
		_ = slf.ws.Close()
	} else if slf.gn != nil {
		_ = slf.gn.Close()
	} else if slf.kcp != nil {
		_ = slf.kcp.Close()
	}
	slf.write = nil
	if slf.packetPool != nil {
		slf.packetPool.Close()
	}
	slf.packetPool = nil
	slf.packets = nil
}

// SetData 设置连接数据
func (slf *Conn) SetData(key, value any) *Conn {
	slf.data[key] = value
	return slf
}

// GetData 获取连接数据
func (slf *Conn) GetData(key any) any {
	return slf.data[key]
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

// WriteString 向连接中写入字符串
//   - 通过转换为[]byte调用 *Conn.Write
func (slf *Conn) WriteString(data string, messageType ...int) {
	slf.Write([]byte(data), messageType...)
}

// Write 向连接中写入数据
//   - messageType: websocket模式中指定消息类型
func (slf *Conn) Write(data []byte, messageType ...int) {
	if slf.packetPool == nil {
		return
	}
	cp := slf.packetPool.Get()
	if len(messageType) > 0 {
		cp.websocketMessageType = messageType[0]
	}
	cp.packet = data
	slf.mutex.Lock()
	slf.packets = append(slf.packets, cp)
	slf.mutex.Unlock()
}

// writeLoop 写循环
func (slf *Conn) writeLoop() {
	slf.packetPool = synchronization.NewPool[*connPacket](10*1024,
		func() *connPacket {
			return &connPacket{}
		}, func(data *connPacket) {
			data.packet = nil
			data.websocketMessageType = 0
		},
	)
	defer func() {
		if err := recover(); err != nil {
			slf.Close()
		}
	}()
	for {
		slf.mutex.Lock()
		if slf.packetPool == nil {
			return
		}
		if len(slf.packets) == 0 {
			slf.mutex.Unlock()
			time.Sleep(50 * time.Millisecond)
			continue
		}
		packets := slf.packets[0:]
		slf.packets = slf.packets[0:0]
		slf.mutex.Unlock()
		for i := 0; i < len(packets); i++ {
			data := packets[i]
			if len(data.packet) == 0 {
				for _, packet := range packets {
					slf.packetPool.Release(packet)
				}
				slf.Close()
				return
			}
			var err error
			if slf.IsWebsocket() {
				if data.websocketMessageType <= 0 {
					data.websocketMessageType = slf.server.websocketWriteMessageType
				}
				err = slf.ws.WriteMessage(data.websocketMessageType, data.packet)
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
			slf.packetPool.Release(data)
			if err != nil {
				panic(err)
			}
		}
	}
}
