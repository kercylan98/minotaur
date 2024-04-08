package server

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/server/internal/v2/message"
	"github.com/kercylan98/minotaur/server/internal/v2/message/messages"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/panjf2000/ants/v2"
	"reflect"
	"time"

	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"github.com/kercylan98/minotaur/utils/network"
)

type Server interface {
	Events
	message.Broker[Producer, string]

	// Run 运行服务器
	Run() error

	// Shutdown 关闭服务器
	Shutdown() error

	// GetStatus 获取服务器状态
	GetStatus() *State

	// PublishSyncMessage 发布同步消息
	PublishSyncMessage(queue string, handler messages.SynchronousHandler[Producer, string, Server])

	// PublishAsyncMessage 发布异步消息，当包含多个 callback 时，仅首个生效
	PublishAsyncMessage(queue string, handler messages.AsynchronousHandler[Producer, string, Server], callback ...messages.AsynchronousCallbackHandler[Producer, string, Server])
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
	reactor *reactor.Reactor[string, message.Message[Producer, string]]
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
	srv.reactor = reactor.NewReactor[string, message.Message[Producer, string]](
		srv.GetActorMessageChannelSize(),
		srv.GetActorMessageBufferInitialSize(),
		srv.onProcessMessage,
		func(message message.Message[Producer, string], err error) {
			if handler := srv.GetMessageErrorHandler(); handler != nil {
				handler(srv, message, err)
			}
		})
	srv.Options.init(srv).Apply(options...)
	srv.reactor.SetLogger(srv.Options.GetLogger())
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
	var queueWait = make(chan struct{})
	go s.reactor.Run(func(queues []*queue.Queue[int, string, message.Message[Producer, string]]) {
		for _, q := range queues {
			s.GetLogger().Debug("Reactor", log.String("action", "run"), log.Any("queue", q.Id()))
		}
		close(queueWait)
	})
	<-queueWait

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
	s.reactor.Close()
	return
}

func (s *server) getSysQueue() string {
	return s.queue
}

func (s *server) PublishMessage(msg message.Message[Producer, string]) {
	s.reactor.Dispatch(msg.GetQueue(), msg)
}

func (s *server) PublishSyncMessage(queue string, handler messages.SynchronousHandler[Producer, string, Server]) {
	s.PublishMessage(messages.Synchronous[Producer, string, Server](
		s, newProducer(s, nil), queue,
		handler,
	))
}

func (s *server) PublishAsyncMessage(queue string, handler messages.AsynchronousHandler[Producer, string, Server], callback ...messages.AsynchronousCallbackHandler[Producer, string, Server]) {
	s.PublishMessage(messages.Asynchronous[Producer, string, Server](
		s, newProducer(s, nil), queue,
		func(ctx context.Context, srv Server, f func(context.Context, Server)) {
			s.ants.Submit(func() {
				f(ctx, s)
			})
		},
		handler,
		collection.FindFirstOrDefaultInSlice(callback, nil),
	))
}

func (s *server) GetStatus() *State {
	return s.state.Status()
}

func (s *server) onProcessMessage(m message.Message[Producer, string]) {
	s.getManyOptions(func(opt *Options) {
		ctx := context.Background()
		m.OnInitialize(ctx)
		m.OnProcess()
	})
}
