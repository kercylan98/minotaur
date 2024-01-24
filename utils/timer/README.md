# Timer

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
|[RegSystemNewDayEvent](#RegSystemNewDayEvent)|注册系统新的一天事件
|[OnSystemNewDayEvent](#OnSystemNewDayEvent)|系统新的一天事件
|[RegOffsetTimeNewDayEvent](#RegOffsetTimeNewDayEvent)|注册偏移时间新的一天事件
|[OnOffsetTimeNewDayEvent](#OnOffsetTimeNewDayEvent)|偏移时间新的一天事件
|[WithCaller](#WithCaller)|通过其他的 handle 执行 Caller
|[WithMark](#WithMark)|通过特定的标记创建定时器
|[NewPool](#NewPool)|创建一个定时器池，当 tickerPoolSize 小于等于 0 时，将会引发 panic，可指定为 DefaultTickerPoolSize
|[SetPoolSize](#SetPoolSize)|设置标准池定时器池大小
|[GetTicker](#GetTicker)|获取标准池中的一个定时器


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[SystemNewDayEventHandle](#struct_SystemNewDayEventHandle)|暂无描述...
|`STRUCT`|[Option](#struct_Option)|暂无描述...
|`STRUCT`|[Pool](#struct_Pool)|定时器池
|`STRUCT`|[Scheduler](#struct_Scheduler)|调度器
|`STRUCT`|[Ticker](#struct_Ticker)|定时器

</details>


***
## 详情信息
#### func RegSystemNewDayEvent(ticker *Ticker, name string, trigger bool, handle SystemNewDayEventHandle)
<span id="RegSystemNewDayEvent"></span>
> 注册系统新的一天事件
>   - 建议全局注册一个事件后再另行拓展
>   - 将特定 name 的定时任务注册到 ticker 中，在系统时间到达每天的 00:00:00 时触发，如果 trigger 为 true，则立即触发一次

***
#### func OnSystemNewDayEvent(name string)
<span id="OnSystemNewDayEvent"></span>
> 系统新的一天事件

***
#### func RegOffsetTimeNewDayEvent(ticker *Ticker, name string, offset *offset.Time, trigger bool, handle OffsetTimeNewDayEventHandle)
<span id="RegOffsetTimeNewDayEvent"></span>
> 注册偏移时间新的一天事件
>   - 建议全局注册一个事件后再另行拓展
>   - 与 RegSystemNewDayEvent 类似，但是触发时间为 offset 时间到达每天的 00:00:00

***
#### func OnOffsetTimeNewDayEvent(name string)
<span id="OnOffsetTimeNewDayEvent"></span>
> 偏移时间新的一天事件

***
#### func WithCaller(handle func (name string, caller func ())) Option
<span id="WithCaller"></span>
> 通过其他的 handle 执行 Caller

***
#### func WithMark(mark string) Option
<span id="WithMark"></span>
> 通过特定的标记创建定时器

***
#### func NewPool(tickerPoolSize int) *Pool
<span id="NewPool"></span>
> 创建一个定时器池，当 tickerPoolSize 小于等于 0 时，将会引发 panic，可指定为 DefaultTickerPoolSize

***
#### func SetPoolSize(size int)
<span id="SetPoolSize"></span>
> 设置标准池定时器池大小
>   - 默认值为 DefaultTickerPoolSize，当定时器池中的定时器不足时，会自动创建新的定时器，当定时器释放时，会将多余的定时器进行释放，否则将放入定时器池中

***
#### func GetTicker(size int, options ...Option) *Ticker
<span id="GetTicker"></span>
> 获取标准池中的一个定时器

***
<span id="struct_SystemNewDayEventHandle"></span>
### SystemNewDayEventHandle `STRUCT`

```go
type SystemNewDayEventHandle func()
```
<span id="struct_Option"></span>
### Option `STRUCT`

```go
type Option func(ticker *Ticker)
```
<span id="struct_Pool"></span>
### Pool `STRUCT`
定时器池
```go
type Pool struct {
	tickers        []*Ticker
	lock           sync.Mutex
	tickerPoolSize int
	closed         bool
}
```
<span id="struct_Pool_ChangePoolSize"></span>

#### func (*Pool) ChangePoolSize(size int)  error
> 改变定时器池大小
>   - 当传入的大小小于或等于 0 时，将会返回错误，并且不会发生任何改变

***
<span id="struct_Pool_GetTicker"></span>

#### func (*Pool) GetTicker(size int, options ...Option)  *Ticker
> 获取一个新的定时器

***
<span id="struct_Pool_Release"></span>

#### func (*Pool) Release()
> 释放定时器池的资源，释放后由其产生的 Ticker 在 Ticker.Release 后将不再回到池中，而是直接释放
>   - 虽然定时器池已被释放，但是依旧可以产出 Ticker

***
<span id="struct_Scheduler"></span>
### Scheduler `STRUCT`
调度器
```go
type Scheduler struct {
	name     string
	after    time.Duration
	interval time.Duration
	total    int
	trigger  int
	kill     bool
	cbFunc   reflect.Value
	cbArgs   []reflect.Value
	timer    *timingwheel.Timer
	ticker   *Ticker
	lock     sync.RWMutex
	expr     *cronexpr.Expression
}
```
<span id="struct_Scheduler_Name"></span>

#### func (*Scheduler) Name()  string
> 获取调度器名称

***
<span id="struct_Scheduler_Next"></span>

#### func (*Scheduler) Next(prev time.Time)  time.Time
> 获取下一次执行的时间

***
<span id="struct_Scheduler_Caller"></span>

#### func (*Scheduler) Caller()
> 可由外部发起调用的执行函数

***
<span id="struct_Ticker"></span>
### Ticker `STRUCT`
定时器
```go
type Ticker struct {
	timer  *Pool
	wheel  *timingwheel.TimingWheel
	timers map[string]*Scheduler
	lock   sync.RWMutex
	handle func(name string, caller func())
	mark   string
}
```
<span id="struct_Ticker_Mark"></span>

#### func (*Ticker) Mark()  string
> 获取定时器的标记
>   - 通常用于鉴别定时器来源

***
<span id="struct_Ticker_Release"></span>

#### func (*Ticker) Release()
> 释放定时器，并将定时器重新放回 Pool 池中

***
<span id="struct_Ticker_StopTimer"></span>

#### func (*Ticker) StopTimer(name string)
> 停止特定名称的调度器

***
<span id="struct_Ticker_IsStopped"></span>

#### func (*Ticker) IsStopped(name string)  bool
> 特定名称的调度器是否已停止

***
<span id="struct_Ticker_GetSchedulers"></span>

#### func (*Ticker) GetSchedulers()  []string
> 获取所有调度器名称

***
<span id="struct_Ticker_Cron"></span>

#### func (*Ticker) Cron(name string, expression string, handleFunc interface {}, args ...interface {})
> 通过 cron 表达式设置一个调度器，当 cron 表达式错误时，将会引发 panic

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestTicker_Cron(t *testing.T) {
	ticker := timer.GetTicker(10)
	ticker.After("1_sec", time.Second, func() {
		t.Log(time.Now().Format(time.DateTime), "1_sec")
	})
	ticker.Loop("1_sec_loop_3", 0, time.Second, 3, func() {
		t.Log(time.Now().Format(time.DateTime), "1_sec_loop_3")
	})
	ticker.Cron("5_sec_cron", "0/5 * * * * * ?", func() {
		t.Log(time.Now().Format(time.DateTime), "5_sec_cron")
	})
	time.Sleep(times.Week)
}

```


</details>


***
<span id="struct_Ticker_CronByInstantly"></span>

#### func (*Ticker) CronByInstantly(name string, expression string, handleFunc interface {}, args ...interface {})
> 与 Cron 相同，但是会立即执行一次

***
<span id="struct_Ticker_After"></span>

#### func (*Ticker) After(name string, after time.Duration, handleFunc interface {}, args ...interface {})
> 设置一个在特定时间后运行一次的调度器

***
<span id="struct_Ticker_Loop"></span>

#### func (*Ticker) Loop(name string, after time.Duration, interval time.Duration, times int, handleFunc interface {}, args ...interface {})
> 设置一个在特定时间后反复运行的调度器

***
