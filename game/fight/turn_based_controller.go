package fight

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"time"
)

type TurnBasedControllerInfo[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] interface {
	// GetRound 获取当前回合数
	GetRound() int
	// GetCamp 获取当前操作阵营
	GetCamp() Camp
	// GetEntity 获取当前操作实体
	GetEntity() Entity
	// GetActionTimeoutDuration 获取当前行动超时时长
	GetActionTimeoutDuration() time.Duration
	// GetActionStartTime 获取当前行动开始时间
	GetActionStartTime() time.Time
	// GetActionEndTime 获取当前行动结束时间
	GetActionEndTime() time.Time
	// Stop 在当前回合执行完毕后停止回合进程
	Stop()
}

type TurnBasedControllerAction[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] interface {
	TurnBasedControllerInfo[CampID, EntityID, Camp, Entity]
	// Finish 结束当前操作，将立即切换到下一个操作实体
	Finish()
}

// TurnBasedController 回合制控制器
type TurnBasedController[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct {
	tb *TurnBased[CampID, EntityID, Camp, Entity]
}

// GetRound 获取当前回合数
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) GetRound() int {
	return slf.tb.round
}

// GetCamp 获取当前操作阵营
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) GetCamp() Camp {
	return slf.tb.currCamp
}

// GetEntity 获取当前操作实体
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) GetEntity() Entity {
	return slf.tb.currEntity
}

// GetActionTimeoutDuration 获取当前行动超时时长
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) GetActionTimeoutDuration() time.Duration {
	return slf.tb.currActionTimeout
}

// GetActionStartTime 获取当前行动开始时间
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) GetActionStartTime() time.Time {
	return slf.tb.currStart
}

// GetActionEndTime 获取当前行动结束时间
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) GetActionEndTime() time.Time {
	return slf.tb.currStart.Add(slf.tb.currActionTimeout)
}

// Finish 结束当前操作，将立即切换到下一个操作实体
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) Finish() {
	slf.tb.actionMutex.Lock()
	defer slf.tb.actionMutex.Unlock()
	if slf.tb.actioning {
		slf.tb.actioning = false
		slf.tb.signal <- signalFinish
	}
}

// Stop 在当前回合执行完毕后停止回合进程
func (slf *TurnBasedController[CampID, EntityID, Camp, Entity]) Stop() {
	slf.tb.Close()
}
