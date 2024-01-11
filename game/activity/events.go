package activity

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection"
	listings2 "github.com/kercylan98/minotaur/utils/collection/listings"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/kercylan98/minotaur/utils/times"
	"reflect"
	"time"
)

type (
	UpcomingEventHandler[ID generic.Basic]            func(activityId ID) // 即将开始的活动事件处理器
	StartedEventHandler[ID generic.Basic]             func(activityId ID) // 活动开始事件处理器
	EndedEventHandler[ID generic.Basic]               func(activityId ID) // 活动结束事件处理器
	ExtendedShowStartedEventHandler[ID generic.Basic] func(activityId ID) // 活动结束后延长展示开始事件处理器
	ExtendedShowEndedEventHandler[ID generic.Basic]   func(activityId ID) // 活动结束后延长展示结束事件处理器
	NewDayEventHandler[ID generic.Basic]              func(activityId ID) // 新的一天事件处理器
)

var (
	upcomingEventHandlers       map[any]*listings2.PrioritySlice[func(activityId any)] // 即将开始的活动事件处理器
	startedEventHandlers        map[any]*listings2.PrioritySlice[func(activityId any)] // 活动开始事件处理器
	endedEventHandlers          map[any]*listings2.PrioritySlice[func(activityId any)] // 活动结束事件处理器
	extShowStartedEventHandlers map[any]*listings2.PrioritySlice[func(activityId any)] // 活动结束后延长展示开始事件处理器
	extShowEndedEventHandlers   map[any]*listings2.PrioritySlice[func(activityId any)] // 活动结束后延长展示结束事件处理器
	newDayEventHandlers         map[any]*listings2.PrioritySlice[func(activityId any)] // 新的一天事件处理器
)

func init() {
	upcomingEventHandlers = make(map[any]*listings2.PrioritySlice[func(activityId any)])
	startedEventHandlers = make(map[any]*listings2.PrioritySlice[func(activityId any)])
	endedEventHandlers = make(map[any]*listings2.PrioritySlice[func(activityId any)])
	extShowStartedEventHandlers = make(map[any]*listings2.PrioritySlice[func(activityId any)])
	extShowEndedEventHandlers = make(map[any]*listings2.PrioritySlice[func(activityId any)])
	newDayEventHandlers = make(map[any]*listings2.PrioritySlice[func(activityId any)])
}

// RegUpcomingEvent 注册即将开始的活动事件处理器
func RegUpcomingEvent[Type, ID generic.Basic](activityType Type, handler UpcomingEventHandler[ID], priority ...int) {
	handlers, exist := upcomingEventHandlers[activityType]
	if !exist {
		handlers = listings2.NewPrioritySlice[func(activityId any)]()
		upcomingEventHandlers[activityType] = handlers
	}
	handlers.Append(func(activityId any) {
		if !reflect.TypeOf(activityId).AssignableTo(reflect.TypeOf(handler).In(0)) {
			return
		}
		handler(activityId.(ID))
	}, collection.FindFirstOrDefaultInSlice(priority, 0))
}

// OnUpcomingEvent 即将开始的活动事件
func OnUpcomingEvent[Type, ID generic.Basic](activity *Activity[Type, ID]) {
	handlers, exist := upcomingEventHandlers[activity.t]
	if !exist {
		return
	}
	handlers.RangeValue(func(index int, value func(activityId any)) bool {
		defer func() {
			if err := recover(); err != nil {
				log.Error("OnUpcomingEvent", log.Any("type", activity.t), log.Any("id", activity.id), log.Any("err", err))
				return
			}
		}()
		value(activity.id)
		return true
	})
}

// RegStartedEvent 注册活动开始事件处理器
func RegStartedEvent[Type, ID generic.Basic](activityType Type, handler StartedEventHandler[ID], priority ...int) {
	handlers, exist := startedEventHandlers[activityType]
	if !exist {
		handlers = listings2.NewPrioritySlice[func(activityId any)]()
		startedEventHandlers[activityType] = handlers
	}
	handlers.Append(func(activityId any) {
		if !reflect.TypeOf(activityId).AssignableTo(reflect.TypeOf(handler).In(0)) {
			return
		}
		handler(activityId.(ID))
	}, collection.FindFirstOrDefaultInSlice(priority, 0))
}

// OnStartedEvent 活动开始事件
func OnStartedEvent[Type, ID generic.Basic](activity *Activity[Type, ID]) {
	handlers, exist := startedEventHandlers[activity.t]
	if !exist {
		return
	}
	handlers.RangeValue(func(index int, value func(activityId any)) bool {
		defer func() {
			if err := recover(); err != nil {
				log.Error("OnStartedEvent", log.Any("type", activity.t), log.Any("id", activity.id), log.Any("err", err))
				return
			}
		}()
		value(activity.id)
		return true
	})

	now := time.Now()
	if !times.IsSameDay(now, activity.getLastNewDayTime()) {
		OnNewDayEvent(activity)
	}
	ticker.Loop(fmt.Sprintf("activity:new_day:%d:%v", reflect.ValueOf(activity.t).Kind(), activity.id), times.GetNextDayInterval(now), times.Day, timer.Forever, func() {
		OnNewDayEvent(activity)
	})
}

// RegEndedEvent 注册活动结束事件处理器
func RegEndedEvent[Type, ID generic.Basic](activityType Type, handler EndedEventHandler[ID], priority ...int) {
	handlers, exist := endedEventHandlers[activityType]
	if !exist {
		handlers = listings2.NewPrioritySlice[func(activityId any)]()
		endedEventHandlers[activityType] = handlers
	}
	handlers.Append(func(activityId any) {
		if !reflect.TypeOf(activityId).AssignableTo(reflect.TypeOf(handler).In(0)) {
			return
		}
		handler(activityId.(ID))
	}, collection.FindFirstOrDefaultInSlice(priority, 0))
}

// OnEndedEvent 活动结束事件
func OnEndedEvent[Type, ID generic.Basic](activity *Activity[Type, ID]) {
	handlers, exist := endedEventHandlers[activity.t]
	if !exist {
		return
	}
	handlers.RangeValue(func(index int, value func(activityId any)) bool {
		defer func() {
			if err := recover(); err != nil {
				log.Error("OnEndedEvent", log.Any("type", activity.t), log.Any("id", activity.id), log.Any("err", err))
				return
			}
		}()
		value(activity.id)
		return true
	})
}

// RegExtendedShowStartedEvent 注册活动结束后延长展示开始事件处理器
func RegExtendedShowStartedEvent[Type, ID generic.Basic](activityType Type, handler ExtendedShowStartedEventHandler[ID], priority ...int) {
	handlers, exist := extShowStartedEventHandlers[activityType]
	if !exist {
		handlers = listings2.NewPrioritySlice[func(activityId any)]()
		extShowStartedEventHandlers[activityType] = handlers
	}
	handlers.Append(func(activityId any) {
		if !reflect.TypeOf(activityId).AssignableTo(reflect.TypeOf(handler).In(0)) {
			return
		}
		handler(activityId.(ID))
	}, collection.FindFirstOrDefaultInSlice(priority, 0))
}

// OnExtendedShowStartedEvent 活动结束后延长展示开始事件
func OnExtendedShowStartedEvent[Type, ID generic.Basic](activity *Activity[Type, ID]) {
	handlers, exist := extShowStartedEventHandlers[activity.t]
	if !exist {
		return
	}
	handlers.RangeValue(func(index int, value func(activityId any)) bool {
		defer func() {
			if err := recover(); err != nil {
				log.Error("OnExtendedShowStartedEvent", log.Any("type", activity.t), log.Any("id", activity.id), log.Any("err", err))
				return
			}
		}()
		value(activity.id)
		return true
	})
}

// RegExtendedShowEndedEvent 注册活动结束后延长展示结束事件处理器
func RegExtendedShowEndedEvent[Type, ID generic.Basic](activityType Type, handler ExtendedShowEndedEventHandler[ID], priority ...int) {
	handlers, exist := extShowEndedEventHandlers[activityType]
	if !exist {
		handlers = listings2.NewPrioritySlice[func(activityId any)]()
		extShowEndedEventHandlers[activityType] = handlers
	}
	handlers.Append(func(activityId any) {
		if !reflect.TypeOf(activityId).AssignableTo(reflect.TypeOf(handler).In(0)) {
			return
		}
		handler(activityId.(ID))
	}, collection.FindFirstOrDefaultInSlice(priority, 0))
}

// OnExtendedShowEndedEvent 活动结束后延长展示结束事件
func OnExtendedShowEndedEvent[Type, ID generic.Basic](activity *Activity[Type, ID]) {
	handlers, exist := extShowEndedEventHandlers[activity.t]
	if !exist {
		return
	}
	handlers.RangeValue(func(index int, value func(activityId any)) bool {
		defer func() {
			if err := recover(); err != nil {
				log.Error("OnExtendedShowEndedEvent", log.Any("type", activity.t), log.Any("id", activity.id), log.Any("err", err))
				return
			}
		}()
		value(activity.id)
		return true
	})
}

// RegNewDayEvent 注册新的一天事件处理器
func RegNewDayEvent[Type, ID generic.Basic](activityType Type, handler NewDayEventHandler[ID], priority ...int) {
	handlers, exist := newDayEventHandlers[activityType]
	if !exist {
		handlers = listings2.NewPrioritySlice[func(activityId any)]()
		newDayEventHandlers[activityType] = handlers
	}
	handlers.Append(func(activityId any) {
		if !reflect.TypeOf(activityId).AssignableTo(reflect.TypeOf(handler).In(0)) {
			return
		}
		handler(activityId.(ID))
	}, collection.FindFirstOrDefaultInSlice(priority, 0))
}

// OnNewDayEvent 新的一天事件
func OnNewDayEvent[Type, ID generic.Basic](activity *Activity[Type, ID]) {
	handlers, exist := newDayEventHandlers[activity.t]
	if !exist {
		return
	}
	handlers.RangeValue(func(index int, value func(activityId any)) bool {
		defer func() {
			if err := recover(); err != nil {
				log.Error("OnNewDayEvent", log.Any("type", activity.t), log.Any("id", activity.id), log.Any("err", err))
				return
			}
		}()
		value(activity.id)
		return true
	})
}
