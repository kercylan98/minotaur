package ecs

import (
	"github.com/kercylan98/minotaur/utils/super"
	"reflect"
)

// NewWorld 创建一个新的世界
func NewWorld() World {
	return World{
		components:     make(map[reflect.Type]int),
		componentTypes: make(map[int]reflect.Type),
		archetypes:     make(map[*super.BitSet[int]]*archetype),
	}
}

type World struct {
	componentIds   []ComponentId                // 已经注册的组件 Id 清单
	components     map[reflect.Type]ComponentId // 已经注册的组件清单
	componentTypes map[ComponentId]reflect.Type // 已经注册的组件类型清单

	archetypes map[*super.BitSet[ComponentId]]*archetype // 已经注册的原型清单
	entityGuid EntityId                                  // 实体的唯一标识当前值
}

// CreateEntity 创建一个新的实体
func (w *World) CreateEntity(componentId ComponentId, componentIds ...ComponentId) Entity {
	mask := super.NewBitSet(append([]ComponentId{componentId}, componentIds...)...)

	var arch *archetype
	for existingMask, existingArch := range w.archetypes {
		if existingMask.Equal(mask) {
			arch = existingArch
			break
		}
	}

	if arch == nil {
		arch = newArchetype(w, mask)
		w.archetypes[mask] = arch
	}

	return arch.addEntity(Entity{
		id: w.entityGuid,
	})
}

// ComponentId 返回一个组件的 Id，如果组件未注册，则注册它
func (w *World) ComponentId(t reflect.Type) ComponentId {
	id, exist := w.components[t]
	if !exist {
		id = len(w.components)
		w.components[t] = id
		w.componentTypes[id] = t
		w.componentIds = append(w.componentIds, id)
	}

	return id
}

// getComponentTypeById 通过 Id 获取一个组件的类型
func (w *World) getComponentTypeById(id ComponentId) reflect.Type {
	if id < 0 || id >= len(w.componentIds) {
		return nil
	}
	return w.componentTypes[id]
}

// unregisterComponentById 通过 Id 注销一个组件
func (w *World) unregisterComponentById(id ComponentId) {
	t := w.componentTypes[id]
	delete(w.components, t)
	delete(w.componentTypes, id)
	w.componentIds = append(w.componentIds[:id], w.componentIds[id+1:]...)
}

// unregisterComponentByType 通过类型注销一个组件
func (w *World) unregisterComponentByType(t reflect.Type) {
	id, exist := w.components[t]
	if !exist {
		return
	}

	w.unregisterComponentById(id)
}

// nextEntityGuid 返回下一个实体的唯一标识
func (w *World) nextEntityGuid() EntityId {
	guid := w.entityGuid
	w.entityGuid++
	return guid
}
