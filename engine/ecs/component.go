package ecs

import "reflect"

type (
	// ComponentId 组件 ID
	ComponentId = uint32

	// ComponentType 组件类型
	ComponentType = reflect.Type
)

func newComponentInfo(id ComponentId, t ComponentType) *componentInfo {
	return &componentInfo{
		id:  id,
		typ: t,
	}
}

// ComponentInfo 组件信息
type componentInfo struct {
	id  ComponentId
	typ ComponentType
}

// instantiate 实例化组件
func (ci *componentInfo) instantiate() any {
	ins := reflect.New(ci.typ).Interface()
	return ins
}
