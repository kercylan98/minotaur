# Activity

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

activity 活动状态管理


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[SetTicker](#SetTicker)|设置自定义定时器，该方法必须在使用活动系统前调用，且只能调用一次
|[LoadGlobalData](#LoadGlobalData)|加载所有活动全局数据
|[LoadEntityData](#LoadEntityData)|加载所有活动实体数据
|[LoadOrRefreshActivity](#LoadOrRefreshActivity)|加载或刷新活动
|[DefineNoneDataActivity](#DefineNoneDataActivity)|声明无数据的活动类型
|[DefineGlobalDataActivity](#DefineGlobalDataActivity)|声明拥有全局数据的活动类型
|[DefineEntityDataActivity](#DefineEntityDataActivity)|声明拥有实体数据的活动类型
|[DefineGlobalAndEntityDataActivity](#DefineGlobalAndEntityDataActivity)|声明拥有全局数据和实体数据的活动类型
|[RegUpcomingEvent](#RegUpcomingEvent)|注册即将开始的活动事件处理器
|[OnUpcomingEvent](#OnUpcomingEvent)|即将开始的活动事件
|[RegStartedEvent](#RegStartedEvent)|注册活动开始事件处理器
|[OnStartedEvent](#OnStartedEvent)|活动开始事件
|[RegEndedEvent](#RegEndedEvent)|注册活动结束事件处理器
|[OnEndedEvent](#OnEndedEvent)|活动结束事件
|[RegExtendedShowStartedEvent](#RegExtendedShowStartedEvent)|注册活动结束后延长展示开始事件处理器
|[OnExtendedShowStartedEvent](#OnExtendedShowStartedEvent)|活动结束后延长展示开始事件
|[RegExtendedShowEndedEvent](#RegExtendedShowEndedEvent)|注册活动结束后延长展示结束事件处理器
|[OnExtendedShowEndedEvent](#OnExtendedShowEndedEvent)|活动结束后延长展示结束事件
|[RegNewDayEvent](#RegNewDayEvent)|注册新的一天事件处理器
|[OnNewDayEvent](#OnNewDayEvent)|新的一天事件
|[NewOptions](#NewOptions)|创建活动选项


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Activity](#struct_Activity)|活动描述
|`STRUCT`|[Controller](#struct_Controller)|活动控制器
|`INTERFACE`|[BasicActivityController](#struct_BasicActivityController)|暂无描述...
|`INTERFACE`|[NoneDataActivityController](#struct_NoneDataActivityController)|无数据活动控制器
|`INTERFACE`|[GlobalDataActivityController](#struct_GlobalDataActivityController)|全局数据活动控制器
|`INTERFACE`|[EntityDataActivityController](#struct_EntityDataActivityController)|实体数据活动控制器
|`INTERFACE`|[GlobalAndEntityDataActivityController](#struct_GlobalAndEntityDataActivityController)|全局数据和实体数据活动控制器
|`STRUCT`|[DataMeta](#struct_DataMeta)|全局活动数据
|`STRUCT`|[EntityDataMeta](#struct_EntityDataMeta)|活动实体数据
|`STRUCT`|[UpcomingEventHandler](#struct_UpcomingEventHandler)|暂无描述...
|`STRUCT`|[Options](#struct_Options)|活动选项

</details>


***
## 详情信息
#### func SetTicker(size int, options ...timer.Option)
<span id="SetTicker"></span>
> 设置自定义定时器，该方法必须在使用活动系统前调用，且只能调用一次

***
#### func LoadGlobalData(handler func (activityType any))
<span id="LoadGlobalData"></span>
> 加载所有活动全局数据

***
#### func LoadEntityData(handler func (activityType any))
<span id="LoadEntityData"></span>
> 加载所有活动实体数据

***
#### func LoadOrRefreshActivity\[Type generic.Basic, ID generic.Basic\](activityType Type, activityId ID, options ...*Options) error
<span id="LoadOrRefreshActivity"></span>
> 加载或刷新活动
>   - 通常在活动配置刷新时候将活动通过该方法注册或刷新

***
#### func DefineNoneDataActivity\[Type generic.Basic, ID generic.Basic\](activityType Type) NoneDataActivityController[Type, ID, *none, none, *none]
<span id="DefineNoneDataActivity"></span>
> 声明无数据的活动类型

***
#### func DefineGlobalDataActivity\[Type generic.Basic, ID generic.Basic, Data any\](activityType Type) GlobalDataActivityController[Type, ID, Data, none, *none]
<span id="DefineGlobalDataActivity"></span>
> 声明拥有全局数据的活动类型

***
#### func DefineEntityDataActivity\[Type generic.Basic, ID generic.Basic, EntityID generic.Basic, EntityData any\](activityType Type) EntityDataActivityController[Type, ID, *none, EntityID, EntityData]
<span id="DefineEntityDataActivity"></span>
> 声明拥有实体数据的活动类型

***
#### func DefineGlobalAndEntityDataActivity\[Type generic.Basic, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any\](activityType Type) GlobalAndEntityDataActivityController[Type, ID, Data, EntityID, EntityData]
<span id="DefineGlobalAndEntityDataActivity"></span>
> 声明拥有全局数据和实体数据的活动类型

***
#### func RegUpcomingEvent\[Type generic.Basic, ID generic.Basic\](activityType Type, handler UpcomingEventHandler[ID], priority ...int)
<span id="RegUpcomingEvent"></span>
> 注册即将开始的活动事件处理器

***
#### func OnUpcomingEvent\[Type generic.Basic, ID generic.Basic\](activity *Activity[Type, ID])
<span id="OnUpcomingEvent"></span>
> 即将开始的活动事件

***
#### func RegStartedEvent\[Type generic.Basic, ID generic.Basic\](activityType Type, handler StartedEventHandler[ID], priority ...int)
<span id="RegStartedEvent"></span>
> 注册活动开始事件处理器

***
#### func OnStartedEvent\[Type generic.Basic, ID generic.Basic\](activity *Activity[Type, ID])
<span id="OnStartedEvent"></span>
> 活动开始事件

***
#### func RegEndedEvent\[Type generic.Basic, ID generic.Basic\](activityType Type, handler EndedEventHandler[ID], priority ...int)
<span id="RegEndedEvent"></span>
> 注册活动结束事件处理器

***
#### func OnEndedEvent\[Type generic.Basic, ID generic.Basic\](activity *Activity[Type, ID])
<span id="OnEndedEvent"></span>
> 活动结束事件

***
#### func RegExtendedShowStartedEvent\[Type generic.Basic, ID generic.Basic\](activityType Type, handler ExtendedShowStartedEventHandler[ID], priority ...int)
<span id="RegExtendedShowStartedEvent"></span>
> 注册活动结束后延长展示开始事件处理器

***
#### func OnExtendedShowStartedEvent\[Type generic.Basic, ID generic.Basic\](activity *Activity[Type, ID])
<span id="OnExtendedShowStartedEvent"></span>
> 活动结束后延长展示开始事件

***
#### func RegExtendedShowEndedEvent\[Type generic.Basic, ID generic.Basic\](activityType Type, handler ExtendedShowEndedEventHandler[ID], priority ...int)
<span id="RegExtendedShowEndedEvent"></span>
> 注册活动结束后延长展示结束事件处理器

***
#### func OnExtendedShowEndedEvent\[Type generic.Basic, ID generic.Basic\](activity *Activity[Type, ID])
<span id="OnExtendedShowEndedEvent"></span>
> 活动结束后延长展示结束事件

***
#### func RegNewDayEvent\[Type generic.Basic, ID generic.Basic\](activityType Type, handler NewDayEventHandler[ID], priority ...int)
<span id="RegNewDayEvent"></span>
> 注册新的一天事件处理器

***
#### func OnNewDayEvent\[Type generic.Basic, ID generic.Basic\](activity *Activity[Type, ID])
<span id="OnNewDayEvent"></span>
> 新的一天事件

***
#### func NewOptions() *Options
<span id="NewOptions"></span>
> 创建活动选项

***
<span id="struct_Activity"></span>
### Activity `STRUCT`
活动描述
```go
type Activity[Type generic.Basic, ID generic.Basic] struct {
	id                ID
	t                 Type
	options           *Options
	state             byte
	lazy              bool
	tickerKey         string
	retention         time.Duration
	retentionKey      string
	mutex             sync.RWMutex
	getLastNewDayTime func() time.Time
	setLastNewDayTime func(time.Time)
	clearData         func()
	initializeData    func()
}
```
<span id="struct_Controller"></span>
### Controller `STRUCT`
活动控制器
```go
type Controller[Type generic.Basic, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] struct {
	t          Type
	activities map[ID]*Activity[Type, ID]
	globalData map[ID]*DataMeta[Data]
	entityData map[ID]map[EntityID]*EntityDataMeta[EntityData]
	entityTof  reflect.Type
	globalInit func(activityId ID, data *DataMeta[Data])
	entityInit func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])
	mutex      sync.RWMutex
}
```
<span id="struct_BasicActivityController"></span>
### BasicActivityController `INTERFACE`

```go
type BasicActivityController[Type generic.Basic, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	IsOpen(activityId ID) bool
	IsShow(activityId ID) bool
	IsOpenOrShow(activityId ID) bool
	Refresh(activityId ID)
}
```
<span id="struct_NoneDataActivityController"></span>
### NoneDataActivityController `INTERFACE`
无数据活动控制器
```go
type NoneDataActivityController[Type generic.Basic, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	InitializeNoneData(handler func(activityId ID, data *DataMeta[Data])) NoneDataActivityController[Type, ID, Data, EntityID, EntityData]
}
```
<span id="struct_GlobalDataActivityController"></span>
### GlobalDataActivityController `INTERFACE`
全局数据活动控制器
```go
type GlobalDataActivityController[Type generic.Basic, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	GetGlobalData(activityId ID) Data
	InitializeGlobalData(handler func(activityId ID, data *DataMeta[Data])) GlobalDataActivityController[Type, ID, Data, EntityID, EntityData]
}
```
<span id="struct_EntityDataActivityController"></span>
### EntityDataActivityController `INTERFACE`
实体数据活动控制器
```go
type EntityDataActivityController[Type generic.Basic, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	GetEntityData(activityId ID, entityId EntityID) EntityData
	InitializeEntityData(handler func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) EntityDataActivityController[Type, ID, Data, EntityID, EntityData]
}
```
<span id="struct_GlobalAndEntityDataActivityController"></span>
### GlobalAndEntityDataActivityController `INTERFACE`
全局数据和实体数据活动控制器
```go
type GlobalAndEntityDataActivityController[Type generic.Basic, ID generic.Basic, Data any, EntityID generic.Basic, EntityData any] interface {
	BasicActivityController[Type, ID, Data, EntityID, EntityData]
	GetGlobalData(activityId ID) Data
	GetEntityData(activityId ID, entityId EntityID) EntityData
	InitializeGlobalAndEntityData(handler func(activityId ID, data *DataMeta[Data]), entityHandler func(activityId ID, entityId EntityID, data *EntityDataMeta[EntityData])) GlobalAndEntityDataActivityController[Type, ID, Data, EntityID, EntityData]
}
```
<span id="struct_DataMeta"></span>
### DataMeta `STRUCT`
全局活动数据
```go
type DataMeta[Data any] struct {
	once       sync.Once
	Data       Data
	LastNewDay time.Time
}
```
<span id="struct_EntityDataMeta"></span>
### EntityDataMeta `STRUCT`
活动实体数据
```go
type EntityDataMeta[Data any] struct {
	once       sync.Once
	Data       Data
	LastNewDay time.Time
}
```
<span id="struct_UpcomingEventHandler"></span>
### UpcomingEventHandler `STRUCT`

```go
type UpcomingEventHandler[ID generic.Basic] func(activityId ID)
```
<span id="struct_Options"></span>
### Options `STRUCT`
活动选项
```go
type Options struct {
	Tl   *times.StateLine[byte]
	Loop time.Duration
}
```
#### func (*Options) WithUpcomingTime(t time.Time)  *Options
> 设置活动预告时间
***
#### func (*Options) WithStartTime(t time.Time)  *Options
> 设置活动开始时间
***
#### func (*Options) WithEndTime(t time.Time)  *Options
> 设置活动结束时间
***
#### func (*Options) WithExtendedShowTime(t time.Time)  *Options
> 设置延长展示时间
***
#### func (*Options) WithLoop(interval time.Duration)  *Options
> 设置活动循环，时间间隔小于等于 0 表示不循环
>   - 当活动状态展示结束后，会根据该选项设置的时间间隔重新开始
***
