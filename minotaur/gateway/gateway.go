package gateway

import (
	"context"
	"sync"
	"sync/atomic"
)

const (
	gatewayStatusNone = status(iota)
	gatewayStatusRunning
	gatewayStatusStopped
)

type (
	status  = int32  // 状态
	Address = string // 完整的网络地址
	Host    = string // 主机地址
	Port    = uint16 // 端口号
)

func New() *Gateway {
	gateway := &Gateway{
		listeners: make(map[Address]Listener),
		wg:        &sync.WaitGroup{},
		status:    gatewayStatusNone,
		events:    make(chan any, 1024),
	}

	gateway.ctx, gateway.cancel = context.WithCancel(context.Background())

	return gateway
}

type Gateway struct {
	ctx       context.Context      // 上下文
	cancel    context.CancelFunc   // 取消函数
	status    status               // 状态
	events    chan any             // 事件
	wg        *sync.WaitGroup      // 等待组
	listeners map[Address]Listener // 监听器
}

func (g *Gateway) BindListener(listener Listener, cb ...func(err error)) {
	g.events <- listenerBindEvent{
		listener: listener,
		callback: cb,
	}
}

func (g *Gateway) Run() {
	if !atomic.CompareAndSwapInt32(&g.status, gatewayStatusNone, gatewayStatusRunning) {
		return
	}
	for {
		select {
		case <-g.ctx.Done():
			return
		case event := <-g.events:
			switch e := event.(type) {
			case listenerBindEvent:
				g.onListenerBind(e)

			}
		}
	}
}

func (g *Gateway) Stop() {
	if !atomic.CompareAndSwapInt32(&g.status, gatewayStatusRunning, gatewayStatusStopped) {
		return
	}
	g.cancel()
	g.wg.Wait()
}

func (g *Gateway) onListenerBind(e listenerBindEvent) {
	if err := e.listener.Start(g.ctx); err != nil {
		for _, cb := range e.callback {
			cb(err)
		}
	} else {
		g.listeners[e.listener.Address()] = e.listener
	}
}
