package game

type GameplayOver interface {
	// GameOver 游戏玩法结束
	GameOver()

	// RegGameplayOverEvent 游戏玩法结束时将立即调用被注册的事件处理函数
	RegGameplayOverEvent(handle GameplayOverEventHandle)
	OnGameplayOverEvent()
}

type (
	GameplayOverEventHandle func()
)
