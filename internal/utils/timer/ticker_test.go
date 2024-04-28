package timer_test

import (
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestTicker_Cron(t *testing.T) {
	ticker := timer.GetTicker(10)
	ticker.After("1_sec", time.Second, func() {
		t.Log(time.Now().Format(time.DateTime), "1_sec")
	})

	ticker.Loop("1_sec_loop_3", 0, time.Second, 3, func() {
		t.Log(time.Now().Format(time.DateTime), "1_sec_loop_3")
	})

	ticker.Cron("5_sec_cron", "0/5 * * * * * ?", func() {
		t.Log(time.Now().Format(time.DateTime), "5_sec_cron")
	})

	time.Sleep(times.Week)
}
