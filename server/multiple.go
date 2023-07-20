package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	servers          []*Server
	addresses        []string
	exitEventHandles []func()
}

func (slf *MultipleServer) Run() {
	var exceptionChannel = make(chan error, 1)
	var runtimeExceptionChannel = make(chan error, 1)
	defer func() {
		close(exceptionChannel)
		close(runtimeExceptionChannel)
	}()
	var running = make([]*Server, 0, len(slf.servers))
	var wait sync.WaitGroup
	for i := 0; i < len(slf.servers); i++ {
		wait.Add(1)
		go func(address string, server *Server) {
			var startFinish bool
			server.startFinishEventHandles = append(server.startFinishEventHandles, func(srv *Server) {
				if !startFinish {
					startFinish = true
					wait.Done()
				}
			})
			server.multiple = slf
			server.multipleRuntimeErrorChan = runtimeExceptionChannel
			if err := server.Run(address); err != nil {
				exceptionChannel <- err
			} else {
				running = append(running, server)
			}
			if !startFinish {
				startFinish = true
				wait.Done()
			}
		}(slf.addresses[i], slf.servers[i])
	}
	wait.Wait()

	log.Info("Server", log.String(serverMultipleMark, "===================================================================="))
	for _, server := range slf.servers {
		log.Info("Server", log.String(serverMultipleMark, "RunningInfo"),
			log.Any("network", server.network),
			log.String("listen", server.addr),
		)
	}
	log.Info("Server", log.String(serverMultipleMark, "===================================================================="))

	systemSignal := make(chan os.Signal, 1)
	signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err := <-exceptionChannel:
		for _, server := range slf.servers {
			server.OnStopEvent()
		}
		for len(slf.servers) > 0 {
			server := slf.servers[0]
			server.shutdown(err)
			slf.servers = slf.servers[1:]
		}
		break
	case <-runtimeExceptionChannel:
		for _, server := range slf.servers {
			server.OnStopEvent()
		}
		for len(slf.servers) > 0 {
			server := slf.servers[0]
			server.multipleRuntimeErrorChan = nil
			server.shutdown(nil)
			slf.servers = slf.servers[1:]
		}
		break
	case <-systemSignal:
		for _, server := range slf.servers {
			server.OnStopEvent()
		}
		for len(slf.servers) > 0 {
			server := slf.servers[0]
			server.multipleRuntimeErrorChan = nil
			server.shutdown(nil)
			slf.servers = slf.servers[1:]
		}
		break
	}

	slf.OnExitEvent()
}

// RegExitEvent 注册退出事件
func (slf *MultipleServer) RegExitEvent(handle func()) {
	slf.exitEventHandles = append(slf.exitEventHandles, handle)
}

func (slf *MultipleServer) OnExitEvent() {
	for _, handle := range slf.exitEventHandles {
		handle()
	}
}
