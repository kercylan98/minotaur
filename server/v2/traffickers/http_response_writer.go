package traffickers

import (
	"bytes"
	"github.com/panjf2000/gnet/v2"
	netHttp "net/http"
	"strconv"
	"sync"
)

type httpResponseWriter struct {
	c          gnet.Conn
	statusCode int
	header     netHttp.Header
}

func (w *httpResponseWriter) init(c gnet.Conn) {
	w.c = c
	w.statusCode = 200
	w.header = make(netHttp.Header)
}

func (w *httpResponseWriter) reset() {
	w.c = nil
	w.statusCode = 200
	w.header = nil
}

func (w *httpResponseWriter) Header() netHttp.Header {
	return w.header
}

func (w *httpResponseWriter) Write(b []byte) (n int, err error) {
	var buf bytes.Buffer
	buf.WriteString("HTTP/1.1 ")
	buf.WriteString(netHttp.StatusText(w.statusCode))
	buf.WriteString("\r\n")
	w.header.Set("Content-Length", strconv.Itoa(len(b)))
	if err = w.header.Write(&buf); err != nil {
		return
	}
	buf.WriteString("\r\n")
	buf.Write(b)
	res := buf.Bytes()
	var wg sync.WaitGroup
	wg.Add(1)
	err = w.c.AsyncWrite(res, func(c gnet.Conn, e error) error {
		err = e
		wg.Done()
		return nil
	})
	wg.Wait()
	return len(res), err
}

func (w *httpResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}
