package game

import "time"

// GameplayTime 为游戏玩法添加游戏时长的特性
type GameplayTime interface {
	Gameplay
	GameplayOver
	// GetEndTime 获取游戏结束时间
	GetEndTime() time.Time
	// SetEndTime 设置游戏结束时间
	SetEndTime(t time.Time)
	// ChangeEndTime 通过相对时间的方式改变游戏结束时间
	ChangeEndTime(d time.Duration)
	// Release 释放资源
	Release()
}
