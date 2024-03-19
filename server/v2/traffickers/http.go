package traffickers

import (
	"bufio"
	"bytes"
	"github.com/kercylan98/minotaur/server/v2"
	"github.com/kercylan98/minotaur/utils/hub"
	"github.com/panjf2000/gnet/v2"
	netHttp "net/http"
)

func Http(handler netHttp.Handler) server.Trafficker {
	return &http{
		handler: handler,
	}
}

type http struct {
	handler netHttp.Handler
	rwp     *hub.ObjectPool[*httpResponseWriter]
}

func (h *http) OnBoot() error {
	h.rwp = hub.NewObjectPool[httpResponseWriter](func() *httpResponseWriter {
		return new(httpResponseWriter)
	}, func(data *httpResponseWriter) {
		data.reset()
	})
	return nil
}

func (h *http) OnTraffic(c gnet.Conn, packet []byte) {
	httpRequest, err := netHttp.ReadRequest(bufio.NewReader(bytes.NewReader(packet)))
	if err != nil {
		return
	}

	responseWriter := h.rwp.Get()
	responseWriter.init(c)

	h.handler.ServeHTTP(responseWriter, httpRequest)
}
