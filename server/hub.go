package server

import (
	"context"
	"github.com/kercylan98/minotaur/utils/collection"
	"sync"
)

type hub struct {
	connections map[string]*Conn // 所有连接

	register   chan *Conn        // 注册连接
	unregister chan string       // 注销连接
	broadcast  chan hubBroadcast // 广播消息

	botCount    int // 机器人数量
	onlineCount int // 在线人数

	chanMutex sync.RWMutex // 避免外界函数导致的并发问题

	closed bool
}

type hubBroadcast struct {
	packet []byte                // 广播的数据包
	filter func(conn *Conn) bool // 过滤掉返回 false 的连接
}

func (h *hub) run(ctx context.Context) {
	h.connections = make(map[string]*Conn)
	h.register = make(chan *Conn, DefaultConnHubBufferSize)
	h.unregister = make(chan string, DefaultConnHubBufferSize)
	h.broadcast = make(chan hubBroadcast, DefaultConnHubBufferSize)
	go func(ctx context.Context, h *hub) {
		for {
			select {
			case conn := <-h.register:
				h.onRegister(conn)
			case id := <-h.unregister:
				h.onUnregister(id)
			case packet := <-h.broadcast:
				h.onBroadcast(packet)
			case <-ctx.Done():
				h.chanMutex.Lock()
				close(h.register)
				close(h.unregister)
				h.closed = true
				h.chanMutex.Unlock()
				return

			}
		}
	}(ctx, h)
}

// registerConn 注册连接
func (h *hub) registerConn(conn *Conn) {
	select {
	case h.register <- conn:
	default:
		h.onRegister(conn)
	}
}

// unregisterConn 注销连接
func (h *hub) unregisterConn(id string) {
	select {
	case h.unregister <- id:
	default:
		h.onUnregister(id)
	}
}

// GetOnlineCount 获取在线人数
func (h *hub) GetOnlineCount() int {
	h.chanMutex.RLock()
	defer h.chanMutex.RUnlock()
	return h.onlineCount
}

// GetOnlineBotCount 获取在线机器人数量
func (h *hub) GetOnlineBotCount() int {
	h.chanMutex.RLock()
	defer h.chanMutex.RUnlock()
	return h.botCount
}

// IsOnline 是否在线
func (h *hub) IsOnline(id string) bool {
	h.chanMutex.RLock()
	_, exist := h.connections[id]
	h.chanMutex.RUnlock()
	return exist
}

// GetOnlineAll 获取所有在线连接
func (h *hub) GetOnlineAll() map[string]*Conn {
	h.chanMutex.RLock()
	cop := collection.CloneMap(h.connections)
	h.chanMutex.RUnlock()
	return cop
}

// GetOnline 获取在线连接
func (h *hub) GetOnline(id string) *Conn {
	h.chanMutex.RLock()
	conn := h.connections[id]
	h.chanMutex.RUnlock()
	return conn
}

// CloseConn 关闭连接
func (h *hub) CloseConn(id string) {
	h.chanMutex.RLock()
	conn := h.connections[id]
	h.chanMutex.RUnlock()
	if conn != nil {
		conn.Close()
	}
}

// Broadcast 广播消息
func (h *hub) Broadcast(packet []byte, filter ...func(conn *Conn) bool) {
	m := hubBroadcast{
		packet: packet,
	}
	if len(filter) > 0 {
		m.filter = filter[0]
	}
	select {
	case h.broadcast <- m:
	default:
		h.onBroadcast(m)
	}
}

func (h *hub) onRegister(conn *Conn) {
	h.chanMutex.Lock()
	if h.closed {
		conn.Close()
		return
	}
	h.connections[conn.GetID()] = conn
	h.onlineCount++
	if conn.IsBot() {
		h.botCount++
	}
	h.chanMutex.Unlock()
}

func (h *hub) onUnregister(id string) {
	h.chanMutex.Lock()
	if conn, ok := h.connections[id]; ok {
		h.onlineCount--
		delete(h.connections, conn.GetID())
		if conn.IsBot() {
			h.botCount--
		}
	}
	h.chanMutex.Unlock()
}

func (h *hub) onBroadcast(packet hubBroadcast) {
	h.chanMutex.RLock()
	defer h.chanMutex.RUnlock()
	for _, conn := range h.connections {
		if packet.filter != nil && !packet.filter(conn) {
			continue
		}
		conn.Write(packet.packet)
	}
}
