package super

import (
	"github.com/kercylan98/minotaur/utils/log"
	"sync"
)

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Exec(f func()) {
	w.WaitGroup.Add(1)
	go func(w *WaitGroup) {
		defer func() {
			w.WaitGroup.Done()
			if err := RecoverTransform(recover()); err != nil {
				log.Error("WaitGroup", log.Err(err))
			}
		}()

		f()
	}(w)
}
