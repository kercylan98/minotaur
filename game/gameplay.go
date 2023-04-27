package game

import (
	"minotaur/utils/offset"
	"time"
)

// Gameplay 游戏玩法
type Gameplay interface {
	// GameStart 游戏玩法开始
	GameStart(handle func() error) error
	// GetTime 获取游戏玩法的时间偏移
	GetTime() *offset.Time
	// GetCurrentTime 获取玩法当前存在偏移的时间
	GetCurrentTime() time.Time
	// SetTimeOffset 设置玩法时间偏移
	SetTimeOffset(offset time.Duration)
	// Release 释放游戏玩法资源
	Release()

	// RegGameplayStartEvent 在游戏玩法开始时将立即执行被注册的事件处理函数
	RegGameplayStartEvent(handle GameplayStartEventHandle)
	OnGameplayStartEvent()
	// RegGameplayTimeChangeEvent 游戏玩法的时间被改变（非自然流逝）时将立刻执行被注册的事件处理函数
	RegGameplayTimeChangeEvent(handle GameplayTimeChangeEventHandle)
	OnGameplayTimeChangeEvent()
	// RegGameplayReleaseEvent 游戏玩法资源被释放前将立即执行被注册的事件处理函数
	RegGameplayReleaseEvent(handle GameplayReleaseEventHandle)
	OnGameplayReleaseEvent()
}

type (
	GameplayStartEventHandle      func(startTime time.Time)
	GameplayTimeChangeEventHandle func(current time.Time)
	GameplayReleaseEventHandle    func()
)
