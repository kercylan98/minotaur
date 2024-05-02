package network

import (
	"context"
	"github.com/kercylan98/minotaur/server"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"time"
)

type httpCore[H http.Handler] struct {
	addr       string
	handler    H
	srv        *http.Server
	controller server.Controller
}

func (h *httpCore[H]) OnSetup(ctx context.Context, controller server.Controller) (err error) {
	h.controller = controller
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

func (h *httpCore[H]) Schema() string {
	return "http(s)"
}

func (h *httpCore[H]) Address() string {
	return h.srv.Addr
}
