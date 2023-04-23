package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"minotaur/utils/log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// New 根据特定网络类型创建一个服务器
func New(network Network) *Server {
	server := &Server{
		network: network,
	}
	server.event = &event{Server: server}

	if network == NetworkHttp {
		server.httpServer = gin.New()
	}
	return server
}

// Server 网络服务器
type Server struct {
	*event
	network            Network
	addr               string
	httpServer         *gin.Engine
	grpcServer         *grpc.Server
	gServer            *gNet
	messageChannel     chan *message
	initMessageChannel bool
	multiple           bool
}

// Run 使用特定地址运行服务器
//
//	server.NetworkTCP (addr:":8888")
//	server.NetworkTCP4 (addr:":8888")
//	server.NetworkTCP6 (addr:":8888")
//	server.NetworkUDP (addr:":8888")
//	server.NetworkUDP4 (addr:":8888")
//	server.NetworkUDP6 (addr:":8888")
//	server.NetworkUnix (addr:"socketPath")
//	server.NetworkHttp (addr:":8888")
//	server.NetworkWebsocket (addr:":8888/ws")
//	server.NetworkKcp (addr:":8888")
func (slf *Server) Run(addr string) error {
	if slf.event == nil {
		return ErrConstructed
	}
	slf.event.check()
	slf.addr = addr
	var protoAddr = fmt.Sprintf("%s://%s", slf.network, slf.addr)
	var connectionInitHandle = func(callback func()) {
		slf.initMessageChannel = true
		slf.messageChannel = make(chan *message, 4096*1000)
		if slf.network != NetworkHttp && slf.network != NetworkWebsocket {
			slf.gServer = &gNet{Server: slf}
		}
		if callback != nil {
			go callback()
		}
		for message := range slf.messageChannel {
			slf.dispatchMessage(message)
		}
	}

	switch slf.network {
	case NetworkGRPC:
		listener, err := net.Listen(string(NetworkTCP), slf.addr)
		if err != nil {
			return err
		}
		slf.grpcServer = grpc.NewServer()
		go func() {
			if err := slf.grpcServer.Serve(listener); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		}()
	case NetworkTCP, NetworkTCP4, NetworkTCP6, NetworkUdp, NetworkUdp4, NetworkUdp6, NetworkUnix:
		go connectionInitHandle(func() {
			if err := gnet.Serve(slf.gServer, protoAddr); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		})
	case NetworkKcp:
		listener, err := kcp.ListenWithOptions(slf.addr, nil, 0, 0)
		if err != nil {
			return err
		}
		go connectionInitHandle(func() {
			for {
				session, err := listener.AcceptKCP()
				if err != nil {
					continue
				}

				conn := newKcpConn(session)
				slf.OnConnectionOpenedEvent(conn)

				go func(conn *Conn) {
					defer func() {
						if err := recover(); err != nil {
							conn.Close()
							slf.OnConnectionClosedEvent(conn)
						}
					}()

					buf := make([]byte, 4096)
					for {
						n, err := conn.kcp.Read(buf)
						if err != nil {
							panic(err)
						}
						slf.PushMessage(MessageTypePacket, conn, buf[:n])
					}
				}(conn)
			}
		})
	case NetworkHttp:
		go func() {
			if err := slf.httpServer.Run(addr); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		}()
	case NetworkWebsocket:
		go connectionInitHandle(nil)
		var pattern string
		var index = strings.Index(addr, "/")
		if index == -1 {
			pattern = "/"
		} else {
			pattern = addr[index:]
		}
		var upgrade = websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
			ip := request.Header.Get("X-Real-IP")
			ws, err := upgrade.Upgrade(writer, request, nil)
			if err != nil {
				return
			}
			if len(ip) == 0 {
				addr := ws.RemoteAddr().String()
				if index := strings.LastIndex(addr, ":"); index != -1 {
					ip = addr[0:index]
				}
			}

			conn := newWebsocketConn(ws)
			conn.ip = ip
			slf.OnConnectionOpenedEvent(conn)

			defer func() {
				if err := recover(); err != nil {
					conn.Close()
					slf.OnConnectionClosedEvent(conn)
				}
			}()

			for {
				if err := ws.SetReadDeadline(time.Now().Add(time.Second * 30)); err != nil {
					panic(err)
				}
				_, packet, err := ws.ReadMessage()
				if err != nil {
					panic(err)
				}
				slf.PushMessage(MessageTypePacket, conn, packet)

			}
		})
		go func() {
			if err := http.ListenAndServe(slf.addr, nil); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		}()
	default:
		return ErrCanNotSupportNetwork
	}

	if !slf.multiple {
		log.Info("Server", zap.String("Minotaur Server", "===================================================================="))
		log.Info("Server", zap.String("Minotaur Server", "RunningInfo"),
			zap.Any("network", slf.network),
			zap.String("listen", slf.addr),
		)
		log.Info("Server", zap.String("Minotaur Server", "===================================================================="))

		systemSignal := make(chan os.Signal, 1)
		signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-systemSignal:
			slf.Shutdown(nil)
		}
	}

	return nil
}

// Shutdown 停止运行服务器
func (slf *Server) Shutdown(err error) {
	if slf.initMessageChannel {
		close(slf.messageChannel)
	}
	if err != nil {
		log.Error("Server", zap.Any("network", slf.network), zap.String("listen", slf.addr),
			zap.String("action", "shutdown"), zap.String("state", "exception"), zap.Error(err))
	} else {
		log.Info("Server", zap.Any("network", slf.network), zap.String("listen", slf.addr),
			zap.String("action", "shutdown"), zap.String("state", "normal"))
	}
}

// HttpRouter 当网络类型为 NetworkHttp 时将被允许获取路由器进行路由注册，否则将会发生 panic
func (slf *Server) HttpRouter() gin.IRouter {
	if slf.httpServer == nil {
		panic(ErrNetworkOnlySupportHttp)
	}
	return slf.httpServer
}

// PushMessage 向服务器中写入特定类型的消息，需严格遵守消息属性要求
func (slf *Server) PushMessage(messageType MessageType, attrs ...any) {
	slf.messageChannel <- &message{
		t:     messageType,
		attrs: attrs,
	}
}

// dispatchMessage 消息分发
func (slf *Server) dispatchMessage(msg *message) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Server", zap.Any("error", err))
		}
	}()
	switch msg.t {
	case MessageTypePacket:
		conn, packet := msg.t.deconstructPacket(msg.attrs...)
		slf.OnConnectionReceivePacketEvent(conn, packet)
	case MessageTypeError:
		err, action := msg.t.deconstructError(msg.attrs...)
		switch action {
		case MessageErrorActionNone:
			log.Error("Server", zap.Error(err))
		case MessageErrorActionShutdown:
			slf.Shutdown(err)
		default:
			log.Warn("Server", zap.String("not support message error action", action.String()))
		}
	default:
		log.Warn("Server", zap.String("not support message type", msg.t.String()))
	}
}