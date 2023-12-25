package server

import (
	"github.com/xtaci/kcp-go/v5"
	"net"
	"sync"
)

type listener struct {
	srv  *Server
	once sync.Once
	net.Listener
	kcpListener *kcp.Listener
	state       chan<- error
}

func (l *listener) init() *listener {
	l.srv.OnStartBeforeEvent()
	return l
}

func (l *listener) Accept() (net.Conn, error) {
	l.once.Do(func() {
		l.state <- nil
	})
	return l.Listener.Accept()
}

func (l *listener) AcceptKCP() (*kcp.UDPSession, error) {
	l.once.Do(func() {
		l.state <- nil
	})
	return l.kcpListener.AcceptKCP()
}
