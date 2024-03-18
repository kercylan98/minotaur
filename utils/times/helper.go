package times

import (
	"time"
)

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
	Day         = time.Hour * 24
	Week        = Day * 7
)

// GetWeekdayDateRelativeToNowWithOffset 获取相对于当前日期所在周的指定星期的日期，可以根据需要进行周数的偏移。
//   - now：当前时间。
//   - week：要获取日期的目标星期。
//   - offsetWeeks：要偏移的周数，正数表示向未来偏移，负数表示向过去偏移。
//
// 该函数通常适用于排行榜等场景，例如排行榜为每周六更新，那么通过该函数可以获取到上周排行榜、本周排行榜的准确日期
//
// 该函数将不会保留 now 的时分秒信息，如果需要，可使用 GetWeekdayTimeRelativeToNowWithOffset 函数
func GetWeekdayDateRelativeToNowWithOffset(now time.Time, week time.Weekday, offsetWeeks int) time.Time {
	w := int(week)
	monday := GetMondayZero(now)
	var curr time.Time
	if WeekDay(now) >= w {
		curr = monday.AddDate(0, 0, (w-1)+offsetWeeks*7)
	} else {
		curr = monday.AddDate(0, 0, (w-1)+offsetWeeks*7)
		curr = curr.AddDate(0, 0, -7)
	}
	return curr
}

// GetWeekdayTimeRelativeToNowWithOffset 获取相对于当前日期所在周的指定星期的日期，并根据传入的 now 参数保留时、分、秒等时间信息。
//   - 参数解释可参考 GetWeekdayDateRelativeToNowWithOffset 函数
func GetWeekdayTimeRelativeToNowWithOffset(now time.Time, week time.Weekday, offsetWeeks int) time.Time {
	curr := GetWeekdayDateRelativeToNowWithOffset(now, week, offsetWeeks)
	return time.Date(curr.Year(), curr.Month(), curr.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
}

// GetMonthDays 获取一个时间当月共有多少天
func GetMonthDays(t time.Time) int {
	t = GetToday(t)
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

// WeekDay 获取一个时间是星期几
//   - 1 ~ 7
func WeekDay(t time.Time) int {
	t = GetToday(t)
	week := int(t.Weekday())
	if week == 0 {
		week = 7
	}

	return week
}

// GetNextDayInterval 获取一个时间到下一天间隔多少秒
func GetNextDayInterval(t time.Time) time.Duration {
	return time.Duration(GetToday(t.AddDate(0, 0, 1)).Unix()-t.Unix()) * time.Second
}

// GetToday 获取一个时间的今天
func GetToday(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// GetSecond 获取共有多少秒
func GetSecond(d time.Duration) int {
	return int(d / time.Second)
}

// IsSameDay 两个时间是否是同一天
func IsSameDay(t1, t2 time.Time) bool {
	t1, t2 = GetToday(t1), GetToday(t2)
	return t1.Unix() == t2.Unix()
}

// IsSameHour 两个时间是否是同一小时
func IsSameHour(t1, t2 time.Time) bool {
	return t1.Hour() == t2.Hour() && t1.Day() == t2.Day() && t1.Month() == t2.Month() && t1.Year() == t2.Year()
}

// GetMondayZero 获取本周一零点
func GetMondayZero(t time.Time) time.Time {
	t = GetToday(t)
	weekDay := WeekDay(t)
	return t.AddDate(0, 0, 1-weekDay)
}

// Date 返回一个特定日期的时间
func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// DateWithHMS 返回一个精确到秒的时间
func DateWithHMS(year int, month time.Month, day, hour, min, sec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, 0, time.Local)
}

// GetDeltaDay 获取两个时间需要加减的天数
func GetDeltaDay(t1, t2 time.Time) int {
	return int(GetToday(t1).Sub(GetToday(t2)) / time.Hour * 24)
}

// GetDeltaWeek 获取两个时间需要加减的周数
func GetDeltaWeek(t1, t2 time.Time) int {
	t1, t2 = GetMondayZero(t1), GetMondayZero(t2)
	return GetDeltaDay(t1, t2) / 7
}

// GetHSMFromString 从时间字符串中获取时分秒
func GetHSMFromString(timeStr, layout string) (hour, min, sec int) {
	t, _ := time.ParseInLocation(layout, timeStr, time.Local)
	return t.Hour(), t.Minute(), t.Second()
}

// GetTimeFromString 将时间字符串转化为时间
func GetTimeFromString(timeStr, layout string) time.Time {
	t, _ := time.ParseInLocation(layout, timeStr, time.Local)
	return t
}

// GetDayZero 获取 t 增加 day 天后的零点时间
func GetDayZero(t time.Time, day int) time.Time {
	return GetToday(t.AddDate(0, 0, day))
}

// GetYesterday 获取昨天
func GetYesterday(t time.Time) time.Time {
	return GetDayZero(t, -1)
}

// GetDayLast 获取某天的最后一刻
//   - 最后一刻即 23:59:59
func GetDayLast(t time.Time) time.Time {
	return GetDayZero(t, 1).Add(-time.Second)
}

// GetYesterdayLast 获取昨天最后一刻
func GetYesterdayLast(t time.Time) time.Time {
	return GetDayLast(GetYesterday(t))
}

// GetMinuteStart 获取一个时间的 0 秒
func GetMinuteStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Local)
}

// GetMinuteEnd 获取一个时间的 59 秒
func GetMinuteEnd(t time.Time) time.Time {
	return GetMinuteStart(t).Add(time.Minute - time.Nanosecond)
}

// GetHourStart 获取一个时间的 0 分 0 秒
func GetHourStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local)
}

// GetHourEnd 获取一个时间的 59 分 59 秒
func GetHourEnd(t time.Time) time.Time {
	return GetHourStart(t).Add(time.Hour - time.Nanosecond)
}

// GetMonthStart 获取一个时间的月初
func GetMonthStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

// GetMonthEnd 获取一个时间的月末
func GetMonthEnd(t time.Time) time.Time {
	return GetMonthStart(t).AddDate(0, 1, -1)
}

// GetYearStart 获取一个时间的年初
func GetYearStart(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
}

// GetYearEnd 获取一个时间的年末
func GetYearEnd(t time.Time) time.Time {
	return GetYearStart(t).AddDate(1, 0, -1)
}
