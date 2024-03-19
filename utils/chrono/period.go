package chrono

import (
	"time"
)

// NewPeriod 创建一个时间段
//   - 如果 start 比 end 晚，则会自动交换两个时间
func NewPeriod(start, end time.Time) Period {
	if start.After(end) {
		start, end = end, start
	}
	return Period{start, end}
}

// NewPeriodWindow 创建一个特定长度的时间窗口
func NewPeriodWindow(t time.Time, size time.Duration) Period {
	start := t.Truncate(size)
	end := start.Add(size)
	return Period{start, end}
}

// NewPeriodWindowWeek 创建一周长度的时间窗口，从周一零点开始至周日 23:59:59 结束
func NewPeriodWindowWeek(t time.Time) Period {
	var start = GetStartOfWeek(t, time.Monday)
	end := start.Add(Week)
	return Period{start, end}
}

// NewPeriodWithTimeArray 创建一个时间段
func NewPeriodWithTimeArray(times [2]time.Time) Period {
	return NewPeriod(times[0], times[1])
}

// NewPeriodWithDayZero 创建一个时间段，从 t 开始，持续到 day 天后的 0 点
func NewPeriodWithDayZero(t time.Time, day int) Period {
	return NewPeriod(t, GetStartOfDay(t.AddDate(0, 0, day)))
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
func (p Period) Start() time.Time {
	return p[0]
}

// End 返回时间段的结束时间
func (p Period) End() time.Time {
	return p[1]
}

// Duration 返回时间段的持续时间
func (p Period) Duration() time.Duration {
	return p[1].Sub(p[0])
}

// Days 返回时间段的持续天数
func (p Period) Days() int {
	return int(p.Duration().Hours() / 24)
}

// Hours 返回时间段的持续小时数
func (p Period) Hours() int {
	return int(p.Duration().Hours())
}

// Minutes 返回时间段的持续分钟数
func (p Period) Minutes() int {
	return int(p.Duration().Minutes())
}

// Seconds 返回时间段的持续秒数
func (p Period) Seconds() int {
	return int(p.Duration().Seconds())
}

// Milliseconds 返回时间段的持续毫秒数
func (p Period) Milliseconds() int {
	return int(p.Duration().Milliseconds())
}

// Microseconds 返回时间段的持续微秒数
func (p Period) Microseconds() int {
	return int(p.Duration().Microseconds())
}

// Nanoseconds 返回时间段的持续纳秒数
func (p Period) Nanoseconds() int {
	return int(p.Duration().Nanoseconds())
}

// IsZero 判断时间段是否为零值
func (p Period) IsZero() bool {
	return p[0].IsZero() && p[1].IsZero()
}

// IsInvalid 判断时间段是否无效
func (p Period) IsInvalid() bool {
	return p[0].IsZero() || p[1].IsZero()
}

// IsBefore 判断时间段是否在指定时间之前
func (p Period) IsBefore(t time.Time) bool {
	return p[1].Before(t)
}

// IsAfter 判断时间段是否在指定时间之后
func (p Period) IsAfter(t time.Time) bool {
	return p[0].After(t)
}

// IsBetween 判断指定时间是否在时间段之间
func (p Period) IsBetween(t time.Time) bool {
	return p[0].Before(t) && p[1].After(t)
}

// IsOngoing 判断指定时间是否正在进行时
//   - 如果时间段的开始时间在指定时间之前或者等于指定时间，且时间段的结束时间在指定时间之后，则返回 true
func (p Period) IsOngoing(t time.Time) bool {
	return (p[0].Before(t) || p[0].Equal(t)) && p[1].After(t)
}

// IsBetweenOrEqual 判断指定时间是否在时间段之间或者等于时间段的开始或结束时间
func (p Period) IsBetweenOrEqual(t time.Time) bool {
	return p.IsBetween(t) || p[0].Equal(t) || p[1].Equal(t)
}

// IsBetweenOrEqualPeriod 判断指定时间是否在时间段之间或者等于时间段的开始或结束时间
func (p Period) IsBetweenOrEqualPeriod(t Period) bool {
	return p.IsBetween(t[0]) || p.IsBetween(t[1]) || p[0].Equal(t[0]) || p[1].Equal(t[1])
}

// IsOverlap 判断时间段是否与指定时间段重叠
func (p Period) IsOverlap(t Period) bool {
	return p.IsBetweenOrEqualPeriod(t) || t.IsBetweenOrEqualPeriod(p)
}
