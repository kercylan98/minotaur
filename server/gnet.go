package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/panjf2000/gnet"
	"go.uber.org/zap"
	"time"
)

type gNet struct {
	*Server
}

func (slf *gNet) OnInitComplete(server gnet.Server) (action gnet.Action) {
	return
}

func (slf *gNet) OnShutdown(server gnet.Server) {
	slf.closeChannel <- struct{}{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := gnet.Stop(ctx, fmt.Sprintf("%s://%s", slf.network, slf.addr)); err != nil {
		log.Error("Server", zap.String("Minotaur GNet Server", "Shutdown"), zap.Error(err))
	}
}

func (slf *gNet) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	conn := newGNetConn(slf.Server, c)
	c.SetContext(conn)
	slf.OnConnectionOpenedEvent(conn)
	return
}

func (slf *gNet) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	slf.OnConnectionClosedEvent(c.Context().(*Conn), err)
	return
}

func (slf *gNet) PreWrite(c gnet.Conn) {
	return
}

func (slf *gNet) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (slf *gNet) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	slf.Server.PushMessage(MessageTypePacket, c.Context().(*Conn), bytes.Clone(packet))
	return nil, gnet.None
}

func (slf *gNet) Tick() (delay time.Duration, action gnet.Action) {
	delay = 1 * time.Second
	if slf.isShutdown.Load() {
		return 0, gnet.Shutdown
	}
	return
}
