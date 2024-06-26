package transport_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/transport"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"testing"
	"time"
)

func TestNewNetwork(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork("127.0.0.1:8800"))
	})

	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case string:
				ctx.Reply(m)
			}
		})
	}, func(options *vivid.ActorOptions) {
		options.WithName("each")
	})

	time.Sleep(time.Second * 1111111)
}

func TestNewNetwork2(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork("127.0.0.1:8888"))
	})

	eachActor := core.NewProcessRef(core.NewAddress("", "", "127.0.0.1", 8800, "/user/each"))
	if result, err := system.Context().FutureAsk(eachActor, "Hello, World!", func(options *vivid.MessageOptions) {
		options.WithFutureTimeout(time.Second * 9999)
	}).Result(); err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

	time.Sleep(time.Second * 1111111)
}

var system = vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
	options.WithModule(transport.NewNetwork("127.0.0.1:8888"))
})

func BenchmarkNetworkAsk(b *testing.B) {
	eachActor := core.NewProcessRef(core.NewAddress("", "", "127.0.0.1", 8800, "/user/each"))
	var send int
	var reply int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		send++
		if err := system.Context().FutureAsk(eachActor, "Hello, World!").Wait(); err != nil {
			panic(err)
		}
		reply++
	}
	b.StopTimer()

	b.Log("send:", send, "reply:", reply)
	system.Shutdown()
}

func TestBenchmarkNetworkAsk(t *testing.T) {
	eachActor := core.NewProcessRef(core.NewAddress("", "", "127.0.0.1", 8800, "/user/each"))

	start := time.Now()
	for i := 0; i < 100000; i++ {
		if err := system.Context().FutureAsk(eachActor, "Hello, World!").Wait(); err != nil {
			panic(err)
		}
	}
	t.Log("time:", time.Since(start))

	system.Shutdown()
}
