package network

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/panjf2000/gnet/v2"
	"time"
)

type gNetGenericHandler struct {
	engine *gnet.Engine
	*gNetCore
}

func (t *gNetGenericHandler) OnInit(core *gNetCore) {
	t.gNetCore = core
}

func (t *gNetGenericHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	t.engine = &eng
	return
}

func (t *gNetGenericHandler) OnShutdown(eng gnet.Engine) {

}

func (t *gNetGenericHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	t.controller.RegisterConnection(c,
		func(packet server.Packet) error {
			return c.AsyncWrite(packet.GetBytes(), func(c gnet.Conn, err error) error {
				t.controller.OnConnectionAsyncWriteError(c.Context().(server.Conn), packet, err)
				return nil
			})
		}, func(conn server.Conn) {
			c.SetContext(conn)
		})
	return
}

func (t *gNetGenericHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	t.controller.EliminateConnection(c, err)
	return
}

func (t *gNetGenericHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	buf, err := c.Next(-1)
	if err != nil {
		return gnet.Close
	}

	var clone = make([]byte, len(buf))
	copy(clone, buf)

	t.controller.ReactPacket(c, server.NewPacket(clone))
	return
}

func (t *gNetGenericHandler) OnTick() (delay time.Duration, action gnet.Action) {
	return
}

func (t *gNetGenericHandler) GetEngine() *gnet.Engine {
	return t.engine
}
