# Moving

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
|[NewTwoDimensional](#NewTwoDimensional)|创建一个用于2D对象移动的实例(TwoDimensional)
|[WithTwoDimensionalTimeUnit](#WithTwoDimensionalTimeUnit)|通过特定时间单位创建
|[WithTwoDimensionalIdleWaitTime](#WithTwoDimensionalIdleWaitTime)|通过特定的空闲等待时间创建
|[WithTwoDimensionalInterval](#WithTwoDimensionalInterval)|通过特定的移动间隔时间创建
|[NewEntity](#NewEntity)|暂无描述...


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[TwoDimensional](#struct_TwoDimensional)|用于2D对象移动的数据结构
|`INTERFACE`|[TwoDimensionalEntity](#struct_TwoDimensionalEntity)|2D移动对象接口定义
|`STRUCT`|[Position2DChangeEventHandle](#struct_Position2DChangeEventHandle)|暂无描述...
|`STRUCT`|[TwoDimensionalOption](#struct_TwoDimensionalOption)|暂无描述...

</details>


***
## 详情信息
#### func NewTwoDimensional\[EID generic.Basic, PosType generic.SignedNumber\](options ...TwoDimensionalOption[EID, PosType]) *TwoDimensional[EID, PosType]
<span id="NewTwoDimensional"></span>
> 创建一个用于2D对象移动的实例(TwoDimensional)

**示例代码：**

```go

func ExampleNewTwoDimensional() {
	m := moving2.NewTwoDimensional[int64, float64]()
	defer func() {
		m.Release()
	}()
	fmt.Println(m != nil)
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestNewTwoDimensional(t *testing.T) {
	m := moving2.NewTwoDimensional[int64, float64]()
	defer func() {
		m.Release()
	}()
}

```


</details>


***
#### func WithTwoDimensionalTimeUnit\[EID generic.Basic, PosType generic.SignedNumber\](duration time.Duration) TwoDimensionalOption[EID, PosType]
<span id="WithTwoDimensionalTimeUnit"></span>
> 通过特定时间单位创建
>   - 默认单位为1毫秒，最小单位也为1毫秒

***
#### func WithTwoDimensionalIdleWaitTime\[EID generic.Basic, PosType generic.SignedNumber\](duration time.Duration) TwoDimensionalOption[EID, PosType]
<span id="WithTwoDimensionalIdleWaitTime"></span>
> 通过特定的空闲等待时间创建
>   - 默认情况下在没有新的移动计划时将限制 100毫秒 + 移动间隔事件(默认100毫秒)

***
#### func WithTwoDimensionalInterval\[EID generic.Basic, PosType generic.SignedNumber\](duration time.Duration) TwoDimensionalOption[EID, PosType]
<span id="WithTwoDimensionalInterval"></span>
> 通过特定的移动间隔时间创建

***
#### func NewEntity(guid int64, speed float64) *MoveEntity
<span id="NewEntity"></span>

***
<span id="struct_TwoDimensional"></span>
### TwoDimensional `STRUCT`
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
#### func (*TwoDimensional) MoveTo(entity TwoDimensionalEntity[EID, PosType], x PosType, y PosType)
> 设置对象移动到特定位置
**示例代码：**

```go

func ExampleTwoDimensional_MoveTo() {
	m := moving2.NewTwoDimensional(moving2.WithTwoDimensionalTimeUnit[int64, float64](time.Second))
	defer func() {
		m.Release()
	}()
	var wait sync.WaitGroup
	m.RegPosition2DDestinationEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println("done")
		wait.Done()
	})
	wait.Add(1)
	entity := NewEntity(1, 100)
	m.MoveTo(entity, 50, 30)
	wait.Wait()
}

```

***
#### func (*TwoDimensional) StopMove(id EID)
> 停止特定对象的移动
**示例代码：**

```go

func ExampleTwoDimensional_StopMove() {
	m := moving2.NewTwoDimensional(moving2.WithTwoDimensionalTimeUnit[int64, float64](time.Second))
	defer func() {
		m.Release()
	}()
	var wait sync.WaitGroup
	m.RegPosition2DChangeEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64], oldX, oldY float64) {
		fmt.Println("move")
	})
	m.RegPosition2DStopMoveEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println("stop")
		wait.Done()
	})
	m.RegPosition2DDestinationEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println("done")
		wait.Done()
	})
	wait.Add(1)
	entity := NewEntity(1, 100)
	m.MoveTo(entity, 50, 300)
	m.StopMove(1)
	wait.Wait()
}

```

<details>
<summary>查看 / 收起单元测试</summary>


```go

func TestTwoDimensional_StopMove(t *testing.T) {
	var wait sync.WaitGroup
	m := moving2.NewTwoDimensional(moving2.WithTwoDimensionalTimeUnit[int64, float64](time.Second))
	defer func() {
		m.Release()
	}()
	m.RegPosition2DChangeEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64], oldX, oldY float64) {
		x, y := entity.GetPosition().GetXY()
		fmt.Println(fmt.Sprintf("%d : %d | %f, %f > %f, %f", entity.GetTwoDimensionalEntityID(), time.Now().UnixMilli(), oldX, oldY, x, y))
	})
	m.RegPosition2DDestinationEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println(fmt.Sprintf("%d : %d | destination", entity.GetTwoDimensionalEntityID(), time.Now().UnixMilli()))
		wait.Done()
	})
	m.RegPosition2DStopMoveEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println(fmt.Sprintf("%d : %d | stop", entity.GetTwoDimensionalEntityID(), time.Now().UnixMilli()))
		wait.Done()
	})
	for i := 0; i < 10; i++ {
		wait.Add(1)
		entity := NewEntity(int64(i)+1, float64(10+i))
		m.MoveTo(entity, 50, 30)
	}
	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {
		m.StopMove(int64(i) + 1)
	}
	wait.Wait()
}

```


</details>


***
#### func (*TwoDimensional) RegPosition2DChangeEvent(handle Position2DChangeEventHandle[EID, PosType])
> 在对象位置改变时将执行注册的事件处理函数
***
#### func (*TwoDimensional) OnPosition2DChangeEvent(entity TwoDimensionalEntity[EID, PosType], oldX PosType, oldY PosType)
***
#### func (*TwoDimensional) RegPosition2DDestinationEvent(handle Position2DDestinationEventHandle[EID, PosType])
> 在对象到达终点时将执行被注册的事件处理函数
***
#### func (*TwoDimensional) OnPosition2DDestinationEvent(entity TwoDimensionalEntity[EID, PosType])
***
#### func (*TwoDimensional) RegPosition2DStopMoveEvent(handle Position2DStopMoveEventHandle[EID, PosType])
> 在对象停止移动时将执行被注册的事件处理函数
***
#### func (*TwoDimensional) OnPosition2DStopMoveEvent(entity TwoDimensionalEntity[EID, PosType])
***
#### func (*TwoDimensional) Release()
> 释放对象移动对象所占用的资源
***
<span id="struct_TwoDimensionalEntity"></span>
### TwoDimensionalEntity `INTERFACE`
2D移动对象接口定义
```go
type TwoDimensionalEntity[EID generic.Basic, PosType generic.SignedNumber] interface {
	GetTwoDimensionalEntityID() EID
	GetSpeed() float64
	GetPosition() geometry.Point[PosType]
	SetPosition(geometry.Point[PosType])
}
```
<span id="struct_Position2DChangeEventHandle"></span>
### Position2DChangeEventHandle `STRUCT`

```go
type Position2DChangeEventHandle[EID generic.Basic, PosType generic.SignedNumber] func(moving *TwoDimensional[EID, PosType], entity TwoDimensionalEntity[EID, PosType], oldX PosType)
```
<span id="struct_TwoDimensionalOption"></span>
### TwoDimensionalOption `STRUCT`

```go
type TwoDimensionalOption[EID generic.Basic, PosType generic.SignedNumber] func(moving *TwoDimensional[EID, PosType])
```
