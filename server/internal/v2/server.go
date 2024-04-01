package server

import (
	"context"
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
}

type server struct {
	*controller
	*events
	*Options
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
		srv.getServerMessageChannelSize(), srv.getActorMessageChannelSize(),
		srv.getServerMessageBufferInitialSize(), srv.getActorMessageBufferInitialSize(),
		func(msg Message) {
			msg.Execute()
		}, func(msg Message, err error) {
			if handler := srv.getMessageErrorHandler(); handler != nil {
				handler(srv, msg, err)
			}
		})
	srv.Options.init(srv).Apply(options...)
	return srv
}

func (s *server) Run() (err error) {
	go s.reactor.Run()
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
	defer s.cancel()
	s.events.onShutdown()
	err = s.network.OnShutdown()
	s.reactor.Close()
	return
}

func (s *server) GetStatus() *State {
	return s.state.Status()
}
