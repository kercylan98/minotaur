package chrono_test

import (
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"testing"
	"time"
)

func TestNewStateLine(t *testing.T) {
	sl := chrono.NewStateLine(0)
	sl.AddState(1, time.Now())
	sl.AddState(2, time.Now().Add(-chrono.Hour))

	sl.Iterate(func(index int, state int, ts time.Time) bool {
		t.Log(index, state, ts)
		return true
	})

	t.Log(sl.GetStateByTime(time.Now()))
}
