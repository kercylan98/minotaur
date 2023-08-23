package survey_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/log/survey"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

func TestRecord(t *testing.T) {
	_ = os.MkdirAll("./test", os.ModePerm)
	survey.Reg("GLOBAL_DATA", "./test/global_data.log")
	now := time.Now()
	//for i := 0; i < 100000000; i++ {
	//	survey.Record("GLOBAL_DATA", map[string]any{
	//		"joinTime": time.Now().Unix(),
	//		"action":   random.Int64(1, 999),
	//	})
	//	// 每500w flush一次
	//	if i%5000000 == 0 {
	//		survey.Flush()
	//	}
	//}
	//survey.Flush()
	//
	var i atomic.Int64
	survey.All("GLOBAL_DATA", time.Now(), func(record survey.R) bool {
		i.Add(record.Get("action").Int())
		return true
	})
	fmt.Println("write cost:", time.Since(now), i.Load())
}

// Line: 30000001, time: 1.45s
