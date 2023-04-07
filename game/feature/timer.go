package feature

import "minotaur/utils/timer"

// Timer 定时器接口定义
type Timer interface {
	// Timer 获取定时器
	Timer() *timer.Manager
}
