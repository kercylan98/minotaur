package fight

import "github.com/kercylan98/minotaur/utils/generic"

type (
	TurnBasedEntitySwitchEventHandler[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]]        func(controller TurnBasedControllerAction[CampID, EntityID, Camp, Entity])
	TurnBasedEntityActionTimeoutEventHandler[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] func(controller TurnBasedControllerInfo[CampID, EntityID, Camp, Entity])
	TurnBasedEntityActionFinishEventHandler[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]]  func(controller TurnBasedControllerInfo[CampID, EntityID, Camp, Entity])
	TurnBasedEntityActionSubmitEventHandler[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]]  func(controller TurnBasedControllerInfo[CampID, EntityID, Camp, Entity])
)

type turnBasedEvents[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct {
	entitySwitchEventHandlers  []TurnBasedEntitySwitchEventHandler[CampID, EntityID, Camp, Entity]
	actionTimeoutEventHandlers []TurnBasedEntityActionTimeoutEventHandler[CampID, EntityID, Camp, Entity]
	actionFinishEventHandlers  []TurnBasedEntityActionFinishEventHandler[CampID, EntityID, Camp, Entity]
	actionSubmitEventHandlers  []TurnBasedEntityActionSubmitEventHandler[CampID, EntityID, Camp, Entity]
}

// RegTurnBasedEntitySwitchEvent 注册回合制实体切换事件处理函数，该处理函数将在切换到实体切换为操作时机时触发
//   - 刚函数通常仅用于告知当前操作实体已经完成切换，适合做一些前置校验，但不应该在该函数中执行长时间阻塞操作
//   - 操作计时将在该函数执行完毕后开始
//
// 场景：
//   - 回合开始，如果该实体被标记为已死亡，则跳过该实体
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) RegTurnBasedEntitySwitchEvent(handler TurnBasedEntitySwitchEventHandler[CampID, EntityID, Camp, Entity]) {
	slf.entitySwitchEventHandlers = append(slf.entitySwitchEventHandlers, handler)
}

// OnTurnBasedEntitySwitchEvent 触发回合制实体切换事件
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) OnTurnBasedEntitySwitchEvent(controller TurnBasedControllerAction[CampID, EntityID, Camp, Entity]) {
	for _, handler := range slf.entitySwitchEventHandlers {
		handler(controller)
	}
}

// RegTurnBasedEntityActionTimeoutEvent 注册回合制实体行动超时事件处理函数，该处理函数将在实体行动超时时触发
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) RegTurnBasedEntityActionTimeoutEvent(handler TurnBasedEntityActionTimeoutEventHandler[CampID, EntityID, Camp, Entity]) {
	slf.actionTimeoutEventHandlers = append(slf.actionTimeoutEventHandlers, handler)
}

// OnTurnBasedEntityActionTimeoutEvent 触发回合制实体行动超时事件
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) OnTurnBasedEntityActionTimeoutEvent(controller TurnBasedControllerInfo[CampID, EntityID, Camp, Entity]) {
	for _, handler := range slf.actionTimeoutEventHandlers {
		handler(controller)
	}
}

// RegTurnBasedEntityActionFinishEvent 注册回合制实体行动结束事件处理函数，该处理函数将在实体行动结束时触发
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) RegTurnBasedEntityActionFinishEvent(handler TurnBasedEntityActionFinishEventHandler[CampID, EntityID, Camp, Entity]) {
	slf.actionFinishEventHandlers = append(slf.actionFinishEventHandlers, handler)
}

// OnTurnBasedEntityActionFinishEvent 触发回合制实体行动结束事件
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) OnTurnBasedEntityActionFinishEvent(controller TurnBasedControllerInfo[CampID, EntityID, Camp, Entity]) {
	for _, handler := range slf.actionFinishEventHandlers {
		handler(controller)
	}
}

// RegTurnBasedEntityActionSubmitEvent 注册回合制实体行动提交事件处理函数，该处理函数将在实体行动提交时触发
//   - 该事件将在实体以任意方式结束行动时触发，包括正常结束、超时结束等
//   - 该事件会在原本的行动结束事件之后触发
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) RegTurnBasedEntityActionSubmitEvent(handler TurnBasedEntityActionSubmitEventHandler[CampID, EntityID, Camp, Entity]) {
	slf.actionSubmitEventHandlers = append(slf.actionSubmitEventHandlers, handler)
}

// OnTurnBasedEntityActionSubmitEvent 触发回合制实体行动提交事件
func (slf *turnBasedEvents[CampID, EntityID, Camp, Entity]) OnTurnBasedEntityActionSubmitEvent(controller TurnBasedControllerInfo[CampID, EntityID, Camp, Entity]) {
	for _, handler := range slf.actionSubmitEventHandlers {
		handler(controller)
	}
}
