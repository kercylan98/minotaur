# Task

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
|[Cond](#Cond)|创建任务条件
|[RegisterRefreshTaskCounterEvent](#RegisterRefreshTaskCounterEvent)|注册特定任务类型的刷新任务计数器事件处理函数
|[OnRefreshTaskCounterEvent](#OnRefreshTaskCounterEvent)|触发特定任务类型的刷新任务计数器事件
|[RegisterRefreshTaskConditionEvent](#RegisterRefreshTaskConditionEvent)|注册特定任务类型的刷新任务条件事件处理函数
|[OnRefreshTaskConditionEvent](#OnRefreshTaskConditionEvent)|触发特定任务类型的刷新任务条件事件
|[WithType](#WithType)|设置任务类型
|[WithCondition](#WithCondition)|设置任务完成条件，当满足条件时，任务状态为完成
|[WithCounter](#WithCounter)|设置任务计数器，当计数器达到要求时，任务状态为完成
|[WithOverflowCounter](#WithOverflowCounter)|设置可溢出的任务计数器，当计数器达到要求时，任务状态为完成
|[WithDeadline](#WithDeadline)|设置任务截止时间，超过截至时间并且任务未完成时，任务状态为失败
|[WithLimitedDuration](#WithLimitedDuration)|设置任务限时，超过限时时间并且任务未完成时，任务状态为失败
|[NewTask](#NewTask)|生成任务


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Condition](#struct_Condition)|任务条件
|`STRUCT`|[RefreshTaskCounterEventHandler](#struct_RefreshTaskCounterEventHandler)|暂无描述...
|`STRUCT`|[Option](#struct_Option)|任务选项
|`STRUCT`|[Status](#struct_Status)|暂无描述...
|`STRUCT`|[Task](#struct_Task)|是对任务信息进行描述和处理的结构体

</details>


***
## 详情信息
#### func Cond(k any, v any) Condition
<span id="Cond"></span>
> 创建任务条件

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestCond(t *testing.T) {
	task := NewTask(WithType("T"), WithCounter(5), WithCondition(Cond("N", 5).Cond("M", 10)))
	task.AssignConditionValueAndRefresh("N", 5)
	task.AssignConditionValueAndRefresh("M", 10)
	RegisterRefreshTaskCounterEvent[*Player](task.Type, func(taskType string, trigger *Player, count int64) {
		fmt.Println("Player", count)
		for _, t := range trigger.tasks[taskType] {
			fmt.Println(t.CurrCount, t.IncrementCounter(count).Status)
		}
	})
	RegisterRefreshTaskConditionEvent[*Player](task.Type, func(taskType string, trigger *Player, condition Condition) {
		fmt.Println("Player", condition)
		for _, t := range trigger.tasks[taskType] {
			fmt.Println(t.CurrCount, t.AssignConditionValueAndRefresh("N", 5).Status)
		}
	})
	RegisterRefreshTaskCounterEvent[*Monster](task.Type, func(taskType string, trigger *Monster, count int64) {
		fmt.Println("Monster", count)
	})
	player := &Player{tasks: map[string][]*Task{task.Type: {task}}}
	OnRefreshTaskCounterEvent(task.Type, player, 1)
	OnRefreshTaskCounterEvent(task.Type, player, 2)
	OnRefreshTaskCounterEvent(task.Type, player, 3)
	OnRefreshTaskCounterEvent(task.Type, new(Monster), 3)
}

```


</details>


***
#### func RegisterRefreshTaskCounterEvent(taskType string, handler RefreshTaskCounterEventHandler[Trigger])
<span id="RegisterRefreshTaskCounterEvent"></span>
> 注册特定任务类型的刷新任务计数器事件处理函数

***
#### func OnRefreshTaskCounterEvent(taskType string, trigger any, count int64)
<span id="OnRefreshTaskCounterEvent"></span>
> 触发特定任务类型的刷新任务计数器事件

***
#### func RegisterRefreshTaskConditionEvent(taskType string, handler RefreshTaskConditionEventHandler[Trigger])
<span id="RegisterRefreshTaskConditionEvent"></span>
> 注册特定任务类型的刷新任务条件事件处理函数

***
#### func OnRefreshTaskConditionEvent(taskType string, trigger any, condition Condition)
<span id="OnRefreshTaskConditionEvent"></span>
> 触发特定任务类型的刷新任务条件事件

***
#### func WithType(taskType string) Option
<span id="WithType"></span>
> 设置任务类型

***
#### func WithCondition(condition Condition) Option
<span id="WithCondition"></span>
> 设置任务完成条件，当满足条件时，任务状态为完成
>   - 任务条件值需要变更时可通过 Task.AssignConditionValueAndRefresh 方法变更
>   - 当多次设置该选项时，后面的设置会覆盖之前的设置

***
#### func WithCounter(counter int64, initCount ...int64) Option
<span id="WithCounter"></span>
> 设置任务计数器，当计数器达到要求时，任务状态为完成
>   - 一些场景下，任务计数器可能会溢出，此时可通过 WithOverflowCounter 设置可溢出的任务计数器
>   - 当多次设置该选项时，后面的设置会覆盖之前的设置
>   - 如果需要初始化计数器的值，可通过 initCount 参数设置

***
#### func WithOverflowCounter(counter int64, initCount ...int64) Option
<span id="WithOverflowCounter"></span>
> 设置可溢出的任务计数器，当计数器达到要求时，任务状态为完成
>   - 当多次设置该选项时，后面的设置会覆盖之前的设置
>   - 如果需要初始化计数器的值，可通过 initCount 参数设置

***
#### func WithDeadline(deadline time.Time) Option
<span id="WithDeadline"></span>
> 设置任务截止时间，超过截至时间并且任务未完成时，任务状态为失败

***
#### func WithLimitedDuration(start time.Time, duration time.Duration) Option
<span id="WithLimitedDuration"></span>
> 设置任务限时，超过限时时间并且任务未完成时，任务状态为失败

***
#### func NewTask(options ...Option) *Task
<span id="NewTask"></span>
> 生成任务

***
<span id="struct_Condition"></span>
### Condition `STRUCT`
任务条件
```go
type Condition map[any]any
```
#### func (Condition) Cond(k any, v any)  Condition
> 创建任务条件
***
#### func (Condition) GetString(key any)  string
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetInt(key any)  int
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetInt8(key any)  int8
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetInt16(key any)  int16
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetInt32(key any)  int32
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetInt64(key any)  int64
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetUint(key any)  uint
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetUint8(key any)  uint8
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetUint16(key any)  uint16
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetUint32(key any)  uint32
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetUint64(key any)  uint64
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetFloat32(key any)  float32
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetFloat64(key any)  float64
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetBool(key any)  bool
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetTime(key any)  time.Time
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetDuration(key any)  time.Duration
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetByte(key any)  byte
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetBytes(key any)  []byte
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetRune(key any)  rune
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetRunes(key any)  []rune
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
#### func (Condition) GetAny(key any)  any
> 获取特定类型的任务条件值，该值必须与预期类型一致，否则返回零值
***
<span id="struct_RefreshTaskCounterEventHandler"></span>
### RefreshTaskCounterEventHandler `STRUCT`

```go
type RefreshTaskCounterEventHandler[Trigger any] func(taskType string, trigger Trigger, count int64)
```
<span id="struct_Option"></span>
### Option `STRUCT`
任务选项
```go
type Option func(task *Task)
```
<span id="struct_Status"></span>
### Status `STRUCT`

```go
type Status byte
```
#### func (Status) String()  string
***
<span id="struct_Task"></span>
### Task `STRUCT`
是对任务信息进行描述和处理的结构体
```go
type Task struct {
	Type            string
	Status          Status
	Cond            Condition
	CondValue       map[any]any
	Counter         int64
	CurrCount       int64
	CurrOverflow    bool
	Deadline        time.Time
	StartTime       time.Time
	LimitedDuration time.Duration
}
```
#### func (*Task) IsComplete()  bool
> 判断任务是否已完成
***
#### func (*Task) IsFailed()  bool
> 判断任务是否已失败
***
#### func (*Task) IsReward()  bool
> 判断任务是否已领取奖励
***
#### func (*Task) ReceiveReward()  bool
> 领取任务奖励，当任务状态为已完成时，才能领取奖励，此时返回 true，并且任务状态变更为已领取奖励
***
#### func (*Task) IncrementCounter(incr int64)  *Task
> 增加计数器的值，当 incr 为负数时，计数器的值不会发生变化
>   - 如果需要溢出计数器，可通过 WithOverflowCounter 设置可溢出的任务计数器
***
#### func (*Task) DecrementCounter(decr int64)  *Task
> 减少计数器的值，当 decr 为负数时，计数器的值不会发生变化
***
#### func (*Task) AssignConditionValueAndRefresh(key any, value any)  *Task
> 分配条件值并刷新任务状态
***
#### func (*Task) AssignConditionValueAndRefreshByCondition(condition Condition)  *Task
> 分配条件值并刷新任务状态
***
#### func (*Task) ResetStatus()  *Task
> 重置任务状态
>   - 该函数会将任务状态重置为已接受状态后，再刷新任务状态
>   - 当任务条件变更，例如任务计数要求为 10，已经完成的情况下，将任务计数要求变更为 5 或 20，此时任务状态由于是已完成或已领取状态，不会自动刷新，需要调用该函数刷新任务状态
***
