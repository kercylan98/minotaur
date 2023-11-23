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
}

// GlobalDataActivityController 全局数据活动控制器
type GlobalDataActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	// GetGlobalData 获取全局数据
	GetGlobalData(activityId ID) Data
}

// EntityDataActivityController 实体数据活动控制器
type EntityDataActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	// GetEntityData 获取实体数据
	GetEntityData(activityId ID, entityId EntityID) EntityData
}

// GlobalAndEntityDataActivityController 全局数据和实体数据活动控制器
type GlobalAndEntityDataActivityController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	// GetGlobalData 获取全局数据
	GetGlobalData(activityId ID) Data
	// GetEntityData 获取实体数据
	GetEntityData(activityId ID, entityId EntityID) EntityData
}
