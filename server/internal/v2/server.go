package server

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"github.com/kercylan98/minotaur/toolkit/nexus/brokers"
	messageEvents "github.com/kercylan98/minotaur/toolkit/nexus/events"
	"github.com/kercylan98/minotaur/toolkit/nexus/queues"
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/panjf2000/ants/v2"
	"reflect"
	"time"

	"github.com/kercylan98/minotaur/utils/network"
)

type Server interface {
	Events

	// Run 运行服务器
	Run() error

	// Shutdown 关闭服务器
	Shutdown() error

	// GetStatus 获取服务器状态
	GetStatus() *State

	// PublishSyncMessage 发布同步消息
	PublishSyncMessage(topic string, handler messageEvents.SynchronousHandler)

	// PublishAsyncMessage 发布异步消息，当包含多个 callback 时，仅首个生效
	PublishAsyncMessage(topic string, handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler)
}

type server struct {
	*controller
	*events
	*Options
	queue   string
	ants    *ants.Pool
	state   *State
	notify  *notify
	ctx     context.Context
	cancel  context.CancelFunc
	network Network
	broker  nexus.Broker[int, string]
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
	srv.broker = brokers.NewSparseGoroutine(func(index int) nexus.Queue[int, string] {
		return queues.NewNonBlockingRW[int, string](index, 1024*8, 1024)
	}, func(handler nexus.EventExecutor) {
		handler()
	})
	srv.Options.init(srv).Apply(options...)
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
	DisableHttpPProf()
	s.events.onShutdown()
	err = s.network.OnShutdown()
	s.broker.Close()
	return
}

func (s *server) getSysQueue() string {
	return s.queue
}

func (s *server) PublishMessage(topic string, event nexus.Event[int, string]) {
	s.broker.Publish(topic, event)
}

func (s *server) PublishSyncMessage(topic string, handler messageEvents.SynchronousHandler) {
	s.PublishMessage(topic, messageEvents.Synchronous[int, string](handler))
}

func (s *server) PublishAsyncMessage(topic string, handler messageEvents.AsynchronousHandler, callback ...messageEvents.AsynchronousCallbackHandler) {
	s.PublishMessage(topic, messageEvents.Asynchronous[int, string](func(ctx context.Context, f func(context.Context)) {
		s.ants.Submit(func() {
			f(ctx)
		})
	}, handler, collection.FindFirstOrDefaultInSlice(callback, nil)))
}

func (s *server) GetStatus() *State {
	return s.state.Status()
}
