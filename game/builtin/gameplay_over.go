package builtin

import (
	"minotaur/game"
)

func NewGameplayOver() *GameplayOver {
	return &GameplayOver{}
}

type GameplayOver struct {
	gameplayOverEventHandles []game.GameplayOverEventHandle
}

func (slf *GameplayOver) GameOver() {
	slf.OnGameplayOverEvent()
}

func (slf *GameplayOver) RegGameplayOverEvent(handle game.GameplayOverEventHandle) {
	slf.gameplayOverEventHandles = append(slf.gameplayOverEventHandles, handle)
}

func (slf *GameplayOver) OnGameplayOverEvent() {
	for _, handle := range slf.gameplayOverEventHandles {
		handle()
	}
}
