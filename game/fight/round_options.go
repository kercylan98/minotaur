package fight

import (
	"github.com/kercylan98/minotaur/utils/timer"
	"time"
)

// RoundOption 回合制游戏选项
type RoundOption[Data RoundData] func(round *Round[Data])

type (
	RoundSwapCampEvent[Data RoundData]      func(round *Round[Data], campId int)
	RoundSwapEntityEvent[Data RoundData]    func(round *Round[Data], campId, entity int)
	RoundGameOverEvent[Data RoundData]      func(round *Round[Data])
	RoundChangeEvent[Data RoundData]        func(round *Round[Data])
	RoundActionTimeoutEvent[Data RoundData] func(round *Round[Data], campId, entity int)
)

// WithRoundTicker 设置游戏的计时器
func WithRoundTicker[Data RoundData](ticker *timer.Ticker) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.ticker = ticker
	}
}

// WithRoundActionTimeout 设置游戏的行动超时时间
func WithRoundActionTimeout[Data RoundData](timeout time.Duration) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.actionTimeout = timeout
	}
}

// WithRoundShareAction 设置游戏的行动是否共享
func WithRoundShareAction[Data RoundData](share bool) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.shareAction = share
	}
}

// WithRoundSwapCampEvent 设置游戏的阵营交换事件
//   - 该事件在触发时已经完成了阵营的交换
func WithRoundSwapCampEvent[Data RoundData](swapCampEventHandle RoundSwapCampEvent[Data]) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.swapCampEventHandles = append(round.swapCampEventHandles, swapCampEventHandle)
	}
}

// WithRoundSwapEntityEvent 设置游戏的实体交换事件
//   - 该事件在触发时已经完成了实体的交换
func WithRoundSwapEntityEvent[Data RoundData](swapEntityEventHandle RoundSwapEntityEvent[Data]) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.swapEntityEventHandles = append(round.swapEntityEventHandles, swapEntityEventHandle)
	}
}

// WithRoundGameOverEvent 设置游戏的结束事件
func WithRoundGameOverEvent[Data RoundData](gameOverEventHandle RoundGameOverEvent[Data]) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.gameOverEventHandles = append(round.gameOverEventHandles, gameOverEventHandle)
	}
}

// WithRoundChangeEvent 设置游戏的回合变更事件
//   - 该事件在触发时已经完成了回合的变更
func WithRoundChangeEvent[Data RoundData](changeEventHandle RoundChangeEvent[Data]) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.changeEventHandles = append(round.changeEventHandles, changeEventHandle)
	}
}

// WithRoundActionTimeoutEvent 设置游戏的超时事件
func WithRoundActionTimeoutEvent[Data RoundData](timeoutEventHandle RoundActionTimeoutEvent[Data]) RoundOption[Data] {
	return func(round *Round[Data]) {
		round.actionTimeoutEventHandles = append(round.actionTimeoutEventHandles, timeoutEventHandle)
	}
}

// WithRoundCampCounterclockwise 设置游戏阵营逆序执行
func WithRoundCampCounterclockwise[Data RoundData]() RoundOption[Data] {
	return func(round *Round[Data]) {
		round.campCounterclockwise = true
	}
}

// WithRoundEntityCounterclockwise 设置游戏实体逆序执行
func WithRoundEntityCounterclockwise[Data RoundData]() RoundOption[Data] {
	return func(round *Round[Data]) {
		round.entityCounterclockwise = true
	}
}
