package chrono_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"testing"
	"time"
)

func TestNewPeriodWindow(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	window := chrono.NewPeriodWindow(cur, chrono.Day)
	fmt.Println(window)
}
