package survey

import (
	"bufio"
	"github.com/kercylan98/minotaur/utils/super"
	"io"
	"os"
	"time"
)

// All 处理特定记录器特定日期的所有记录，当 handle 返回 false 时停止处理
func All(name string, t time.Time, handle func(record map[string]any) bool) {
	logger := survey[name]
	if logger == nil {
		return
	}
	fp := logger.filePath(t.Format(logger.layout))
	logger.wl.Lock()
	defer logger.wl.Unlock()

	f, err := os.Open(fp)
	if err != nil {
		return
	}
	defer func() {
		_ = f.Close()
	}()
	reader := bufio.NewReader(f)
	var m = make(map[string]any)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if err = super.UnmarshalJSON(line[logger.dataLayoutLen:], &m); err != nil {
			panic(err)
		}
		if !handle(m) {
			break
		}
	}
}

// Sum 处理特定记录器特定日期的所有记录，根据指定的字段进行汇总
func Sum(name string, t time.Time, field string) float64 {
	var res float64
	All(name, t, func(record map[string]any) bool {
		v, exist := record[field]
		if !exist {
			return true
		}
		switch value := v.(type) {
		case float64:
			res += value
		case int:
			res += float64(value)
		case int64:
			res += float64(value)
		case string:
			res += super.StringToFloat64(value)
		}
		return true
	})
	return res
}
