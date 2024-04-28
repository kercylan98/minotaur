package activity

import "github.com/kercylan98/minotaur/utils/generic"

type BasicActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	// IsOpen 活动是否开启
	IsOpen(activityId ID) bool
	// IsShow 活动是否展示
	IsShow(activityId ID) bool
	// IsOpenOrShow 活动是否开启或展示
	IsOpenOrShow(activityId ID) bool
	// Refresh 刷新活动
	Refresh(activityId ID)
}

// NoneDataActivityController 无数据活动控制器
type NoneDataActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	// InitializeNoneData 初始化活动
	//  - 该函数提供了一个操作活动数据的入口，可以在该函数中对传入的活动数据进行初始化
	//
	// 对于无数据活动，该函数的意义在于，可以在该函数中对活动进行初始化，比如设置活动的状态等，虽然为无数据活动，但是例如活动本身携带的状态数据也是需要加载的
	InitializeNoneData(handler func(activityId ID, data *DataMeta[Data])) NoneDataActivityController[Type, ID, Data, EntityID, EntityData]
}

// GlobalDataActivityController 全局数据活动控制器
type GlobalDataActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	// GetGlobalData 获取全局数据
	GetGlobalData(activityId ID) Data
	// InitializeGlobalData 初始化活动
	//  - 该函数提供了一个操作活动数据的入口，可以在该函数中对传入的活动数据进行初始化
	InitializeGlobalData(handler func(activityId ID, data *DataMeta[Data])) GlobalDataActivityController[Type, ID, Data, EntityID, EntityData]
}

// EntityDataActivityController 实体数据活动控制器
type EntityDataActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	// GetEntityData 获取实体数据
	GetEntityData(activityId ID, entityId EntityID) EntityData
	// InitializeEntityData 初始化活动
	//  - 该函数提供了一个操作活动数据的入口，可以在该函数中对传入的活动数据进行初始化
	InitializeEntityData(handler func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) EntityDataActivityController[Type, ID, Data, EntityID, EntityData]
}

// GlobalAndEntityDataActivityController 全局数据和实体数据活动控制器
type GlobalAndEntityDataActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	// GetGlobalData 获取全局数据
	GetGlobalData(activityId ID) Data
	// GetEntityData 获取实体数据
	GetEntityData(activityId ID, entityId EntityID) EntityData
	// InitializeGlobalAndEntityData 初始化活动
	//  - 该函数提供了一个操作活动数据的入口，可以在该函数中对传入的活动数据进行初始化
	InitializeGlobalAndEntityData(handler func(activityId ID, data *DataMeta[Data]), entityHandler func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) GlobalAndEntityDataActivityController[Type, ID, Data, EntityID, EntityData]
}
