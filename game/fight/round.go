package fight

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/random"
	"github.com/kercylan98/minotaur/utils/slice"
	"github.com/kercylan98/minotaur/utils/timer"
	"time"
)

// RoundGameOverVerifyHandle 回合制游戏结束验证函数
type RoundGameOverVerifyHandle[Data RoundData] func(round *Round[Data]) bool

// NewRound 创建一个新的回合制游戏
//   - data 游戏数据
//   - camps 阵营
//   - roundGameOverVerifyHandle 游戏结束验证函数
//   - options 选项
func NewRound[Data RoundData](data Data, camps []*RoundCamp, roundGameOverVerifyHandle RoundGameOverVerifyHandle[Data], options ...RoundOption[Data]) *Round[Data] {
	mark := random.HostName()
	round := &Round[Data]{
		mark:                      mark,
		data:                      data,
		camps:                     make(map[int][]int),
		actionTimeoutTickerName:   fmt.Sprintf("round_action_timeout_%s", mark),
		roundGameOverVerifyHandle: roundGameOverVerifyHandle,
	}
	for _, camp := range camps {
		round.camps[camp.campId] = camp.entities
		round.campOrder = append(round.campOrder, camp.campId)
	}
	for _, option := range options {
		option(round)
	}
	if round.ticker == nil {
		round.ticker = timer.GetTicker(5, timer.WithMark(mark))
	}
	return round
}

// Round 回合制游戏结构
type Round[Data RoundData] struct {
	mark                      string                          // 标记
	data                      Data                            // 游戏数据
	ticker                    *timer.Ticker                   // 计时器
	camps                     map[int][]int                   // 阵营
	campOrder                 []int                           // 阵营顺序
	actionTimeout             time.Duration                   // 行动超时时间
	actionTimeoutTickerName   string                          // 行动超时计时器名称
	round                     int                             // 回合数
	roundCount                int                             // 回合计数
	currentCamp               int                             // 当前行动阵营
	currentEntity             int                             // 当前行动的阵营实体索引
	currentEndTime            int64                           // 当前行动结束时间
	shareAction               bool                            // 是否共享行动（同阵营共享行动时间）
	roundGameOverVerifyHandle RoundGameOverVerifyHandle[Data] // 游戏结束验证函数
	campCounterclockwise      bool                            // 是否阵营逆时针
	entityCounterclockwise    bool                            // 是否对象逆时针

	swapCampEventHandles      []RoundSwapCampEvent[Data]      // 阵营交换事件
	swapEntityEventHandles    []RoundSwapEntityEvent[Data]    // 实体交换事件
	gameOverEventHandles      []RoundGameOverEvent[Data]      // 游戏结束事件
	changeEventHandles        []RoundChangeEvent[Data]        // 游戏回合变更事件
	actionTimeoutEventHandles []RoundActionTimeoutEvent[Data] // 行动超时事件
}

// GetData 获取游戏数据
func (slf *Round[Data]) GetData() Data {
	return slf.data
}

// Run 运行游戏
//   - 将通过传入的 Camp 进行初始化，Camp 为一个二维数组，每个数组内的元素都是一个行动标识
func (slf *Round[Data]) Run() {
	slf.currentEntity = -1
	slf.round = 1
	slf.loop(false)
}

// Release 释放游戏
func (slf *Round[Data]) Release() {
	if slf.mark == slf.ticker.Mark() {
		slf.ticker.Release()
		return
	}
	slf.ticker.StopTimer(slf.actionTimeoutTickerName)
}

func (slf *Round[Data]) loop(timeout bool) {
	slf.ticker.StopTimer(slf.actionTimeoutTickerName)
	if timeout {
		slf.OnActionTimeoutEvent()
	}
	if slf.roundGameOverVerifyHandle(slf) {
		slf.OnGameOverEvent()
		return
	} else {
		slf.ActionRefresh()
	}
	if slf.currentEntity == -1 || slf.currentEntity >= len(slf.camps[slf.currentCamp])-1 {
		if !slf.campCounterclockwise {
			slf.currentCamp = slf.campOrder[0]
			slf.campOrder = append(slf.campOrder[1:], slf.currentCamp)
		} else {
			slf.currentCamp = slf.campOrder[len(slf.campOrder)-1]
			slf.campOrder = append([]int{slf.currentCamp}, slf.campOrder[:len(slf.campOrder)-1]...)
		}

		slf.currentEntity = -1
		slf.roundCount++
		if slf.roundCount > len(slf.camps) {
			slf.round++
			slf.roundCount = 1
			slf.OnChangeEvent()
		}
		slf.OnSwapCampEvent()
	}
	slf.currentEntity++
	slf.OnSwapEntityEvent()
}

// SetCurrent 设置当前行动对象
func (slf *Round[Data]) SetCurrent(campId int, entityId int) {
	camp, exist := slf.camps[campId]
	if !exist || !slice.Contains(camp, entityId) {
		return
	}
	slf.currentCamp = campId
	slf.currentEntity = slice.GetIndex(camp, entityId)
}

// SkipCamp 跳过当前阵营剩余对象的行动
func (slf *Round[Data]) SkipCamp() {
	slf.currentEntity = -1
}

// ActionRefresh 刷新行动超时时间
func (slf *Round[Data]) ActionRefresh() {
	slf.currentEndTime = time.Now().Unix()
	slf.ticker.After(slf.actionTimeoutTickerName, slf.actionTimeout, slf.loop, true)
}

// ActionFinish 结束行动
func (slf *Round[Data]) ActionFinish() {
	slf.ticker.StopTimer(slf.actionTimeoutTickerName)
	if slf.shareAction {
		slf.currentEntity = -1
	}
	slf.loop(false)
}

// GetRound 获取当前回合数
func (slf *Round[Data]) GetRound() int {
	return slf.round
}

// AllowAction 是否允许行动
func (slf *Round[Data]) AllowAction(camp, entity int) bool {
	return (slf.currentCamp == camp && slf.currentEntity == entity) || slf.shareAction && camp == slf.currentCamp
}

// CampAllowAction 阵容是否允许行动
func (slf *Round[Data]) CampAllowAction(camp int) bool {
	return slf.currentCamp == camp
}

// GetCurrentCamp 获取当前行动的阵营
func (slf *Round[Data]) GetCurrentCamp() int {
	return slf.currentCamp
}

// GetCurrentRoundProgressRate 获取当前回合进度
func (slf *Round[Data]) GetCurrentRoundProgressRate() float64 {
	return float64(slf.roundCount / len(slf.camps))
}

// GetCurrent 获取当前行动的阵营和对象
func (slf *Round[Data]) GetCurrent() (camp, entity int) {
	return slf.currentCamp, slf.camps[slf.currentCamp][slf.currentEntity]
}

// OnSwapCampEvent 触发阵营交换事件
func (slf *Round[Data]) OnSwapCampEvent() {
	for _, handle := range slf.swapCampEventHandles {
		handle(slf, slf.currentCamp)
	}
}

// OnSwapEntityEvent 触发实体交换事件
func (slf *Round[Data]) OnSwapEntityEvent() {
	for _, handle := range slf.swapEntityEventHandles {
		if slf.entityCounterclockwise {
			handle(slf, slf.currentCamp, slf.camps[slf.currentCamp][len(slf.camps[slf.currentCamp])-slf.currentEntity-1])
		} else {
			handle(slf, slf.currentCamp, slf.camps[slf.currentCamp][slf.currentEntity])
		}
	}
}

// OnGameOverEvent 触发游戏结束事件
func (slf *Round[Data]) OnGameOverEvent() {
	for _, handle := range slf.gameOverEventHandles {
		handle(slf)
	}
}

// OnChangeEvent 触发回合变更事件
func (slf *Round[Data]) OnChangeEvent() {
	for _, handle := range slf.changeEventHandles {
		handle(slf)
	}
}

// OnActionTimeoutEvent 触发行动超时事件
func (slf *Round[Data]) OnActionTimeoutEvent() {
	for _, handle := range slf.actionTimeoutEventHandles {
		handle(slf, slf.currentCamp, slf.camps[slf.currentCamp][slf.currentEntity])
	}
}
