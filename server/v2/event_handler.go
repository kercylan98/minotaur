package server

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet/v2"
	"time"
)

func newEventHandler(trafficker Trafficker) (handler *eventHandler, err error) {
	var wp *ants.Pool
	if wp, err = ants.NewPool(ants.DefaultAntsPoolSize, ants.WithNonblocking(true)); err != nil {
		return
	}
	handler = &eventHandler{
		trafficker: trafficker,
		workerPool: wp,
	}
	return
}

type (
	Trafficker interface {
		OnBoot() error
		OnTraffic(c gnet.Conn, packet []byte)
	}
	eventHandler struct {
		trafficker Trafficker
		workerPool *ants.Pool
	}
)

func (e *eventHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	if err := e.trafficker.OnBoot(); err != nil {
		action = gnet.Shutdown
	}
	return
}

func (e *eventHandler) OnShutdown(eng gnet.Engine) {
	return
}

func (e *eventHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (e *eventHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println("断开")
	return
}

func (e *eventHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	buf, err := c.Next(-1)
	if err != nil {
		return
	}
	var packet = make([]byte, len(buf))
	copy(packet, buf)
	err = e.workerPool.Submit(func() {
		e.trafficker.OnTraffic(c, packet)
	})

	return
}

func (e *eventHandler) OnTick() (delay time.Duration, action gnet.Action) {
	return
}
