# Times

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

暂无介绍...


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[CalcNextSec](#CalcNextSec)|计算下一个N秒在多少秒之后
|[CalcNextSecWithTime](#CalcNextSecWithTime)|计算下一个N秒在多少秒之后
|[CalcNextTimeWithRefer](#CalcNextTimeWithRefer)|根据参考时间计算下一个整点时间
|[IntervalFormatSet](#IntervalFormatSet)|针对 IntervalFormat 函数设置格式化内容
|[IntervalFormat](#IntervalFormat)|返回指定时间戳之间的间隔
|[GetMonthDays](#GetMonthDays)|获取一个时间当月共有多少天
|[WeekDay](#WeekDay)|获取一个时间是星期几
|[GetNextDayInterval](#GetNextDayInterval)|获取一个时间到下一天间隔多少秒
|[GetToday](#GetToday)|获取一个时间的今天
|[GetSecond](#GetSecond)|获取共有多少秒
|[IsSameDay](#IsSameDay)|两个时间是否是同一天
|[IsSameHour](#IsSameHour)|两个时间是否是同一小时
|[GetMondayZero](#GetMondayZero)|获取本周一零点
|[Date](#Date)|返回一个特定日期的时间
|[DateWithHMS](#DateWithHMS)|返回一个精确到秒的时间
|[GetDeltaDay](#GetDeltaDay)|获取两个时间需要加减的天数
|[GetDeltaWeek](#GetDeltaWeek)|获取两个时间需要加减的周数
|[GetHSMFromString](#GetHSMFromString)|从时间字符串中获取时分秒
|[GetTimeFromString](#GetTimeFromString)|将时间字符串转化为时间
|[GetDayZero](#GetDayZero)|获取 t 增加 day 天后的零点时间
|[GetYesterday](#GetYesterday)|获取昨天
|[GetDayLast](#GetDayLast)|获取某天的最后一刻
|[GetYesterdayLast](#GetYesterdayLast)|获取昨天最后一刻
|[GetMinuteStart](#GetMinuteStart)|获取一个时间的 0 秒
|[GetMinuteEnd](#GetMinuteEnd)|获取一个时间的 59 秒
|[GetHourStart](#GetHourStart)|获取一个时间的 0 分 0 秒
|[GetHourEnd](#GetHourEnd)|获取一个时间的 59 分 59 秒
|[GetMonthStart](#GetMonthStart)|获取一个时间的月初
|[GetMonthEnd](#GetMonthEnd)|获取一个时间的月末
|[GetYearStart](#GetYearStart)|获取一个时间的年初
|[GetYearEnd](#GetYearEnd)|获取一个时间的年末
|[NewStateLine](#NewStateLine)|创建一个从左向右由早到晚的状态时间线
|[SetGlobalTimeOffset](#SetGlobalTimeOffset)|设置全局时间偏移量
|[NowByNotOffset](#NowByNotOffset)|获取未偏移的当前时间
|[GetGlobalTimeOffset](#GetGlobalTimeOffset)|获取全局时间偏移量
|[ResetGlobalTimeOffset](#ResetGlobalTimeOffset)|重置全局时间偏移量
|[NewPeriod](#NewPeriod)|创建一个时间段
|[NewPeriodWindow](#NewPeriodWindow)|创建一个特定长度的时间窗口
|[NewPeriodWindowWeek](#NewPeriodWindowWeek)|创建一周长度的时间窗口，从周一零点开始至周日 23:59:59 结束
|[NewPeriodWithTimeArray](#NewPeriodWithTimeArray)|创建一个时间段
|[NewPeriodWithDayZero](#NewPeriodWithDayZero)|创建一个时间段，从 t 开始，持续到 day 天后的 0 点
|[NewPeriodWithDay](#NewPeriodWithDay)|创建一个时间段，从 t 开始，持续 day 天
|[NewPeriodWithHour](#NewPeriodWithHour)|创建一个时间段，从 t 开始，持续 hour 小时
|[NewPeriodWithMinute](#NewPeriodWithMinute)|创建一个时间段，从 t 开始，持续 minute 分钟
|[NewPeriodWithSecond](#NewPeriodWithSecond)|创建一个时间段，从 t 开始，持续 second 秒
|[NewPeriodWithMillisecond](#NewPeriodWithMillisecond)|创建一个时间段，从 t 开始，持续 millisecond 毫秒
|[NewPeriodWithMicrosecond](#NewPeriodWithMicrosecond)|创建一个时间段，从 t 开始，持续 microsecond 微秒
|[NewPeriodWithNanosecond](#NewPeriodWithNanosecond)|创建一个时间段，从 t 开始，持续 nanosecond 纳秒


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[StateLine](#stateline)|状态时间线
|`STRUCT`|[Period](#period)|表示一个时间段

</details>


***
## 详情信息
#### func CalcNextSec(sec int)  int
<span id="CalcNextSec"></span>
> 计算下一个N秒在多少秒之后

***
#### func CalcNextSecWithTime(t time.Time, sec int)  int
<span id="CalcNextSecWithTime"></span>
> 计算下一个N秒在多少秒之后

示例代码：
```go

func ExampleCalcNextSecWithTime() {
	now := time.Date(2023, 9, 20, 0, 0, 3, 0, time.Local)
	fmt.Println(times.CalcNextSecWithTime(now, 10))
}

```

***
#### func CalcNextTimeWithRefer(now time.Time, refer time.Duration)  time.Time
<span id="CalcNextTimeWithRefer"></span>
> 根据参考时间计算下一个整点时间
>   - 假设当 now 为 14:15:16 ， 参考时间为 10 分钟， 则返回 14:20:00
>   - 假设当 now 为 14:15:16 ， 参考时间为 20 分钟， 则返回 14:20:00
> 
> 当 refer 小于 1 分钟时，将会返回当前时间

***
#### func IntervalFormatSet(intervalType int, str string)
<span id="IntervalFormatSet"></span>
> 针对 IntervalFormat 函数设置格式化内容

***
#### func IntervalFormat(time1 time.Time, time2 time.Time)  string
<span id="IntervalFormat"></span>
> 返回指定时间戳之间的间隔
>   - 使用传入的时间进行计算换算，将结果体现为几年前、几天前、几小时前、几分钟前、几秒前。

***
#### func GetMonthDays(t time.Time)  int
<span id="GetMonthDays"></span>
> 获取一个时间当月共有多少天

***
#### func WeekDay(t time.Time)  int
<span id="WeekDay"></span>
> 获取一个时间是星期几
>   - 1 ~ 7

***
#### func GetNextDayInterval(t time.Time)  time.Duration
<span id="GetNextDayInterval"></span>
> 获取一个时间到下一天间隔多少秒

***
#### func GetToday(t time.Time)  time.Time
<span id="GetToday"></span>
> 获取一个时间的今天

***
#### func GetSecond(d time.Duration)  int
<span id="GetSecond"></span>
> 获取共有多少秒

***
#### func IsSameDay(t1 time.Time, t2 time.Time)  bool
<span id="IsSameDay"></span>
> 两个时间是否是同一天

***
#### func IsSameHour(t1 time.Time, t2 time.Time)  bool
<span id="IsSameHour"></span>
> 两个时间是否是同一小时

***
#### func GetMondayZero(t time.Time)  time.Time
<span id="GetMondayZero"></span>
> 获取本周一零点

***
#### func Date(year int, month time.Month, day int)  time.Time
<span id="Date"></span>
> 返回一个特定日期的时间

***
#### func DateWithHMS(year int, month time.Month, day int, hour int, min int, sec int)  time.Time
<span id="DateWithHMS"></span>
> 返回一个精确到秒的时间

***
#### func GetDeltaDay(t1 time.Time, t2 time.Time)  int
<span id="GetDeltaDay"></span>
> 获取两个时间需要加减的天数

***
#### func GetDeltaWeek(t1 time.Time, t2 time.Time)  int
<span id="GetDeltaWeek"></span>
> 获取两个时间需要加减的周数

***
#### func GetHSMFromString(timeStr string, layout string) (hour int, min int, sec int)
<span id="GetHSMFromString"></span>
> 从时间字符串中获取时分秒

***
#### func GetTimeFromString(timeStr string, layout string)  time.Time
<span id="GetTimeFromString"></span>
> 将时间字符串转化为时间

***
#### func GetDayZero(t time.Time, day int)  time.Time
<span id="GetDayZero"></span>
> 获取 t 增加 day 天后的零点时间

***
#### func GetYesterday(t time.Time)  time.Time
<span id="GetYesterday"></span>
> 获取昨天

***
#### func GetDayLast(t time.Time)  time.Time
<span id="GetDayLast"></span>
> 获取某天的最后一刻
>   - 最后一刻即 23:59:59

***
#### func GetYesterdayLast(t time.Time)  time.Time
<span id="GetYesterdayLast"></span>
> 获取昨天最后一刻

***
#### func GetMinuteStart(t time.Time)  time.Time
<span id="GetMinuteStart"></span>
> 获取一个时间的 0 秒

***
#### func GetMinuteEnd(t time.Time)  time.Time
<span id="GetMinuteEnd"></span>
> 获取一个时间的 59 秒

***
#### func GetHourStart(t time.Time)  time.Time
<span id="GetHourStart"></span>
> 获取一个时间的 0 分 0 秒

***
#### func GetHourEnd(t time.Time)  time.Time
<span id="GetHourEnd"></span>
> 获取一个时间的 59 分 59 秒

***
#### func GetMonthStart(t time.Time)  time.Time
<span id="GetMonthStart"></span>
> 获取一个时间的月初

***
#### func GetMonthEnd(t time.Time)  time.Time
<span id="GetMonthEnd"></span>
> 获取一个时间的月末

***
#### func GetYearStart(t time.Time)  time.Time
<span id="GetYearStart"></span>
> 获取一个时间的年初

***
#### func GetYearEnd(t time.Time)  time.Time
<span id="GetYearEnd"></span>
> 获取一个时间的年末

***
#### func NewStateLine(zero State)  *StateLine[State]
<span id="NewStateLine"></span>
> 创建一个从左向右由早到晚的状态时间线

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewStateLine(t *testing.T) {
	sl := times.NewStateLine(0)
	sl.AddState(1, time.Now())
	sl.AddState(2, time.Now().Add(-times.Hour))
	sl.Range(func(index int, state int, ts time.Time) bool {
		t.Log(index, state, ts)
		return true
	})
	t.Log(sl.GetStateByTime(time.Now()))
}

```


</details>


***
#### func SetGlobalTimeOffset(offset time.Duration)
<span id="SetGlobalTimeOffset"></span>
> 设置全局时间偏移量

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestSetGlobalTimeOffset(t *testing.T) {
	fmt.Println(time.Now())
	times.SetGlobalTimeOffset(-times.Hour)
	fmt.Println(time.Now())
	times.SetGlobalTimeOffset(times.Hour)
	fmt.Println(time.Now())
	fmt.Println(times.NowByNotOffset())
}

```


</details>


***
#### func NowByNotOffset()  time.Time
<span id="NowByNotOffset"></span>
> 获取未偏移的当前时间

***
#### func GetGlobalTimeOffset()  time.Duration
<span id="GetGlobalTimeOffset"></span>
> 获取全局时间偏移量

***
#### func ResetGlobalTimeOffset()
<span id="ResetGlobalTimeOffset"></span>
> 重置全局时间偏移量

***
#### func NewPeriod(start time.Time, end time.Time)  Period
<span id="NewPeriod"></span>
> 创建一个时间段
>   - 如果 start 比 end 晚，则会自动交换两个时间

***
#### func NewPeriodWindow(t time.Time, size time.Duration)  Period
<span id="NewPeriodWindow"></span>
> 创建一个特定长度的时间窗口

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewPeriodWindow(t *testing.T) {
	cur := time.Now()
	fmt.Println(cur)
	window := times.NewPeriodWindow(cur, times.Day)
	fmt.Println(window)
}

```


</details>


***
#### func NewPeriodWindowWeek(t time.Time)  Period
<span id="NewPeriodWindowWeek"></span>
> 创建一周长度的时间窗口，从周一零点开始至周日 23:59:59 结束

***
#### func NewPeriodWithTimeArray(times [2]time.Time)  Period
<span id="NewPeriodWithTimeArray"></span>
> 创建一个时间段

***
#### func NewPeriodWithDayZero(t time.Time, day int)  Period
<span id="NewPeriodWithDayZero"></span>
> 创建一个时间段，从 t 开始，持续到 day 天后的 0 点

***
#### func NewPeriodWithDay(t time.Time, day int)  Period
<span id="NewPeriodWithDay"></span>
> 创建一个时间段，从 t 开始，持续 day 天

***
#### func NewPeriodWithHour(t time.Time, hour int)  Period
<span id="NewPeriodWithHour"></span>
> 创建一个时间段，从 t 开始，持续 hour 小时

***
#### func NewPeriodWithMinute(t time.Time, minute int)  Period
<span id="NewPeriodWithMinute"></span>
> 创建一个时间段，从 t 开始，持续 minute 分钟

***
#### func NewPeriodWithSecond(t time.Time, second int)  Period
<span id="NewPeriodWithSecond"></span>
> 创建一个时间段，从 t 开始，持续 second 秒

***
#### func NewPeriodWithMillisecond(t time.Time, millisecond int)  Period
<span id="NewPeriodWithMillisecond"></span>
> 创建一个时间段，从 t 开始，持续 millisecond 毫秒

***
#### func NewPeriodWithMicrosecond(t time.Time, microsecond int)  Period
<span id="NewPeriodWithMicrosecond"></span>
> 创建一个时间段，从 t 开始，持续 microsecond 微秒

***
#### func NewPeriodWithNanosecond(t time.Time, nanosecond int)  Period
<span id="NewPeriodWithNanosecond"></span>
> 创建一个时间段，从 t 开始，持续 nanosecond 纳秒

***
### StateLine `STRUCT`
状态时间线
```go
type StateLine[State generic.Basic] struct {
	states  []State
	points  []time.Time
	trigger [][]func()
}
```
#### func (*StateLine) Check(missingAllowed bool, states ...State)  bool
> 根据状态顺序检查时间线是否合法
>   - missingAllowed: 是否允许状态缺失，如果为 true，则状态可以不连续，如果为 false，则状态必须连续
> 
> 状态不连续表示时间线中存在状态缺失，例如：
>   - 状态为 [1, 2, 3, 4, 5] 的时间线，如果 missingAllowed 为 true，则状态为 [1, 3, 5] 也是合法的
>   - 状态为 [1, 2, 3, 4, 5] 的时间线，如果 missingAllowed 为 false，则状态为 [1, 3, 5] 是不合法的
***
#### func (*StateLine) GetMissingStates(states ...State)  []State
> 获取缺失的状态
***
#### func (*StateLine) HasState(state State)  bool
> 检查时间线中是否包含指定状态
***
#### func (*StateLine) String()  string
> 获取时间线的字符串表示
***
#### func (*StateLine) AddState(state State, t time.Time, onTrigger ...func ())  *StateLine[State]
> 添加一个状态到时间线中，状态不能与任一时间点重合，否则将被忽略
>   - onTrigger: 该状态绑定的触发器，该触发器不会被主动执行，需要主动获取触发器执行
***
#### func (*StateLine) GetTimeByState(state State)  time.Time
> 获取指定状态的时间点
***
#### func (*StateLine) GetNextTimeByState(state State)  time.Time
> 获取指定状态的下一个时间点
***
#### func (*StateLine) GetLastState()  State
> 获取最后一个状态
***
#### func (*StateLine) GetPrevTimeByState(state State)  time.Time
> 获取指定状态的上一个时间点
***
#### func (*StateLine) GetIndexByState(state State)  int
> 获取指定状态的索引
***
#### func (*StateLine) GetStateByTime(t time.Time)  State
> 获取指定时间点的状态
***
#### func (*StateLine) GetTimeByIndex(index int)  time.Time
> 获取指定索引的时间点
***
#### func (*StateLine) Move(d time.Duration)  *StateLine[State]
> 时间线整体移动
***
#### func (*StateLine) GetNextStateTimeByIndex(index int)  time.Time
> 获取指定索引的下一个时间点
***
#### func (*StateLine) GetPrevStateTimeByIndex(index int)  time.Time
> 获取指定索引的上一个时间点
***
#### func (*StateLine) GetStateIndexByTime(t time.Time)  int
> 获取指定时间点的索引
***
#### func (*StateLine) GetStateCount()  int
> 获取状态数量
***
#### func (*StateLine) GetStateByIndex(index int)  State
> 获取指定索引的状态
***
#### func (*StateLine) GetTriggerByTime(t time.Time)  []func ()
> 获取指定时间点的触发器
***
#### func (*StateLine) GetTriggerByIndex(index int)  []func ()
> 获取指定索引的触发器
***
#### func (*StateLine) GetTriggerByState(state State)  []func ()
> 获取指定状态的触发器
***
#### func (*StateLine) AddTriggerToState(state State, onTrigger ...func ())  *StateLine[State]
> 给指定状态添加触发器
***
#### func (*StateLine) Range(handler func (index int, state State, t time.Time)  bool)
> 按照时间顺序遍历时间线
***
#### func (*StateLine) RangeReverse(handler func (index int, state State, t time.Time)  bool)
> 按照时间逆序遍历时间线
***
### Period `STRUCT`
表示一个时间段
```go
type Period [2]time.Time
```
#### func (Period) Start()  time.Time
> 返回时间段的开始时间
***
#### func (Period) End()  time.Time
> 返回时间段的结束时间
***
#### func (Period) Duration()  time.Duration
> 返回时间段的持续时间
***
#### func (Period) Day()  int
> 返回时间段的持续天数
***
#### func (Period) Hour()  int
> 返回时间段的持续小时数
***
#### func (Period) Minute()  int
> 返回时间段的持续分钟数
***
#### func (Period) Seconds()  int
> 返回时间段的持续秒数
***
#### func (Period) Milliseconds()  int
> 返回时间段的持续毫秒数
***
#### func (Period) Microseconds()  int
> 返回时间段的持续微秒数
***
#### func (Period) Nanoseconds()  int
> 返回时间段的持续纳秒数
***
#### func (Period) IsZero()  bool
> 判断时间段是否为零值
***
#### func (Period) IsInvalid()  bool
> 判断时间段是否无效
***
#### func (Period) IsBefore(t time.Time)  bool
> 判断时间段是否在指定时间之前
***
#### func (Period) IsAfter(t time.Time)  bool
> 判断时间段是否在指定时间之后
***
#### func (Period) IsBetween(t time.Time)  bool
> 判断指定时间是否在时间段之间
***
#### func (Period) IsOngoing(t time.Time)  bool
> 判断指定时间是否正在进行时
>   - 如果时间段的开始时间在指定时间之前或者等于指定时间，且时间段的结束时间在指定时间之后，则返回 true
***
#### func (Period) IsBetweenOrEqual(t time.Time)  bool
> 判断指定时间是否在时间段之间或者等于时间段的开始或结束时间
***
#### func (Period) IsBetweenOrEqualPeriod(t Period)  bool
> 判断指定时间是否在时间段之间或者等于时间段的开始或结束时间
***
#### func (Period) IsOverlap(t Period)  bool
> 判断时间段是否与指定时间段重叠
***
