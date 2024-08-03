package cluster

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/cm"
)

func newDrillmasterActor(system *ActorSystem) *drillmasterActor {
	d := &drillmasterActor{
		members: make(map[string]map[string]vivid.ActorRef),
		system:  system,
	}
	for s := range system.config.abilities {
		d.members[s] = make(map[string]vivid.ActorRef)
	}
	return d
}

type drillmasterActor struct {
	system *ActorSystem

	members map[string]map[string]vivid.ActorRef // 集群创建的能力成员列表 ability => identity => ref
}

func (d *drillmasterActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *cm.ActorOf:
		d.onActorOf(ctx, m)
	}
}

func (d *drillmasterActor) onActorOf(ctx vivid.ActorContext, m *cm.ActorOf) {
	ability, exist := d.system.config.abilities[m.Ability]
	if !exist {
		ctx.Reply(fmt.Errorf("the ability %s does not support", m.Ability))
		return
	}

	ref, exist := d.members[m.Ability][m.Identity]
	if !exist {
		ref = ctx.ActorOf(
			vivid.FunctionalActorProvider(func() vivid.Actor {
				return newActor(d.system, ability.provider)
			}),
			append(ability.configurator, vivid.FunctionalActorDescriptorConfigurator(func(descriptor *vivid.ActorDescriptor) {
				descriptor.WithNamePrefix(m.Identity).WithName(m.Ability)
			}))...,
		)
	}

	ctx.Reply(ref)
}
