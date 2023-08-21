package concurrent_test

import (
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestPool_EAC(t *testing.T) {
	var p = concurrent.NewPool[int](10, func() int {
		return 0
	}, func(data int) {
	})

	go func() {
		for i := 0; i < 1000; i++ {
			go func() {
				for {
					p.Release(p.Get())
				}
			}()
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		p.EAC(2048)
	}()

	time.Sleep(100 * times.Day)
}
