# Moving



[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/moving)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewTwoDimensional](#NewTwoDimensional)|创建一个用于2D对象移动的实例(TwoDimensional)
|[WithTwoDimensionalTimeUnit](#WithTwoDimensionalTimeUnit)|通过特定时间单位创建
|[WithTwoDimensionalIdleWaitTime](#WithTwoDimensionalIdleWaitTime)|通过特定的空闲等待时间创建
|[WithTwoDimensionalInterval](#WithTwoDimensionalInterval)|通过特定的移动间隔时间创建
|[NewEntity](#NewEntity)|暂无描述...


> 结构体定义

|结构体|描述
|:--|:--
|[TwoDimensional](#twodimensional)|用于2D对象移动的数据结构
|[TwoDimensionalEntity](#twodimensionalentity)|2D移动对象接口定义
|[Position2DChangeEventHandle](#position2dchangeeventhandle)|暂无描述...
|[TwoDimensionalOption](#twodimensionaloption)|暂无描述...

</details>


#### func NewTwoDimensional(options ...TwoDimensionalOption[EID, PosType])  *TwoDimensional[EID, PosType]
<span id="NewTwoDimensional"></span>
> 创建一个用于2D对象移动的实例(TwoDimensional)
***
#### func WithTwoDimensionalTimeUnit(duration time.Duration)  TwoDimensionalOption[EID, PosType]
<span id="WithTwoDimensionalTimeUnit"></span>
> 通过特定时间单位创建
>   - 默认单位为1毫秒，最小单位也为1毫秒
***
#### func WithTwoDimensionalIdleWaitTime(duration time.Duration)  TwoDimensionalOption[EID, PosType]
<span id="WithTwoDimensionalIdleWaitTime"></span>
> 通过特定的空闲等待时间创建
>   - 默认情况下在没有新的移动计划时将限制 100毫秒 + 移动间隔事件(默认100毫秒)
***
#### func WithTwoDimensionalInterval(duration time.Duration)  TwoDimensionalOption[EID, PosType]
<span id="WithTwoDimensionalInterval"></span>
> 通过特定的移动间隔时间创建
***
#### func NewEntity(guid int64, speed float64)  *MoveEntity
<span id="NewEntity"></span>
***
### TwoDimensional
用于2D对象移动的数据结构
  - 通过对象调用 MoveTo 方法后将开始执行该对象的移动
  - 移动将在根据设置的每次移动间隔时间(WithTwoDimensionalInterval)进行移动，当无对象移动需要移动时将会进入短暂的休眠
  - 当对象移动速度永久为0时，将会导致永久无法完成的移动
```go
type TwoDimensional[EID generic.Basic, PosType generic.SignedNumber] struct {
	rw                                sync.RWMutex
	entities                          map[EID]*moving2DTarget[EID, PosType]
	timeUnit                          float64
	idle                              time.Duration
	interval                          time.Duration
	close                             bool
	position2DChangeEventHandles      []Position2DChangeEventHandle[EID, PosType]
	position2DDestinationEventHandles []Position2DDestinationEventHandle[EID, PosType]
	position2DStopMoveEventHandles    []Position2DStopMoveEventHandle[EID, PosType]
}
```
### TwoDimensionalEntity
2D移动对象接口定义
```go
type TwoDimensionalEntity[EID generic.Basic, PosType generic.SignedNumber] struct{}
```
### Position2DChangeEventHandle

```go
type Position2DChangeEventHandle[EID generic.Basic, PosType generic.SignedNumber] struct{}
```
### TwoDimensionalOption

```go
type TwoDimensionalOption[EID generic.Basic, PosType generic.SignedNumber] struct{}
```
