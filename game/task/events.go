package task

import (
	"reflect"
)

type (
	RefreshTaskCounterEventHandler[Trigger any]   func(taskType string, trigger Trigger, count int64)         // 刷新任务计数器事件处理函数
	RefreshTaskConditionEventHandler[Trigger any] func(taskType string, trigger Trigger, condition Condition) // 刷新任务条件事件处理函数
)

var (
	refreshTaskCounterEventHandlers = make(map[string][]struct {
		t reflect.Type
		h func(taskType string, trigger any, count int64)
	})
	refreshTaskConditionEventHandlers = make(map[string][]struct {
		t reflect.Type
		h func(taskType string, trigger any, condition Condition)
	})
)

// RegisterRefreshTaskCounterEvent 注册特定任务类型的刷新任务计数器事件处理函数
func RegisterRefreshTaskCounterEvent[Trigger any](taskType string, handler RefreshTaskCounterEventHandler[Trigger]) {
	if refreshTaskCounterEventHandlers == nil {
		refreshTaskCounterEventHandlers = make(map[string][]struct {
			t reflect.Type
			h func(taskType string, trigger any, count int64)
		})
	}
	refreshTaskCounterEventHandlers[taskType] = append(refreshTaskCounterEventHandlers[taskType], struct {
		t reflect.Type
		h func(taskType string, trigger any, count int64)
	}{reflect.TypeOf(handler).In(1), func(taskType string, trigger any, count int64) {
		handler(taskType, trigger.(Trigger), count)
	}})
}

// OnRefreshTaskCounterEvent 触发特定任务类型的刷新任务计数器事件
func OnRefreshTaskCounterEvent(taskType string, trigger any, count int64) {
	if handlers, exist := refreshTaskCounterEventHandlers[taskType]; exist {
		for _, handler := range handlers {
			if !reflect.TypeOf(trigger).AssignableTo(handler.t) {
				continue
			}
			handler.h(taskType, trigger, count)
		}
	}
}

// RegisterRefreshTaskConditionEvent 注册特定任务类型的刷新任务条件事件处理函数
func RegisterRefreshTaskConditionEvent[Trigger any](taskType string, handler RefreshTaskConditionEventHandler[Trigger]) {
	if refreshTaskConditionEventHandlers == nil {
		refreshTaskConditionEventHandlers = make(map[string][]struct {
			t reflect.Type
			h func(taskType string, trigger any, condition Condition)
		})
	}
	refreshTaskConditionEventHandlers[taskType] = append(refreshTaskConditionEventHandlers[taskType], struct {
		t reflect.Type
		h func(taskType string, trigger any, condition Condition)
	}{reflect.TypeOf(handler).In(1), func(taskType string, trigger any, condition Condition) {
		handler(taskType, trigger.(Trigger), condition)
	}})
}

// OnRefreshTaskConditionEvent 触发特定任务类型的刷新任务条件事件
func OnRefreshTaskConditionEvent(taskType string, trigger any, condition Condition) {
	if handlers, exist := refreshTaskConditionEventHandlers[taskType]; exist {
		for _, handler := range handlers {
			if !reflect.TypeOf(trigger).AssignableTo(handler.t) {
				continue
			}
			handler.h(taskType, trigger, condition)
		}
	}
}
