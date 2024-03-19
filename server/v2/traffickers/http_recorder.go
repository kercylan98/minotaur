package traffickers

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"io"
	"net"
	netHttp "net/http"
	"net/textproto"
	"strconv"
	"strings"

	"golang.org/x/net/http/httpguts"
)

type httpResponseWriter struct {
	Code      int
	HeaderMap netHttp.Header
	Body      *bytes.Buffer
	Flushed   bool

	conn        *websocketConn
	result      *netHttp.Response
	snapHeader  netHttp.Header
	wroteHeader bool
	isHijack    bool
}

func (rw *httpResponseWriter) init(c gnet.Conn) {
	rw.conn = &websocketConn{Conn: c}
	rw.Code = 200
	rw.Body = new(bytes.Buffer)
	rw.HeaderMap = make(netHttp.Header)
	rw.isHijack = false
}

func (rw *httpResponseWriter) reset() {
	rw.conn = nil
	rw.Code = 200
	rw.Body = nil
	rw.HeaderMap = nil
	rw.isHijack = false
}

func (rw *httpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if !rw.isHijack {
		return rw.conn, bufio.NewReadWriter(bufio.NewReader(rw.conn), bufio.NewWriter(rw.conn)), nil
	}
	return nil, nil, netHttp.ErrHijacked
}

func (rw *httpResponseWriter) Header() netHttp.Header {
	m := rw.HeaderMap
	if m == nil {
		m = make(netHttp.Header)
		rw.HeaderMap = m
	}
	return m
}

func (rw *httpResponseWriter) writeHeader(b []byte, str string) {
	if rw.wroteHeader {
		return
	}
	if len(str) > 512 {
		str = str[:512]
	}

	m := rw.Header()

	_, hasType := m["Content-Type"]
	hasTE := m.Get("Transfer-Encoding") != ""
	if !hasType && !hasTE {
		if b == nil {
			b = []byte(str)
		}
		m.Set("Content-Type", netHttp.DetectContentType(b))
	}

	rw.WriteHeader(200)
}

func (rw *httpResponseWriter) Write(buf []byte) (n int, err error) {
	if rw.isHijack {
		n = len(buf)
		var wait = make(chan error)
		if err = rw.conn.AsyncWrite(buf, func(c gnet.Conn, err error) error {
			if err != nil {
				wait <- err
			}
			return nil
		}); err != nil {
			return
		}
		err = <-wait
		return
	}
	rw.writeHeader(buf, "")
	if rw.Body != nil {
		rw.Body.Write(buf)
	}
	return len(buf), nil
}

func (rw *httpResponseWriter) WriteString(str string) (int, error) {
	rw.writeHeader(nil, str)
	if rw.Body != nil {
		rw.Body.WriteString(str)
	}
	return len(str), nil
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}

func (rw *httpResponseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	checkWriteHeaderCode(code)
	rw.Code = code
	rw.wroteHeader = true
	if rw.HeaderMap == nil {
		rw.HeaderMap = make(netHttp.Header)
	}
	rw.snapHeader = rw.HeaderMap.Clone()
}

func (rw *httpResponseWriter) Flush() {
	if !rw.wroteHeader {
		rw.WriteHeader(200)
	}
	rw.Flushed = true
}

func (rw *httpResponseWriter) Result() *netHttp.Response {
	if rw.result != nil {
		return rw.result
	}
	if rw.snapHeader == nil {
		rw.snapHeader = rw.HeaderMap.Clone()
	}
	res := &netHttp.Response{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		StatusCode: rw.Code,
		Header:     rw.snapHeader,
	}
	rw.result = res
	if res.StatusCode == 0 {
		res.StatusCode = 200
	}
	res.Status = fmt.Sprintf("%03d %s", res.StatusCode, netHttp.StatusText(res.StatusCode))
	if rw.Body != nil {
		res.Body = io.NopCloser(bytes.NewReader(rw.Body.Bytes()))
	} else {
		res.Body = netHttp.NoBody
	}
	res.ContentLength = parseContentLength(res.Header.Get("Content-Length"))

	if trailers, ok := rw.snapHeader["Trailer"]; ok {
		res.Trailer = make(netHttp.Header, len(trailers))
		for _, k := range trailers {
			for _, k := range strings.Split(k, ",") {
				k = netHttp.CanonicalHeaderKey(textproto.TrimString(k))
				if !httpguts.ValidTrailerHeader(k) {
					// Ignore since forbidden by RFC 7230, section 4.1.2.
					continue
				}
				vv, ok := rw.HeaderMap[k]
				if !ok {
					continue
				}
				vv2 := make([]string, len(vv))
				copy(vv2, vv)
				res.Trailer[k] = vv2
			}
		}
	}
	for k, vv := range rw.HeaderMap {
		if !strings.HasPrefix(k, netHttp.TrailerPrefix) {
			continue
		}
		if res.Trailer == nil {
			res.Trailer = make(netHttp.Header)
		}
		for _, v := range vv {
			res.Trailer.Add(strings.TrimPrefix(k, netHttp.TrailerPrefix), v)
		}
	}
	return res
}

func parseContentLength(cl string) int64 {
	cl = textproto.TrimString(cl)
	if cl == "" {
		return -1
	}
	n, err := strconv.ParseUint(cl, 10, 63)
	if err != nil {
		return -1
	}
	return int64(n)
}
