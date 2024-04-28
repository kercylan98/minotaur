package super

import (
	"sync"
)

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Exec(f func()) {
	w.WaitGroup.Add(1)
	go func(w *WaitGroup) {
		defer w.WaitGroup.Done()
		f()
	}(w)
}
