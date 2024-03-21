package network

import (
	"context"
	"github.com/kercylan98/minotaur/server/v2"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"time"
)

func Http(addr string) server.Network {
	return HttpWithHandler(addr, &HttpServe{ServeMux: http.NewServeMux()})
}

func HttpWithHandler[H http.Handler](addr string, handler H) server.Network {
	c := &httpCore[H]{
		addr:    addr,
		handler: handler,
		srv: &http.Server{
			Addr:                         addr,
			Handler:                      handler,
			DisableGeneralOptionsHandler: false,
		},
	}
	return c
}

type httpCore[H http.Handler] struct {
	addr    string
	handler H
	srv     *http.Server
	event   server.NetworkCore
}

func (h *httpCore[H]) OnSetup(ctx context.Context, event server.NetworkCore) (err error) {
	h.event = event
	h.srv.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}
	return
}

func (h *httpCore[H]) OnRun() (err error) {
	if err = h.srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		err = nil
	}
	return
}

func (h *httpCore[H]) OnShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return h.srv.Shutdown(ctx)
}
