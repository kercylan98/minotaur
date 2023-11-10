package times_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/times"
	"testing"
	"time"
)

func TestSetGlobalTimeOffset(t *testing.T) {
	fmt.Println(time.Now())
	times.SetGlobalTimeOffset(-times.Hour)
	fmt.Println(time.Now())
	times.SetGlobalTimeOffset(times.Hour)
	fmt.Println(time.Now())
	fmt.Println(times.NowByNotOffset())
}
