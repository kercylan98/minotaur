package transport_test

import (
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
		options.WithModule(transport.NewNetwork("127.0.0.1:8899"))
	})

	eachActor := core.NewProcessRef(core.NewAddress("", "", "127.0.0.1", 8800, "/user/each"))
	v := "Hello, World!"
	start := time.Now()
	for i := 0; i < 10000; i++ {
		if _, err := system.Context().FutureAsk(eachActor, v, func(options *vivid.MessageOptions) {
			options.WithFutureTimeout(0)
		}).Result(); err != nil {
			panic(err)
		} else {

		}
	}
	cost := time.Now().Sub(start)
	t.Log("cost:", cost, "second:", cost.Seconds())

	time.Sleep(time.Second * 1111111)
}
