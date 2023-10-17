package fight

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"sync"
	"time"
)

const (
	signalFinish = 1 + iota // 操作结束信号
	signalStop              // 停止回合制信号
)

// NewTurnBased 创建一个新的回合制
//   - calcNextTurnDuration 将返回下一次行动时间间隔，适用于按照速度计算下一次行动时间间隔的情况。当返回 0 时，将使用默认的行动超时时间
func NewTurnBased[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]](calcNextTurnDuration func(Camp, Entity) time.Duration) *TurnBased[CampID, EntityID, Camp, Entity] {
	tb := &TurnBased[CampID, EntityID, Camp, Entity]{
		turnBasedEvents:      &turnBasedEvents[CampID, EntityID, Camp, Entity]{},
		campRel:              make(map[EntityID]Camp),
		calcNextTurnDuration: calcNextTurnDuration,
		actionTimeoutHandler: func(camp Camp, entity Entity) time.Duration {
			return 0
		},
	}
	tb.controller = &TurnBasedController[CampID, EntityID, Camp, Entity]{tb: tb}
	return tb
}

// TurnBased 回合制
type TurnBased[CampID, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct {
	*turnBasedEvents[CampID, EntityID, Camp, Entity]
	controller           *TurnBasedController[CampID, EntityID, Camp, Entity] // 控制器
	ticker               *time.Ticker                                         // 计时器
	actionWaitTicker     *time.Ticker                                         // 行动等待计时器
	actioning            bool                                                 // 是否正在行动
	actionMutex          sync.RWMutex                                         // 行动锁
	entities             []Entity                                             // 所有阵容实体顺序
	campRel              map[EntityID]Camp                                    // 实体与阵营的关系
	calcNextTurnDuration func(Camp, Entity) time.Duration                     // 下一次行动时间间隔
	actionTimeoutHandler func(Camp, Entity) time.Duration                     // 行动超时时间

	signal            chan byte     // 信号
	round             int           // 当前回合数
	currCamp          Camp          // 当前操作阵营
	currEntity        Entity        // 当前操作实体
	currActionTimeout time.Duration // 当前行动超时时间
	currStart         time.Time     // 当前回合开始时间

	closeMutex sync.RWMutex // 关闭锁
	closed     bool
}

// Close 关闭回合制
func (slf *TurnBased[CampID, EntityID, Camp, Entity]) Close() {
	slf.closeMutex.Lock()
	defer slf.closeMutex.Unlock()
	slf.closed = true
}

// AddCamp 添加阵营
func (slf *TurnBased[CampID, EntityID, Camp, Entity]) AddCamp(camp Camp, entity Entity, entities ...Entity) {
	for _, e := range append([]Entity{entity}, entities...) {
		slf.entities = append(slf.entities, e)
		slf.campRel[e.GetId()] = camp
	}
}

// SetActionTimeout 设置行动超时时间处理函数
//   - 默认情况下行动超时时间函数将始终返回 0
func (slf *TurnBased[CampID, EntityID, Camp, Entity]) SetActionTimeout(actionTimeoutHandler func(Camp, Entity) time.Duration) {
	if actionTimeoutHandler == nil {
		panic("actionTimeoutHandler can not be nil")
	}
	slf.actionTimeoutHandler = actionTimeoutHandler
}

// Run 运行
func (slf *TurnBased[CampID, EntityID, Camp, Entity]) Run() {
	slf.round = 1
	slf.signal = make(chan byte, 1)
	var actionDuration = make(map[EntityID]time.Duration)
	var actionSubmit = func() {
		slf.actionMutex.Lock()
		slf.actioning = false
		if slf.actionWaitTicker != nil {
			slf.actionWaitTicker.Stop()
		}
		slf.actionMutex.Unlock()
	}
	for {
		slf.closeMutex.RLock()
		if slf.closed {
			slf.closeMutex.RUnlock()
			break
		}
		slf.closeMutex.RUnlock()

		var minDuration *time.Duration
		var delay time.Duration
		for _, entity := range slf.entities {
			camp := slf.campRel[entity.GetId()]
			next := slf.calcNextTurnDuration(camp, entity)
			accumulate := next + actionDuration[entity.GetId()]
			if minDuration == nil || accumulate < *minDuration {
				minDuration = &accumulate
				slf.currEntity = entity
				slf.currCamp = camp
				delay = next
			}
		}
		if *minDuration == 0 {
			*minDuration = 1 // 防止永远是第一对象行动
		}
		actionDuration[slf.currEntity.GetId()] = *minDuration
		if len(actionDuration) == len(slf.entities) {
			for key := range actionDuration {
				delete(actionDuration, key)
			}
		}

		if delay > 0 {
			if slf.ticker == nil {
				slf.ticker = time.NewTicker(delay)
			} else {
				slf.ticker.Reset(delay)
			}
			<-slf.ticker.C
		}

		// 进入回合操作阶段
		slf.currActionTimeout = slf.actionTimeoutHandler(slf.currCamp, slf.currEntity)
		slf.currStart = time.Now()
		slf.actionMutex.Lock()
		slf.actioning = true
		slf.actionMutex.Unlock()
		slf.OnTurnBasedEntitySwitchEvent(slf.controller)
		if slf.actionWaitTicker == nil {
			slf.actionWaitTicker = time.NewTicker(slf.currActionTimeout)
		} else {
			slf.actionWaitTicker.Reset(slf.currActionTimeout)
		}
	breakListen:
		for {
			select {
			case <-slf.actionWaitTicker.C:
				actionSubmit()
				slf.OnTurnBasedEntityActionTimeoutEvent(slf.controller)
				break breakListen
			case sign := <-slf.signal:
				switch sign {
				case signalFinish:
					actionSubmit()
					slf.OnTurnBasedEntityActionFinishEvent(slf.controller)
					break breakListen
				}
			}
		}
		slf.OnTurnBasedEntityActionSubmitEvent(slf.controller)

		slf.closeMutex.Lock()
		if slf.closed {
			if slf.ticker != nil {
				slf.ticker.Stop()
				slf.ticker = nil
			}
			if slf.actionWaitTicker != nil {
				slf.actionWaitTicker.Stop()
				slf.actionWaitTicker = nil
			}
			if slf.signal != nil {
				close(slf.signal)
				slf.signal = nil
			}
			slf.closeMutex.Unlock()
			break
		}
		slf.closeMutex.Unlock()

		if len(actionDuration) == 0 {
			slf.round++
		}

	}
}
