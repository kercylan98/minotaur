package server

import (
	"context"
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"github.com/panjf2000/ants/v2"
	"time"

	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
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

	// PushSyncMessage 推送同步消息
	PushSyncMessage(handler func(srv Server))

	// PushAsyncMessage 推送异步消息
	PushAsyncMessage(handler func(srv Server) error, callbacks ...func(srv Server, err error))
}

type server struct {
	*controller
	*events
	*Options
	ants    *ants.Pool
	state   *State
	notify  *notify
	ctx     context.Context
	cancel  context.CancelFunc
	network Network
	reactor *reactor.Reactor[Message]
}

func NewServer(network Network, options ...*Options) Server {
	srv := &server{
		network: network,
		Options: DefaultOptions(),
	}
	srv.ctx, srv.cancel = context.WithCancel(context.Background())
	srv.notify = new(notify).init(srv)
	srv.controller = new(controller).init(srv)
	srv.events = new(events).init(srv)
	srv.state = new(State).init(srv)
	srv.reactor = reactor.NewReactor[Message](
		srv.GetServerMessageChannelSize(), srv.GetActorMessageChannelSize(),
		srv.GetServerMessageBufferInitialSize(), srv.GetActorMessageBufferInitialSize(),
		func(message queue.MessageWrapper[int, string, Message]) {
			message.Message().OnExecute()
		}, func(message queue.MessageWrapper[int, string, Message], err error) {
			if handler := srv.GetMessageErrorHandler(); handler != nil {
				handler(srv, message.Message(), err)
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
	go s.reactor.Run(func(queues []*queue.Queue[int, string, Message]) {
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

func (s *server) PushSyncMessage(handler func(srv Server)) {
	if err := s.reactor.DispatchWithSystem(SyncMessage(s, func(srv *server) {
		handler(srv)
	})); err != nil {
		panic(err)
	}
}

func (s *server) PushAsyncMessage(handler func(srv Server) error, callbacks ...func(srv Server, err error)) {
	if err := s.reactor.DispatchWithSystem(AsyncMessage(s, reactor.SysIdent, func(srv *server) error {
		return handler(srv)
	}, func(srv *server, err error) {
		for _, callback := range callbacks {
			callback(srv, err)
		}
	})); err != nil {
		panic(err)
	}
}

func (s *server) GetStatus() *State {
	return s.state.Status()
}
