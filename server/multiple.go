package server

import (
	"github.com/xtaci/kcp-go/v5"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func NewMultipleServer(serverHandle ...func() (addr string, srv *Server)) *MultipleServer {
	ms := &MultipleServer{
		servers:       make([]*Server, len(serverHandle), len(serverHandle)),
		addresses:     make([]string, len(serverHandle), len(serverHandle)),
		serverHandles: serverHandle,
	}
	return ms
}

type MultipleServer struct {
	servers                  []*Server
	addresses                []string
	exitEventHandles         []func()
	startFinishEventHandlers []func()
	services                 []func()
	preload                  []func()
	serverHandles            []func() (addr string, srv *Server)
}

func (slf *MultipleServer) Run() {
	var exceptionChannel = make(chan error, 1)
	var runtimeExceptionChannel = make(chan error, 1)
	defer func() {
		close(exceptionChannel)
		close(runtimeExceptionChannel)
	}()
	for _, service := range slf.services {
		service()
	}
	for _, preload := range slf.preload {
		preload()
	}
	for i := 0; i < len(slf.serverHandles); i++ {
		slf.addresses[i], slf.servers[i] = slf.serverHandles[i]()
	}
	var wait sync.WaitGroup
	var hasKcp bool
	for i := 0; i < len(slf.servers); i++ {
		wait.Add(1)
		if slf.servers[i].network == NetworkKcp {
			hasKcp = true
		}
		go func(address string, server *Server) {
			var lock sync.Mutex
			var startFinish bool
			server.startFinishEventHandlers.Append(func(srv *Server) {
				lock.Lock()
				defer lock.Unlock()
				if !startFinish {
					startFinish = true
					wait.Done()
				}
			}, math.MaxInt)
			server.multiple = slf
			server.multipleRuntimeErrorChan = runtimeExceptionChannel
			if err := server.Run(address); err != nil {
				exceptionChannel <- err
			}
			lock.Lock()
			defer lock.Unlock()
			if !startFinish {
				startFinish = true
				wait.Done()
			}
		}(slf.addresses[i], slf.servers[i])
	}
	wait.Wait()
	if !hasKcp {
		kcp.SystemTimedSched.Close()
	}

	slf.OnStartFinishEvent()
	showServersInfo(serverMultipleMark, slf.servers...)

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

// RegStartFinishEvent 注册启动完成事件
func (slf *MultipleServer) RegStartFinishEvent(handle func()) {
	slf.startFinishEventHandlers = append(slf.startFinishEventHandlers, handle)
}

func (slf *MultipleServer) OnStartFinishEvent() {
	for _, handle := range slf.startFinishEventHandlers {
		handle()
	}
}
