package builtin

import (
	"fmt"
	"minotaur/utils/timer"
)

type GameplayTimeOption func(time *GameplayTime)

func WithGameplayTimeWheelSize(size int) GameplayTimeOption {
	return func(time *GameplayTime) {
		if time.ticker != nil {
			time.ticker.Release()
			time.id = 0
		}
		time.ticker = timer.GetTicker(size)
		time.afterName = gameplayTimeTickerEndTime
	}
}

func WithGameplayTimeTicker(id int64, ticker *timer.Ticker) GameplayTimeOption {
	return func(time *GameplayTime) {
		if time.ticker != nil {
			time.ticker.Release()
		}
		time.id = id
		time.ticker = ticker
		time.afterName = fmt.Sprintf("%s_%d", gameplayTimeTickerEndTime, id)
	}
}
