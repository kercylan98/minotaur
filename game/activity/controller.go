package activity

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"reflect"
	"sync"
)

type none byte

// DefineNoneDataActivity 声明无数据的活动类型
func DefineNoneDataActivity[Type, ID generic.Basic](activityType Type) NoneDataActivityController[Type, ID, *none, none, *none] {
	return regController(&Controller[Type, ID, *none, none, *none]{
		t: activityType,
	})
}

// DefineGlobalDataActivity 声明拥有全局数据的活动类型
func DefineGlobalDataActivity[Type, ID generic.Basic, Data any](activityType Type) GlobalDataActivityController[Type, ID, Data, none, *none] {
	return regController(&Controller[Type, ID, Data, none, *none]{
		t: activityType,
	})
}

// DefineEntityDataActivity 声明拥有实体数据的活动类型
func DefineEntityDataActivity[Type, ID, EntityID generic.Basic, EntityData any](activityType Type) EntityDataActivityController[Type, ID, *none, EntityID, EntityData] {
	return regController(&Controller[Type, ID, *none, EntityID, EntityData]{
		t: activityType,
	})
}

// DefineGlobalAndEntityDataActivity 声明拥有全局数据和实体数据的活动类型
func DefineGlobalAndEntityDataActivity[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any](activityType Type) GlobalAndEntityDataActivityController[Type, ID, Data, EntityID, EntityData] {
	return regController(&Controller[Type, ID, Data, EntityID, EntityData]{
		t: activityType,
	})
}

// Controller 活动控制器
type Controller[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] struct {
	t          Type                                                                     // 活动类型
	activities map[ID]*Activity[Type, ID]                                               // 活动列表
	globalData map[ID]*DataMeta[Data]                                                   // 全局数据
	entityData map[ID]map[EntityID]*EntityDataMeta[EntityData]                          // 实体数据
	entityTof  reflect.Type                                                             // 实体数据类型
	globalInit func(activityId ID, data *DataMeta[Data])                                // 全局数据初始化函数
	entityInit func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData]) // 实体数据初始化函数
	mutex      sync.RWMutex
}

// GetGlobalData 获取特定活动全局数据
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) GetGlobalData(activityId ID) Data {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	global := slf.globalData[activityId]
	if slf.globalInit != nil {
		global.once.Do(func() {
			slf.globalInit(activityId, global)
		})
	}
	return global.Data
}

// GetEntityData 获取特定活动实体数据
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) GetEntityData(activityId ID, entityId EntityID) EntityData {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	entities, exist := slf.entityData[activityId]
	if !exist {
		entities = make(map[EntityID]*EntityDataMeta[EntityData])
		slf.entityData[activityId] = entities
	}
	entity, exist := entities[entityId]
	if !exist {
		entity = &EntityDataMeta[EntityData]{
			Data: reflect.New(slf.entityTof).Interface().(EntityData),
		}
		entities[entityId] = entity
	}
	if slf.entityInit != nil {
		entity.once.Do(func() {
			slf.entityInit(activityId, entityId, entity)
		})
	}
	return entity.Data
}

// IsOpen 活动是否开启
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) IsOpen(activityId ID) bool {
	slf.mutex.RLock()
	activity, exist := slf.activities[activityId]
	slf.mutex.RUnlock()
	if !exist {
		return false
	}
	activity.mutex.RLock()
	defer activity.mutex.RUnlock()
	return activity.state == stateStarted
}

// IsShow 活动是否展示
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) IsShow(activityId ID) bool {
	slf.mutex.RLock()
	activity, exist := slf.activities[activityId]
	slf.mutex.RUnlock()
	if !exist {
		return false
	}
	activity.mutex.RLock()
	defer activity.mutex.RUnlock()
	return activity.state == stateUpcoming || (activity.state == stateEnded && activity.options.Tl.HasState(stateExtendedShowEnded))
}

// IsOpenOrShow 活动是否开启或展示
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) IsOpenOrShow(activityId ID) bool {
	slf.mutex.RLock()
	activity, exist := slf.activities[activityId]
	slf.mutex.RUnlock()
	if !exist {
		return false
	}

	activity.mutex.RLock()
	defer activity.mutex.RUnlock()
	return activity.state == stateStarted || activity.state == stateUpcoming || (activity.state == stateEnded && activity.options.Tl.HasState(stateExtendedShowEnded))
}

// Refresh 刷新活动
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) Refresh(activityId ID) {
	slf.mutex.RLock()
	activity, exist := slf.activities[activityId]
	slf.mutex.RUnlock()
	if !exist {
		return
	}
	activity.refresh()
}

func (slf *Controller[Type, ID, Data, EntityID, EntityData]) InitializeNoneData(handler func(activityId ID, data *DataMeta[Data])) NoneDataActivityController[Type, ID, Data, EntityID, EntityData] {
	slf.globalInit = handler
	return slf
}

func (slf *Controller[Type, ID, Data, EntityID, EntityData]) InitializeGlobalData(handler func(activityId ID, data *DataMeta[Data])) GlobalDataActivityController[Type, ID, Data, EntityID, EntityData] {
	slf.globalInit = handler
	return slf
}

func (slf *Controller[Type, ID, Data, EntityID, EntityData]) InitializeEntityData(handler func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) EntityDataActivityController[Type, ID, Data, EntityID, EntityData] {
	slf.entityInit = handler
	return slf
}

func (slf *Controller[Type, ID, Data, EntityID, EntityData]) InitializeGlobalAndEntityData(handler func(activityId ID, data *DataMeta[Data]), entityHandler func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) GlobalAndEntityDataActivityController[Type, ID, Data, EntityID, EntityData] {
	slf.globalInit = handler
	slf.entityInit = entityHandler
	return slf
}
