package transport_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/transport"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"sync"
	"testing"
	"time"
)

type BenchmarkActor struct {
	Expect  int64
	Curr    int64
	Once    sync.Once
	StartAt time.Time
	EndAt   time.Time
	Wait    *sync.WaitGroup
	Sender  bool
}

func (b *BenchmarkActor) OnReceive(ctx vivid.ActorContext) {
	switch v := ctx.Message().(type) {
	case vivid.OnLaunch:
		if b.Sender {
			target := core.NewProcessRef(core.NewAddress("", "", "127.0.0.1", 8800, "/user/1"))
			ctx.Tell(target, int64(0))
		}
	case int64:
		b.Once.Do(func() {
			b.StartAt = time.Now()
		})
		ctx.Reply(v)
		b.Curr++
		if b.Curr == b.Expect {
			b.EndAt = time.Now()
			b.Wait.Done()

			fmt.Println("end:", b.EndAt.String(), "cost:", b.EndAt.Sub(b.StartAt))
		}
	}
}

func TestNetwork_Node_Server(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork("127.0.0.1:8800"))
	})

	system.ActorOf(func() vivid.Actor {
		return &BenchmarkActor{
			Expect: 1000000,
			Wait:   wait,
		}
	})

	wait.Wait()
}

func TestNetwork_Node_Client(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork("127.0.0.1:7700"))
	})

	system.ActorOf(func() vivid.Actor {
		return &BenchmarkActor{
			Expect: 1000000,
			Wait:   wait,
			Sender: true,
		}
	})
	wait.Add(1)

	wait.Wait()
}
