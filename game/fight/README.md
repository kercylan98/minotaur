# Fight



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/fight)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewTurnBased](#NewTurnBased)|创建一个新的回合制


> 结构体定义

|结构体|描述
|:--|:--
|[TurnBased](#turnbased)|回合制
|[TurnBasedControllerInfo](#turnbasedcontrollerinfo)|暂无描述...
|[TurnBasedControllerAction](#turnbasedcontrolleraction)|暂无描述...
|[TurnBasedController](#turnbasedcontroller)|回合制控制器
|[TurnBasedEntitySwitchEventHandler](#turnbasedentityswitcheventhandler)|暂无描述...

</details>


#### func NewTurnBased(calcNextTurnDuration func ( Camp,  Entity)  time.Duration)  *TurnBased[CampID, EntityID, Camp, Entity]
<span id="NewTurnBased"></span>
> 创建一个新的回合制
>   - calcNextTurnDuration 将返回下一次行动时间间隔，适用于按照速度计算下一次行动时间间隔的情况。当返回 0 时，将使用默认的行动超时时间
***
### TurnBased
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
### TurnBasedControllerInfo

```go
type TurnBasedControllerInfo[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct{}
```
### TurnBasedControllerAction

```go
type TurnBasedControllerAction[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct{}
```
### TurnBasedController
回合制控制器
```go
type TurnBasedController[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct {
	tb *TurnBased[CampID, EntityID, Camp, Entity]
}
```
### TurnBasedEntitySwitchEventHandler

```go
type TurnBasedEntitySwitchEventHandler[CampID comparable, EntityID comparable, Camp generic.IdR[CampID], Entity generic.IdR[EntityID]] struct{}
```
