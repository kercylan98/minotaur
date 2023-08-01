package fight

import (
	"fmt"
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
			fmt.Println(campId, entity)
			if campId == 1 && entity == 2 {
				round.SetCurrent(1, 1)
			}
		}),
	)

	wait.Add(1)
	r.Run()
	wait.Wait()
}
