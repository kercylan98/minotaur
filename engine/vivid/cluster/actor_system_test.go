package cluster_test

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"testing"
	"time"
)

type MyFixedProvider struct {
	sys *cluster.ActorSystem
}

func (m *MyFixedProvider) ProvideActor() (actor vivid.Actor) {
	return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
	})
}

func (m *MyFixedProvider) ProvideConfigurator() (configurator vivid.ActorDescriptorConfigurator) {
	return vivid.FunctionalActorDescriptorConfigurator(func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithNamePrefix(m.sys.Name())
	})
}

func TestActorSystem(t *testing.T) {
	fap1 := new(MyFixedProvider)
	fap2 := new(MyFixedProvider)
	fap3 := new(MyFixedProvider)

	system1 := cluster.NewActorSystem("127.0.0.1:6666", []prc.PhysicalAddress{"127.0.0.1:6666"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithFixedActorProvider("fixed", fap1)
	}))
	system2 := cluster.NewActorSystem("127.0.0.1:6667", []prc.PhysicalAddress{"127.0.0.1:6666"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithFixedActorProvider("fixed", fap2)
	}))
	system3 := cluster.NewActorSystem("127.0.0.1:6668", []prc.PhysicalAddress{"127.0.0.1:6666"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithFixedActorProvider("fixed2", fap3)
	}))

	fap1.sys = system1
	fap2.sys = system2
	fap3.sys = system3

	time.Sleep(time.Second * 3)

	for i := 0; i < 100; i++ {
		_, err := system1.SpawnFixedActor("fixed")
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 10; i++ {
		_, err := system1.SpawnFixedActor("fixed2")
		if err != nil {
			panic(err)
		}
	}

	time.Sleep(time.Hour)
}
