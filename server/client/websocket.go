package client

import (
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"sync"
	"time"
)

// NewWebsocket 创建 websocket 客户端
func NewWebsocket(addr string) *Websocket {
	client := &Websocket{
		websocketEvents: new(websocketEvents),
		addr:            addr,
		data:            map[string]any{},
	}
	return client
}

// Websocket websocket 客户端
type Websocket struct {
	*websocketEvents
	conn *websocket.Conn
	addr string
	data map[string]any

	mutex      sync.Mutex
	packetPool *concurrent.Pool[*websocketPacket]
	packets    []*websocketPacket

	accumulate []server.Packet
}

// Run 启动
func (slf *Websocket) Run() error {
	ws, _, err := websocket.DefaultDialer.Dial(slf.addr, nil)
	if err != nil {
		return err
	}
	slf.conn = ws
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go slf.writeLoop(wait)
	wait.Wait()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				slf.Close()
				slf.OnConnectionClosedEvent(slf, err)
			}
		}()
		slf.OnConnectionOpenedEvent(slf)
		for slf.packetPool != nil {
			messageType, packet, readErr := ws.ReadMessage()
			if readErr != nil {
				panic(readErr)
			}
			slf.OnConnectionReceivePacketEvent(slf, server.NewWSPacket(messageType, packet))
		}
	}()
	return nil
}

// Close 关闭
func (slf *Websocket) Close() {
	if slf.packetPool != nil {
		slf.packetPool.Close()
		slf.packetPool = nil
	}
	slf.packets = nil
}

// IsConnected 是否已连接
func (slf *Websocket) IsConnected() bool {
	return slf.packetPool != nil
}

// GetData 获取数据
func (slf *Websocket) GetData(key string) any {
	return slf.data[key]
}

// SetData 设置数据
func (slf *Websocket) SetData(key string, value any) {
	slf.data[key] = value
}

// Write 向连接中写入数据
//   - messageType: websocket模式中指定消息类型
func (slf *Websocket) Write(packet server.Packet) {
	if slf.packetPool == nil {
		slf.accumulate = append(slf.accumulate, packet)
		return
	}
	cp := slf.packetPool.Get()
	cp.websocketMessageType = packet.WebsocketType
	cp.packet = packet.Data
	slf.mutex.Lock()
	slf.packets = append(slf.packets, cp)
	slf.mutex.Unlock()
}

// writeLoop 写循环
func (slf *Websocket) writeLoop(wait *sync.WaitGroup) {
	slf.packetPool = concurrent.NewPool[*websocketPacket](10*1024,
		func() *websocketPacket {
			return &websocketPacket{}
		}, func(data *websocketPacket) {
			data.packet = nil
			data.websocketMessageType = 0
			data.callback = nil
		},
	)
	slf.mutex.Lock()
	for _, packet := range slf.accumulate {
		cp := slf.packetPool.Get()
		cp.websocketMessageType = packet.WebsocketType
		cp.packet = packet.Data
		slf.packets = append(slf.packets, cp)
	}
	slf.accumulate = nil
	slf.mutex.Unlock()
	defer func() {
		if err := recover(); err != nil {
			slf.Close()
		}
	}()
	wait.Done()
	for {
		slf.mutex.Lock()
		if slf.packetPool == nil {
			slf.mutex.Unlock()
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
			var err = slf.conn.WriteMessage(data.websocketMessageType, data.packet)
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
}
