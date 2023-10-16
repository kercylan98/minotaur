package space

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/slice"
	"sync"
)

func newRoomController[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]](manager *RoomManager[EntityID, RoomID, Entity, Room], room Room, options *RoomControllerOptions) *RoomController[EntityID, RoomID, Entity, Room] {
	controller := &RoomController[EntityID, RoomID, Entity, Room]{
		manager:  manager,
		options:  options,
		entities: make(map[EntityID]Entity),
		room:     room,
	}

	manager.roomsRWMutex.Lock()
	defer manager.roomsRWMutex.Unlock()
	manager.rooms[room.GetId()] = controller

	return controller
}

// RoomController 对房间进行操作的控制器，由 RoomManager 接管后返回
type RoomController[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	manager         *RoomManager[EntityID, RoomID, Entity, Room]
	options         *RoomControllerOptions
	room            Room
	entities        map[EntityID]Entity
	entitiesRWMutex sync.RWMutex

	vacancy []int       // 空缺的座位
	seat    []*EntityID // 座位上的玩家
}

// JoinSeat 设置特定对象加入座位，当具体的座位不存在的时候，将会自动分配座位
//   - 当目标座位存在玩家或未添加到房间中的时候，将会返回错误
func (slf *RoomController[EntityID, RoomID, Entity, Room]) JoinSeat(entityId EntityID, seat ...int) error {
	slf.entitiesRWMutex.Lock()
	defer slf.entitiesRWMutex.Unlock()
	_, exist := slf.entities[entityId]
	if !exist {
		return ErrNotInRoom
	}
	var targetSeat int
	if len(seat) > 0 {
		targetSeat = seat[0]
		if targetSeat < len(slf.seat) && slf.seat[targetSeat] != nil {
			return ErrSeatNotEmpty
		}
	} else {
		if len(slf.vacancy) > 0 {
			targetSeat = slf.vacancy[0]
			slf.vacancy = slf.vacancy[1:]
		} else {
			targetSeat = len(slf.seat)
		}
	}

	if targetSeat >= len(slf.seat) {
		slf.seat = append(slf.seat, make([]*EntityID, targetSeat-len(slf.seat)+1)...)
	}

	slf.seat[targetSeat] = &entityId
	return nil
}

// LeaveSeat 离开座位
func (slf *RoomController[EntityID, RoomID, Entity, Room]) LeaveSeat(entityId EntityID) {
	slf.entitiesRWMutex.Lock()
	defer slf.entitiesRWMutex.Unlock()
	slf.leaveSeat(entityId)
}

// leaveSeat 离开座位（无锁）
func (slf *RoomController[EntityID, RoomID, Entity, Room]) leaveSeat(entityId EntityID) {
	for i, seat := range slf.seat {
		if seat != nil && *seat == entityId {
			slf.seat[i] = nil
			slf.vacancy = append(slf.vacancy, i)
			break
		}
	}
}

// GetSeat 获取座位
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetSeat(entityId EntityID) int {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	for i, seat := range slf.seat {
		if seat != nil && *seat == entityId {
			return i
		}
	}
	return -1
}

// GetNotEmptySeat 获取非空座位
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetNotEmptySeat() []int {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	var seats []int
	for i, player := range slf.seat {
		if player != nil {
			seats = append(seats, i)
		}
	}
	return seats
}

// GetEmptySeat 获取空座位
//   - 空座位需要在有对象离开座位后才可能出现
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetEmptySeat() []int {
	return slice.Copy(slf.vacancy)
}

// HasSeat 判断是否有座位
func (slf *RoomController[EntityID, RoomID, Entity, Room]) HasSeat(entityId EntityID) bool {
	return slf.GetSeat(entityId) != -1
}

// GetSeatEntityCount 获取座位上的实体数量
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntityCount() int {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	var count int
	for _, seat := range slf.seat {
		if seat != nil {
			count++
		}
	}
	return count
}

// GetSeatEntities 获取座位上的实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntities() map[EntityID]Entity {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	var entities = make(map[EntityID]Entity)
	for _, entityId := range slf.seat {
		if entityId != nil {
			entities[*entityId] = slf.entities[*entityId]
		}
	}
	return entities
}

// GetSeatEntitiesByOrdered 有序的获取座位上的实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntitiesByOrdered() []Entity {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	var entities = make([]Entity, 0, len(slf.seat))
	for _, entityId := range slf.seat {
		if entityId != nil {
			entities = append(entities, slf.entities[*entityId])
		}
	}
	return entities
}

// GetSeatEntitiesByOrderedAndContainsEmpty 获取有序的座位上的实体，包含空座位
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntitiesByOrderedAndContainsEmpty() []Entity {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	var entities = make([]Entity, len(slf.seat))
	for i, entityId := range slf.seat {
		if entityId != nil {
			entities[i] = slf.entities[*entityId]
		}
	}
	return entities
}

// GetSeatEntity 获取座位上的实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntity(seat int) (entity Entity) {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	if seat < len(slf.seat) {
		eid := slf.seat[seat]
		if eid != nil {
			return slf.entities[*eid]
		}
	}
	return entity
}

// ContainEntity 房间内是否包含实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) ContainEntity(id EntityID) bool {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	_, exist := slf.entities[id]
	return exist
}

// GetRoom 获取原始房间实例
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetRoom() Room {
	return slf.room
}

// GetEntities 获取所有实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetEntities() map[EntityID]Entity {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	return hash.Copy(slf.entities)
}

// HasEntity 判断是否有实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) HasEntity(id EntityID) bool {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	_, exist := slf.entities[id]
	return exist
}

// GetEntity 获取实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetEntity(id EntityID) Entity {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	return slf.entities[id]
}

// GetEntityIDs 获取所有实体ID
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetEntityIDs() []EntityID {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	return hash.KeyToSlice(slf.entities)
}

// GetEntityCount 获取实体数量
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetEntityCount() int {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	return len(slf.entities)
}

// ChangePassword 修改房间密码
//   - 当房间密码为 nil 时，将会取消密码
func (slf *RoomController[EntityID, RoomID, Entity, Room]) ChangePassword(password *string) {
	old := slf.options.password
	slf.options.password = password
	slf.manager.OnRoomChangePasswordEvent(slf, old, slf.options.password)
}

// AddEntity 添加实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) AddEntity(entity Entity) error {
	if slf.options.password != nil {
		return ErrRoomPasswordNotMatch
	}
	slf.entitiesRWMutex.Lock()
	defer slf.entitiesRWMutex.Unlock()

	if slf.options.maxEntityCount != nil && len(slf.entities) > *slf.options.maxEntityCount {
		return ErrRoomFull
	}
	slf.entities[entity.GetId()] = entity

	slf.manager.OnRoomAddEntityEvent(slf, entity)
	return nil
}

// AddEntityByPassword 通过房间密码添加实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) AddEntityByPassword(entity Entity, password string) error {
	if slf.options.password == nil || *slf.options.password != password {
		return ErrRoomPasswordNotMatch
	}
	slf.entitiesRWMutex.Lock()
	defer slf.entitiesRWMutex.Unlock()

	if slf.options.maxEntityCount != nil && len(slf.entities) > *slf.options.maxEntityCount {
		return ErrRoomFull
	}
	slf.entities[entity.GetId()] = entity

	slf.manager.OnRoomAddEntityEvent(slf, entity)
	return nil
}

// RemoveEntity 移除实体
//   - 当实体被移除时如果实体在座位上，将会自动离开座位
func (slf *RoomController[EntityID, RoomID, Entity, Room]) RemoveEntity(id EntityID) {
	slf.entitiesRWMutex.RLock()
	defer slf.entitiesRWMutex.RUnlock()
	slf.removeEntity(id)
}

// removeEntity 移除实体（无锁）
func (slf *RoomController[EntityID, RoomID, Entity, Room]) removeEntity(id EntityID) {
	slf.leaveSeat(id)
	entity, exist := slf.entities[id]
	delete(slf.entities, id)
	if !exist {
		return
	}
	slf.manager.OnRoomRemoveEntityEvent(slf, entity)
}

// RemoveAllEntities 移除所有实体
func (slf *RoomController[EntityID, RoomID, Entity, Room]) RemoveAllEntities() {
	slf.entitiesRWMutex.Lock()
	defer slf.entitiesRWMutex.Unlock()
	for id := range slf.entities {
		slf.removeEntity(id)
		delete(slf.entities, id)
	}
}

// Destroy 销毁房间
func (slf *RoomController[EntityID, RoomID, Entity, Room]) Destroy() {
	slf.manager.roomsRWMutex.Lock()
	defer slf.manager.roomsRWMutex.Unlock()

	delete(slf.manager.rooms, slf.room.GetId())
	slf.manager.OnRoomDestroyEvent(slf)

	slf.entitiesRWMutex.Lock()
	defer slf.entitiesRWMutex.Unlock()

	for eid := range slf.entities {
		slf.removeEntity(eid)
		delete(slf.entities, eid)
	}

	slf.entities = make(map[EntityID]Entity)
	slf.seat = slf.seat[:]
	slf.vacancy = slf.vacancy[:]
}

// GetRoomManager 获取房间管理器
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetRoomManager() *RoomManager[EntityID, RoomID, Entity, Room] {
	return slf.manager
}

// GetRoomID 获取房间ID
func (slf *RoomController[EntityID, RoomID, Entity, Room]) GetRoomID() RoomID {
	return slf.room.GetId()
}

// Broadcast 广播消息
func (slf *RoomController[EntityID, RoomID, Entity, Room]) Broadcast(handler func(Entity), conditions ...func(Entity) bool) {
	slf.entitiesRWMutex.RLock()
	entities := hash.Copy(slf.entities)
	slf.entitiesRWMutex.RUnlock()
	for _, entity := range entities {
		for _, condition := range conditions {
			if !condition(entity) {
				continue
			}
		}
		handler(entity)
	}
}
