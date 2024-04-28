package times_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestNewPeriodWindow(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	window := times.NewPeriodWindow(cur, times.Day)
	fmt.Println(window)
}
