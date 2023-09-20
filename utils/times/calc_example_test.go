package times_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/times"
	"time"
)

func ExampleCalcNextSecWithTime() {
	now := time.Date(2023, 9, 20, 0, 0, 3, 0, time.Local)
	fmt.Println(times.CalcNextSecWithTime(now, 10))

	// Output:
	// 7
}
