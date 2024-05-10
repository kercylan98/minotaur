package chrono_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"testing"
	"time"
)

func TestRegisterCronTask(t *testing.T) {
	chrono.RegisterDayMomentTask("newday", time.Now().Add(time.Minute*-2), 0, 0, 0, 0, func() {
		fmt.Println("newday")
	})
}
