package space

import "github.com/kercylan98/minotaur/utils/generic"

type (
	RoomAssumeControlEventHandle[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]]  func(controller *RoomController[EntityID, RoomID, Entity, Room])
	RoomDestroyEventHandle[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]]        func(controller *RoomController[EntityID, RoomID, Entity, Room])
	RoomAddEntityEventHandle[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]]      func(controller *RoomController[EntityID, RoomID, Entity, Room], entity Entity)
	RoomRemoveEntityEventHandle[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]]   func(controller *RoomController[EntityID, RoomID, Entity, Room], entity Entity)
	RoomChangePasswordEventHandle[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] func(controller *RoomController[EntityID, RoomID, Entity, Room], oldPassword, password *string)
)

type roomManagerEvents[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	roomAssumeControlEventHandles  []RoomAssumeControlEventHandle[EntityID, RoomID, Entity, Room]
	roomDestroyEventHandles        []RoomDestroyEventHandle[EntityID, RoomID, Entity, Room]
	roomAddEntityEventHandles      []RoomAddEntityEventHandle[EntityID, RoomID, Entity, Room]
	roomRemoveEntityEventHandles   []RoomRemoveEntityEventHandle[EntityID, RoomID, Entity, Room]
	roomChangePasswordEventHandles []RoomChangePasswordEventHandle[EntityID, RoomID, Entity, Room]
}

// RegRoomAssumeControlEvent 注册房间接管事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) RegRoomAssumeControlEvent(handle RoomAssumeControlEventHandle[EntityID, RoomID, Entity, Room]) {
	rme.roomAssumeControlEventHandles = append(rme.roomAssumeControlEventHandles, handle)
}

// OnRoomAssumeControlEvent 房间接管事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) OnRoomAssumeControlEvent(controller *RoomController[EntityID, RoomID, Entity, Room]) {
	for _, handle := range rme.roomAssumeControlEventHandles {
		handle(controller)
	}
}

// RegRoomDestroyEvent 注册房间销毁事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) RegRoomDestroyEvent(handle RoomDestroyEventHandle[EntityID, RoomID, Entity, Room]) {
	rme.roomDestroyEventHandles = append(rme.roomDestroyEventHandles, handle)
}

// OnRoomDestroyEvent 房间销毁事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) OnRoomDestroyEvent(controller *RoomController[EntityID, RoomID, Entity, Room]) {
	for _, handle := range rme.roomDestroyEventHandles {
		handle(controller)
	}
}

// RegRoomAddEntityEvent 注册房间添加对象事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) RegRoomAddEntityEvent(handle RoomAddEntityEventHandle[EntityID, RoomID, Entity, Room]) {
	rme.roomAddEntityEventHandles = append(rme.roomAddEntityEventHandles, handle)
}

// OnRoomAddEntityEvent 房间添加对象事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) OnRoomAddEntityEvent(controller *RoomController[EntityID, RoomID, Entity, Room], entity Entity) {
	for _, handle := range rme.roomAddEntityEventHandles {
		handle(controller, entity)
	}
}

// RegRoomRemoveEntityEvent 注册房间移除对象事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) RegRoomRemoveEntityEvent(handle RoomRemoveEntityEventHandle[EntityID, RoomID, Entity, Room]) {
	rme.roomRemoveEntityEventHandles = append(rme.roomRemoveEntityEventHandles, handle)
}

// OnRoomRemoveEntityEvent 房间移除对象事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) OnRoomRemoveEntityEvent(controller *RoomController[EntityID, RoomID, Entity, Room], entity Entity) {
	for _, handle := range rme.roomRemoveEntityEventHandles {
		handle(controller, entity)
	}
}

// RegRoomChangePasswordEvent 注册房间修改密码事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) RegRoomChangePasswordEvent(handle RoomChangePasswordEventHandle[EntityID, RoomID, Entity, Room]) {
	rme.roomChangePasswordEventHandles = append(rme.roomChangePasswordEventHandles, handle)
}

// OnRoomChangePasswordEvent 房间修改密码事件
func (rme *roomManagerEvents[EntityID, RoomID, Entity, Room]) OnRoomChangePasswordEvent(controller *RoomController[EntityID, RoomID, Entity, Room], oldPassword, password *string) {
	for _, handle := range rme.roomChangePasswordEventHandles {
		handle(controller, oldPassword, password)
	}
}
