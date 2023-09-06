package client

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"sync"
)

// NewClient 创建客户端
func NewClient(core Core) *Client {
	client := &Client{
		events: new(events),
		core:   core,
	}
	return client
}

// CloneClient 克隆客户端
func CloneClient(client *Client) *Client {
	return NewClient(client.core.Clone())
}

// Client 客户端
type Client struct {
	*events
	core       Core
	mutex      sync.Mutex
	packetPool *concurrent.Pool[*Packet]
	packets    chan *Packet

	accumulate   chan *Packet
	accumulation int // 积压消息数
}

func (slf *Client) Run() error {
	var runState = make(chan error)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				slf.Close(err.(error))
			}
		}()
		slf.core.Run(runState, slf.onReceive)
	}()
	err := <-runState
	if err != nil {
		slf.mutex.Lock()
		if slf.packetPool != nil {
			slf.packetPool.Close()
			slf.packetPool = nil
		}
		slf.mutex.Unlock()
		return err
	}
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go slf.writeLoop(wait)
	wait.Wait()
	slf.OnConnectionOpenedEvent(slf)
	return nil
}

// IsConnected 是否已连接
func (slf *Client) IsConnected() bool {
	return slf.packetPool != nil
}

// Close 关闭
func (slf *Client) Close(err ...error) {
	slf.mutex.Lock()
	var unlock bool
	defer func() {
		if !unlock {
			slf.mutex.Unlock()
		}
	}()
	slf.core.Close()
	if slf.packetPool != nil {
		slf.packetPool.Close()
		slf.packetPool = nil
	}
	if slf.packets != nil {
		close(slf.packets)
		slf.packets = nil
	}
	if slf.accumulate != nil {
		close(slf.accumulate)
		slf.accumulate = nil
	}
	slf.packets = nil
	unlock = true
	slf.mutex.Unlock()
	if len(err) > 0 {
		slf.OnConnectionClosedEvent(slf, err[0])
	} else {
		slf.OnConnectionClosedEvent(slf, nil)
	}
}

// WriteWS 向连接中写入指定 websocket 数据类型
//   - wst: websocket模式中指定消息类型
func (slf *Client) WriteWS(wst int, packet []byte, callback ...func(err error)) {
	slf.write(wst, packet, callback...)
}

// Write 向连接中写入数据
func (slf *Client) Write(packet []byte, callback ...func(err error)) {
	slf.write(0, packet, callback...)
}

// write 向连接中写入数据
//   - messageType: websocket模式中指定消息类型
func (slf *Client) write(wst int, packet []byte, callback ...func(err error)) {
	slf.mutex.Lock()
	if slf.packetPool == nil || slf.packets == nil {
		var p = &Packet{
			wst:  wst,
			data: packet,
		}
		if len(callback) > 0 {
			p.callback = callback[0]
		}
		if slf.accumulate == nil {
			slf.accumulate = make(chan *Packet, 1024*10)
		}
		slf.accumulate <- p
	} else {
		cp := slf.packetPool.Get()
		cp.wst = wst
		cp.data = packet
		if len(callback) > 0 {
			cp.callback = callback[0]
		}
		slf.packets <- cp
		slf.accumulation = len(slf.accumulate) + len(slf.packets)
	}
	slf.mutex.Unlock()
}

// writeLoop 写循环
func (slf *Client) writeLoop(wait *sync.WaitGroup) {
	slf.packets = make(chan *Packet, 1024*10)
	slf.packetPool = concurrent.NewPool[*Packet](10*1024,
		func() *Packet {
			return &Packet{}
		}, func(data *Packet) {
			data.wst = 0
			data.data = nil
			data.callback = nil
		},
	)
	go func() {
		for packet := range slf.accumulate {
			slf.packets <- packet
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			err, isErr := err.(error)
			if !isErr {
				err = fmt.Errorf("%v", err)
			}
			slf.Close(err)
		}
	}()
	wait.Done()

	for packet := range slf.packets {
		data := packet
		var err = slf.core.Write(data)
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

func (slf *Client) onReceive(wst int, packet []byte) {
	slf.OnConnectionReceivePacketEvent(slf, wst, packet)
}

// GetServerAddr 获取服务器地址
func (slf *Client) GetServerAddr() string {
	return slf.core.GetServerAddr()
}

// GetMessageAccumulationTotal 获取消息积压总数
func (slf *Client) GetMessageAccumulationTotal() int {
	return slf.accumulation
}
