package network

import (
	"errors"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/kercylan98/minotaur/server/internal/v2"
	"github.com/kercylan98/minotaur/utils/log"
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
	w.controller.RegisterConnection(c, func(packet server.Packet) error {
		return wsutil.WriteServerMessage(c, packet.GetContext().(ws.OpCode), packet.GetBytes())
	})
	return
}

func (w *websocketHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	w.controller.EliminateConnection(c, err)
	return
}

func (w *websocketHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	wrapper := c.Context().(*websocketWrapper)

	// read to buffer
	if err := wrapper.readToBuffer(); err != nil {
		log.Error("websocket", log.Err(err))
		return gnet.Close
	}

	// check or upgrade
	if err := wrapper.upgrade(w.upgrader); err != nil {
		log.Error("websocket", log.Err(err))
		return gnet.Close
	}
	wrapper.active = time.Now()

	// decode
	messages, err := wrapper.decode()
	if err != nil {
		log.Error("websocket", log.Err(err))
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
