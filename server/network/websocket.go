package network

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/panjf2000/gnet/v2"
	"time"
)

func WebSocket(addr string, pattern ...string) server.Network {
	ws := &websocketCore{
		addr:    addr,
		pattern: collection.FindFirstOrDefaultInSlice(pattern, "/"),
	}
	return ws
}

type websocketCore struct {
	ctx        context.Context
	controller server.Controller
	handler    *websocketHandler
	addr       string
	pattern    string
}

func (w *websocketCore) OnSetup(ctx context.Context, controller server.Controller) (err error) {
	w.ctx = ctx
	w.handler = newWebsocketHandler(w)
	w.controller = controller
	return
}

func (w *websocketCore) OnRun() (err error) {
	err = gnet.Run(w.handler, fmt.Sprintf("tcp://%s", w.addr))
	return
}

func (w *websocketCore) OnShutdown() error {
	if w.handler.engine != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return w.handler.engine.Stop(ctx)
	}
	return nil
}

func (w *websocketCore) Schema() string {
	return "ws"
}

func (w *websocketCore) Address() string {
	if w.pattern == "/" {
		return w.addr
	}
	return fmt.Sprintf("%s:%s", w.addr, w.pattern)
}
