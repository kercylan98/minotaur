package standard

import "minotaur/utils/timer"

// Timer 附加定时器功能
type Timer struct {
	*timer.Manager
}

func (slf *Timer) Timer() *timer.Manager {
	if slf.Manager == nil {
		slf.Manager = timer.GetManager(64)
	}
	return slf.Manager
}
