package times

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	IntervalYear = iota
	IntervalDay
	IntervalHour
	IntervalMinute
	IntervalSecond
	IntervalNow
)

var (
	intervalFormat = map[int]string{
		IntervalYear:   "年前",
		IntervalDay:    "天前",
		IntervalHour:   "小时前",
		IntervalMinute: "分钟前",
		IntervalSecond: "秒钟前",
		IntervalNow:    "刚刚",
	}
)

// IntervalFormatSet 针对 IntervalFormat 函数设置格式化内容
func IntervalFormatSet(intervalType int, str string) {
	if intervalType < IntervalYear || intervalType > IntervalSecond {
		return
	}
	intervalFormat[intervalType] = str
}

// IntervalFormat 返回指定时间戳之间的间隔
//   - 使用传入的时间进行计算换算，将结果体现为几年前、几天前、几小时前、几分钟前、几秒前。
func IntervalFormat(time1, time2 time.Time) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	cur := time1.Unix()
	ct := cur - time2.Unix()
	if ct <= 0 {
		return intervalFormat[IntervalNow]
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = fmt.Sprint(tempStr, intervalFormat[i])
		}
		break
	}
	return res
}
