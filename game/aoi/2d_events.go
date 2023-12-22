package aoi

import "github.com/kercylan98/minotaur/utils/generic"

type (
	EntityJoinVisionEventHandle[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]]  func(entity, target E)
	EntityLeaveVisionEventHandle[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]] func(entity, target E)
)

type event[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]] struct {
	entityJoinVisionEventHandles  []EntityJoinVisionEventHandle[EID, PosType, E]
	entityLeaveVisionEventHandles []EntityLeaveVisionEventHandle[EID, PosType, E]
}

// RegEntityJoinVisionEvent 在新对象进入视野时将会立刻执行被注册的事件处理函数
func (slf *event[EID, PosType, E]) RegEntityJoinVisionEvent(handle EntityJoinVisionEventHandle[EID, PosType, E]) {
	slf.entityJoinVisionEventHandles = append(slf.entityJoinVisionEventHandles, handle)
}

// OnEntityJoinVisionEvent 在新对象进入视野时将会立刻执行被注册的事件处理函数
func (slf *event[EID, PosType, E]) OnEntityJoinVisionEvent(entity, target E) {
	for _, handle := range slf.entityJoinVisionEventHandles {
		handle(entity, target)
	}
}

// RegEntityLeaveVisionEvent 在新对象离开视野时将会立刻执行被注册的事件处理函数
func (slf *event[EID, PosType, E]) RegEntityLeaveVisionEvent(handle EntityLeaveVisionEventHandle[EID, PosType, E]) {
	slf.entityLeaveVisionEventHandles = append(slf.entityLeaveVisionEventHandles, handle)
}

// OnEntityLeaveVisionEvent 在新对象离开视野时将会立刻执行被注册的事件处理函数
func (slf *event[EID, PosType, E]) OnEntityLeaveVisionEvent(entity, target E) {
	for _, handle := range slf.entityLeaveVisionEventHandles {
		handle(entity, target)
	}
}
