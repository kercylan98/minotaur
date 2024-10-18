package gossip

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"testing"
	"time"
)

func TestGossip(t *testing.T) {
	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared("127.0.0.1:8080")
	}))

	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared("127.0.0.1:8081")
	}))

	_ = system1
	_ = system2

	system1.ActorOfF(func() vivid.Actor {
		return NewGossiperActor([]string{"127.0.0.1:8080"})
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("gossip")
	})

	system2.ActorOfF(func() vivid.Actor {
		return NewGossiperActor([]string{"127.0.0.1:8080", "127.0.0.1:8081"})
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("gossip")
	})

	time.Sleep(chrono.Hour)
}

func TestGossip2(t *testing.T) {
	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared("127.0.0.1:7777")
	}))

	_ = system1

	ref := system1.ActorOfF(func() vivid.Actor {
		return NewGossiperActor([]string{"127.0.0.1:8080"})
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("gossip")
	})

	_ = ref

	//time.Sleep(chrono.Second * 3)
	//
	//system1.Tell(ref, &GossipActorLeaveClusterMessage{})

	time.Sleep(chrono.Hour)
}
