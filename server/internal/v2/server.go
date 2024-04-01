package server

import (
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"github.com/kercylan98/minotaur/utils/network"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
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
}

type server struct {
	*controller
	*events
	state   *State
	ctx     context.Context
	cancel  context.CancelFunc
	network Network
	reactor *reactor.Reactor[Message]
}

func NewServer(network Network) Server {
	srv := &server{
		network: network,
	}
	srv.ctx, srv.cancel = context.WithCancel(context.Background())
	srv.controller = new(controller).init(srv)
	srv.events = new(events).init(srv)
	srv.reactor = reactor.NewReactor[Message](1024*8, 1024, func(msg Message) {
		msg.Execute()
	}, nil)
	srv.state = new(State).init(srv)
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

	var systemSignal = make(chan os.Signal)
	signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-systemSignal
	if err := s.Shutdown(); err != nil {
		panic(err)
	}
	return
}

func (s *server) Shutdown() (err error) {
	defer s.cancel()
	err = s.network.OnShutdown()
	s.reactor.Close()
	return
}

func (s *server) GetStatus() *State {
	return s.state.Status()
}
