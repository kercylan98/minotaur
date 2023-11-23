package activity

import "github.com/kercylan98/minotaur/utils/generic"

type none byte

// DefineNoneDataActivity 声明无数据的活动类型
func DefineNoneDataActivity[Type, ID generic.Basic](activityType Type) NoneDataActivityController[Type, ID, none, none, none] {
	return regController(activityType, &Controller[Type, ID, none, none, none]{
		t: activityType,
	})
}

// DefineGlobalDataActivity 声明拥有全局数据的活动类型
func DefineGlobalDataActivity[Type, ID generic.Basic, Data any](activityType Type, initializer func(activityId ID, data *DataMeta[Data])) GlobalDataActivityController[Type, ID, Data, none, none] {
	return regController(activityType, &Controller[Type, ID, Data, none, none]{
		t:          activityType,
		globalData: make(map[ID]*DataMeta[Data]),
		globalInit: initializer,
	})
}

// DefineEntityDataActivity 声明拥有实体数据的活动类型
func DefineEntityDataActivity[Type, ID, EntityID generic.Basic, EntityData any](activityType Type, initializer func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) EntityDataActivityController[Type, ID, none, EntityID, EntityData] {
	return regController(activityType, &Controller[Type, ID, none, EntityID, EntityData]{
		t:          activityType,
		entityData: make(map[ID]map[EntityID]*EntityDataMeta[EntityData]),
		entityInit: initializer,
	})
}

// DefineGlobalAndEntityDataActivity 声明拥有全局数据和实体数据的活动类型
func DefineGlobalAndEntityDataActivity[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any](activityType Type, globalInitializer func(activityId ID, data *DataMeta[Data]), entityInitializer func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) GlobalAndEntityDataActivityController[Type, ID, Data, EntityID, EntityData] {
	return regController(activityType, &Controller[Type, ID, Data, EntityID, EntityData]{
		t:          activityType,
		globalData: make(map[ID]*DataMeta[Data]),
		entityData: make(map[ID]map[EntityID]*EntityDataMeta[EntityData]),
		globalInit: globalInitializer,
		entityInit: entityInitializer,
	})
}

// Controller 活动控制器
type Controller[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] struct {
	t                 Type                                                                                                                            // 活动类型
	activities        map[ID]*Activity[Type, ID]                                                                                                      // 活动
	globalInit        func(activityId ID, data *DataMeta[Data])                                                                                       // 全局初始化器
	entityInit        func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])                                                        // 实体初始化器
	globalDataLoader  func(activityId any)                                                                                                            // 全局数据加载器
	entityDataLoader  func(activityId any, entityId any)                                                                                              // 实体数据加载器
	globalInitializer func(activityType Type, activityId ID, data map[ID]*DataMeta[Data])                                                             // 全局数据初始化器
	entityInitializer func(activityType Type, activityId ID, data map[ID]*DataMeta[Data], entityData map[ID]map[EntityID]*EntityDataMeta[EntityData]) // 实体数据初始化器
	globalData        map[ID]*DataMeta[Data]                                                                                                          // 全局数据
	entityData        map[ID]map[EntityID]*EntityDataMeta[EntityData]                                                                                 // 实体数据
}

// GetGlobalData 获取特定活动全局数据
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) GetGlobalData(activityId ID) Data {
	slf.globalDataLoader(activityId)
	return slf.globalData[activityId].Data
}

// GetEntityData 获取特定活动实体数据
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) GetEntityData(activityId ID, entityId EntityID) EntityData {
	slf.entityDataLoader(activityId, entityId)
	return slf.entityData[activityId][entityId].Data
}

// GetAllEntityData 获取特定活动所有实体数据
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) GetAllEntityData(activityId ID) map[EntityID]EntityData {
	var entities = make(map[EntityID]EntityData)
	for k, v := range slf.entityData[activityId] {
		entities[k] = v.Data
	}
	return entities
}

// IsOpen 活动是否开启
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) IsOpen(activityId ID) bool {
	activity, exist := slf.activities[activityId]
	if !exist {
		return false
	}
	return activity.state == stateStarted
}

// IsShow 活动是否展示
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) IsShow(activityId ID) bool {
	activity, exist := slf.activities[activityId]
	if !exist {
		return false
	}
	return activity.state == stateUpcoming || (activity.state == stateEnded && activity.tl.HasState(stateExtendedShowEnded))
}

// IsOpenOrShow 活动是否开启或展示
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) IsOpenOrShow(activityId ID) bool {
	activity, exist := slf.activities[activityId]
	if !exist {
		return false
	}
	return activity.state == stateStarted || activity.state == stateUpcoming || (activity.state == stateEnded && activity.tl.HasState(stateExtendedShowEnded))
}

// Refresh 刷新活动
func (slf *Controller[Type, ID, Data, EntityID, EntityData]) Refresh(activityId ID) {
	activity, exist := slf.activities[activityId]
	if !exist {
		return
	}
	activity.refresh()
}
