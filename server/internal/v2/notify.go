package server

import (
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type notify struct {
	server *server

	systemSignal   chan os.Signal
	lifeCycleLimit chan struct{}
	lifeCycleTimer *time.Timer
	lifeCycleTime  chan time.Duration
}

func (n *notify) init(srv *server) *notify {
	n.server = srv
	n.systemSignal = make(chan os.Signal, 1)
	n.lifeCycleLimit = make(chan struct{}, 1)
	n.lifeCycleTime = make(chan time.Duration, 1)
	n.lifeCycleTimer = time.NewTimer(math.MaxInt64)
	n.lifeCycleTimer.Stop()
	signal.Notify(n.systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	return n
}

func (n *notify) run() {
	defer func() {
		if err := n.server.Shutdown(); err != nil {
			panic(err)
		}
	}()
	for {
		select {
		case <-n.systemSignal:
			return
		case <-n.lifeCycleLimit:
			n.systemSignal <- syscall.SIGQUIT
		case <-n.lifeCycleTimer.C:
			n.systemSignal <- syscall.SIGQUIT
		case d := <-n.lifeCycleTime:
			n.lifeCycleTimer.Stop()
			if d > 0 {
				n.lifeCycleTimer.Reset(d)
			}
		}
	}
}
