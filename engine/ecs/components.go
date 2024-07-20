package ecs

import "reflect"

func newComponents() *components {
	return &components{
		componentIds:   make(map[ComponentId]*componentInfo),
		componentTypes: make(map[ComponentType]*componentInfo),
	}
}

type components struct {
	componentIds   map[ComponentId]*componentInfo   // 组件 id 到组件信息的映射
	componentTypes map[ComponentType]*componentInfo // 组件类型到组件信息的映射
}

func (c *components) reg(component any) ComponentId {
	componentType := reflect.TypeOf(component).Elem()
	info, exists := c.componentTypes[componentType]
	if exists {
		return info.id
	}
	info = newComponentInfo(ComponentId(len(c.componentTypes)+1), componentType)

	c.componentIds[info.id] = info
	c.componentTypes[info.typ] = info

	return info.id
}

func (c *components) getInfo(id ComponentId) *componentInfo {
	return c.componentIds[id]
}
