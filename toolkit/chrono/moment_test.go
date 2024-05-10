package chrono_test

import (
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"testing"
	"time"
)

func TestGetNextMoment(t *testing.T) {
	var cases = []struct {
		name string
		now  time.Time
		hour int
		min  int
		sec  int
		want time.Time
	}{
		{
			name: "today 00:00:00, next should be today 12:00:00",
			now:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
			hour: 12, min: 0, sec: 0,
			want: time.Date(2021, 1, 1, 12, 0, 0, 0, time.Local),
		},
		{
			name: "today 12:00:00, next should be tomorrow 12:00:00",
			now:  time.Date(2021, 1, 1, 12, 0, 0, 0, time.Local),
			hour: 12, min: 0, sec: 0,
			want: time.Date(2021, 1, 2, 12, 0, 0, 0, time.Local),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := chrono.GetNextMoment(c.now, c.hour, c.min, c.sec)
			if got != c.want {
				t.Errorf("GetNextMoment(%v, %d, %d, %d) = %v, want %v", c.now, c.hour, c.min, c.sec, got, c.want)
			}
		})
	}
}
