package client

import (
	"github.com/kercylan98/minotaur/utils/concurrent"
	"sync"
	"time"
)

// NewClient 创建客户端
func NewClient(core Core) *Client {
	client := &Client{
		events: new(events),
		core:   core,
	}
	return client
}

// Client 客户端
type Client struct {
	*events
	core       Core
	mutex      sync.Mutex
	packetPool *concurrent.Pool[*Packet]
	packets    []*Packet

	accumulate []*Packet
}

func (slf *Client) Run() error {
	var wait = new(sync.WaitGroup)
	wait.Add(1)
	go slf.writeLoop(wait)
	wait.Wait()
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
		slf.Close()
		return err
	}
	slf.OnConnectionOpenedEvent(slf)
	return nil
}

// IsConnected 是否已连接
func (slf *Client) IsConnected() bool {
	return slf.packetPool != nil
}

// Close 关闭
func (slf *Client) Close(err ...error) {
	slf.core.Close()
	if slf.packetPool != nil {
		slf.packetPool.Close()
		slf.packetPool = nil
	}
	slf.packets = nil
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
	cp := slf.packetPool.Get()
	cp.wst = wst
	cp.data = packet
	if len(callback) > 0 {
		cp.callback = callback[0]
	}
	if slf.packetPool == nil {
		slf.accumulate = append(slf.accumulate, cp)
		return
	}
	slf.mutex.Lock()
	slf.packets = append(slf.packets, cp)
	slf.mutex.Unlock()
}

// writeLoop 写循环
func (slf *Client) writeLoop(wait *sync.WaitGroup) {
	slf.packetPool = concurrent.NewPool[*Packet](10*1024,
		func() *Packet {
			return &Packet{}
		}, func(data *Packet) {
			data.wst = 0
			data.data = nil
			data.callback = nil
		},
	)
	slf.mutex.Lock()
	slf.packets = append(slf.packets, slf.accumulate...)
	slf.accumulate = nil
	slf.mutex.Unlock()
	defer func() {
		if err := recover(); err != nil {
			slf.Close(err.(error))
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
}

func (slf *Client) onReceive(wst int, packet []byte) {
	slf.OnConnectionReceivePacketEvent(slf, wst, packet)
}

// GetServerAddr 获取服务器地址
func (slf *Client) GetServerAddr() string {
	return slf.core.GetServerAddr()
}
