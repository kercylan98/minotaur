package times

import "time"

// NewPeriod 创建一个时间段
//   - 如果 start 比 end 晚，则会自动交换两个时间
func NewPeriod(start, end time.Time) Period {
	if start.After(end) {
		start, end = end, start
	}
	return Period{start, end}
}

// NewPeriodWithTimeArray 创建一个时间段
func NewPeriodWithTimeArray(times [2]time.Time) Period {
	return NewPeriod(times[0], times[1])
}

// NewPeriodWithDayZero 创建一个时间段，从 t 开始，持续到 day 天后的 0 点
func NewPeriodWithDayZero(t time.Time, day int) Period {
	return NewPeriod(t, GetToday(t.AddDate(0, 0, day)))
}

// NewPeriodWithDay 创建一个时间段，从 t 开始，持续 day 天
func NewPeriodWithDay(t time.Time, day int) Period {
	return NewPeriod(t, t.AddDate(0, 0, day))
}

// NewPeriodWithHour 创建一个时间段，从 t 开始，持续 hour 小时
func NewPeriodWithHour(t time.Time, hour int) Period {
	return NewPeriod(t, t.Add(time.Duration(hour)*time.Hour))
}

// NewPeriodWithMinute 创建一个时间段，从 t 开始，持续 minute 分钟
func NewPeriodWithMinute(t time.Time, minute int) Period {
	return NewPeriod(t, t.Add(time.Duration(minute)*time.Minute))
}

// NewPeriodWithSecond 创建一个时间段，从 t 开始，持续 second 秒
func NewPeriodWithSecond(t time.Time, second int) Period {
	return NewPeriod(t, t.Add(time.Duration(second)*time.Second))
}

// NewPeriodWithMillisecond 创建一个时间段，从 t 开始，持续 millisecond 毫秒
func NewPeriodWithMillisecond(t time.Time, millisecond int) Period {
	return NewPeriod(t, t.Add(time.Duration(millisecond)*time.Millisecond))
}

// NewPeriodWithMicrosecond 创建一个时间段，从 t 开始，持续 microsecond 微秒
func NewPeriodWithMicrosecond(t time.Time, microsecond int) Period {
	return NewPeriod(t, t.Add(time.Duration(microsecond)*time.Microsecond))
}

// NewPeriodWithNanosecond 创建一个时间段，从 t 开始，持续 nanosecond 纳秒
func NewPeriodWithNanosecond(t time.Time, nanosecond int) Period {
	return NewPeriod(t, t.Add(time.Duration(nanosecond)*time.Nanosecond))
}

// Period 表示一个时间段
type Period [2]time.Time

// Start 返回时间段的开始时间
func (slf Period) Start() time.Time {
	return slf[0]
}

// End 返回时间段的结束时间
func (slf Period) End() time.Time {
	return slf[1]
}

// Duration 返回时间段的持续时间
func (slf Period) Duration() time.Duration {
	return slf[1].Sub(slf[0])
}

// Day 返回时间段的持续天数
func (slf Period) Day() int {
	return int(slf.Duration().Hours() / 24)
}

// Hour 返回时间段的持续小时数
func (slf Period) Hour() int {
	return int(slf.Duration().Hours())
}

// Minute 返回时间段的持续分钟数
func (slf Period) Minute() int {
	return int(slf.Duration().Minutes())
}

// Seconds 返回时间段的持续秒数
func (slf Period) Seconds() int {
	return int(slf.Duration().Seconds())
}

// Milliseconds 返回时间段的持续毫秒数
func (slf Period) Milliseconds() int {
	return int(slf.Duration().Milliseconds())
}

// Microseconds 返回时间段的持续微秒数
func (slf Period) Microseconds() int {
	return int(slf.Duration().Microseconds())
}

// Nanoseconds 返回时间段的持续纳秒数
func (slf Period) Nanoseconds() int {
	return int(slf.Duration().Nanoseconds())
}

// IsZero 判断时间段是否为零值
func (slf Period) IsZero() bool {
	return slf[0].IsZero() && slf[1].IsZero()
}

// IsInvalid 判断时间段是否无效
func (slf Period) IsInvalid() bool {
	return slf[0].IsZero() || slf[1].IsZero()
}

// IsBefore 判断时间段是否在指定时间之前
func (slf Period) IsBefore(t time.Time) bool {
	return slf[1].Before(t)
}

// IsAfter 判断时间段是否在指定时间之后
func (slf Period) IsAfter(t time.Time) bool {
	return slf[0].After(t)
}

// IsBetween 判断指定时间是否在时间段之间
func (slf Period) IsBetween(t time.Time) bool {
	return slf[0].Before(t) && slf[1].After(t)
}

// IsOngoing 判断指定时间是否正在进行时
//   - 如果时间段的开始时间在指定时间之前或者等于指定时间，且时间段的结束时间在指定时间之后，则返回 true
func (slf Period) IsOngoing(t time.Time) bool {
	return (slf[0].Before(t) || slf[0].Equal(t)) && slf[1].After(t)
}

// IsBetweenOrEqual 判断指定时间是否在时间段之间或者等于时间段的开始或结束时间
func (slf Period) IsBetweenOrEqual(t time.Time) bool {
	return slf.IsBetween(t) || slf[0].Equal(t) || slf[1].Equal(t)
}

// IsBetweenOrEqualPeriod 判断指定时间是否在时间段之间或者等于时间段的开始或结束时间
func (slf Period) IsBetweenOrEqualPeriod(t Period) bool {
	return slf.IsBetween(t[0]) || slf.IsBetween(t[1]) || slf[0].Equal(t[0]) || slf[1].Equal(t[1])
}

// IsOverlap 判断时间段是否与指定时间段重叠
func (slf Period) IsOverlap(t Period) bool {
	return slf.IsBetweenOrEqualPeriod(t) || t.IsBetweenOrEqualPeriod(slf)
}
