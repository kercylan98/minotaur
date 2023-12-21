package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/server/internal/logger"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/network"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/kercylan98/minotaur/utils/super"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// New 根据特定网络类型创建一个服务器
func New(network Network, options ...Option) *Server {
	server := &Server{
		runtime: &runtime{
			messagePoolSize: DefaultMessageBufferSize,
			packetWarnSize:  DefaultPacketWarnSize,
		},
		option:           &option{},
		network:          network,
		online:           concurrent.NewBalanceMap[string, *Conn](),
		closeChannel:     make(chan struct{}, 1),
		systemSignal:     make(chan os.Signal, 1),
		ctx:              context.Background(),
		dispatchers:      make(map[string]*dispatcher),
		dispatcherMember: map[string]map[string]*Conn{},
		currDispatcher:   map[string]*dispatcher{},
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
		server.ants, err = ants.NewPool(server.antsPoolSize, ants.WithLogger(new(logger.Ants)))
		if err != nil {
			panic(err)
		}
	}

	server.option = nil
	return server
}

// Server 网络服务器
type Server struct {
	*event                                                         // 事件
	*runtime                                                       // 运行时
	*option                                                        // 可选项
	ginServer                *gin.Engine                           // HTTP模式下的路由器
	httpServer               *http.Server                          // HTTP模式下的服务器
	grpcServer               *grpc.Server                          // GRPC模式下的服务器
	gServer                  *gNet                                 // TCP或UDP模式下的服务器
	multiple                 *MultipleServer                       // 多服务器模式下的服务器
	ants                     *ants.Pool                            // 协程池
	messagePool              *concurrent.Pool[*Message]            // 消息池
	ctx                      context.Context                       // 上下文
	online                   *concurrent.BalanceMap[string, *Conn] // 在线连接
	systemDispatcher         *dispatcher                           // 系统消息分发器
	network                  Network                               // 网络类型
	addr                     string                                // 侦听地址
	systemSignal             chan os.Signal                        // 系统信号
	closeChannel             chan struct{}                         // 关闭信号
	multipleRuntimeErrorChan chan error                            // 多服务器模式下的运行时错误
	messageLock              sync.RWMutex                          // 消息锁
	dispatcherLock           sync.RWMutex                          // 消息分发器锁
	isShutdown               atomic.Bool                           // 是否已关闭
	messageCounter           atomic.Int64                          // 消息计数器
	isRunning                bool                                  // 是否正在运行
	dispatchers              map[string]*dispatcher                // 消息分发器集合
	dispatcherMember         map[string]map[string]*Conn           // 消息分发器包含的连接
	currDispatcher           map[string]*dispatcher                // 当前连接所处消息分发器
}

// Run 使用特定地址运行服务器
//   - server.NetworkTcp (addr:":8888")
//   - server.NetworkTcp4 (addr:":8888")
//   - server.NetworkTcp6 (addr:":8888")
//   - server.NetworkUdp (addr:":8888")
//   - server.NetworkUdp4 (addr:":8888")
//   - server.NetworkUdp6 (addr:":8888")
//   - server.NetworkUnix (addr:"socketPath")
//   - server.NetworkHttp (addr:":8888")
//   - server.NetworkWebsocket (addr:":8888/ws")
//   - server.NetworkKcp (addr:":8888")
//   - server.NetworkNone (addr:"")
func (slf *Server) Run(addr string) error {
	if slf.network == NetworkNone {
		addr = "-"
	}
	if slf.event == nil {
		return ErrConstructed
	}
	slf.event.check()
	slf.addr = addr
	slf.systemDispatcher = generateDispatcher(serverSystemDispatcher, slf.dispatchMessage)
	var protoAddr = fmt.Sprintf("%s://%s", slf.network, slf.addr)
	var messageInitFinish = make(chan struct{}, 1)
	var connectionInitHandle = func(callback func()) {
		slf.messageLock.Lock()
		slf.messagePool = concurrent.NewPool[*Message](slf.messagePoolSize,
			func() *Message {
				return &Message{}
			},
			func(data *Message) {
				data.reset()
			},
		)
		slf.messageLock.Unlock()
		if slf.network != NetworkHttp && slf.network != NetworkWebsocket && slf.network != NetworkGRPC {
			slf.gServer = &gNet{Server: slf}
		}
		if callback != nil {
			go callback()
		}
		go func() {
			messageInitFinish <- struct{}{}
			slf.systemDispatcher.start()
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
				slf.PushErrorMessage(err, MessageErrorActionShutdown)
			}
		}()
	case NetworkTcp, NetworkTcp4, NetworkTcp6, NetworkUdp, NetworkUdp4, NetworkUdp6, NetworkUnix:
		go connectionInitHandle(func() {
			slf.isRunning = true
			slf.OnStartBeforeEvent()
			if err := gnet.Serve(slf.gServer, protoAddr,
				gnet.WithLogger(new(logger.GNet)),
				gnet.WithTicker(true),
				gnet.WithMulticore(true),
			); err != nil {
				slf.isRunning = false
				slf.PushErrorMessage(err, MessageErrorActionShutdown)
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

				go func(conn *Conn) {
					defer func() {
						if err := recover(); err != nil {
							e, ok := err.(error)
							if !ok {
								e = fmt.Errorf("%v", err)
							}
							conn.Close(e)
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
						slf.PushPacketMessage(conn, 0, buf[:n])
					}
				}(conn)
			}
		})
	case NetworkHttp:
		go func() {
			slf.isRunning = true
			slf.OnStartBeforeEvent()
			slf.httpServer.Addr = slf.addr
			gin.SetMode(gin.ReleaseMode)
			slf.ginServer.Use(func(c *gin.Context) {
				t := time.Now()
				c.Next()
				log.Info("Server", log.String("type", "http"),
					log.String("method", c.Request.Method), log.Int("status", c.Writer.Status()),
					log.String("ip", c.ClientIP()), log.String("path", c.Request.URL.Path),
					log.Duration("cost", time.Since(t)))
			})
			go connectionInitHandle(nil)
			if len(slf.certFile)+len(slf.keyFile) > 0 {
				if err := slf.httpServer.ListenAndServeTLS(slf.certFile, slf.keyFile); err != nil {
					slf.isRunning = false
					slf.PushErrorMessage(err, MessageErrorActionShutdown)
				}
			} else {
				if err := slf.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					slf.isRunning = false
					slf.PushErrorMessage(err, MessageErrorActionShutdown)
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
				conn.SetData(wsRequestKey, request)
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
						e, ok := err.(error)
						if !ok {
							e = fmt.Errorf("%v", err)
						}
						conn.Close(e)
					}
				}()
				for !conn.IsClosed() {
					if slf.websocketReadDeadline > 0 {
						if err := ws.SetReadDeadline(time.Now().Add(slf.websocketReadDeadline)); err != nil {
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
					if len(slf.supportMessageTypes) > 0 && !slf.supportMessageTypes[messageType] {
						panic(ErrWebsocketIllegalMessageType)
					}
					slf.PushPacketMessage(conn, messageType, packet)
				}
			})
			go func() {
				slf.isRunning = true
				slf.OnStartBeforeEvent()
				if len(slf.certFile)+len(slf.keyFile) > 0 {
					if err := http.ListenAndServeTLS(slf.addr, slf.certFile, slf.keyFile, nil); err != nil {
						slf.isRunning = false
						slf.PushErrorMessage(err, MessageErrorActionShutdown)
					}
				} else {
					if err := http.ListenAndServe(slf.addr, nil); err != nil {
						slf.isRunning = false
						slf.PushErrorMessage(err, MessageErrorActionShutdown)
					}
				}

			}()
		})
	default:
		return ErrCanNotSupportNetwork
	}

	if slf.multiple == nil && slf.network != NetworkKcp {
		kcp.SystemTimedSched.Close()
	}

	<-messageInitFinish
	close(messageInitFinish)
	messageInitFinish = nil
	if slf.multiple == nil {
		ip, _ := network.IP()
		log.Info("Server", log.String(serverMark, "===================================================================="))
		log.Info("Server", log.String(serverMark, "RunningInfo"),
			log.Any("network", slf.network),
			log.String("ip", ip.String()),
			log.String("listen", slf.addr),
		)
		log.Info("Server", log.String(serverMark, "===================================================================="))
		slf.OnStartFinishEvent()
		time.Sleep(time.Second)
		if !slf.isShutdown.Load() {
			slf.OnMessageReadyEvent()
		}

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
		time.Sleep(time.Second)
		if !slf.isShutdown.Load() {
			slf.OnMessageReadyEvent()
		}
	}

	return nil
}

// IsSocket 是否是 Socket 模式
func (slf *Server) IsSocket() bool {
	return slf.network == NetworkTcp || slf.network == NetworkTcp4 || slf.network == NetworkTcp6 ||
		slf.network == NetworkUdp || slf.network == NetworkUdp4 || slf.network == NetworkUdp6 ||
		slf.network == NetworkUnix || slf.network == NetworkKcp || slf.network == NetworkWebsocket
}

// RunNone 是 Run("") 的简写，仅适用于运行 NetworkNone 服务器
func (slf *Server) RunNone() error {
	return slf.Run(str.None)
}

// Context 获取服务器上下文
func (slf *Server) Context() context.Context {
	return slf.ctx
}

// TimeoutContext 获取服务器超时上下文，context.WithTimeout 的简写
func (slf *Server) TimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(slf.ctx, timeout)
}

// GetOnlineCount 获取在线人数
func (slf *Server) GetOnlineCount() int {
	return slf.online.Size()
}

// GetOnlineBotCount 获取在线机器人数量
func (slf *Server) GetOnlineBotCount() int {
	var count int
	slf.online.Range(func(id string, conn *Conn) bool {
		if conn.IsBot() {
			count++
		}
		return true
	})
	return count
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
	if err != nil {
		log.Error("Server", log.String("state", "shutdown"), log.Err(err))
	}
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
	if slf.gServer != nil && slf.isRunning {
		if shutdownErr := gnet.Stop(context.Background(), fmt.Sprintf("%s://%s", slf.network, slf.addr)); err != nil {
			log.Error("Server", log.Err(shutdownErr))
		}
	}
	if slf.ticker != nil {
		slf.ticker.Release()
	}
	if slf.ants != nil {
		slf.ants.Release()
		slf.ants = nil
	}
	slf.dispatcherLock.Lock()
	for s, d := range slf.dispatchers {
		d.close()
		delete(slf.dispatchers, s)
	}
	slf.dispatcherLock.Unlock()
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

// UseShunt 切换连接所使用的消息分流渠道，当分流渠道 name 不存在时将会创建一个新的分流渠道，否则将会加入已存在的分流渠道
//   - 默认情况下，所有连接都使用系统通道进行消息分发，当指定消息分流渠道时，将会使用指定的消息分流渠道进行消息分发
func (slf *Server) UseShunt(conn *Conn, name string) {
	slf.dispatcherLock.Lock()
	defer slf.dispatcherLock.Unlock()
	d, exist := slf.dispatchers[name]
	if !exist {
		d = generateDispatcher(name, slf.dispatchMessage)
		go d.start()
		slf.dispatchers[name] = d
	}

	curr, exist := slf.currDispatcher[conn.GetID()]
	if exist {
		if curr.name == name {
			return
		}

		delete(slf.dispatcherMember[curr.name], conn.GetID())
		if len(slf.dispatcherMember[curr.name]) == 0 {
			curr.close()
			delete(slf.dispatchers, curr.name)
		}
	}
	slf.currDispatcher[conn.GetID()] = d

	member, exist := slf.dispatcherMember[name]
	if !exist {
		member = map[string]*Conn{}
		slf.dispatcherMember[name] = member
	}

	member[conn.GetID()] = conn
}

// getConnDispatcher 获取连接所使用的消息分发器
func (slf *Server) getConnDispatcher(conn *Conn) *dispatcher {
	if conn == nil {
		return slf.systemDispatcher
	}
	slf.dispatcherLock.RLock()
	defer slf.dispatcherLock.RUnlock()
	d, exist := slf.currDispatcher[conn.GetID()]
	if exist {
		return d
	}
	return slf.systemDispatcher
}

// releaseDispatcher 关闭消息分发器
func (slf *Server) releaseDispatcher(conn *Conn) {
	if conn == nil {
		return
	}
	slf.dispatcherLock.Lock()
	defer slf.dispatcherLock.Unlock()
	d, exist := slf.currDispatcher[conn.GetID()]
	if exist {
		delete(slf.dispatcherMember[d.name], conn.GetID())
		if len(slf.dispatcherMember[d.name]) == 0 {
			d.close()
			delete(slf.dispatchers, d.name)
		}
		delete(slf.currDispatcher, conn.GetID())
	}
}

// pushMessage 向服务器中写入特定类型的消息，需严格遵守消息属性要求
func (slf *Server) pushMessage(message *Message) {
	if slf.messagePool.IsClose() || !slf.OnMessageExecBeforeEvent(message) {
		slf.messagePool.Release(message)
		return
	}
	var dispatcher *dispatcher
	switch message.t {
	case MessageTypePacket,
		MessageTypeShuntTicker, MessageTypeShuntAsync, MessageTypeShuntAsyncCallback,
		MessageTypeUniqueShuntAsync, MessageTypeUniqueShuntAsyncCallback,
		MessageTypeShunt:
		dispatcher = slf.getConnDispatcher(message.conn)
	case MessageTypeSystem, MessageTypeAsync, MessageTypeUniqueAsync, MessageTypeAsyncCallback, MessageTypeUniqueAsyncCallback, MessageTypeError, MessageTypeTicker:
		dispatcher = slf.systemDispatcher
	}
	if dispatcher == nil {
		return
	}
	if (message.t == MessageTypeUniqueShuntAsync || message.t == MessageTypeUniqueAsync) && dispatcher.unique(message.name) {
		slf.messagePool.Release(message)
		return
	}
	slf.messageCounter.Add(1)
	dispatcher.put(message)
}

func (slf *Server) low(message *Message, present time.Time, expect time.Duration, messageReplace ...string) {
	cost := time.Since(present)
	if cost > expect {
		if len(messageReplace) > 0 {
			for i, s := range messageReplace {
				message.marks = append(message.marks, log.String(fmt.Sprintf("Other-%d", i+1), s))
			}
		}
		var fields = make([]log.Field, 0, len(message.marks)+4)
		fields = append(fields, log.String("type", messageNames[message.t]), log.String("cost", cost.String()), log.String("message", message.String()))
		fields = append(fields, message.marks...)
		//fields = append(fields, log.Stack("stack"))
		log.Warn("ServerLowMessage", fields...)
		slf.OnMessageLowExecEvent(message, cost)
	}
}

// dispatchMessage 消息分发
func (slf *Server) dispatchMessage(dispatcher *dispatcher, msg *Message) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	if slf.deadlockDetect > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), slf.deadlockDetect)
		go func(ctx context.Context, msg *Message) {
			select {
			case <-ctx.Done():
				if err := ctx.Err(); errors.Is(err, context.DeadlineExceeded) {
					log.Warn("Server", log.String("MessageType", messageNames[msg.t]), log.String("Info", msg.String()), log.Any("SuspectedDeadlock", msg))
					slf.OnDeadlockDetectEvent(msg)
				}
			}
		}(ctx, msg)
	}

	present := time.Now()
	if msg.t != MessageTypeAsync && msg.t != MessageTypeUniqueAsync && msg.t != MessageTypeShuntAsync && msg.t != MessageTypeUniqueShuntAsync {
		defer func(msg *Message) {
			super.Handle(cancel)
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				log.Error("Server", log.String("MessageType", messageNames[msg.t]), log.String("Info", msg.String()), log.Any("error", err), log.String("stack", stack))
				fmt.Println(stack)
				e, ok := err.(error)
				if !ok {
					e = fmt.Errorf("%v", err)
				}
				slf.OnMessageErrorEvent(msg, e)
			}
			if msg.t == MessageTypeUniqueAsyncCallback || msg.t == MessageTypeUniqueShuntAsyncCallback {
				dispatcher.antiUnique(msg.name)
			}

			slf.low(msg, present, time.Millisecond*100)
			slf.messageCounter.Add(-1)

			if !slf.isShutdown.Load() {
				slf.messagePool.Release(msg)
			}
		}(msg)
	} else {
		if cancel != nil {
			defer cancel()
		}
	}

	switch msg.t {
	case MessageTypePacket:
		if !slf.OnConnectionPacketPreprocessEvent(msg.conn, msg.packet, func(newPacket []byte) {
			msg.packet = newPacket
		}) {
			slf.OnConnectionReceivePacketEvent(msg.conn, msg.packet)
		}
	case MessageTypeError:
		switch msg.errAction {
		case MessageErrorActionNone:
			log.Panic("Server", log.Err(msg.err))
		case MessageErrorActionShutdown:
			slf.shutdown(msg.err)
		default:
			log.Warn("Server", log.String("not support message error action", msg.errAction.String()))
		}
	case MessageTypeTicker, MessageTypeShuntTicker:
		msg.ordinaryHandler()
	case MessageTypeAsync, MessageTypeShuntAsync, MessageTypeUniqueAsync, MessageTypeUniqueShuntAsync:
		if err := slf.ants.Submit(func() {
			defer func() {
				if err := recover(); err != nil {
					if msg.t == MessageTypeUniqueAsync || msg.t == MessageTypeUniqueShuntAsync {
						dispatcher.antiUnique(msg.name)
					}
					stack := string(debug.Stack())
					log.Error("Server", log.String("MessageType", messageNames[msg.t]), log.Any("error", err), log.String("stack", stack))
					fmt.Println(stack)
					e, ok := err.(error)
					if !ok {
						e = fmt.Errorf("%v", err)
					}
					slf.OnMessageErrorEvent(msg, e)
				}
				super.Handle(cancel)
				slf.low(msg, present, time.Second)
				slf.messageCounter.Add(-1)

				if !slf.isShutdown.Load() {
					slf.messagePool.Release(msg)
				}
			}()
			var err error
			if msg.exceptionHandler != nil {
				err = msg.exceptionHandler()
			}
			if msg.errHandler != nil {
				if msg.conn == nil {
					if msg.t == MessageTypeUniqueAsync {
						slf.PushUniqueAsyncCallbackMessage(msg.name, err, msg.errHandler)
						return
					}
					slf.PushAsyncCallbackMessage(err, msg.errHandler)
					return
				}
				if msg.t == MessageTypeUniqueShuntAsync {
					slf.PushUniqueShuntAsyncCallbackMessage(msg.conn, msg.name, err, msg.errHandler)
					return
				}
				slf.PushShuntAsyncCallbackMessage(msg.conn, err, msg.errHandler)
				return
			}
			dispatcher.antiUnique(msg.name)
			if err != nil {
				log.Error("Server", log.String("MessageType", messageNames[msg.t]), log.Any("error", err), log.String("stack", string(debug.Stack())))
			}
		}); err != nil {
			panic(err)
		}
	case MessageTypeAsyncCallback, MessageTypeShuntAsyncCallback, MessageTypeUniqueAsyncCallback, MessageTypeUniqueShuntAsyncCallback:
		msg.errHandler(msg.err)
	case MessageTypeSystem, MessageTypeShunt:
		msg.ordinaryHandler()
	default:
		log.Warn("Server", log.String("not support message type", msg.t.String()))
	}
}

// PushSystemMessage 向服务器中推送 MessageTypeSystem 消息
//   - 系统消息仅包含一个可执行函数，将在系统分发器中执行
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Server) PushSystemMessage(handler func(), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToSystemMessage(handler, mark...))
}

// PushAsyncMessage 向服务器中推送 MessageTypeAsync 消息
//   - 异步消息将在服务器的异步消息队列中进行处理，处理完成 caller 的阻塞操作后，将会通过系统消息执行 callback 函数
//   - callback 函数将在异步消息处理完成后进行调用，无论过程是否产生 err，都将被执行，允许为 nil
//   - 需要注意的是，为了避免并发问题，caller 函数请仅处理阻塞操作，其他操作应该在 callback 函数中进行
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Server) PushAsyncMessage(caller func() error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToAsyncMessage(caller, callback, mark...))
}

// PushAsyncCallbackMessage 向服务器中推送 MessageTypeAsyncCallback 消息
//   - 异步消息回调将会通过一个接收 error 的函数进行处理，该函数将在系统分发器中执行
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Server) PushAsyncCallbackMessage(err error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToAsyncCallbackMessage(err, callback, mark...))
}

// PushShuntAsyncMessage 向特定分发器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
//   - 需要注意的是，当未指定 WithShunt 时，将会通过 PushAsyncMessage 进行转发
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Server) PushShuntAsyncMessage(conn *Conn, caller func() error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToShuntAsyncMessage(conn, caller, callback, mark...))
}

// PushShuntAsyncCallbackMessage 向特定分发器中推送 MessageTypeAsyncCallback 消息，消息执行与 MessageTypeAsyncCallback 一致
//   - 需要注意的是，当未指定 WithShunt 时，将会通过 PushAsyncCallbackMessage 进行转发
func (slf *Server) PushShuntAsyncCallbackMessage(conn *Conn, err error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToShuntAsyncCallbackMessage(conn, err, callback, mark...))
}

// PushPacketMessage 向服务器中推送 MessageTypePacket 消息
//   - 当存在 WithShunt 的选项时，将会根据选项中的 shuntMatcher 进行分发，否则将在系统分发器中处理消息
func (slf *Server) PushPacketMessage(conn *Conn, wst int, packet []byte, mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToPacketMessage(
		&Conn{wst: wst, connection: conn.connection},
		packet,
	))
}

// PushTickerMessage 向服务器中推送 MessageTypeTicker 消息
//   - 通过该函数推送定时消息，当消息触发时将在系统分发器中处理消息
//   - 可通过 timer.Ticker 或第三方定时器将执行函数(caller)推送到该消息中进行处理，可有效的避免线程安全问题
//   - 参数 name 仅用作标识该定时器名称
//
// 定时消息执行不会有特殊的处理，仅标记为定时任务，也就是允许将各类函数通过该消息发送处理，但是并不建议
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Server) PushTickerMessage(name string, caller func(), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToTickerMessage(name, caller, mark...))
}

// PushShuntTickerMessage 向特定分发器中推送 MessageTypeTicker 消息，消息执行与 MessageTypeTicker 一致
//   - 需要注意的是，当未指定 UseShunt 时，将会通过 PushTickerMessage 进行转发
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Server) PushShuntTickerMessage(conn *Conn, name string, caller func(), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToShuntTickerMessage(conn, name, caller, mark...))
}

// PushUniqueAsyncMessage 向服务器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
//   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息
func (slf *Server) PushUniqueAsyncMessage(unique string, caller func() error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToUniqueAsyncMessage(unique, caller, callback, mark...))
}

// PushUniqueAsyncCallbackMessage 向服务器中推送 MessageTypeAsyncCallback 消息，消息执行与 MessageTypeAsyncCallback 一致
func (slf *Server) PushUniqueAsyncCallbackMessage(unique string, err error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToUniqueAsyncCallbackMessage(unique, err, callback, mark...))
}

// PushUniqueShuntAsyncMessage 向特定分发器中推送 MessageTypeAsync 消息，消息执行与 MessageTypeAsync 一致
//   - 需要注意的是，当未指定 UseShunt 时，将会通过系统分流渠道进行转发
//   - 不同的是当上一个相同的 unique 消息未执行完成时，将会忽略该消息
func (slf *Server) PushUniqueShuntAsyncMessage(conn *Conn, unique string, caller func() error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToUniqueShuntAsyncMessage(conn, unique, caller, callback, mark...))
}

// PushUniqueShuntAsyncCallbackMessage 向特定分发器中推送 MessageTypeAsyncCallback 消息，消息执行与 MessageTypeAsyncCallback 一致
//   - 需要注意的是，当未指定 UseShunt 时，将会通过系统分流渠道进行转发
func (slf *Server) PushUniqueShuntAsyncCallbackMessage(conn *Conn, unique string, err error, callback func(err error), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToUniqueShuntAsyncCallbackMessage(conn, unique, err, callback, mark...))
}

// PushErrorMessage 向服务器中推送 MessageTypeError 消息
//   - 通过该函数推送错误消息，当消息触发时将在系统分发器中处理消息
//   - 参数 errAction 用于指定错误消息的处理方式，可选值为 MessageErrorActionNone 和 MessageErrorActionShutdown
//   - 参数 errAction 为 MessageErrorActionShutdown 时，将会停止服务器的运行
//   - mark 为可选的日志标记，当发生异常时，将会在日志中进行体现
func (slf *Server) PushErrorMessage(err error, errAction MessageErrorAction, mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToErrorMessage(err, errAction, mark...))
}

// PushShuntMessage 向特定分发器中推送 MessageTypeShunt 消息，消息执行与 MessageTypeSystem 一致，不同的是将会在特定分发器中执行
func (slf *Server) PushShuntMessage(conn *Conn, caller func(), mark ...log.Field) {
	slf.pushMessage(slf.messagePool.Get().castToShuntMessage(conn, caller, mark...))
}
