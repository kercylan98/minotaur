package ecs

import (
	"github.com/kercylan98/minotaur/utils/super"
	"reflect"
)

func QueryEntity[C any](world *World, entity Entity) *C {
	t := reflect.TypeOf((*C)(nil)).Elem()
	id, exist := world.components[t]
	if !exist {
		var ids []ComponentId
		for i := range t.NumField() {
			ids = append(ids, world.ComponentId(t.Field(i).Type.Elem()))
		}
		mask := super.NewBitSet(ids...)
		for _, arch := range world.archetypes {
			if !arch.mask.ContainsAll(mask) {
				continue
			}

			values := arch.getEntityData(entity)
			fields := make(map[reflect.Type]reflect.Value)
			for _, value := range values {
				fields[value.Type()] = value
			}

			result := reflect.New(t)
			for i := range t.NumField() {
				f := result.Elem().Field(i)
				f.Set(fields[f.Type()])
			}
			return result.Interface().(*C)
		}
		return nil
	}

	for _, arch := range world.archetypes {
		if !arch.mask.Has(id) {
			continue
		}

		for _, e := range arch.entities {
			if e == entity {
				return arch.getEntityComponentData(entity, id).Interface().(*C)
			}
		}
	}

	return nil
}

func QueryEntitiesByComponentId[T any](world *World, id ComponentId) []*T {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if world.components[t] != id {
		return nil
	}

	var cs []*T
	for _, arch := range world.archetypes {
		if arch.mask.Has(id) {
			for _, entity := range arch.entities {
				arch.getEntityComponentData(entity, id)
				cs = append(cs, arch.getEntityComponentData(entity, id).Interface().(*T))
			}
		}
	}
	return cs
}
