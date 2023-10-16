package space

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
	"sync"
)

// NewRoomManager 创建房间管理器
func NewRoomManager[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]]() *RoomManager[EntityID, RoomID, Entity, Room] {
	return &RoomManager[EntityID, RoomID, Entity, Room]{
		roomManagerEvents: new(roomManagerEvents[EntityID, RoomID, Entity, Room]),
		rooms:             make(map[RoomID]*RoomController[EntityID, RoomID, Entity, Room]),
	}
}

// RoomManager 房间管理器
type RoomManager[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	*roomManagerEvents[EntityID, RoomID, Entity, Room]
	roomsRWMutex sync.RWMutex
	rooms        map[RoomID]*RoomController[EntityID, RoomID, Entity, Room]
}

// AssumeControl 将房间控制权交由 RoomManager 接管
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) AssumeControl(room Room, options ...*RoomControllerOptions) *RoomController[EntityID, RoomID, Entity, Room] {
	controller := newRoomController(slf, room, mergeRoomControllerOptions(options...))
	slf.OnRoomAssumeControlEvent(controller)
	return controller
}

// DestroyRoom 销毁房间
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) DestroyRoom(id RoomID) {
	slf.roomsRWMutex.Lock()
	room, exist := slf.rooms[id]
	slf.roomsRWMutex.Unlock()
	if !exist {
		return
	}
	room.Destroy()
}

// GetRoom 获取房间
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) GetRoom(id RoomID) *RoomController[EntityID, RoomID, Entity, Room] {
	slf.roomsRWMutex.RLock()
	defer slf.roomsRWMutex.RUnlock()
	return slf.rooms[id]
}

// GetRooms 获取所有房间
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) GetRooms() map[RoomID]*RoomController[EntityID, RoomID, Entity, Room] {
	slf.roomsRWMutex.RLock()
	defer slf.roomsRWMutex.RUnlock()
	return hash.Copy(slf.rooms)
}

// GetRoomCount 获取房间数量
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) GetRoomCount() int {
	slf.roomsRWMutex.RLock()
	defer slf.roomsRWMutex.RUnlock()
	return len(slf.rooms)
}

// GetRoomIDs 获取所有房间ID
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) GetRoomIDs() []RoomID {
	slf.roomsRWMutex.RLock()
	defer slf.roomsRWMutex.RUnlock()
	return hash.KeyToSlice(slf.rooms)
}

// HasEntity 判断特定对象是否在任一房间中
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) HasEntity(entityId EntityID) bool {
	slf.roomsRWMutex.RLock()
	rooms := hash.Copy(slf.rooms)
	slf.roomsRWMutex.RUnlock()
	for _, room := range rooms {
		if room.HasEntity(entityId) {
			return true
		}
	}
	return false
}

// GetEntityRooms 获取特定对象所在的房间
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) GetEntityRooms(entityId EntityID) map[RoomID]*RoomController[EntityID, RoomID, Entity, Room] {
	slf.roomsRWMutex.RLock()
	rooms := hash.Copy(slf.rooms)
	slf.roomsRWMutex.RUnlock()
	var result = make(map[RoomID]*RoomController[EntityID, RoomID, Entity, Room])
	for id, room := range rooms {
		if room.HasEntity(entityId) {
			result[id] = room
		}
	}
	return result
}

// Broadcast 向所有房间对象广播消息
func (slf *RoomManager[EntityID, RoomID, Entity, Room]) Broadcast(handler func(Entity), conditions ...func(Entity) bool) {
	slf.roomsRWMutex.RLock()
	rooms := hash.Copy(slf.rooms)
	slf.roomsRWMutex.RUnlock()
	for _, room := range rooms {
		room.Broadcast(handler, conditions...)
	}
}
