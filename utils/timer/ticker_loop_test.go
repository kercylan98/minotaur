package timer

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestTicker_Loop(t *testing.T) {
	r := gin.Default()
	pprof.Register(r)

	go func() {
		r.Run(":9999")
	}()

	ticker := GetTicker(10, WithCaller(func(name string, caller func()) {
		caller()
	}))

	ticker.After("stop", time.Second, func() {
		ticker.StopTimer("stop")
		t.Log("success")
		ticker.After("stop1", time.Second, func() {
			ticker.StopTimer("stop1")
			t.Log("success1")
			ticker.After("stop2", time.Second, func() {
				ticker.StopTimer("stop2")
				t.Log("success2")
			})
		})
	})

	time.Sleep(times.Week)
}
