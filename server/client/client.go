package client

import (
	"github.com/kercylan98/minotaur/utils/concurrent"
	"sync"
)

// NewClient 创建客户端
func NewClient(core Core) *Client {
	client := &Client{
		cond:   sync.NewCond(&sync.Mutex{}),
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
	cond       *sync.Cond
	packetPool *concurrent.Pool[*Packet]
	packets    []*Packet

	accumulate   []*Packet
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
		slf.cond.L.Lock()
		if slf.packetPool != nil {
			slf.packetPool.Close()
			slf.packetPool = nil
		}
		slf.accumulate = append(slf.accumulate, slf.packets...)
		slf.packets = nil
		slf.cond.L.Unlock()
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
	if slf.packetPool == nil {
		var p = &Packet{
			wst:  wst,
			data: packet,
		}
		if len(callback) > 0 {
			p.callback = callback[0]
		}
		slf.cond.L.Lock()
		slf.accumulate = append(slf.accumulate, p)
		slf.accumulation = len(slf.accumulate) + len(slf.packets)
		slf.cond.L.Unlock()
		return
	}
	cp := slf.packetPool.Get()
	cp.wst = wst
	cp.data = packet
	if len(callback) > 0 {
		cp.callback = callback[0]
	}
	slf.cond.L.Lock()
	slf.packets = append(slf.packets, cp)
	slf.accumulation = len(slf.accumulate) + len(slf.packets)
	slf.cond.Signal()
	slf.cond.L.Unlock()
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
	slf.cond.L.Lock()
	slf.packets = append(slf.packets, slf.accumulate...)
	slf.accumulate = nil
	slf.cond.L.Unlock()
	defer func() {
		if err := recover(); err != nil {
			slf.Close(err.(error))
		}
	}()
	wait.Done()

	for {
		slf.cond.L.Lock()
		if slf.packetPool == nil {
			slf.cond.L.Unlock()
			return
		}
		if len(slf.packets) == 0 {
			slf.cond.Wait()
		}
		packets := slf.packets[0:]
		slf.packets = slf.packets[0:0]
		slf.cond.L.Unlock()
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

// GetMessageAccumulationTotal 获取消息积压总数
func (slf *Client) GetMessageAccumulationTotal() int {
	return slf.accumulation
}
