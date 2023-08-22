package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/kercylan98/minotaur/utils/super"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/kercylan98/minotaur/utils/times"
	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pkg/logging"
	"github.com/xtaci/kcp-go/v5"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
)

// New 根据特定网络类型创建一个服务器
func New(network Network, options ...Option) *Server {
	server := &Server{
		runtime:      &runtime{messagePoolSize: DefaultMessageBufferSize, messageChannelSize: DefaultMessageChannelSize},
		option:       &option{},
		network:      network,
		online:       concurrent.NewBalanceMap[string, *Conn](),
		closeChannel: make(chan struct{}, 1),
		systemSignal: make(chan os.Signal, 1),
	}
	server.event = newEvent(server)

	switch network {
	case NetworkHttp:
		server.ginServer = gin.New()
		server.httpServer = &http.Server{
			Handler: server.ginServer,
		}
	case NetworkGRPC:
		server.grpcServer = grpc.NewServer()
	case NetworkWebsocket:
		server.websocketReadDeadline = DefaultWebsocketReadDeadline
	}

	for _, option := range options {
		option(server)
	}

	if !server.disableAnts {
		if server.antsPoolSize <= 0 {
			server.antsPoolSize = DefaultAsyncPoolSize
		}
		var err error
		server.ants, err = ants.NewPool(server.antsPoolSize, ants.WithLogger(log.GetLogger()))
		if err != nil {
			panic(err)
		}
	}

	server.option = nil
	return server
}

// Server 网络服务器
type Server struct {
	*event                                                                     // 事件
	*runtime                                                                   // 运行时
	*option                                                                    // 可选项
	network                  Network                                           // 网络类型
	addr                     string                                            // 侦听地址
	systemSignal             chan os.Signal                                    // 系统信号
	online                   *concurrent.BalanceMap[string, *Conn]             // 在线连接
	ginServer                *gin.Engine                                       // HTTP模式下的路由器
	httpServer               *http.Server                                      // HTTP模式下的服务器
	grpcServer               *grpc.Server                                      // GRPC模式下的服务器
	gServer                  *gNet                                             // TCP或UDP模式下的服务器
	isRunning                bool                                              // 是否正在运行
	isShutdown               atomic.Bool                                       // 是否已关闭
	closeChannel             chan struct{}                                     // 关闭信号
	ants                     *ants.Pool                                        // 协程池
	messagePool              *concurrent.Pool[*Message]                        // 消息池
	messageChannel           chan *Message                                     // 消息管道
	multiple                 *MultipleServer                                   // 多服务器模式下的服务器
	multipleRuntimeErrorChan chan error                                        // 多服务器模式下的运行时错误
	runMode                  RunMode                                           // 运行模式
	shuntChannels            *concurrent.BalanceMap[int64, chan *Message]      // 分流管道
	channelGenerator         func(guid int64) chan *Message                    // 消息管道生成器
	shuntMatcher             func(conn *Conn) (guid int64, allowToCreate bool) // 分流管道匹配器
	messageCounter           atomic.Int64                                      // 消息计数器
}

// Run 使用特定地址运行服务器
//
//		server.NetworkTcp (addr:":8888")
//		server.NetworkTcp4 (addr:":8888")
//		server.NetworkTcp6 (addr:":8888")
//		server.NetworkUdp (addr:":8888")
//		server.NetworkUdp4 (addr:":8888")
//		server.NetworkUdp6 (addr:":8888")
//		server.NetworkUnix (addr:"socketPath")
//		server.NetworkHttp (addr:":8888")
//		server.NetworkWebsocket (addr:":8888/ws")
//		server.NetworkKcp (addr:":8888")
//	 server.NetworkNone (addr:"")
func (slf *Server) Run(addr string) error {
	if slf.network == NetworkNone {
		addr = "-"
	}
	if slf.event == nil {
		return ErrConstructed
	}
	slf.event.check()
	slf.addr = addr
	var protoAddr = fmt.Sprintf("%s://%s", slf.network, slf.addr)
	var messageInitFinish = make(chan struct{}, 1)
	var connectionInitHandle = func(callback func()) {
		slf.messagePool = concurrent.NewPool[*Message](slf.messagePoolSize,
			func() *Message {
				return &Message{}
			},
			func(data *Message) {
				data.t = 0
				data.attrs = nil
			},
		)
		slf.messageChannel = make(chan *Message, slf.messageChannelSize)
		if slf.network != NetworkHttp && slf.network != NetworkWebsocket && slf.network != NetworkGRPC {
			slf.gServer = &gNet{Server: slf}
		}
		if callback != nil {
			go callback()
		}
		go func() {
			messageInitFinish <- struct{}{}
			for message := range slf.messageChannel {
				slf.dispatchMessage(message)
			}
		}()
	}

	switch slf.network {
	case NetworkNone:
		go connectionInitHandle(func() {
			slf.isRunning = true
			slf.OnStartBeforeEvent()
		})
	case NetworkGRPC:
		listener, err := net.Listen(string(NetworkTcp), slf.addr)
		if err != nil {
			return err
		}
		go connectionInitHandle(nil)
		go func() {
			slf.isRunning = true
			slf.OnStartBeforeEvent()
			if err := slf.grpcServer.Serve(listener); err != nil {
				slf.isRunning = false
				PushErrorMessage(slf, err, MessageErrorActionShutdown)
			}
		}()
	case NetworkTcp, NetworkTcp4, NetworkTcp6, NetworkUdp, NetworkUdp4, NetworkUdp6, NetworkUnix:
		go connectionInitHandle(func() {
			slf.isRunning = true
			slf.OnStartBeforeEvent()
			if err := gnet.Serve(slf.gServer, protoAddr,
				gnet.WithLogger(log.GetLogger()),
				gnet.WithLogLevel(super.If(slf.runMode == RunModeProd, logging.ErrorLevel, logging.DebugLevel)),
				gnet.WithTicker(true),
				gnet.WithMulticore(true),
			); err != nil {
				slf.isRunning = false
				PushErrorMessage(slf, err, MessageErrorActionShutdown)
			}
		})
	case NetworkKcp:
		listener, err := kcp.ListenWithOptions(slf.addr, nil, 0, 0)
		if err != nil {
			return err
		}
		go connectionInitHandle(func() {
			slf.isRunning = true
			slf.OnStartBeforeEvent()
			for {
				session, err := listener.AcceptKCP()
				if err != nil {
					continue
				}

				conn := newKcpConn(slf, session)
				slf.OnConnectionOpenedEvent(conn)
				slf.OnConnectionOpenedAfterEvent(conn)

				go func(conn *Conn) {
					defer func() {
						if err := recover(); err != nil {
							slf.OnConnectionClosedEvent(conn, err)
						}
					}()

					buf := make([]byte, 4096)
					for {
						n, err := conn.kcp.Read(buf)
						if err != nil {
							panic(err)
						}
						PushPacketMessage(slf, conn, buf[:n])
					}
				}(conn)
			}
		})
	case NetworkHttp:
		switch slf.runMode {
		case RunModeDev:
			gin.SetMode(gin.DebugMode)
		case RunModeTest:
			gin.SetMode(gin.TestMode)
		case RunModeProd:
			gin.SetMode(gin.ReleaseMode)
		}
		go func() {
			slf.isRunning = true
			slf.OnStartBeforeEvent()
			slf.httpServer.Addr = slf.addr
			go connectionInitHandle(nil)
			if len(slf.certFile)+len(slf.keyFile) > 0 {
				if err := slf.httpServer.ListenAndServeTLS(slf.certFile, slf.keyFile); err != nil {
					slf.isRunning = false
					PushErrorMessage(slf, err, MessageErrorActionShutdown)
				}
			} else {
				if err := slf.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					slf.isRunning = false
					PushErrorMessage(slf, err, MessageErrorActionShutdown)
				}
			}

		}()
	case NetworkWebsocket:
		go connectionInitHandle(func() {
			var pattern string
			var index = strings.Index(addr, "/")
			if index == -1 {
				pattern = "/"
			} else {
				pattern = addr[index:]
				slf.addr = slf.addr[:index]
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
				if slf.websocketCompression > 0 {
					_ = ws.SetCompressionLevel(slf.websocketCompression)
				}
				ws.EnableWriteCompression(slf.websocketWriteCompression)
				conn := newWebsocketConn(slf, ws, ip)
				for k, v := range request.URL.Query() {
					if len(v) == 1 {
						conn.SetData(k, v[0])
					} else {
						conn.SetData(k, v)
					}
				}
				slf.OnConnectionOpenedEvent(conn)

				defer func() {
					if err := recover(); err != nil {
						slf.OnConnectionClosedEvent(conn, err)
					}
				}()
				for {
					if err := ws.SetReadDeadline(super.If(slf.websocketReadDeadline <= 0, times.Zero, time.Now().Add(slf.websocketReadDeadline))); err != nil {
						panic(err)
					}
					messageType, packet, readErr := ws.ReadMessage()
					if readErr != nil {
						panic(readErr)
					}
					if len(slf.supportMessageTypes) > 0 && !slf.supportMessageTypes[messageType] {
						panic(ErrWebsocketIllegalMessageType)
					}
					PushPacketMessage(slf, conn, append(packet, byte(messageType)))
				}
			})
			go func() {
				slf.isRunning = true
				slf.OnStartBeforeEvent()
				if len(slf.certFile)+len(slf.keyFile) > 0 {
					if err := http.ListenAndServeTLS(slf.addr, slf.certFile, slf.keyFile, nil); err != nil {
						slf.isRunning = false
						PushErrorMessage(slf, err, MessageErrorActionShutdown)
					}
				} else {
					if err := http.ListenAndServe(slf.addr, nil); err != nil {
						slf.isRunning = false
						PushErrorMessage(slf, err, MessageErrorActionShutdown)
					}
				}

			}()
		})
	default:
		return ErrCanNotSupportNetwork
	}

	<-messageInitFinish
	close(messageInitFinish)
	messageInitFinish = nil
	if slf.multiple == nil {
		log.Info("Server", log.String(serverMark, "===================================================================="))
		log.Info("Server", log.String(serverMark, "RunningInfo"),
			log.Any("network", slf.network),
			log.String("listen", slf.addr),
		)
		log.Info("Server", log.String(serverMark, "===================================================================="))
		slf.OnStartFinishEvent()

		signal.Notify(slf.systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-slf.systemSignal:
			slf.shutdown(nil)
		}

		select {
		case <-slf.closeChannel:
			close(slf.closeChannel)
		}
	} else {
		slf.OnStartFinishEvent()
	}

	return nil
}

// RunNone 是 Run("") 的简写，仅适用于运行 NetworkNone 服务器
func (slf *Server) RunNone() error {
	return slf.Run(str.None)
}

// GetOnlineCount 获取在线人数
func (slf *Server) GetOnlineCount() int {
	return slf.online.Size()
}

// GetOnline 获取在线连接
func (slf *Server) GetOnline(id string) *Conn {
	return slf.online.Get(id)
}

// GetOnlineAll 获取所有在线连接
func (slf *Server) GetOnlineAll() map[string]*Conn {
	return slf.online.Map()
}

// IsOnline 是否在线
func (slf *Server) IsOnline(id string) bool {
	return slf.online.Exist(id)
}

// CloseConn 关闭连接
func (slf *Server) CloseConn(id string) {
	if conn, exist := slf.online.GetExist(id); exist {
		conn.Close()
	}
}

// GetID 获取服务器id
func (slf *Server) GetID() int64 {
	if slf.cross == nil {
		panic(ErrNoSupportCross)
	}
	return slf.id
}

// Ticker 获取服务器定时器
func (slf *Server) Ticker() *timer.Ticker {
	if slf.ticker == nil {
		panic(ErrNoSupportTicker)
	}
	return slf.ticker
}

// Shutdown 主动停止运行服务器
func (slf *Server) Shutdown() {
	slf.systemSignal <- syscall.SIGQUIT
}

// shutdown 停止运行服务器
func (slf *Server) shutdown(err error) {
	slf.isShutdown.Store(true)
	for slf.messageCounter.Load() > 0 {
		log.Info("Server", log.Any("network", slf.network), log.String("listen", slf.addr),
			log.String("action", "shutdown"), log.String("state", "waiting"), log.Int64("message", slf.messageCounter.Load()))
		time.Sleep(time.Second)
	}
	if slf.multiple == nil {
		slf.OnStopEvent()
	}
	defer func() {
		if slf.multipleRuntimeErrorChan != nil {
			slf.multipleRuntimeErrorChan <- err
		}
	}()
	if slf.ticker != nil {
		slf.ticker.Release()
	}
	if slf.ants != nil {
		slf.ants.Release()
		slf.ants = nil
	}
	for _, cross := range slf.cross {
		cross.Release()
	}
	if slf.messageChannel != nil {
		close(slf.messageChannel)
		slf.messagePool.Close()
		slf.messageChannel = nil
	}
	if slf.shuntChannels != nil {
		slf.shuntChannels.Range(func(key int64, c chan *Message) bool {
			close(c)
			return false
		})
		slf.shuntChannels.Clear()
		slf.shuntChannels = nil
	}
	if slf.grpcServer != nil && slf.isRunning {
		slf.grpcServer.GracefulStop()
	}
	if slf.httpServer != nil && slf.isRunning {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if shutdownErr := slf.httpServer.Shutdown(ctx); shutdownErr != nil {
			log.Error("Server", log.Err(shutdownErr))
		}
	}
	if slf.gServer != nil && slf.isRunning {
		if shutdownErr := gnet.Stop(context.Background(), fmt.Sprintf("%s://%s", slf.network, slf.addr)); err != nil {
			log.Error("Server", log.Err(shutdownErr))
		}
	}

	if err != nil {
		if slf.multiple != nil {
			slf.multiple.RegExitEvent(func() {
				log.Panic("Server", log.Any("network", slf.network), log.String("listen", slf.addr),
					log.String("action", "shutdown"), log.String("state", "exception"), log.Err(err))
			})
			for i, server := range slf.multiple.servers {
				if server.addr == slf.addr {
					slf.multiple.servers = append(slf.multiple.servers[:i], slf.multiple.servers[i+1:]...)
					break
				}
			}
		} else {
			log.Panic("Server", log.Any("network", slf.network), log.String("listen", slf.addr),
				log.String("action", "shutdown"), log.String("state", "exception"), log.Err(err))
		}
	} else {
		log.Info("Server", log.Any("network", slf.network), log.String("listen", slf.addr),
			log.String("action", "shutdown"), log.String("state", "normal"))
	}
	slf.closeChannel <- struct{}{}
}

// GRPCServer 当网络类型为 NetworkGRPC 时将被允许获取 grpc 服务器，否则将会发生 panic
func (slf *Server) GRPCServer() *grpc.Server {
	if slf.grpcServer == nil {
		panic(ErrNetworkOnlySupportGRPC)
	}
	return slf.grpcServer
}

// HttpRouter 当网络类型为 NetworkHttp 时将被允许获取路由器进行路由注册，否则将会发生 panic
//   - 通过该函数注册的路由将无法在服务器关闭时正常等待请求结束
//
// Deprecated: 从 Minotaur 0.0.29 开始，由于设计原因已弃用，该函数将直接返回 *gin.Server 对象，导致无法正常的对请求结束时进行处理
func (slf *Server) HttpRouter() gin.IRouter {
	if slf.ginServer == nil {
		panic(ErrNetworkOnlySupportHttp)
	}
	return slf.ginServer
}

// HttpServer 替代 HttpRouter 的函数，返回一个 *Http[*HttpContext] 对象
//   - 通过该函数注册的路由将在服务器关闭时正常等待请求结束
//   - 如果需要自行包装 Context 对象，可以使用 NewHttpHandleWrapper 方法
func (slf *Server) HttpServer() *Http[*HttpContext] {
	if slf.ginServer == nil {
		panic(ErrNetworkOnlySupportHttp)
	}
	return NewHttpHandleWrapper(slf, func(ctx *gin.Context) *HttpContext {
		return NewHttpContext(ctx)
	})
}

// GetMessageCount 获取当前服务器中消息的数量
func (slf *Server) GetMessageCount() int64 {
	return slf.messageCounter.Load()
}

// ShuntChannelFreed 释放分流通道
func (slf *Server) ShuntChannelFreed(channelGuid int64) {
	if slf.shuntChannels == nil {
		return
	}
	channel, exist := slf.shuntChannels.GetExist(channelGuid)
	if exist {
		close(channel)
		slf.shuntChannels.Delete(channelGuid)
		slf.OnShuntChannelClosedEvent(channelGuid)
	}
}

// pushMessage 向服务器中写入特定类型的消息，需严格遵守消息属性要求
func (slf *Server) pushMessage(message *Message) {
	if slf.messagePool.IsClose() {
		slf.messagePool.Release(message)
		return
	}
	if slf.isShutdown.Load() {
		return
	}
	if slf.shuntChannels != nil && message.t == MessageTypePacket {
		conn := message.attrs[0].(*Conn)
		channelGuid, allowToCreate := slf.shuntMatcher(conn)
		channel, exist := slf.shuntChannels.GetExist(channelGuid)
		if !exist && allowToCreate {
			channel = slf.channelGenerator(channelGuid)
			slf.shuntChannels.Set(channelGuid, channel)
			go func(channel chan *Message) {
				for message := range channel {
					slf.dispatchMessage(message)
				}
			}(channel)
			defer slf.OnShuntChannelCreatedEvent(channelGuid)
		}
		if channel != nil {
			channel <- message
			return
		}
	}
	slf.messageChannel <- message
}

func (slf *Server) low(message *Message, present time.Time, expect time.Duration, messageReplace ...string) {
	cost := time.Since(present)
	if cost > expect {
		var m = "unknown"
		if message != nil {
			m = message.String()
		} else if len(messageReplace) > 0 {
			m = messageReplace[0]
		}
		log.Warn("Server", log.String("type", "low-message"), log.String("cost", cost.String()), log.String("message", m), log.Stack("stack"))
		slf.OnMessageLowExecEvent(message, cost)
	}
}

// dispatchMessage 消息分发
func (slf *Server) dispatchMessage(msg *Message) {
	slf.messageCounter.Add(1)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	if slf.deadlockDetect > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), slf.deadlockDetect)
		go func() {
			select {
			case <-ctx.Done():
				if err := ctx.Err(); err == context.DeadlineExceeded {
					log.Warn("Server", log.String("MessageType", messageNames[msg.t]), log.Any("SuspectedDeadlock", msg.attrs))
				}
			}
		}()
	}

	present := time.Now()
	defer func() {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			log.Error("Server", log.String("MessageType", messageNames[msg.t]), log.Any("MessageAttrs", msg.attrs), log.Any("error", err), log.String("stack", stack))
			fmt.Println(stack)
			if e, ok := err.(error); ok {
				slf.OnMessageErrorEvent(msg, e)
			}
		}

		if msg.t == MessageTypeAsync {
			return
		}

		super.Handle(cancel)
		slf.low(msg, present, time.Millisecond*100)
		slf.messageCounter.Add(-1)

		if !slf.isShutdown.Load() {
			slf.messagePool.Release(msg)
		}

	}()
	var attrs = msg.attrs
	switch msg.t {
	case MessageTypePacket:
		var conn = attrs[0].(*Conn)
		var packet = attrs[1].([]byte)
		var wst = int(packet[len(packet)-1])
		if len(packet) >= 2 {
			var ct = packet[len(packet)-2]
			if ct == 0xff {
				var gp GP
				if err := super.UnmarshalJSON(packet[:len(packet)-2], &gp); err != nil {
					panic(err)
				}
				packet = gp.D
				conn = newGatewayConn(conn, gp.C)
			} else {
				packet = packet[:len(packet)-1]
			}
		} else {
			packet = packet[:len(packet)-1]
		}

		if !slf.OnConnectionPacketPreprocessEvent(conn, packet, func(newPacket []byte) { packet = newPacket }) {
			slf.OnConnectionReceivePacketEvent(conn, Packet{Data: packet, WebsocketType: wst})
		}
	case MessageTypeError:
		err, action := attrs[0].(error), attrs[1].(MessageErrorAction)
		switch action {
		case MessageErrorActionNone:
			log.Panic("Server", log.Err(err))
		case MessageErrorActionShutdown:
			slf.shutdown(err)
		default:
			log.Warn("Server", log.String("not support message error action", action.String()))
		}
	case MessageTypeCross:
		slf.OnReceiveCrossPacketEvent(attrs[0].(int64), attrs[1].([]byte))
	case MessageTypeTicker:
		attrs[0].(func())()
	case MessageTypeAsync:
		handle := attrs[0].(func() error)
		callback, cb := attrs[1].(func(err error))
		if err := slf.ants.Submit(func() {
			defer func() {
				if err := recover(); err != nil {
					stack := string(debug.Stack())
					log.Error("Server", log.String("MessageType", messageNames[msg.t]), log.Any("error", err), log.String("stack", stack))
					fmt.Println(stack)
					if e, ok := err.(error); ok {
						slf.OnMessageErrorEvent(msg, e)
					}
				}
				super.Handle(cancel)
				slf.low(msg, present, time.Second)
				slf.messageCounter.Add(-1)

				if !slf.isShutdown.Load() {
					slf.messagePool.Release(msg)
				}
			}()
			err := handle()
			if cb && callback != nil {
				acm := slf.messagePool.Get()
				acm.t = MessageTypeAsyncCallback
				if len(attrs) > 2 {
					acm.attrs = append([]any{func() { callback(err) }}, attrs[2:]...)
				} else {
					acm.attrs = []any{func() { callback(err) }}
				}
				slf.pushMessage(acm)
			} else if err != nil {
				log.Error("Server", log.String("MessageType", messageNames[msg.t]), log.Any("error", err), log.String("stack", string(debug.Stack())))
			}
		}); err != nil {
			panic(err)
		}
	case MessageTypeAsyncCallback:
		attrs[0].(func())()
	case MessageTypeSystem:
		attrs[0].(func())()
	default:
		log.Warn("Server", log.String("not support message type", msg.t.String()))
	}
}
