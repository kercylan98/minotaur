package timer

import (
	"sync"

	"github.com/RussellLuo/timingwheel"
)

var timer = new(Timer)

func GetManager(size int) *Manager {
	return timer.NewManager(size)
}

type Timer struct {
	Managers []*Manager
	lock     sync.Mutex
}

func (slf *Timer) NewManager(size int) *Manager {
	slf.lock.Lock()
	defer slf.lock.Unlock()

	var manager *Manager
	if len(slf.Managers) > 0 {
		manager = slf.Managers[0]
		slf.Managers = slf.Managers[1:]
		return manager
	}

	manager = &Manager{
		timer:  slf,
		wheel:  timingwheel.NewTimingWheel(timingWheelTick, int64(size)),
		timers: make(map[string]*Scheduler),
	}
	manager.wheel.Start()
	return manager
}
