package client

import (
	"bufio"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"net"
	"sync"
	"time"
)

// NewUnixDomainSocket 创建 unix domain socket 客户端
func NewUnixDomainSocket(addr string) *UnixDomainSocket {
	return &UnixDomainSocket{
		udsEvents: new(udsEvents),
		addr:      addr,
		data:      map[string]any{},
	}
}

// UnixDomainSocket unix domain socket 客户端
type UnixDomainSocket struct {
	*udsEvents
	conn net.Conn
	addr string
	data map[string]any

	mutex      sync.Mutex
	packetPool *concurrent.Pool[*Packet]
	packets    []*Packet

	accumulate []server.Packet
}

// Run 启动
func (slf *UnixDomainSocket) Run() error {
	c, err := net.Dial("unix", slf.addr)
	if err != nil {
		return err
	}
	slf.conn = c
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go slf.writeLoop(wait)
	wait.Wait()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				slf.Close()
				slf.OnUDSConnectionClosedEvent(slf, err)
			}
		}()
		slf.OnUDSConnectionOpenedEvent(slf)
		for slf.packetPool != nil {
			reader := bufio.NewReader(slf.conn)
			packet, readErr := reader.ReadBytes('\n')
			if readErr != nil {
				panic(readErr)
			}
			slf.OnUDSConnectionReceivePacketEvent(slf, server.NewPacket(packet))
		}
	}()
	return nil
}

// Close 关闭
func (slf *UnixDomainSocket) Close() {
	if slf.packetPool != nil {
		slf.packetPool.Close()
		slf.packetPool = nil
	}
	slf.packets = nil
}

// IsConnected 是否已连接
func (slf *UnixDomainSocket) IsConnected() bool {
	return slf.packetPool != nil
}

// GetData 获取数据
func (slf *UnixDomainSocket) GetData(key string) any {
	return slf.data[key]
}

// SetData 设置数据
func (slf *UnixDomainSocket) SetData(key string, value any) {
	slf.data[key] = value
}

// Write 向连接中写入数据
func (slf *UnixDomainSocket) Write(packet server.Packet) {
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
func (slf *UnixDomainSocket) writeLoop(wait *sync.WaitGroup) {
	slf.packetPool = concurrent.NewPool[*Packet](10*1024,
		func() *Packet {
			return &Packet{}
		}, func(data *Packet) {
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
			var _, err = slf.conn.Write(data.packet)
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
