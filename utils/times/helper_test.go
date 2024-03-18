package times_test

import (
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestGetCurrWeekDate(t *testing.T) {
	now := time.Now()
	date := time.Date(2024, 3, 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)

	for i := 0; i < 31; i++ {
		target := date.AddDate(0, 0, i)
		t.Logf(target.Format(time.DateTime) + " -> " + times.GetWeekdayTimeRelativeToNowWithOffset(target, 6, 1).Format(time.DateTime))
	}
}
