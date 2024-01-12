package space

import (
	"github.com/kercylan98/minotaur/utils/collection"
	"github.com/kercylan98/minotaur/utils/generic"
	"sync"
)

// NewRoomManager 创建房间管理器 RoomManager 的实例
func NewRoomManager[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]]() *RoomManager[EntityID, RoomID, Entity, Room] {
	return &RoomManager[EntityID, RoomID, Entity, Room]{
		roomManagerEvents: new(roomManagerEvents[EntityID, RoomID, Entity, Room]),
		rooms:             make(map[RoomID]*RoomController[EntityID, RoomID, Entity, Room]),
	}
}

// RoomManager 房间管理器是用于对房间进行管理的基本单元，通过该实例可以对房间进行增删改查等操作
//   - 该实例是线程安全的
type RoomManager[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	*roomManagerEvents[EntityID, RoomID, Entity, Room]
	roomsRWMutex sync.RWMutex
	rooms        map[RoomID]*RoomController[EntityID, RoomID, Entity, Room]
}

// AssumeControl 将房间控制权交由 RoomManager 接管，返回 RoomController 实例
//   - 当任何房间需要被 RoomManager 管理时，都应该调用该方法获取到 RoomController 实例后进行操作
//   - 房间被接管后需要在释放房间控制权时调用 RoomController.Destroy 方法，否则将会导致 RoomManager 一直持有房间资源
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) AssumeControl(room Room, options ...*RoomControllerOptions[EntityID, RoomID, Entity, Room]) *RoomController[EntityID, RoomID, Entity, Room] {
	controller := newRoomController(rm, room, mergeRoomControllerOptions(options...))
	rm.OnRoomAssumeControlEvent(controller)
	return controller
}

// DestroyRoom 销毁房间，该函数为 RoomController.Destroy 的快捷方式
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) DestroyRoom(id RoomID) {
	rm.roomsRWMutex.Lock()
	room, exist := rm.rooms[id]
	rm.roomsRWMutex.Unlock()
	if !exist {
		return
	}
	room.Destroy()
}

// GetRoom 通过房间 ID 获取对应房间的控制器 RoomController，当房间不存在时将返回 nil
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) GetRoom(id RoomID) *RoomController[EntityID, RoomID, Entity, Room] {
	rm.roomsRWMutex.RLock()
	defer rm.roomsRWMutex.RUnlock()
	return rm.rooms[id]
}

// GetRooms 获取包含所有房间 ID 到对应控制器 RoomController 的映射
//   - 返回值的 map 为拷贝对象，可安全的对其进行增删等操作
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) GetRooms() map[RoomID]*RoomController[EntityID, RoomID, Entity, Room] {
	rm.roomsRWMutex.RLock()
	defer rm.roomsRWMutex.RUnlock()
	return collection.CloneMap(rm.rooms)
}

// GetRoomCount 获取房间管理器接管的房间数量
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) GetRoomCount() int {
	rm.roomsRWMutex.RLock()
	defer rm.roomsRWMutex.RUnlock()
	return len(rm.rooms)
}

// GetRoomIDs 获取房间管理器接管的所有房间 ID
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) GetRoomIDs() []RoomID {
	rm.roomsRWMutex.RLock()
	defer rm.roomsRWMutex.RUnlock()
	return collection.ConvertMapKeysToSlice(rm.rooms)
}

// HasEntity 判断特定对象是否在任一房间中，当对象不在任一房间中时将返回 false
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) HasEntity(entityId EntityID) bool {
	rm.roomsRWMutex.RLock()
	rooms := collection.CloneMap(rm.rooms)
	rm.roomsRWMutex.RUnlock()
	for _, room := range rooms {
		if room.HasEntity(entityId) {
			return true
		}
	}
	return false
}

// GetEntityRooms 获取特定对象所在的房间，返回值为房间 ID 到对应控制器 RoomController 的映射
//   - 由于一个对象可能在多个房间中，因此返回值为 map 类型
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) GetEntityRooms(entityId EntityID) map[RoomID]*RoomController[EntityID, RoomID, Entity, Room] {
	rm.roomsRWMutex.RLock()
	rooms := collection.CloneMap(rm.rooms)
	rm.roomsRWMutex.RUnlock()
	var result = make(map[RoomID]*RoomController[EntityID, RoomID, Entity, Room])
	for id, room := range rooms {
		if room.HasEntity(entityId) {
			result[id] = room
		}
	}
	return result
}

// Broadcast 向所有房间对象广播消息，该方法将会遍历所有房间控制器并调用 RoomController.Broadcast 方法
func (rm *RoomManager[EntityID, RoomID, Entity, Room]) Broadcast(handler func(Entity), conditions ...func(Entity) bool) {
	rm.roomsRWMutex.RLock()
	rooms := collection.CloneMap(rm.rooms)
	rm.roomsRWMutex.RUnlock()
	for _, room := range rooms {
		room.Broadcast(handler, conditions...)
	}
}
