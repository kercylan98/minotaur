package network

import (
	server "github.com/kercylan98/minotaur/server3"
	"net/http"
)

// Http 创建一个基于 http.ServeMux 的 HTTP 的网络
func Http(addr string) server.Network {
	return HttpWithHandler(addr, &HttpServe{ServeMux: http.NewServeMux()})
}

// HttpWithHandler 创建一个基于 http.Handler 的 HTTP 的网络
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

// WebSocket 创建一个基于 TCP 的 WebSocket 网络
//   - addr 期望为类似于 127.0.0.1:1234 或 :1234 的地址
//   - pattern 期望为 WebSocket 的路径，如果为空则默认为 /
func WebSocket(addr string, pattern ...string) server.Network {
	return newGnetEngine(schemaWebSocket, addr, pattern...)
}

// Tcp 创建一个 TCP 网络
//   - addr 期望为类似于 127.0.0.1:1234 或 :1234 的地址
func Tcp(addr string) server.Network {
	return newGnetEngine(schemaTcp, addr)
}

// Tcp4 创建一个 IPv4 TCP 网络
//   - addr 期望为类似于 127.0.0.1:1234 或 :1234 的地址
func Tcp4(addr string) server.Network {
	return newGnetEngine(schemaTcp4, addr)
}

// Tcp6 创建一个 IPv6 TCP 网络
//   - addr 期望为类似于 [::1]:1234 的地址
func Tcp6(addr string) server.Network {
	return newGnetEngine(schemaTcp6, addr)
}

// Udp 创建一个 UDP 网络
//   - addr 期望为类似于 127.0.0.1:1234 或 :1234 的地址
func Udp(addr string) server.Network {
	return newGnetEngine(schemaUdp, addr)
}

// Udp4 创建一个 IPv4 UDP 网络
//   - addr 期望为类似于 127.0.0.1:1234 或 :1234 的地址
func Udp4(addr string) server.Network {
	return newGnetEngine(schemaUdp4, addr)
}

// Udp6 创建一个 IPv6 UDP 网络
//   - addr 期望为类似于 [::1]:1234 的地址
func Udp6(addr string) server.Network {
	return newGnetEngine(schemaUdp6, addr)
}

// Unix 创建一个 Unix Domain Socket 网络
//   - addr 期望为类似于 /tmp/xxx.sock 的文件地址
func Unix(addr string) server.Network {
	return newGnetEngine(schemaUnix, addr)
}

// Kcp 创建一个 KCP 网络
//   - addr 期望为类似于 127.0.0.1:1234 或 :1234 的地址
func Kcp(addr string) server.Network {
	return newKcpCore(addr)
}
