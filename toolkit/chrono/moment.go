package chrono

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"math"
	"time"
)

var zero = time.Time{}

// GetNextMoment 获取下一个指定时刻发生的时间。
func GetNextMoment(now time.Time, hour, min, sec int) time.Time {
	moment := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, time.Local)
	// 如果要检查的时刻已经过了，则返回明天的这个时刻
	if now.After(moment) || now.Equal(moment) {
		moment = moment.AddDate(0, 0, 1)
	}
	return moment
}

// IsMomentPassed 检查指定时刻是否已经过去
func IsMomentPassed(now time.Time, hour, min, sec int) bool {
	moment := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, time.Local)
	return now.After(moment)
}

// IsMomentFuture 检查指定时刻是否在未来
func IsMomentFuture(now time.Time, hour, min, sec int) bool {
	return !IsMomentPassed(now, hour, min, sec)
}

// GetStartOfDay 获取指定时间的当天第一刻，即 00:00:00
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetEndOfDay 获取指定时间的当天最后一刻，即 23:59:59
func GetEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// GetRelativeStartOfDay 获取相对于指定时间减去或加上指定天数后的当天开始时间
//   - offsetDays: 要偏移的天数，负数表示过去的某一天，正数表示未来的某一天
func GetRelativeStartOfDay(t time.Time, offsetDays int) time.Time {
	return GetStartOfDay(GetStartOfDay(t.AddDate(0, 0, offsetDays)))
}

// GetRelativeEndOfDay 获取相对于指定时间减去或加上指定天数后的当天结束时间
//   - offsetDays: 要偏移的天数，负数表示过去的某一天，正数表示未来的某一天
func GetRelativeEndOfDay(t time.Time, offsetDays int) time.Time {
	return GetEndOfDay(GetEndOfDay(t.AddDate(0, 0, offsetDays)))
}

// GetStartOfWeek 获取指定时间所在周的特定周的开始时刻，即 00:00:00
func GetStartOfWeek(t time.Time, weekday time.Weekday) time.Time {
	t = GetStartOfDay(t)
	tw := t.Weekday()
	if tw == 0 {
		tw = 7
	}
	d := 1 - int(tw)
	switch weekday {
	case time.Sunday:
		d += 6
	default:
		d += int(weekday) - 1
	}
	return t.AddDate(0, 0, d)
}

// GetEndOfWeek 获取指定时间所在周的特定周的最后时刻，即 23:59:59
func GetEndOfWeek(t time.Time, weekday time.Weekday) time.Time {
	return GetEndOfDay(GetStartOfWeek(t, weekday))
}

// GetRelativeStartOfWeek  获取相对于当前时间的本周开始时间，以指定的星期作为一周的开始，并根据需要进行周数的偏移
//   - now：当前时间
//   - week：以哪一天作为一周的开始
//   - offsetWeeks：要偏移的周数，正数表示向未来偏移，负数表示向过去偏移
//
// 该函数返回以指定星期作为一周的开始时间，然后根据偏移量进行周数偏移，得到相对于当前时间的周的开始时间
//
// 假设 week 为 time.Saturday 且 offsetWeeks 为 -1，则表示获取上周六的开始时间，一下情况中第一个时间为 now，第二个时间为函数返回值
//   - 2024-03-01 00:00:00 --相对时间-> 2024-02-24 00:00:00 --偏移时间--> 2024-02-17 00:00:00
//   - 2024-03-02 00:00:00 --相对时间-> 2024-02-24 00:00:00 --偏移时间--> 2024-02-17 00:00:00
//   - 2024-03-03 00:00:00 --相对时间-> 2024-03-02 00:00:00 --偏移时间--> 2024-02-24 00:00:00
func GetRelativeStartOfWeek(now time.Time, week time.Weekday, offsetWeeks int) time.Time {
	nowWeekday, weekday := int(now.Weekday()), int(week)
	if nowWeekday == 0 {
		nowWeekday = 7
	}
	if weekday == 0 {
		weekday = 7
	}
	if nowWeekday < weekday {
		now = now.Add(-Week)
	}
	moment := GetStartOfWeek(now, week)
	return moment.Add(Week * time.Duration(offsetWeeks))
}

// GetRelativeEndOfWeek 获取相对于当前时间的本周结束时间，以指定的星期作为一周的开始，并根据需要进行周数的偏移
//   - 该函数详细解释参考 GetRelativeEndOfWeek 函数，其中不同的是，该函数返回的是这一天最后一刻的时间，即 23:59:59
func GetRelativeEndOfWeek(now time.Time, week time.Weekday, offsetWeeks int) time.Time {
	return GetEndOfDay(GetRelativeStartOfWeek(now, week, offsetWeeks))
}

// GetRelativeTimeOfWeek 获取相对于当前时间的本周指定星期的指定时刻，以指定的星期作为一周的开始，并根据需要进行周数的偏移
//   - 该函数详细解释参考 GetRelativeStartOfWeek 函数，其中不同的是，该函数返回的是这一天对应 now 的时间
func GetRelativeTimeOfWeek(now time.Time, week time.Weekday, offsetWeeks int) time.Time {
	moment := GetRelativeStartOfWeek(now, week, offsetWeeks)
	return time.Date(moment.Year(), moment.Month(), moment.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
}

// Zero 获取一个零值的时间
func Zero() time.Time {
	return zero
}

// IsZero 检查一个时间是否为零值
func IsZero(t time.Time) bool {
	return t.IsZero()
}

// Max 获取两个时间中的最大值
func Max(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t1
	}
	return t2
}

// Min 获取两个时间中的最小值
func Min(t1, t2 time.Time) time.Time {
	if t1.Before(t2) {
		return t1
	}
	return t2
}

// SmallerFirst 将两个时间按照从小到大的顺序排列
func SmallerFirst(t1, t2 time.Time) (time.Time, time.Time) {
	if t1.Before(t2) {
		return t1, t2
	}
	return t2, t1
}

// SmallerLast 将两个时间按照从大到小的顺序排列
func SmallerLast(t1, t2 time.Time) (time.Time, time.Time) {
	if t1.Before(t2) {
		return t2, t1
	}
	return t1, t2
}

// Delta 获取两个时间之间的时间差
func Delta(t1, t2 time.Time) time.Duration {
	if t1.Before(t2) {
		return t2.Sub(t1)
	}
	return t1.Sub(t2)
}

// FloorDeltaDays 计算两个时间之间的天数差异，并向下取整
func FloorDeltaDays(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Day)
}

// CeilDeltaDays 计算两个时间之间的天数差异，并向上取整
func CeilDeltaDays(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(math.Ceil(float64(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Day)))
}

// RoundDeltaDays 计算两个时间之间的天数差异，并四舍五入
func RoundDeltaDays(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(math.Round(float64(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Day)))
}

// FloorDeltaHours 计算两个时间之间的小时数差异，并向下取整
func FloorDeltaHours(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Hour)
}

// CeilDeltaHours 计算两个时间之间的小时数差异，并向上取整
func CeilDeltaHours(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(math.Ceil(float64(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Hour)))
}

// RoundDeltaHours 计算两个时间之间的小时数差异，并四舍五入
func RoundDeltaHours(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(math.Round(float64(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Hour)))
}

// FloorDeltaMinutes 计算两个时间之间的分钟数差异，并向下取整
func FloorDeltaMinutes(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Minute)
}

// CeilDeltaMinutes 计算两个时间之间的分钟数差异，并向上取整
func CeilDeltaMinutes(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(math.Ceil(float64(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Minute)))
}

// RoundDeltaMinutes 计算两个时间之间的分钟数差异，并四舍五入
func RoundDeltaMinutes(t1, t2 time.Time) int {
	t1, t2 = SmallerFirst(t1, t2)
	return int(math.Round(float64(GetStartOfDay(t2).Sub(GetStartOfDay(t1)) / Minute)))
}

// IsSameSecond 检查两个时间是否在同一秒
func IsSameSecond(t1, t2 time.Time) bool {
	return t1.Unix() == t2.Unix()
}

// IsSameMinute 检查两个时间是否在同一分钟
func IsSameMinute(t1, t2 time.Time) bool {
	return t1.Minute() == t2.Minute() && IsSameHour(t1, t2)
}

// IsSameHour 检查两个时间是否在同一小时
func IsSameHour(t1, t2 time.Time) bool {
	return t1.Hour() == t2.Hour() && IsSameDay(t1, t2)
}

// IsSameDay 检查两个时间是否在同一天
func IsSameDay(t1, t2 time.Time) bool {
	return GetStartOfDay(t1).Equal(GetStartOfDay(t2))
}

// IsSameWeek 检查两个时间是否在同一周
func IsSameWeek(t1, t2 time.Time) bool {
	return GetStartOfWeek(t1, time.Monday).Equal(GetStartOfWeek(t2, time.Monday))
}

// IsSameMonth 检查两个时间是否在同一月
func IsSameMonth(t1, t2 time.Time) bool {
	return t1.Month() == t2.Month() && t1.Year() == t2.Year()
}

// IsSameYear 检查两个时间是否在同一年
func IsSameYear(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year()
}

// GetMonthDays 获取指定时间所在月的天数
func GetMonthDays(t time.Time) int {
	year, month, _ := t.Date()
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			return 30
		}
		return 31
	}
	if ((year%4 == 0) && (year%100 != 0)) || year%400 == 0 {
		return 29
	}
	return 28
}

// ToDuration 将一个数值转换为 time.Duration 类型，当 unit 为空时，默认为纳秒单位
func ToDuration[V constraints.Number](v V, unit ...time.Duration) time.Duration {
	var u = Nanosecond
	if len(unit) > 0 {
		u = unit[0]
	}
	return time.Duration(v) * u
}

// ToDurationSecond 将一个数值转换为秒的 time.Duration 类型
func ToDurationSecond[V constraints.Number](v V) time.Duration {
	return ToDuration(v, Second)
}

// ToDurationMinute 将一个数值转换为分钟的 time.Duration 类型
func ToDurationMinute[V constraints.Number](v V) time.Duration {
	return ToDuration(v, Minute)
}

// ToDurationHour 将一个数值转换为小时的 time.Duration 类型
func ToDurationHour[V constraints.Number](v V) time.Duration {
	return ToDuration(v, Hour)
}

// ToDurationDay 将一个数值转换为天的 time.Duration 类型
func ToDurationDay[V constraints.Number](v V) time.Duration {
	return ToDuration(v, Day)
}

// ToDurationWeek 将一个数值转换为周的 time.Duration 类型
func ToDurationWeek[V constraints.Number](v V) time.Duration {
	return ToDuration(v, Week)
}
