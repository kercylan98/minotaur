package server

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"github.com/kercylan98/minotaur/toolkit/nexus/brokers"
	messageEvents "github.com/kercylan98/minotaur/toolkit/nexus/events"
	"github.com/kercylan98/minotaur/toolkit/nexus/queues"
	"github.com/kercylan98/minotaur/toolkit/random"
	"github.com/panjf2000/ants/v2"
	"reflect"
	"time"
)

type Server interface {
	Events

	// Run 运行服务器
	Run() error

	// Shutdown 关闭服务器
	Shutdown() error

	// GetStatus 获取服务器状态
	GetStatus() *State

	// PublishMessage 发布消息到指定 topic，该函数将被允许使用自定义的消息事件
	PublishMessage(topic string, event nexus.Event[int, string])

	// PublishSystemMessage 发布消息到系统队列，该函数将被允许使用自定义的消息事件
	PublishSystemMessage(event nexus.Event[int, string])

	// PublishSyncMessage 发布同步消息到指定 topic
	PublishSyncMessage(topic string, handler messageEvents.SynchronousHandler)

	// PublishAsyncMessage 发布异步消息到指定 topic，当包含多个 callback 时，仅首个生效
	PublishAsyncMessage(topic string, handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler)

	// PublishSystemSyncMessage 发布同步消息到系统队列
	PublishSystemSyncMessage(handler messageEvents.SynchronousHandler)

	// PublishSystemAsyncMessage 发布异步消息到系统队列
	PublishSystemAsyncMessage(handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler)
}

type server struct {
	*controller                           // 实现了 Controller 接口的控制器数据结构，用于暴露 Server 对用户非公开的接口信息，适用于功能性的拓展
	*events                               // 实现了 Events 接口的事件数据结构，用于处理服务器事件
	*Options                              // 服务器配置
	queue       string                    // 系统消息队列名称，该队列将随机产生
	ants        *ants.Pool                // 异步池
	state       *State                    // 服务器状态
	notify      *notify                   // 通知器
	ctx         context.Context           // 服务器全局上下文
	cancel      context.CancelFunc        // 服务器全局上下文取消函数
	network     Network                   // 服务器使用的网络接口
	broker      nexus.Broker[int, string] // 服务器使用的消息队列
}

func NewServer(network Network, options ...*Options) Server {
	srv := &server{
		network: network,
		Options: DefaultOptions(),
		queue:   fmt.Sprintf("%s:%s", reflect.TypeOf(new(server)).String(), random.HostName()),
	}
	srv.ctx, srv.cancel = context.WithCancel(context.Background())
	srv.notify = new(notify).init(srv)
	srv.controller = new(controller).init(srv)
	srv.events = new(events).init(srv)
	srv.state = new(State).init(srv)
	srv.Options.init(srv).Apply(options...)
	srv.broker = brokers.NewSparseGoroutine(srv.Options.GetSparseGoroutineNum(), func(index int) nexus.Queue[int, string] {
		return queues.NewNonBlockingRW[int, string](index, srv.Options.GetServerMessageChannelSize(), srv.Options.GetServerMessageBufferInitialSize())
	}, func(handler nexus.EventExecutor) {
		handler()
	})
	antsPool, err := ants.NewPool(ants.DefaultAntsPoolSize, ants.WithOptions(ants.Options{
		ExpiryDuration: 10 * time.Second,
		Nonblocking:    true,
		//Logger:         &antsLogger{logging.GetDefaultLogger()},
		//PanicHandler: func(i interface{}) {
		//logging.Errorf("goroutine pool panic: %v", i)
		//},
	}))
	if err != nil {
		panic(err)
	}
	srv.ants = antsPool
	return srv
}

func (s *server) Run() (err error) {
	go s.broker.Run()
	if err = s.network.OnSetup(s.ctx, s); err != nil {
		return
	}

	ip, _ := network.IP()
	s.state.onLaunched(ip.String(), time.Now())
	s.Options.getManyOptions(func(opt *Options) {
		opt.logger.Info("Minotaur Server", log.String("", "============================================================================"))
		opt.logger.Info("Minotaur Server", log.String("", "RunningInfo"), log.String("network", reflect.TypeOf(s.network).String()), log.String("listen", fmt.Sprintf("%s://%s%s", s.network.Schema(), ip.String(), s.network.Address())))
		opt.logger.Info("Minotaur Server", log.String("", "============================================================================"))
		opt.logger.Info("Minotaur Server", log.String("", reflect.TypeOf(s.broker).String()), log.Int("num", opt.sparseGoroutineNum))
	})
	go func(s *server) {
		if err = s.network.OnRun(); err != nil {
			panic(err)
		}
	}(s)

	s.Options.active()
	s.notify.run()
	return
}

func (s *server) Shutdown() (err error) {
	s.GetLogger().Info("Minotaur Server", log.String("", "ShutdownInfo"), log.String("state", "start"))
	defer func(startAt time.Time) {
		s.GetLogger().Info("Minotaur Server", log.String("", "ShutdownInfo"), log.String("state", "done"), log.String("cost", time.Since(startAt).String()))
	}(time.Now())

	defer s.cancel()
	s.events.onShutdown()
	DisableHttpPProf()
	err = s.network.OnShutdown()
	s.broker.Close()
	return
}

func (s *server) getSysQueue() string {
	return s.queue
}

func (s *server) PublishMessage(topic string, event nexus.Event[int, string]) {
	if err := s.broker.Publish(topic, event); err != nil {
		s.GetLogger().Warn("Minotaur Server", log.String("topic", topic), log.String("error", err.Error()))
	}
}

func (s *server) PublishSystemMessage(event nexus.Event[int, string]) {
	s.PublishMessage(s.getSysQueue(), event)
}

func (s *server) PublishSyncMessage(topic string, handler messageEvents.SynchronousHandler) {
	s.PublishMessage(topic, messageEvents.Synchronous[int, string](handler))
}

func (s *server) PublishAsyncMessage(topic string, handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler) {
	s.PublishMessage(topic, messageEvents.Asynchronous[int, string](func(ctx context.Context, f func(context.Context)) {
		if err := s.ants.Submit(func() {
			f(ctx)
		}); err != nil {
			s.GetLogger().Warn("Minotaur Server", log.String("topic", topic), log.String("error", err.Error()))
			go f(ctx)
		}
	}, handler, collection.FindFirstOrDefaultInSlice(callback, nil)))
}

func (s *server) PublishSystemSyncMessage(handler messageEvents.SynchronousHandler) {
	s.PublishSyncMessage(s.getSysQueue(), handler)
}

func (s *server) PublishSystemAsyncMessage(handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler) {
	s.PublishAsyncMessage(s.getSysQueue(), handler, callback...)
}

func (s *server) GetStatus() *State {
	return s.state.Status()
}
