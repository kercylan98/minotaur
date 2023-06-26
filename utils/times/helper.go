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
