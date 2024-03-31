package reactor_test

import (
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
	"time"
)

func BenchmarkReactor_Dispatch(b *testing.B) {

	var r = reactor.NewReactor(1024*16, 1024, func(msg func()) {
		msg()
	}, func(msg func(), err error) {
		b.Error(err)
	}).SetDebug(false)

	go r.Run()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := r.Dispatch(random.HostName(), func() {

			}); err != nil {

			}
		}
	})
}

func TestReactor_Dispatch(t *testing.T) {
	var r = reactor.NewReactor(1024*16, 1024, func(msg func()) {
		msg()
	}, func(msg func(), err error) {
		t.Error(err)
	}).SetDebug(false)

	go r.Run()

	for i := 0; i < 10000; i++ {
		go func() {
			id := random.HostName()
			for {
				// 每秒 50 次
				time.Sleep(time.Millisecond * 20)
				if err := r.Dispatch(id, func() {

				}); err != nil {
					t.Error(err)
				}
			}
		}()
	}

	time.Sleep(time.Second * 10)
}
