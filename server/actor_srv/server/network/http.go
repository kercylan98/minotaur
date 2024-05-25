package network

import (
	"context"
	"errors"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/vivids"
	"net/http"
)

type HttpServe struct {
	*http.ServeMux
}

type httpCore[H http.Handler] struct {
	vivid.BasicActor
	addr       string
	handler    H
	srv        *http.Server
	controller server.Controller
}

func (h *httpCore[H]) OnPreStart(ctx vivids.ActorContext) (err error) {
	ctx.Future(func() vivids.Message {
		return h.srv.ListenAndServe()
	})

	return
}

func (h *httpCore[H]) OnReceived(ctx vivids.MessageContext) (err error) {
	switch v := ctx.GetMessage().(type) {
	case error:
		switch {
		case errors.Is(v, http.ErrServerClosed):
			ctx.NotifyTerminated()
		}
	}

	return
}

func (h *httpCore[H]) OnDestroy(ctx vivids.ActorContext) (err error) {
	return h.srv.Shutdown(context.TODO())
}
