package builtin

import (
	"minotaur/game"
	"minotaur/utils/offset"
	"time"
)

func NewGameplay() *Gameplay {
	return &Gameplay{
		Time: offset.NewTime(0),
	}
}

type Gameplay struct {
	*offset.Time
	startTime time.Time

	gameplayStartEventHandles      []game.GameplayStartEventHandle
	gameplayTimeChangeEventHandles []game.GameplayTimeChangeEventHandle
	gameplayReleaseEventHandles    []game.GameplayReleaseEventHandle
}

func (slf *Gameplay) GameStart(handle func() error) error {
	if err := handle(); err != nil {
		return err
	}
	slf.startTime = slf.Time.Now()
	slf.OnGameplayStartEvent()
	return nil
}

func (slf *Gameplay) GetTime() *offset.Time {
	return slf.Time
}

func (slf *Gameplay) GetCurrentTime() time.Time {
	return slf.Time.Now()
}

func (slf *Gameplay) SetTimeOffset(offset time.Duration) {
	slf.Time.SetOffset(offset)
	slf.OnGameplayTimeChangeEvent()
}

func (slf *Gameplay) Release() {
	slf.OnGameplayReleaseEvent()
	slf.gameplayStartEventHandles = nil
	slf.gameplayTimeChangeEventHandles = nil
	slf.gameplayReleaseEventHandles = nil
}

func (slf *Gameplay) RegGameplayStartEvent(handle game.GameplayStartEventHandle) {
	slf.gameplayStartEventHandles = append(slf.gameplayStartEventHandles, handle)
}

func (slf *Gameplay) OnGameplayStartEvent() {
	for _, handle := range slf.gameplayStartEventHandles {
		handle(slf.startTime)
	}
}

func (slf *Gameplay) RegGameplayTimeChangeEvent(handle game.GameplayTimeChangeEventHandle) {
	slf.gameplayTimeChangeEventHandles = append(slf.gameplayTimeChangeEventHandles, handle)
}

func (slf *Gameplay) OnGameplayTimeChangeEvent() {
	if len(slf.gameplayTimeChangeEventHandles) == 0 {
		return
	}
	current := slf.Time.Now()
	for _, handle := range slf.gameplayTimeChangeEventHandles {
		handle(current)
	}
}

func (slf *Gameplay) RegGameplayReleaseEvent(handle game.GameplayReleaseEventHandle) {
	slf.gameplayReleaseEventHandles = append(slf.gameplayReleaseEventHandles, handle)
}

func (slf *Gameplay) OnGameplayReleaseEvent() {
	for _, handle := range slf.gameplayReleaseEventHandles {
		handle()
	}
}
