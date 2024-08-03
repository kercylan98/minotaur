package cluster_test

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"testing"
)

func TestNewActorSystem(t *testing.T) {

	systemA := cluster.NewActorSystem("127.0.0.1:8080", "127.0.0.1:1267", cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithAbility("calc", cluster.FunctionalActorProvider(func() cluster.Actor {
			state := 0
			return cluster.FunctionalActor(func(ctx cluster.ActorContext) {
				switch ctx.Message().(type) {
				case vivid.ActorRef:
					state++
					ctx.Reply(ctx.Ref())
				}
			})
		}))
	}))

	systemB := cluster.NewActorSystem("127.0.0.1:8081", "127.0.0.1:1268", cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithSeedNodes("127.0.0.1:1267")
	}))

	defer systemA.Shutdown(true)
	defer systemB.Shutdown(true)

	ref := systemB.ActorOfC("user", "calc")
	for i := 0; i < 5; i++ {
		if _, err := vivid.FutureAsk[vivid.ActorRef](systemB, ref, ref).Result(); err != nil {
			panic(err)
		} else {
			t.Log(i + 1)
		}
	}
}
