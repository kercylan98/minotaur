package counter_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/counter"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestCounter_Add(t *testing.T) {
	c := counter.NewCounter[string, int64]()

	c.Add("login_count", 1, times.GetNextDayInterval(time.Now()))
	c.Add("login_count", 1, times.GetNextDayInterval(time.Now()))
	c.SubCounter("live").Add("login_count", 1, times.GetNextDayInterval(time.Now()))

	fmt.Println(c)
}
