package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"go.uber.org/zap"
	"minotaur/utils/log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

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

type Server struct {
	*event
	network        Network
	addr           string
	httpServer     *gin.Engine
	gServer        *gServer
	messageChannel chan *message
}

// Run 使用特定网络模式运行服务器
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
	var connectionInitHandle = func() {
		slf.messageChannel = make(chan *message, 4096*1000)
		if slf.network != NetworkHttp && slf.network != NetworkWebsocket {
			slf.gServer = &gServer{Server: slf}
		}
		for message := range slf.messageChannel {
			slf.dispatchMessage(message)
		}
	}

	switch slf.network {
	case NetworkTCP, NetworkTCP4, NetworkTCP6, NetworkUdp, NetworkUdp4, NetworkUdp6, NetworkUnix:
		go connectionInitHandle()
		go func() {
			if err := gnet.Serve(slf.gServer, protoAddr); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		}()
	case NetworkKcp:
		listener, err := kcp.ListenWithOptions(slf.addr, nil, 0, 0)
		if err != nil {
			return err
		}
		go connectionInitHandle()
		go func() {
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
		}()
	case NetworkHttp:
		go func() {
			if err := slf.httpServer.Run(addr); err != nil {
				slf.PushMessage(MessageTypeError, err, MessageErrorActionShutdown)
			}
		}()
	case NetworkWebsocket:
		go connectionInitHandle()
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
	return nil
}

func (slf *Server) Shutdown(err error) {
	close(slf.messageChannel)
	if err != nil {
		log.Error("Server", zap.String("action", "shutdown"), zap.String("state", "exception"), zap.Error(err))
	} else {
		log.Info("Server", zap.String("action", "shutdown"), zap.String("state", "normal"))
	}
}

func (slf *Server) HttpRouter() gin.IRouter {
	if slf.httpServer == nil {
		panic(ErrNetworkOnlySupportHttp)
	}
	return slf.httpServer
}

func (slf *Server) PushMessage(messageType MessageType, attrs ...any) {
	slf.messageChannel <- &message{
		t:     messageType,
		attrs: attrs,
	}
}

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
