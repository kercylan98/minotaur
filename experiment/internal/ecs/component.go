package ecs

import "reflect"

// ComponentId 组件 ID
type ComponentId = uint32

// ComponentType 组件类型
type ComponentType = reflect.Type

func newComponentInfo(id ComponentId, t ComponentType) *componentInfo {
	return &componentInfo{
		id:        id,
		typ:       t,
		entityIdx: map[EntityId]int{},
	}
}

// ComponentInfo 组件信息
type componentInfo struct {
	id  ComponentId
	typ ComponentType

	entityIdx map[EntityId]int
	data      []any
}

func (c *componentInfo) addEntity(id EntityId) {
	c.entityIdx[id] = len(c.data)
	ins := reflect.New(c.typ).Interface()
	c.data = append(c.data, ins)
}

func (c *componentInfo) getData(id EntityId) any {
	return c.data[c.entityIdx[id]]
}
