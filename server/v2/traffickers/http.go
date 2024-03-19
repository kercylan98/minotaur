package traffickers

import (
	"bufio"
	"bytes"
	"github.com/kercylan98/minotaur/server/v2"
	"github.com/kercylan98/minotaur/utils/hub"
	"github.com/panjf2000/gnet/v2"
	netHttp "net/http"
)

func Http[H netHttp.Handler](handler H) server.Trafficker {
	return &http[H]{
		handler: handler,
		ncb: func(c gnet.Conn, err error) error {
			return nil
		},
	}
}

type http[H netHttp.Handler] struct {
	handler H
	rwp     *hub.ObjectPool[*httpResponseWriter]
	ncb     func(c gnet.Conn, err error) error
}

func (h *http[H]) OnBoot(options *server.Options) error {
	h.rwp = hub.NewObjectPool[httpResponseWriter](func() *httpResponseWriter {
		return new(httpResponseWriter)
	}, func(data *httpResponseWriter) {
		data.reset()
	})
	return nil
}

func (h *http[H]) OnTraffic(c gnet.Conn, packet []byte) {
	var responseWriter *httpResponseWriter
	defer func() {
		if responseWriter == nil || !responseWriter.isHijack {
			_ = c.Close()
		}
	}()
	httpRequest, err := netHttp.ReadRequest(bufio.NewReader(bytes.NewReader(packet)))
	if err != nil {
		return
	}

	responseWriter = h.rwp.Get()
	responseWriter.init(c)

	h.handler.ServeHTTP(responseWriter, httpRequest)
	if responseWriter.isHijack {
		return
	}
	_ = responseWriter.Result().Write(c)
}
