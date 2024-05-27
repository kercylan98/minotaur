package network

import (
	"context"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"net/http"
)

type HttpServe struct {
	*http.ServeMux
}

type httpCore[H http.Handler] struct {
	addr    string
	handler H
	srv     *http.Server
}

func (h *httpCore[H]) Launch(ctx context.Context, srv transport.ServerCore) error {
	return h.srv.ListenAndServe()
}

func (h *httpCore[H]) Shutdown() error {
	return h.srv.Shutdown(context.TODO())
}

func (h *httpCore[H]) Schema() string {
	return "http(s)"
}

func (h *httpCore[H]) Address() string {
	return h.addr
}
