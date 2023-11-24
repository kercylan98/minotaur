package activity

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/times"
	"reflect"
	"time"
)

var (
	activityRegister         map[any]func(activityId any, options *Options) (activity any)      // type -> register （控制活动注册到特定类型控制器的注册机）
	activityGlobalDataLoader []func(handler func(activityType, activityId, data any))           // 全局数据加载器
	activityEntityDataLoader []func(handler func(activityType, activityId, entityId, data any)) // 实体数据加载器
)

func init() {
	activityRegister = make(map[any]func(activityId any, options *Options) (activity any))
}

// regController 注册活动类型控制器
func regController[Type, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any](controller *Controller[Type, ID, Data, EntityID, EntityData]) *Controller[Type, ID, Data, EntityID, EntityData] {
	var entityZero EntityData
	controller.activities = make(map[ID]*Activity[Type, ID])
	controller.globalData = make(map[ID]*DataMeta[Data])
	controller.entityTof = reflect.TypeOf(entityZero)
	if controller.entityTof.Kind() == reflect.Pointer {
		controller.entityTof = controller.entityTof.Elem()
	}

	// 反射类型
	var (
		zero Data
		tof  = reflect.TypeOf(zero)
	)
	if tof.Kind() == reflect.Pointer {
		tof = tof.Elem()
	}

	// 活动注册机（注册机内不加载活动数据，仅定义基本活动信息）
	activityRegister[controller.t] = func(aid any, options *Options) any {
		activityId := aid.(ID)
		controller.mutex.Lock()
		activity, exist := controller.activities[activityId]
		if !exist {
			activity = &Activity[Type, ID]{
				t:         controller.t,
				id:        activityId,
				options:   options,
				tickerKey: fmt.Sprintf("activity:%d:%v", reflect.ValueOf(controller.t).Kind(), activityId),
				getLastNewDayTime: func() time.Time {
					return controller.globalData[activityId].LastNewDay
				},
				setLastNewDayTime: func(t time.Time) {
					controller.globalData[activityId].LastNewDay = t
				},
				clearData: func() {
					controller.mutex.Lock()
					defer controller.mutex.Unlock()
					delete(controller.globalData, activityId)
					delete(controller.entityData, activityId)
				},
				initializeData: func() {
					controller.mutex.Lock()
					defer controller.mutex.Unlock()
					controller.globalData[activityId] = &DataMeta[Data]{
						Data: reflect.New(tof).Interface().(Data),
					}
					if controller.entityData == nil {
						controller.entityData = make(map[ID]map[EntityID]*EntityDataMeta[EntityData])
					}
					controller.entityData[activityId] = make(map[EntityID]*EntityDataMeta[EntityData])
				},
			}
			if activity.options == nil {
				activity.options = NewOptions()
			}
			if activity.options.Tl == nil || activity.options.Tl.GetStateCount() == 0 {
				activity.options.Tl = times.NewStateLine[byte](stateClosed)
			}
			controller.activities[activityId] = activity
		}
		controller.mutex.Unlock()

		// 全局数据加载器
		activityGlobalDataLoader = append(activityGlobalDataLoader, func(handler func(activityType any, activityId any, data any)) {
			controller.mutex.RLock()
			data := controller.globalData[activityId]
			controller.mutex.RUnlock()
			handler(controller.t, activityId, data)
		})

		// 实体数据加载器
		activityEntityDataLoader = append(activityEntityDataLoader, func(handler func(activityType any, activityId any, entityId any, data any)) {
			controller.mutex.RLock()
			entities := hash.Copy(controller.entityData[activityId])
			controller.mutex.RUnlock()
			for entityId, data := range entities {
				handler(controller.t, activityId, entityId, data)
			}
		})
		return activity
	}

	return controller
}
