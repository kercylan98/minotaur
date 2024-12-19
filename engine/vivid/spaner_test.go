package vivid_test

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
)

type MyFixedActorProvider struct {
	t *testing.T
}

func (m *MyFixedActorProvider) ProvideActor() (actor vivid.Actor) {
	return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
		switch ctx.Message().(type) {
		case *vivid.OnLaunch:
			m.t.Log("Launch")
		}
	})
}

func (m *MyFixedActorProvider) ProvideConfigurator() (configurator vivid.ActorDescriptorConfigurator) {
	return vivid.FunctionalActorDescriptorConfigurator(func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("fixed")
	})
}

func TestSpawnActorFromFixedProvider(t *testing.T) {
	sys := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithFixedActorProvider("fixed", &MyFixedActorProvider{t: t})
	}))

	if _, err := vivid.SpawnActorFromFixedProvider(sys, sys, "fixed2"); err == nil {
		panic("expect error")
	} else {
		t.Log(err)
	}

	if _, err := vivid.SpawnActorFromFixedProvider(sys, sys, "fixed"); err != nil {
		panic(err)
	}
	sys.Shutdown(true)
}
