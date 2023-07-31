package fight

import (
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var wait sync.WaitGroup
	var camps []*RoundCamp
	camps = append(camps, NewRoundCamp(1, 1, 2, 3))
	camps = append(camps, NewRoundCamp(2, 4, 5, 6))
	camps = append(camps, NewRoundCamp(3, 7, 8, 9))
	r := NewRound("", camps, func(round *Round[string]) bool {
		return round.GetRound() == 2
	},
		WithRoundActionTimeout[string](time.Second),
		WithRoundSwapEntityEvent[string](func(round *Round[string], campId, entity int) {
			t.Log(time.Now(), "swap entity", round.GetRound(), campId, entity)
		}),
		WithRoundGameOverEvent[string](func(round *Round[string]) {
			t.Log(time.Now(), "game over", round.GetRound())
			wait.Done()
		}),
		WithRoundCampCounterclockwise[string](),
		WithRoundEntityCounterclockwise[string](),
	)

	wait.Add(1)
	r.Run()
	wait.Wait()
}
