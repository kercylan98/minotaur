package builtin

import (
	"minotaur/game"
	"minotaur/utils/timer"
	"time"
)

const (
	gameplayTimeTickerEndTime = "GameplayTimeTickerEndTime"
)

func NewGameplayTime(gameplay game.Gameplay, gameplayOver game.GameplayOver, options ...GameplayTimeOption) *GameplayTime {
	gameplayTime := &GameplayTime{
		Gameplay:     gameplay,
		GameplayOver: gameplayOver,
	}
	for _, option := range options {
		option(gameplayTime)
	}
	if gameplayTime.ticker == nil {
		gameplayTime.afterName = gameplayTimeTickerEndTime
		gameplayTime.ticker = timer.GetTicker(10)
	}
	return gameplayTime
}

type GameplayTime struct {
	game.Gameplay
	game.GameplayOver
	id        int64
	afterName string
	endTime   time.Time
	ticker    *timer.Ticker
}

func (slf *GameplayTime) GetEndTime() time.Time {
	return slf.endTime
}

func (slf *GameplayTime) SetEndTime(t time.Time) {
	compare := t.Compare(slf.endTime)
	if compare == 0 {
		return
	}

	slf.ticker.StopTimer(slf.afterName)
	current := slf.GetCurrentTime()
	if compare < 0 && t.Compare(current) < 0 {
		slf.GameplayOver.GameOver()
		return
	}

	slf.endTime = t
	slf.ticker.After(slf.afterName, slf.endTime.Sub(current), func() {
		slf.GameplayOver.GameOver()
	})
}

func (slf *GameplayTime) ChangeEndTime(d time.Duration) {
	slf.SetEndTime(slf.endTime.Add(d))
}

func (slf *GameplayTime) Release() {
	slf.ticker.Release()
	slf.ticker = nil
}
