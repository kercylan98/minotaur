package transport_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"testing"
	"time"
)

func TestNewNetwork(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork(":8800"))
	})

	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *transport.ConnectionOpened:
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
		options.WithModule(transport.NewNetwork(":8899"))
	})

	eachActor := core.NewProcessRef(core.NewAddress("", "", "127.0.0.1", 8800, "/user/each"))
	v := &transport.ConnectionOpened{}
	start := time.Now()
	for i := 0; i < 10000; i++ {
		if _, err := system.Context().FutureAsk(eachActor, v).Result(); err != nil {
			panic(err)
		}
	}
	cost := time.Now().Sub(start)
	t.Log("cost:", cost, "second:", cost.Seconds())

	time.Sleep(time.Second * 1111111)
}

func TestNewNetworkKindOf(t *testing.T) {
	system := transport.NewTestActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork(":8800"))
	})

	system.RegKind("test", func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnLaunch:
				fmt.Println("launch")
			case string:
				fmt.Println("receive", ctx.Message())
				panic("test panic")
			}
		})
	})

	time.Sleep(time.Second * 1111111)
}

func TestNewNetworkKindOf2(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork(":8899"))
	})

	root := core.NewProcessRef(core.NewAddress("", "", "127.0.0.1", 8800, "/user"))
	ref := system.KindOf("test", root)

	system.Context().Tell(ref, "123")
	time.Sleep(time.Second * 1111111)
}
