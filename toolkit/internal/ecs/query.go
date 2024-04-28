package ecs

// QueryComponent 查询特定世界中的特定实体的特定组件，如果组件未注册或实体不存在，则返回 nil
func QueryComponent[C any](world *World, id EntityId) (component *C) {
	componentId := GetComponentId[C](world)

	arch, exist := world.entities[id]
	if !exist {
		return nil
	}
	ptr := arch.getEntityData(id, componentId)
	return (*C)(ptr)
}
