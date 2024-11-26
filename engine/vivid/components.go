package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"reflect"
)

func newComponents(actorSystem *ActorSystem) *components {
	var cs = &components{}
	for _, comp := range actorSystem.config.components {
		tryBindComponent[ShutdownComponent](&cs.shutdown, comp)
		tryBindComponent[ActorSpawnBeforeComponent](&cs.actorSpawn, comp)
	}

	cs.onInitialize()
	return cs
}
func tryBindComponent[C Component](slice *[]C, comp Component) {
	if c, ok := comp.(C); ok {
		*slice = append(*slice, c)
	}
}

type components struct {
	actorSystem *ActorSystem
	shutdown    []ShutdownComponent
	actorSpawn  []ActorSpawnBeforeComponent
}

func (cs *components) OnActorSpawnBefore(provider ActorProvider, descriptor *ActorDescriptor) {
	if cs == nil {
		return
	}
	for _, c := range cs.actorSpawn {
		c.OnActorSpawnBefore(cs.actorSystem, provider, descriptor)
	}
}

func (cs *components) onInitialize() {
	if cs == nil {
		return
	}
	for _, c := range cs.actorSystem.config.components {
		if err := c.OnInitialize(cs.actorSystem); err != nil {
			panic(err)
		} else {
			cs.actorSystem.Logger().Info("component", log.String("name", reflect.TypeOf(c).Name()), log.String("status", "initialized"))
		}
	}
}

func (cs *components) onShutdown() {
	if cs == nil {
		return
	}
	for _, c := range cs.shutdown {
		if err := c.OnShutdown(cs.actorSystem); err != nil {
			cs.actorSystem.Logger().Error("component", log.String("name", reflect.TypeOf(c).Name()), log.String("status", "shutdown failed"), log.Err(err))
		} else {
			cs.actorSystem.Logger().Info("component", log.String("name", reflect.TypeOf(c).Name()), log.String("status", "shutdown success"))
		}
	}
}
