package client

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/server/writeloop"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"sync"
)

// NewClient 创建客户端
func NewClient(core Core) *Client {
	client := &Client{
		events: new(events),
		core:   core,
		closed: true,
	}
	return client
}

// CloneClient 克隆客户端
func CloneClient(client *Client) *Client {
	cli := NewClient(client.core.Clone())
	return cli
}

// Client 客户端
type Client struct {
	*events
	core   Core
	mutex  sync.Mutex
	closed bool                          // 是否已关闭
	pool   *concurrent.Pool[*Packet]     // 数据包缓冲池
	loop   *writeloop.WriteLoop[*Packet] // 写入循环
	block  chan struct{}                 // 以阻塞方式运行
}

// Run 运行客户端，当客户端已运行时，会先关闭客户端再重新运行
//   - block 以阻塞方式运行
func (slf *Client) Run(block ...bool) error {
	slf.mutex.Lock()
	if !slf.closed {
		slf.mutex.Unlock()
		slf.Close()
		slf.mutex.Lock()
	}
	if len(block) > 0 && block[0] {
		slf.block = make(chan struct{})
	}
	var runState = make(chan error)
	go func(runState chan<- error) {
		defer func() {
			if err := recover(); err != nil {
				slf.Close(err.(error))
			}
		}()
		slf.core.Run(runState, slf.onReceive)
	}(runState)
	err := <-runState
	if err != nil {
		slf.mutex.Unlock()
		return err
	}
	slf.closed = false
	slf.pool = concurrent.NewPool[Packet](func() *Packet {
		return new(Packet)
	}, func(data *Packet) {
		data.wst = 0
		data.data = nil
		data.callback = nil
	})
	slf.loop = writeloop.NewWriteLoop[*Packet](slf.pool, func(message *Packet) error {
		err := slf.core.Write(message)
		if message.callback != nil {
			message.callback(err)
		}
		return err
	}, func(err any) {
		slf.Close(errors.New(fmt.Sprint(err)))
	})
	slf.mutex.Unlock()

	slf.OnConnectionOpenedEvent(slf)
	if slf.block != nil {
		<-slf.block
	}
	return nil
}

// IsConnected 是否已连接
func (slf *Client) IsConnected() bool {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	return !slf.closed
}

// Close 关闭
func (slf *Client) Close(err ...error) {
	slf.mutex.Lock()
	slf.closed = true
	slf.core.Close()
	slf.loop.Close()
	slf.mutex.Unlock()
	if len(err) > 0 {
		slf.OnConnectionClosedEvent(slf, err[0])
	} else {
		slf.OnConnectionClosedEvent(slf, nil)
	}
	if slf.block != nil {
		slf.block <- struct{}{}
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
	defer slf.mutex.Unlock()
	if slf.closed {
		return
	}

	cp := slf.pool.Get()
	cp.wst = wst
	cp.data = packet
	if len(callback) > 0 {
		cp.callback = callback[0]
	}
	slf.loop.Put(cp)
}

func (slf *Client) onReceive(wst int, packet []byte) {
	slf.OnConnectionReceivePacketEvent(slf, wst, packet)
}

// GetServerAddr 获取服务器地址
func (slf *Client) GetServerAddr() string {
	return slf.core.GetServerAddr()
}
