package ecs

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"reflect"
)

func NewWorld() *World {
	return &World{
		entityManager:  newEntityManager(),
		componentIds:   make(map[ComponentId]*componentInfo),
		componentTypes: make(map[ComponentType]*componentInfo),
		logger: func() *log.Logger {
			return log.GetDefault()
		},
	}
}

type World struct {
	entityManager  *entityManager
	componentIds   map[ComponentId]*componentInfo
	componentTypes map[ComponentType]*componentInfo
	logger         vivid.LoggerProvider
}

func (w *World) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case WorldRegisterComponentMessage:
		w.onRegisterComponent(ctx, m)
	case WorldCreateEntityMessage:
		w.onCreateEntity(ctx, m)
	case WorldLoadComponentMessage:
		w.onLoadComponent(ctx, m)
	}
}

func (w *World) onRegisterComponent(ctx vivid.ActorContext, m WorldRegisterComponentMessage) {
	if m.Kind() == reflect.Pointer {
		m = m.Elem()
	}

	info, exists := w.componentTypes[m]
	if exists {
		ctx.Reply(info)
		return
	}
	info = newComponentInfo(ComponentId(len(w.componentTypes)+1), m)

	w.componentIds[info.id] = info
	w.componentTypes[info.typ] = info

	ctx.Reply(info.id)

	w.logger().Debug("ECS", log.String("info", "register component"), log.Uint32("id", info.id), log.String("type", info.typ.String()))
}

func (w *World) onCreateEntity(ctx vivid.ActorContext, m WorldCreateEntityMessage) {
	eid := w.entityManager.createEntity()
	for _, id := range m {
		w.componentIds[id].addEntity(eid)
	}

	ctx.Reply(eid)

	w.logger().Debug("ECS", log.String("info", "create entity"), log.Uint32("id", uint32(eid)))
}

func (w *World) onLoadComponent(ctx vivid.ActorContext, m WorldLoadComponentMessage) {
	if m.ComponentType.Kind() == reflect.Pointer {
		m.ComponentType = m.ComponentType.Elem()
	}

	info, exists := w.componentTypes[m.ComponentType]
	if !exists {
		ctx.Reply(nil)
		return
	}

	ctx.Reply(info.getData(m.EntityId))
}
