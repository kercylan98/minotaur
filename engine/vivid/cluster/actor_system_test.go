package cluster_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"github.com/kercylan98/minotaur/toolkit"
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

type MyOnlyActorProvider struct {
}

func (m *MyOnlyActorProvider) ProvideActor() (actor vivid.Actor) {
	return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
		switch ctx.Message().(type) {
		case *vivid.OnLaunch:
			fmt.Println(ctx.System().Name(), "Only Actor Launch")
		case *vivid.OnTerminated:
			fmt.Println(ctx.System().Name(), "Only Actor Terminated")
		case vivid.ActorRef:
			fmt.Println(ctx.System().Name(), "Only Actor Receive Message")
		}
	})
}

func (m *MyOnlyActorProvider) ProvideConfigurator() (configurator vivid.ActorDescriptorConfigurator) {
	return vivid.FunctionalActorDescriptorConfigurator(func(descriptor *vivid.ActorDescriptor) {

	})
}

func TestActorSystemFixedActor(t *testing.T) {
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

func TestActorSystemOnlyActor(t *testing.T) {

	system1 := cluster.NewActorSystem("127.0.0.1:6666", []prc.PhysicalAddress{"127.0.0.1:6666"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithOnlyActorProvider("only", new(MyOnlyActorProvider))
	}))
	system2 := cluster.NewActorSystem("127.0.0.1:6667", []prc.PhysicalAddress{"127.0.0.1:6666"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithOnlyActorProvider("only", new(MyOnlyActorProvider))
	}))
	system3 := cluster.NewActorSystem("127.0.0.1:6668", []prc.PhysicalAddress{"127.0.0.1:6666"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithOnlyActorProvider("only", new(MyOnlyActorProvider))
	}))

	_ = system1
	_ = system2
	_ = system3

	time.Sleep(time.Hour)
}

func TestActorSystemOnlyActorA(t *testing.T) {

	system1 := cluster.NewActorSystem("127.0.0.1:6666", []prc.PhysicalAddress{"127.0.0.1:6666", "127.0.0.1:6667"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithOnlyActorProvider("only", new(MyOnlyActorProvider))
	}))

	_ = system1
	time.Sleep(time.Hour)
}

func TestActorSystemOnlyActorB(t *testing.T) {
	system2 := cluster.NewActorSystem("127.0.0.1:6667", []prc.PhysicalAddress{"127.0.0.1:6666", "127.0.0.1:6667"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithOnlyActorProvider("only", new(MyOnlyActorProvider))
	}))

	_ = system2

	time.Sleep(time.Hour)
}

func TestActorSystemGetOnlyActor(t *testing.T) {
	system3 := cluster.NewActorSystem("127.0.0.1:6668", []prc.PhysicalAddress{"127.0.0.1:6666", "127.0.0.1:6667"}, cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
		config.WithOnlyActorProvider("only", new(MyOnlyActorProvider))
	}))

	_ = system3
	system3.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *vivid.OnLaunch:
				ctx.Subscribe(vivid.AbyssTopic)
			case *vivid.OnAbyssMessageEvent:
				fmt.Println(string(toolkit.MarshalJSON(m)))
			}
		})
	})

	ref := system3.GetOnlyActor("only")
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Second)
		system3.Tell(ref, ref)
	}
}
