package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kercylan98/minotaur/server/internal/logger"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/slice"
	"github.com/kercylan98/minotaur/utils/super"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strings"
	"time"
)

type Network string

const (
	// NetworkNone 该模式下不监听任何网络端口，仅开启消息队列，适用于纯粹的跨服服务器等情况
	NetworkNone Network = "none"
	NetworkTcp  Network = "tcp"
	NetworkTcp4 Network = "tcp4"
	NetworkTcp6 Network = "tcp6"
	NetworkUdp  Network = "udp"
	NetworkUdp4 Network = "udp4"
	NetworkUdp6 Network = "udp6"
	NetworkUnix Network = "unix"
	NetworkHttp Network = "http"
	// NetworkWebsocket 该模式下需要获取url参数值时，可以通过连接的GetData函数获取
	//  - 当有多个同名参数时，获取到的值为切片类型
	NetworkWebsocket Network = "websocket"
	NetworkKcp       Network = "kcp"
	NetworkGRPC      Network = "grpc"
)

var (
	networkNameMap map[string]struct{}
	networks       = []Network{
		NetworkNone, NetworkTcp, NetworkTcp4, NetworkTcp6, NetworkUdp, NetworkUdp4, NetworkUdp6, NetworkUnix, NetworkHttp, NetworkWebsocket, NetworkKcp, NetworkGRPC,
	}
)

func init() {
	networkNameMap = make(map[string]struct{}, len(networks))
	for _, network := range networks {
		networkNameMap[string(network)] = struct{}{}
	}
}

// GetNetworks 获取所有支持的网络模式
func GetNetworks() []Network {
	return slice.Copy(networks)
}

// check 检查网络模式是否支持
func (n Network) check() {
	if !hash.Exist(networkNameMap, string(n)) {
		panic(fmt.Errorf("unsupported network mode: %s", n))
	}
}

// preprocessing 服务器预处理
func (n Network) preprocessing(srv *Server) {
	switch n {
	case NetworkNone:
	case NetworkTcp:
	case NetworkTcp4:
	case NetworkTcp6:
	case NetworkUdp:
	case NetworkUdp4:
	case NetworkUdp6:
	case NetworkUnix:
	case NetworkHttp:
		srv.ginServer = gin.New()
		srv.httpServer = &http.Server{
			Handler: srv.ginServer,
		}
	case NetworkWebsocket:
		srv.websocketReadDeadline = DefaultWebsocketReadDeadline
	case NetworkKcp:
	case NetworkGRPC:
		srv.grpcServer = grpc.NewServer()
	}
}

// adaptation 服务器适配
func (n Network) adaptation(srv *Server) <-chan error {
	state := make(chan error, 1)
	switch n {
	case NetworkNone:
		srv.addr = "-"
	case NetworkTcp:
		n.gNetMode(state, srv)
	case NetworkTcp4:
		n.gNetMode(state, srv)
	case NetworkTcp6:
		n.gNetMode(state, srv)
	case NetworkUdp:
		n.gNetMode(state, srv)
	case NetworkUdp4:
		n.gNetMode(state, srv)
	case NetworkUdp6:
		n.gNetMode(state, srv)
	case NetworkUnix:
		n.gNetMode(state, srv)
	case NetworkHttp:
		n.httpMode(state, srv)
	case NetworkWebsocket:
		n.websocketMode(state, srv)
	case NetworkKcp:
		n.kcpMode(state, srv)
	case NetworkGRPC:
		n.grpcMode(state, srv)
	}
	return state
}

// gNetMode gNet模式
func (n Network) gNetMode(state chan<- error, srv *Server) {
	srv.gServer = &gNet{Server: srv, state: state}
	go func(srv *Server) {
		if err := gnet.Serve(srv.gServer, fmt.Sprintf("%s://%s", srv.network, srv.addr),
			gnet.WithLogger(new(logger.GNet)),
			gnet.WithTicker(true),
			gnet.WithMulticore(true),
		); err != nil {
			srv.gServer.state <- err
		}
	}(srv)
}

// grpcMode grpc模式
func (n Network) grpcMode(state chan<- error, srv *Server) {
	l, err := net.Listen(string(NetworkTcp), srv.addr)
	if err != nil {
		state <- err
		return
	}
	lis := (&listener{srv: srv, Listener: l, state: state}).init()
	go func(srv *Server, lis *listener) {
		if err = srv.grpcServer.Serve(lis); err != nil {
			lis.state <- err
		}
	}(srv, lis)
}

// kcpMode kcp模式
func (n Network) kcpMode(state chan<- error, srv *Server) {
	l, err := kcp.ListenWithOptions(srv.addr, nil, 0, 0)
	if err != nil {
		state <- err
		return
	}
	lis := (&listener{srv: srv, kcpListener: l, state: state}).init()
	go func(lis *listener) {
		for {
			session, err := lis.AcceptKCP()
			if err != nil {
				continue
			}

			conn := newKcpConn(lis.srv, session)
			lis.srv.OnConnectionOpenedEvent(conn)

			go func(conn *Conn) {
				defer func() {
					if err := super.RecoverTransform(recover()); err != nil {
						conn.Close(err)
					}
				}()

				buf := make([]byte, 4096)
				for !conn.IsClosed() {
					n, err := conn.kcp.Read(buf)
					if err != nil {
						if conn.IsClosed() {
							break
						}
						panic(err)
					}
					lis.srv.PushPacketMessage(conn, 0, buf[:n])
				}
			}(conn)
		}
	}(lis)
	return
}

// httpMode http模式
func (n Network) httpMode(state chan<- error, srv *Server) {
	srv.httpServer.Addr = srv.addr
	l, err := net.Listen(string(NetworkTcp), srv.addr)
	if err != nil {
		state <- err
		return
	}
	gin.SetMode(gin.ReleaseMode)
	srv.ginServer.Use(func(c *gin.Context) {
		t := time.Now()
		c.Next()
		log.Info("Server", log.String("type", "http"),
			log.String("method", c.Request.Method), log.Int("status", c.Writer.Status()),
			log.String("ip", c.ClientIP()), log.String("path", c.Request.URL.Path),
			log.Duration("cost", time.Since(t)))
	})
	go func(lis *listener) {
		var err error
		if len(lis.srv.certFile)+len(srv.keyFile) > 0 {
			err = lis.srv.httpServer.ServeTLS(lis, lis.srv.certFile, lis.srv.keyFile)
		} else {
			err = lis.srv.httpServer.Serve(lis)
		}
		if err != nil {
			lis.state <- err
		}
	}((&listener{srv: srv, Listener: l, state: state}).init())
}

// websocketMode websocket模式
func (n Network) websocketMode(state chan<- error, srv *Server) {
	l, err := net.Listen(string(NetworkTcp), srv.addr)
	if err != nil {
		state <- err
		return
	}
	var pattern string
	var index = strings.Index(srv.addr, "/")
	if index == -1 {
		pattern = "/"
	} else {
		pattern = srv.addr[index:]
		srv.addr = srv.addr[:index]
	}
	if srv.websocketUpgrader == nil {
		srv.websocketUpgrader = DefaultWebsocketUpgrader()
	}
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		ip := request.Header.Get("X-Real-IP")
		ws, err := srv.websocketUpgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}
		if srv.websocketConnInitializer != nil {
			if err = srv.websocketConnInitializer(writer, request, ws); err != nil {
				return
			}
		}
		if len(ip) == 0 {
			addr := ws.RemoteAddr().String()
			if index := strings.LastIndex(addr, ":"); index != -1 {
				ip = addr[0:index]
			}
		}
		if srv.websocketCompression > 0 {
			_ = ws.SetCompressionLevel(srv.websocketCompression)
		}
		ws.EnableWriteCompression(srv.websocketWriteCompression)
		conn := newWebsocketConn(srv, ws, ip)
		conn.SetData(wsRequestKey, request)
		for k, v := range request.URL.Query() {
			if len(v) == 1 {
				conn.SetData(k, v[0])
			} else {
				conn.SetData(k, v)
			}
		}
		srv.OnConnectionOpenedEvent(conn)

		defer func() {
			if err := super.RecoverTransform(recover()); err != nil {
				conn.Close(err)
			}
		}()
		for !conn.IsClosed() {
			if srv.websocketReadDeadline > 0 {
				if err := ws.SetReadDeadline(time.Now().Add(srv.websocketReadDeadline)); err != nil {
					panic(err)
				}
			}
			messageType, packet, readErr := ws.ReadMessage()
			if readErr != nil {
				if conn.IsClosed() {
					break
				}
				panic(readErr)
			}
			if len(srv.supportMessageTypes) > 0 && !srv.supportMessageTypes[messageType] {
				panic(ErrWebsocketIllegalMessageType)
			}
			srv.PushPacketMessage(conn, messageType, packet)
		}
	})
	go func(lis *listener) {
		var err error
		if len(lis.srv.certFile)+len(lis.srv.keyFile) > 0 {
			err = http.ServeTLS(lis, nil, lis.srv.certFile, lis.srv.keyFile)
		} else {
			err = http.Serve(lis, nil)
		}
		if err != nil {
			lis.state <- err
		}
	}((&listener{srv: srv, Listener: l, state: state}).init())
}
