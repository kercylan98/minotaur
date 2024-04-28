package space

import (
"github.com/kercylan98/minotaur/toolkit/collection"
"github.com/kercylan98/minotaur/utils/generic"
"sync"
)

const UnknownSeat = -1 // 未知座位号

func newRoomController[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]](manager *RoomManager[EntityID, RoomID, Entity, Room], room Room, options *RoomControllerOptions[EntityID, RoomID, Entity, Room]) *RoomController[EntityID, RoomID, Entity, Room] {
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
	options         *RoomControllerOptions[EntityID, RoomID, Entity, Room]
	room            Room
	entities        map[EntityID]Entity
	entitiesRWMutex sync.RWMutex

	vacancy []int       // 空缺的座位
	seat    []*EntityID // 座位上的玩家
	owner   *EntityID   // 房主
}

// HasOwner 判断是否有房主
func (rc *RoomController[EntityID, RoomID, Entity, Room]) HasOwner() bool {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	return rc.owner != nil
}

// IsOwner 判断是否为房主
func (rc *RoomController[EntityID, RoomID, Entity, Room]) IsOwner(entityId EntityID) bool {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	return rc.owner != nil && *rc.owner == entityId
}

// GetOwner 获取房主
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetOwner() Entity {
	return rc.GetEntity(*rc.owner)
}

// GetOwnerID 获取房主 ID
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetOwnerID() EntityID {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	return *rc.owner
}

// GetOwnerExist 获取房间，并返回房主是否存在的状态
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetOwnerExist() (Entity, bool) {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	entity, exist := rc.entities[*rc.owner]
	return entity, exist
}

// SetOwner 设置房主
func (rc *RoomController[EntityID, RoomID, Entity, Room]) SetOwner(entityId EntityID) {
	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()
	old := *rc.owner
	rc.owner = &entityId
	rc.manager.OnRoomOwnerChangeEvent(rc, &old, &entityId)
}

// DelOwner 删除房主，将房间设置为无主的状态
func (rc *RoomController[EntityID, RoomID, Entity, Room]) DelOwner() {
	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()
	old := *rc.owner
	rc.owner = nil
	rc.manager.OnRoomOwnerChangeEvent(rc, &old, nil)
}

// JoinSeat 设置特定对象加入座位，当具体的座位不存在的时候，将会自动分配座位
//   - 当目标座位存在玩家或未添加到房间中的时候，将会返回错误
func (rc *RoomController[EntityID, RoomID, Entity, Room]) JoinSeat(entityId EntityID, seat ...int) error {
	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()
	_, exist := rc.entities[entityId]
	if !exist {
		return ErrNotInRoom
	}
	var targetSeat int
	if len(seat) > 0 {
		targetSeat = seat[0]
		if targetSeat < len(rc.seat) && rc.seat[targetSeat] != nil {
			return ErrSeatNotEmpty
		}
	} else {
		if len(rc.vacancy) > 0 {
			targetSeat = rc.vacancy[0]
			rc.vacancy = rc.vacancy[1:]
		} else {
			targetSeat = len(rc.seat)
		}
	}

	if targetSeat >= len(rc.seat) {
		rc.seat = append(rc.seat, make([]*EntityID, targetSeat-len(rc.seat)+1)...)
	}

	rc.seat[targetSeat] = &entityId
	return nil
}

// LeaveSeat 离开座位
func (rc *RoomController[EntityID, RoomID, Entity, Room]) LeaveSeat(entityId EntityID) {
	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()
	rc.leaveSeat(entityId)
}

// leaveSeat 离开座位（无锁）
func (rc *RoomController[EntityID, RoomID, Entity, Room]) leaveSeat(entityId EntityID) {
	for i, seat := range rc.seat {
		if seat != nil && *seat == entityId {
			rc.seat[i] = nil
			rc.vacancy = append(rc.vacancy, i)
			break
		}
	}
}

// GetSeat 获取座位
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetSeat(entityId EntityID) int {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	for i, seat := range rc.seat {
		if seat != nil && *seat == entityId {
			return i
		}
	}
	return UnknownSeat
}

// GetFirstNotEmptySeat 获取第一个非空座位号，如果没有非空座位，将返回 UnknownSeat
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetFirstNotEmptySeat() int {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	for i, seat := range rc.seat {
		if seat != nil {
			return i
		}
	}
	return UnknownSeat
}

// GetFirstEmptySeatEntity 获取第一个空座位上的实体，如果没有空座位，将返回空实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetFirstEmptySeatEntity() (entity Entity) {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	for _, seat := range rc.seat {
		if seat == nil {
			return rc.entities[*seat]
		}
	}
	return entity
}

// GetRandomEntity 获取随机实体，如果房间中没有实体，将返回空实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetRandomEntity() (entity Entity) {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	for _, entity = range rc.entities {
		return entity
	}
	return entity
}

// GetNotEmptySeat 获取非空座位
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetNotEmptySeat() []int {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	var seats []int
	for i, player := range rc.seat {
		if player != nil {
			seats = append(seats, i)
		}
	}
	return seats
}

// GetEmptySeat 获取空座位
//   - 空座位需要在有对象离开座位后才可能出现
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetEmptySeat() []int {
	return collection.CloneSlice(rc.vacancy)
}

// HasSeat 判断是否有座位
func (rc *RoomController[EntityID, RoomID, Entity, Room]) HasSeat(entityId EntityID) bool {
	return rc.GetSeat(entityId) != UnknownSeat
}

// GetSeatEntityCount 获取座位上的实体数量
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntityCount() int {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	var count int
	for _, seat := range rc.seat {
		if seat != nil {
			count++
		}
	}
	return count
}

// GetSeatEntities 获取座位上的实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntities() map[EntityID]Entity {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	var entities = make(map[EntityID]Entity)
	for _, entityId := range rc.seat {
		if entityId != nil {
			entities[*entityId] = rc.entities[*entityId]
		}
	}
	return entities
}

// GetSeatEntitiesByOrdered 有序的获取座位上的实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntitiesByOrdered() []Entity {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	var entities = make([]Entity, 0, len(rc.seat))
	for _, entityId := range rc.seat {
		if entityId != nil {
			entities = append(entities, rc.entities[*entityId])
		}
	}
	return entities
}

// GetSeatEntitiesByOrderedAndContainsEmpty 获取有序的座位上的实体，包含空座位
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntitiesByOrderedAndContainsEmpty() []Entity {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	var entities = make([]Entity, len(rc.seat))
	for i, entityId := range rc.seat {
		if entityId != nil {
			entities[i] = rc.entities[*entityId]
		}
	}
	return entities
}

// GetSeatEntity 获取座位上的实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetSeatEntity(seat int) (entity Entity) {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	if seat < len(rc.seat) {
		eid := rc.seat[seat]
		if eid != nil {
			return rc.entities[*eid]
		}
	}
	return entity
}

// ContainEntity 房间内是否包含实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) ContainEntity(id EntityID) bool {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	_, exist := rc.entities[id]
	return exist
}

// GetRoom 获取原始房间实例，该实例为被接管的房间的原始实例
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetRoom() Room {
	return rc.room
}

// GetEntities 获取所有实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetEntities() map[EntityID]Entity {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	return collection.CloneMap(rc.entities)
}

// HasEntity 判断是否有实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) HasEntity(id EntityID) bool {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	_, exist := rc.entities[id]
	return exist
}

// GetEntity 获取实体
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetEntity(id EntityID) Entity {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	return rc.entities[id]
}

// GetEntityExist 获取实体，并返回实体是否存在的状态
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetEntityExist(id EntityID) (Entity, bool) {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	entity, exist := rc.entities[id]
	return entity, exist
}

// GetEntityIDs 获取所有实体ID
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetEntityIDs() []EntityID {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	return collection.ConvertMapKeysToSlice(rc.entities)
}

// GetEntityCount 获取实体数量
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetEntityCount() int {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	return len(rc.entities)
}

// ChangePassword 修改房间密码
//   - 当房间密码为 nil 时，将会取消密码
func (rc *RoomController[EntityID, RoomID, Entity, Room]) ChangePassword(password *string) {
	old := rc.options.password
	rc.options.password = password
	rc.manager.OnRoomChangePasswordEvent(rc, old, rc.options.password)
}

// AddEntity 添加实体，如果房间存在密码，应使用 AddEntityByPassword 函数进行添加，否则将始终返回 ErrRoomPasswordNotMatch 错误
//   - 当房间已满时，将会返回 ErrRoomFull 错误
func (rc *RoomController[EntityID, RoomID, Entity, Room]) AddEntity(entity Entity) error {
	if rc.options.password != nil {
		return ErrRoomPasswordNotMatch
	}
	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()

	if rc.options.maxEntityCount != nil && len(rc.entities) > *rc.options.maxEntityCount {
		return ErrRoomFull
	}
	rc.entities[entity.GetId()] = entity

	rc.manager.OnRoomAddEntityEvent(rc, entity)
	return nil
}

// AddEntityByPassword 通过房间密码添加实体到该房间中
//   - 当未设置房间密码时，password 参数将会被忽略
//   - 当房间密码不匹配时，将会返回 ErrRoomPasswordNotMatch 错误
//   - 当房间已满时，将会返回 ErrRoomFull 错误
func (rc *RoomController[EntityID, RoomID, Entity, Room]) AddEntityByPassword(entity Entity, password string) error {
	if rc.options.password == nil || *rc.options.password != password {
		return ErrRoomPasswordNotMatch
	}
	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()

	if rc.options.maxEntityCount != nil && len(rc.entities) > *rc.options.maxEntityCount {
		return ErrRoomFull
	}
	rc.entities[entity.GetId()] = entity

	rc.manager.OnRoomAddEntityEvent(rc, entity)
	return nil
}

// RemoveEntity 移除实体
//   - 当实体被移除时如果实体在座位上，将会自动离开座位
//   - 如果实体为房主，将会根据 RoomControllerOptions.WithOwnerInherit 函数的设置进行继承
func (rc *RoomController[EntityID, RoomID, Entity, Room]) RemoveEntity(id EntityID) {
	rc.entitiesRWMutex.RLock()
	defer rc.entitiesRWMutex.RUnlock()
	rc.removeEntity(id)
}

// removeEntity 移除实体（无锁）
func (rc *RoomController[EntityID, RoomID, Entity, Room]) removeEntity(id EntityID) {
	rc.leaveSeat(id)
	entity, exist := rc.entities[id]
	if !exist {
		return
	}
	delete(rc.entities, id)
	if !rc.options.ownerInherit {
		if rc.owner != nil && *rc.owner == id {
			rc.owner = nil
		}
	} else {
		if rc.owner != nil && *rc.owner == id {
			rc.owner = rc.options.ownerInheritHandler(rc)
			defer rc.manager.OnRoomOwnerChangeEvent(rc, &id, rc.owner)
		}
	}
	rc.manager.OnRoomRemoveEntityEvent(rc, entity)
}

// RemoveAllEntities 移除该房间中的所有实体
//   - 当实体被移除时如果实体在座位上，将会自动离开座位
//   - 如果实体为房主，将会根据 RoomControllerOptions.WithOwnerInherit 函数的设置进行继承
func (rc *RoomController[EntityID, RoomID, Entity, Room]) RemoveAllEntities() {
	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()
	for id := range rc.entities {
		rc.removeEntity(id)
		delete(rc.entities, id)
	}
}

// Destroy 销毁房间，房间会从 RoomManager 中移除，同时所有房间的实体、座位等数据都会被清空
//   - 该函数与 RoomManager.DestroyRoom 相同，RoomManager.DestroyRoom 函数为该函数的快捷方式
func (rc *RoomController[EntityID, RoomID, Entity, Room]) Destroy() {
	rc.manager.roomsRWMutex.Lock()
	defer rc.manager.roomsRWMutex.Unlock()

	delete(rc.manager.rooms, rc.room.GetId())
	rc.manager.OnRoomDestroyEvent(rc)

	rc.entitiesRWMutex.Lock()
	defer rc.entitiesRWMutex.Unlock()

	for eid := range rc.entities {
		rc.removeEntity(eid)
		delete(rc.entities, eid)
	}

	rc.entities = make(map[EntityID]Entity)
	rc.seat = rc.seat[:]
	rc.vacancy = rc.vacancy[:]
}

// GetRoomManager 获取该房间控制器所属的房间管理器
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetRoomManager() *RoomManager[EntityID, RoomID, Entity, Room] {
	return rc.manager
}

// GetRoomID 获取房间 ID
func (rc *RoomController[EntityID, RoomID, Entity, Room]) GetRoomID() RoomID {
	return rc.room.GetId()
}

// Broadcast 广播，该函数会将所有房间中满足 conditions 的对象传入 handler 中进行处理
func (rc *RoomController[EntityID, RoomID, Entity, Room]) Broadcast(handler func(Entity), conditions ...func(Entity) bool) {
	rc.entitiesRWMutex.RLock()
	entities := collection.CloneMap(rc.entities)
	rc.entitiesRWMutex.RUnlock()
	for _, entity := range entities {
		for _, condition := range conditions {
			if !condition(entity) {
				continue
			}
		}
		handler(entity)
	}
}
