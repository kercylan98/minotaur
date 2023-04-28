package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"minotaur/utils/log"
	"minotaur/utils/synchronization"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// New 根据特定网络类型创建一个服务器
func New(network Network, options ...Option) *Server {
	server := &Server{
		event:   &event{},
		network: network,
		options: options,
	}
	server.event.Server = server

	if network == NetworkHttp {
		server.ginServer = gin.New()
		server.httpServer = &http.Server{
			Handler: server.ginServer,
		}
	} else if network == NetworkGRPC {
		server.grpcServer = grpc.NewServer()
	}
	for _, option := range options {
		option(server)
	}
	return server
}

// Server 网络服务器
type Server struct {
	*event
	network    Network      // 网络类型
	addr       string       // 侦听地址
	options    []Option     // 选项
	ginServer  *gin.Engine  // HTTP模式下的路由器
	httpServer *http.Server // HTTP模式下的服务器
	grpcServer *grpc.Server // GRPC模式下的服务器

	connections        *synchronization.Map[string, *Conn] // 所有在线的连接
	gServer            *gNet                               // TCP或UDP模式下的服务器
	messagePool        *synchronization.Pool[*message]     // 消息池
	messagePoolSize    int                                 // 消息池大小
	messageChannel     chan *message                       // 消息管道
	initMessageChannel bool                                // 消息管道是否已经初始化
	multiple           bool                                // 是否为多服务器模式下运行
	prod               bool                                // 是否为生产模式
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
		slf.connections = synchronization.NewMap[string, *Conn]()
		slf.initMessageChannel = true
		if slf.messagePoolSize <= 0 {
			slf.messagePoolSize = 1024
		}
		slf.messagePool = synchronization.NewPool[*message](slf.messagePoolSize,
			func() *message {
				return &message{}
			}, func(data *message) {
				data.t = 0
				data.attrs = nil
			},
		)
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
		go func() {
			slf.OnStartBeforeEvent()
			if err := slf.grpcServer.Serve(listener); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		}()
	case NetworkTCP, NetworkTCP4, NetworkTCP6, NetworkUdp, NetworkUdp4, NetworkUdp6, NetworkUnix:
		go connectionInitHandle(func() {
			slf.OnStartBeforeEvent()
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
			slf.OnStartBeforeEvent()
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
		if slf.prod {
			log.SetProd()
			gin.SetMode(gin.ReleaseMode)
		}
		go func() {
			slf.OnStartBeforeEvent()
			slf.httpServer.Addr = slf.addr
			if err := slf.httpServer.ListenAndServe(); err != nil {
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
			slf.OnStartBeforeEvent()
			if err := http.ListenAndServe(slf.addr, nil); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		}()
	default:
		return ErrCanNotSupportNetwork
	}

	if !slf.multiple {
		time.Sleep(500 * time.Millisecond)
		log.Info("Server", zap.String("Minotaur Server", "===================================================================="))
		log.Info("Server", zap.String("Minotaur Server", "RunningInfo"),
			zap.Any("network", slf.network),
			zap.String("listen", slf.addr),
		)
		log.Info("Server", zap.String("Minotaur Server", "===================================================================="))
		slf.OnStartFinishEvent()
		systemSignal := make(chan os.Signal, 1)
		signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-systemSignal:
			slf.Shutdown(nil)
		}
	} else {
		slf.OnStartFinishEvent()
	}

	return nil
}

// IsProd 是否为生产模式
func (slf *Server) IsProd() bool {
	return slf.prod
}

// IsDev 是否为开发模式
func (slf *Server) IsDev() bool {
	return !slf.prod
}

// Shutdown 停止运行服务器
func (slf *Server) Shutdown(err error) {
	if slf.initMessageChannel {
		if slf.gServer != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			if shutdownErr := gnet.Stop(ctx, fmt.Sprintf("%s://%s", slf.network, slf.addr)); shutdownErr != nil {
				log.Error("Server", zap.Error(shutdownErr))
			}
		}
		slf.connections.Range(func(connId string, conn *Conn) {
			conn.Close()
		})
		close(slf.messageChannel)
		slf.messagePool.Close()
		slf.initMessageChannel = false
	}
	if slf.grpcServer != nil {
		slf.grpcServer.Stop()
	}
	if slf.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if shutdownErr := slf.httpServer.Shutdown(ctx); shutdownErr != nil {
			log.Error("Server", zap.Error(shutdownErr))
		}
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
	if slf.ginServer == nil {
		panic(ErrNetworkOnlySupportHttp)
	}
	return slf.ginServer
}

// PushMessage 向服务器中写入特定类型的消息，需严格遵守消息属性要求
func (slf *Server) PushMessage(messageType MessageType, attrs ...any) {
	msg := slf.messagePool.Get()
	msg.t = messageType
	msg.attrs = attrs
	slf.messageChannel <- msg
}

// dispatchMessage 消息分发
func (slf *Server) dispatchMessage(msg *message) {
	defer func() {
		slf.messagePool.Release(msg)
		if err := recover(); err != nil {
			log.Error("Server", zap.String("MessageType", messageNames[msg.t]), zap.Any("MessageAttrs", msg.attrs), zap.Any("error", err))
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
