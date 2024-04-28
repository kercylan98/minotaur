package times_test

import (
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestNewStateLine(t *testing.T) {
	sl := times.NewStateLine(0)
	sl.AddState(1, time.Now())
	sl.AddState(2, time.Now().Add(-times.Hour))

	sl.Range(func(index int, state int, ts time.Time) bool {
		t.Log(index, state, ts)
		return true
	})

	t.Log(sl.GetStateByTime(time.Now()))
}
