# Fight

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
|[NewTurnBased](#NewTurnBased)|创建一个新的回合制


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[TurnBased](#struct_TurnBased)|回合制
|`INTERFACE`|[TurnBasedControllerInfo](#struct_TurnBasedControllerInfo)|暂无描述...
|`INTERFACE`|[TurnBasedControllerAction](#struct_TurnBasedControllerAction)|暂无描述...
|`STRUCT`|[TurnBasedController](#struct_TurnBasedController)|回合制控制器
|`STRUCT`|[TurnBasedEntitySwitchEventHandler](#struct_TurnBasedEntitySwitchEventHandler)|暂无描述...

</details>


***
## 详情信息
#### func NewTurnBased\[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]\](calcNextTurnDuration func ( Camp,  Entity)  time.Duration) *TurnBased[CampID, EntityID, Camp, Entity]
<span id="NewTurnBased"></span>
> 创建一个新的回合制
>   - calcNextTurnDuration 将返回下一次行动时间间隔，适用于按照速度计算下一次行动时间间隔的情况。当返回 0 时，将使用默认的行动超时时间

***
<span id="struct_TurnBased"></span>
### TurnBased `STRUCT`
回合制
```go
type TurnBased[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct {
	*turnBasedEvents[CampID, EntityID, Camp, Entity]
	controller           *TurnBasedController[CampID, EntityID, Camp, Entity]
	ticker               *time.Ticker
	actionWaitTicker     *time.Ticker
	actioning            bool
	actionMutex          sync.RWMutex
	entities             []Entity
	campRel              map[EntityID]Camp
	calcNextTurnDuration func(Camp, Entity) time.Duration
	actionTimeoutHandler func(Camp, Entity) time.Duration
	signal               chan signal
	round                int
	currCamp             Camp
	currEntity           Entity
	currActionTimeout    time.Duration
	currStart            time.Time
	closeMutex           sync.RWMutex
	closed               bool
}
```
<span id="struct_TurnBased_Close"></span>

#### func (*TurnBased) Close()
> 关闭回合制

***
<span id="struct_TurnBased_AddCamp"></span>

#### func (*TurnBased) AddCamp(camp Camp, entity Entity, entities ...Entity)
> 添加阵营

***
<span id="struct_TurnBased_SetActionTimeout"></span>

#### func (*TurnBased) SetActionTimeout(actionTimeoutHandler func ( Camp,  Entity)  time.Duration)
> 设置行动超时时间处理函数
>   - 默认情况下行动超时时间函数将始终返回 0

***
<span id="struct_TurnBased_Run"></span>

#### func (*TurnBased) Run()
> 运行

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestTurnBased_Run(t *testing.T) {
	tbi := fight.NewTurnBased[string, string, *Camp, *Entity](func(camp *Camp, entity *Entity) time.Duration {
		return time.Duration(float64(time.Second) / entity.speed)
	})
	tbi.SetActionTimeout(func(camp *Camp, entity *Entity) time.Duration {
		return time.Second * 5
	})
	tbi.RegTurnBasedEntityActionTimeoutEvent(func(controller fight.TurnBasedControllerInfo[string, string, *Camp, *Entity]) {
		t.Log("时间", time.Now().Unix(), "回合", controller.GetRound(), "阵营", controller.GetCamp().GetId(), "实体", controller.GetEntity().GetId(), "超时")
	})
	tbi.RegTurnBasedRoundChangeEvent(func(controller fight.TurnBasedControllerInfo[string, string, *Camp, *Entity]) {
		t.Log("时间", time.Now().Unix(), "回合", controller.GetRound(), "回合切换")
	})
	tbi.RegTurnBasedEntitySwitchEvent(func(controller fight.TurnBasedControllerAction[string, string, *Camp, *Entity]) {
		switch controller.GetEntity().GetId() {
		case "1":
			go func() {
				time.Sleep(time.Second * 2)
				controller.Finish()
			}()
		case "2":
			controller.Refresh(time.Second)
		case "4":
			controller.Stop()
		}
		t.Log("时间", time.Now().Unix(), "回合", controller.GetRound(), "阵营", controller.GetCamp().GetId(), "实体", controller.GetEntity().GetId(), "开始行动")
	})
	tbi.AddCamp(&Camp{id: "1"}, &Entity{id: "1", speed: 1}, &Entity{id: "2", speed: 1})
	tbi.AddCamp(&Camp{id: "2"}, &Entity{id: "3", speed: 1}, &Entity{id: "4", speed: 1})
	tbi.Run()
}

```


</details>


***
<span id="struct_TurnBasedControllerInfo"></span>
### TurnBasedControllerInfo `INTERFACE`

```go
type TurnBasedControllerInfo[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] interface {
	GetRound() int
	GetCamp() Camp
	GetEntity() Entity
	GetActionTimeoutDuration() time.Duration
	GetActionStartTime() time.Time
	GetActionEndTime() time.Time
	Stop()
}
```
<span id="struct_TurnBasedControllerAction"></span>
### TurnBasedControllerAction `INTERFACE`

```go
type TurnBasedControllerAction[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] interface {
	TurnBasedControllerInfo[CampID, EntityID, Camp, Entity]
	Finish()
	Refresh(duration time.Duration) time.Time
}
```
<span id="struct_TurnBasedController"></span>
### TurnBasedController `STRUCT`
回合制控制器
```go
type TurnBasedController[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct {
	tb *TurnBased[CampID, EntityID, Camp, Entity]
}
```
<span id="struct_TurnBasedController_GetRound"></span>

#### func (*TurnBasedController) GetRound()  int
> 获取当前回合数

***
<span id="struct_TurnBasedController_GetCamp"></span>

#### func (*TurnBasedController) GetCamp()  Camp
> 获取当前操作阵营

***
<span id="struct_TurnBasedController_GetEntity"></span>

#### func (*TurnBasedController) GetEntity()  Entity
> 获取当前操作实体

***
<span id="struct_TurnBasedController_GetActionTimeoutDuration"></span>

#### func (*TurnBasedController) GetActionTimeoutDuration()  time.Duration
> 获取当前行动超时时长

***
<span id="struct_TurnBasedController_GetActionStartTime"></span>

#### func (*TurnBasedController) GetActionStartTime()  time.Time
> 获取当前行动开始时间

***
<span id="struct_TurnBasedController_GetActionEndTime"></span>

#### func (*TurnBasedController) GetActionEndTime()  time.Time
> 获取当前行动结束时间

***
<span id="struct_TurnBasedController_Finish"></span>

#### func (*TurnBasedController) Finish()
> 结束当前操作，将立即切换到下一个操作实体

***
<span id="struct_TurnBasedController_Stop"></span>

#### func (*TurnBasedController) Stop()
> 在当前回合执行完毕后停止回合进程

***
<span id="struct_TurnBasedController_Refresh"></span>

#### func (*TurnBasedController) Refresh(duration time.Duration)  time.Time
> 刷新当前操作实体的行动超时时间
>   - 当不在行动阶段时，将返回 time.Time 零值

***
<span id="struct_TurnBasedEntitySwitchEventHandler"></span>
### TurnBasedEntitySwitchEventHandler `STRUCT`

```go
type TurnBasedEntitySwitchEventHandler[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] func(controller TurnBasedControllerAction[CampID, EntityID, Camp, Entity])
```
