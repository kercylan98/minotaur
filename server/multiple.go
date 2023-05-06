package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewMultipleServer(serverHandle ...func() (addr string, srv *Server)) *MultipleServer {
	ms := &MultipleServer{
		servers:   make([]*Server, len(serverHandle), len(serverHandle)),
		addresses: make([]string, len(serverHandle), len(serverHandle)),
	}
	for i := 0; i < len(serverHandle); i++ {
		ms.addresses[i], ms.servers[i] = serverHandle[i]()
	}
	return ms
}

type MultipleServer struct {
	servers   []*Server
	addresses []string
}

func (slf *MultipleServer) Run() {
	var exceptionChannel = make(chan error, 1)
	defer close(exceptionChannel)
	var running = make([]*Server, 0, len(slf.servers))
	for i := 0; i < len(slf.servers); i++ {
		go func(address string, server *Server) {
			server.multiple = true
			if err := server.Run(address); err != nil {
				exceptionChannel <- err
			} else {
				running = append(running, server)
			}
		}(slf.addresses[i], slf.servers[i])
	}

	time.Sleep(500 * time.Millisecond)

	log.Info("Server", zap.String("Minotaur Multiple Server", "===================================================================="))
	for _, server := range slf.servers {
		log.Info("Server", zap.String("Minotaur Multiple Server", "RunningInfo"),
			zap.Any("network", server.network),
			zap.String("listen", server.addr),
		)
	}
	log.Info("Server", zap.String("Minotaur Multiple Server", "===================================================================="))

	systemSignal := make(chan os.Signal, 1)
	signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err := <-exceptionChannel:
		for _, server := range slf.servers {
			server.Shutdown(err)
		}
		return
	case <-systemSignal:
		for _, server := range slf.servers {
			server.Shutdown(nil)
		}
		return
	}
}
