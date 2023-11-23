package activity

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/times"
	"reflect"
	"sync"
	"time"
)

var (
	controllers                map[any]any                                                                                 // type -> controller （特定类型活动控制器）
	controllerRegisters        map[any]func(activityType, activityId any) (act any, optionInitCallback func(activity any)) // type -> register （控制活动注册到特定类型控制器的注册机）
	controllerGlobalDataReader []func(handler func(activityType, activityId, data any))                                    // 活动全局数据读取器
	controllerEntityDataReader []func(handler func(activityType, activityId, entityId, data any))                          // 活动实体数据读取器
	controllerReset            map[any]func(activityId any)                                                                // type -> reset （活动数据重置器）
	controllersLock            sync.RWMutex
)

func init() {
	controllers = make(map[any]any)
	controllerRegisters = make(map[any]func(activityType, activityId any) (act any, optionInitCallback func(activity any)))
	controllerGlobalDataReader = make([]func(handler func(activityType, activityId, data any)), 0)
	controllerEntityDataReader = make([]func(handler func(activityType, activityId, entityId, data any)), 0)
	controllerReset = make(map[any]func(activityId any))
}

// regController 注册活动类型
func regController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any](activityType Type, controller *Controller[Type, ID, Data, EntityID, EntityData]) *Controller[Type, ID, Data, EntityID, EntityData] {
	controllersLock.Lock()
	defer controllersLock.Unlock()
	controller.activities = make(map[ID]*Activity[Type, ID])
	controllerGlobalDataReader = append(controllerGlobalDataReader, func(handler func(activityType, activityId, data any)) {
		for activityId, data := range controller.globalData {
			handler(activityType, activityId, data)
		}
	})
	controllerEntityDataReader = append(controllerEntityDataReader, func(handler func(activityType, activityId, entityId, data any)) {
		for activityId, entities := range controller.entityData {
			for entityId, data := range entities {
				handler(activityType, activityId, entityId, data)
			}
		}
	})

	if controller.globalData != nil {
		var d Data
		var tof = reflect.TypeOf(d)
		if tof.Kind() == reflect.Pointer {
			tof = tof.Elem()
		}
		controller.globalDataLoader = func(activityId any) {
			var id = activityId.(ID)
			if _, exist := controller.globalData[id]; exist {
				return
			}
			data := &DataMeta[Data]{
				Data: reflect.New(tof).Interface().(Data),
			}
			if controller.globalInit != nil {
				controller.globalInit(id, data)
			}
			controller.globalData[id] = data
		}
	}
	if controller.entityData != nil {
		var d Data
		var tof = reflect.TypeOf(d)
		if tof.Kind() == reflect.Pointer {
			tof = tof.Elem()
		}
		controller.entityDataLoader = func(activityId any, entityId any) {
			var id, eid = activityId.(ID), entityId.(EntityID)
			entities, exist := controller.entityData[id]
			if !exist {
				entities = make(map[EntityID]*EntityDataMeta[EntityData])
				controller.entityData[id] = entities
			}
			if _, exist = entities[eid]; exist {
				return
			}
			data := &EntityDataMeta[EntityData]{
				Data: reflect.New(tof).Interface().(EntityData),
			}
			if controller.entityInit != nil {
				controller.entityInit(id, eid, data)
			}
			entities[eid] = data
		}
	}
	controllers[activityType] = controller
	controllerRegisters[activityType] = func(activityType, activityId any) (act any, optionInitCallback func(activity any)) {
		var at, ai = activityType.(Type), activityId.(ID)
		activity, exist := controller.activities[ai]
		if !exist {
			activity = &Activity[Type, ID]{
				t:         at,
				id:        ai,
				tl:        times.NewStateLine[byte](stateClosed),
				tickerKey: fmt.Sprintf("activity:%d:%v", reflect.ValueOf(at).Kind(), ai),
				getLastNewDayTime: func() time.Time {
					return controller.globalData[ai].LastNewDay
				},
				setLastNewDayTime: func(t time.Time) {
					controller.globalData[ai].LastNewDay = t
				},
			}
		}
		controller.activities[activity.id] = activity
		return activity, func(activity any) {
			act := activity.(*Activity[Type, ID])
			if !act.lazy {
				controller.GetGlobalData(ai)
			}
			if act.retention > 0 {
				act.retentionKey = fmt.Sprintf("%s:retention", act.tickerKey)
			}
		}
	}
	controllerReset[activityType] = func(activityId any) {
		var id = activityId.(ID)
		delete(controller.globalData, id)
		delete(controller.entityData, id)
	}
	return controller
}

// getControllerRegister 获取活动类型注册机
func getControllerRegister[Type generic.Basic](activityType Type) func(activityType, activityId any) (act any, optionInitCallback func(activity any)) {
	controllersLock.RLock()
	defer controllersLock.RUnlock()
	return controllerRegisters[activityType]
}

// resetActivityData 重置活动数据
func resetActivityData[Type, ID generic.Basic](activityType Type, activityId ID) {
	controllersLock.RLock()
	defer controllersLock.RUnlock()
	reset, exist := controllerReset[activityType]
	if !exist {
		return
	}
	reset(activityId)
}
