package network

import (
	"errors"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/kercylan98/minotaur/server"
	"github.com/panjf2000/gnet/v2"
	"time"
)

func newWebsocketHandler(core *websocketCore) *websocketHandler {
	return &websocketHandler{
		websocketCore: core,
	}
}

type websocketHandler struct {
	engine   *gnet.Engine
	upgrader ws.Upgrader
	*websocketCore
}

func (w *websocketHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	w.engine = &eng
	w.initUpgrader()
	return
}

func (w *websocketHandler) OnShutdown(eng gnet.Engine) {

}

func (w *websocketHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	wrapper := newWebsocketWrapper(c)
	c.SetContext(wrapper)
	return
}

func (w *websocketHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	w.controller.EliminateConnection(c, err)
	return
}

func (w *websocketHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	wrapper := c.Context().(*websocketWrapper)

	if err := wrapper.readToBuffer(); err != nil {
		return gnet.Close
	}

	if err := wrapper.upgrade(w.upgrader, func() {
		// 协议升级成功后视为连接建立
		w.controller.RegisterConnection(c, func(packet server.Packet) error {
			return wsutil.WriteServerMessage(c, packet.GetContext().(ws.OpCode), packet.GetBytes())
		})
	}); err != nil {
		return gnet.Close
	}
	wrapper.active = time.Now()

	// decode
	messages, err := wrapper.decode()
	if err != nil {
		return gnet.Close
	}

	for _, message := range messages {
		packet := server.NewPacket(message.Payload)
		packet.SetContext(message.OpCode)
		w.controller.ReactPacket(c, packet)
	}
	return
}

func (w *websocketHandler) OnTick() (delay time.Duration, action gnet.Action) {
	return
}

func (w *websocketHandler) initUpgrader() {
	w.upgrader = ws.Upgrader{
		OnRequest: func(uri []byte) (err error) {
			if string(uri) != w.pattern {
				err = errors.New("bad request")
			}
			return
		},
	}
}
