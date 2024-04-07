package reactor_test

import (
	"github.com/kercylan98/minotaur/server/internal/v2/queue"
	"github.com/kercylan98/minotaur/server/internal/v2/reactor"
	"github.com/kercylan98/minotaur/utils/log/v2"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/times"
	"os"
	"testing"
	"time"
)

func BenchmarkReactor_Dispatch(b *testing.B) {
	var r = reactor.NewReactor(1024*16, 1024, 1024, 1024, func(message queue.MessageWrapper[int, string, func()]) {
		message.Message()
	}, func(message queue.MessageWrapper[int, string, func()], err error) {

	})

	r.SetLogger(log.NewLogger(log.NewHandler(os.Stdout, log.DefaultOptions().WithLevel(log.LevelInfo))))

	go r.Run()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := r.IdentDispatch(random.HostName(), func() {
			}); err != nil {
				return
			}
		}
	})
}

func TestReactor_Dispatch(t *testing.T) {
	var r = reactor.NewReactor(1024*16, 1024, 1024, 1024, func(message queue.MessageWrapper[int, string, func()]) {
		message.Message()
	}, func(message queue.MessageWrapper[int, string, func()], err error) {

	})

	go r.Run()

	for i := 0; i < 10000; i++ {
		go func() {
			id := random.HostName()
			for {
				time.Sleep(time.Millisecond * 20)
				if err := r.IdentDispatch(id, func() {

				}); err != nil {
					return
				}
			}
		}()
	}

	time.Sleep(times.Second)
	r.Close()
}
