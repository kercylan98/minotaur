package network

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/server/v2"
	"net/http"
)

func WebSocket(addr, pattern string) server.Network {
	return WebSocketWithHandler[*HttpServe](addr, &HttpServe{ServeMux: http.NewServeMux()}, func(handler *HttpServe, ws http.HandlerFunc) {
		handler.Handle(fmt.Sprintf("GET %s", pattern), ws)
	})
}

func WebSocketWithHandler[H http.Handler](addr string, handler H, upgraderHandlerFunc WebSocketUpgraderHandlerFunc[H]) server.Network {
	c := &websocketCore[H]{
		httpCore:            HttpWithHandler(addr, handler).(*httpCore[H]),
		upgraderHandlerFunc: upgraderHandlerFunc,
	}
	return c
}

type WebSocketUpgraderHandlerFunc[H http.Handler] func(handler H, ws http.HandlerFunc)

type websocketCore[H http.Handler] struct {
	*httpCore[H]
	upgraderHandlerFunc WebSocketUpgraderHandlerFunc[H]
	core                server.Core
}

func (w *websocketCore[H]) OnSetup(ctx context.Context, core server.Core) (err error) {
	w.core = core
	if err = w.httpCore.OnSetup(ctx, core); err != nil {
		return
	}
	w.upgraderHandlerFunc(w.handler, w.onUpgrade)
	return
}

func (w *websocketCore[H]) OnRun(ctx context.Context) error {
	return w.httpCore.OnRun(ctx)
}

func (w *websocketCore[H]) OnShutdown() error {
	return w.httpCore.OnShutdown()
}

func (w *websocketCore[H]) onUpgrade(writer http.ResponseWriter, request *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(request, writer)
	if err != nil {
		return
	}

	w.core.Event() <- conn
}
