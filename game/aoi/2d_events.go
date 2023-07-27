package aoi

type (
	EntityJoinVisionEventHandle[E TwoDimensionalEntity]  func(entity, target E)
	EntityLeaveVisionEventHandle[E TwoDimensionalEntity] func(entity, target E)
)

type event[E TwoDimensionalEntity] struct {
	entityJoinVisionEventHandles  []EntityJoinVisionEventHandle[E]
	entityLeaveVisionEventHandles []EntityLeaveVisionEventHandle[E]
}

// RegEntityJoinVisionEvent 在新对象进入视野时将会立刻执行被注册的事件处理函数
func (slf *event[E]) RegEntityJoinVisionEvent(handle EntityJoinVisionEventHandle[E]) {
	slf.entityJoinVisionEventHandles = append(slf.entityJoinVisionEventHandles, handle)
}

// OnEntityJoinVisionEvent 在新对象进入视野时将会立刻执行被注册的事件处理函数
func (slf *event[E]) OnEntityJoinVisionEvent(entity, target E) {
	for _, handle := range slf.entityJoinVisionEventHandles {
		handle(entity, target)
	}
}

// RegEntityLeaveVisionEvent 在新对象离开视野时将会立刻执行被注册的事件处理函数
func (slf *event[E]) RegEntityLeaveVisionEvent(handle EntityLeaveVisionEventHandle[E]) {
	slf.entityLeaveVisionEventHandles = append(slf.entityLeaveVisionEventHandles, handle)
}

// OnEntityLeaveVisionEvent 在新对象离开视野时将会立刻执行被注册的事件处理函数
func (slf *event[E]) OnEntityLeaveVisionEvent(entity, target E) {
	for _, handle := range slf.entityLeaveVisionEventHandles {
		handle(entity, target)
	}
}
