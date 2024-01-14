# Arrangement

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/arrangement)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

arrangement 包提供了一些有用的函数来处理数组的排列。

更多的详细信息和使用示例，可以参考每个函数的文档。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[WithAreaConstraint](#WithAreaConstraint)|设置编排区域的约束条件
|[WithAreaConflict](#WithAreaConflict)|设置编排区域的冲突条件，冲突处理函数需要返回造成冲突的成员列表
|[WithAreaEvaluate](#WithAreaEvaluate)|设置编排区域的评估函数
|[NewArrangement](#NewArrangement)|创建一个新的编排
|[WithItemFixed](#WithItemFixed)|设置成员的固定编排区域
|[WithItemPriority](#WithItemPriority)|设置成员的优先级
|[WithItemNotAllow](#WithItemNotAllow)|设置成员不允许的编排区域
|[WithRetryThreshold](#WithRetryThreshold)|设置编排时的重试阈值
|[WithConstraintHandle](#WithConstraintHandle)|设置编排时触发约束时的处理函数
|[WithConflictHandle](#WithConflictHandle)|设置编排时触发冲突时的处理函数


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Area](#area)|编排区域
|`STRUCT`|[AreaOption](#areaoption)|编排区域选项
|`STRUCT`|[AreaConstraintHandle](#areaconstrainthandle)|暂无描述...
|`STRUCT`|[Arrangement](#arrangement)|用于针对多条数据进行合理编排的数据结构
|`STRUCT`|[Editor](#editor)|提供了大量辅助函数的编辑器
|`INTERFACE`|[Item](#item)|编排成员
|`STRUCT`|[ItemOption](#itemoption)|编排成员选项
|`STRUCT`|[ItemFixedAreaHandle](#itemfixedareahandle)|暂无描述...
|`STRUCT`|[Option](#option)|编排选项
|`STRUCT`|[ConstraintHandle](#constrainthandle)|暂无描述...

</details>


***
## 详情信息
#### func WithAreaConstraint(constraint AreaConstraintHandle[ID, AreaInfo])  AreaOption[ID, AreaInfo]
<span id="WithAreaConstraint"></span>
> 设置编排区域的约束条件
>   - 该约束用于判断一个成员是否可以被添加到该编排区域中
>   - 与 WithAreaConflict 不同的是，约束通常用于非成员关系导致的硬性约束，例如：成员的等级过滤、成员的性别等

***
#### func WithAreaConflict(conflict AreaConflictHandle[ID, AreaInfo])  AreaOption[ID, AreaInfo]
<span id="WithAreaConflict"></span>
> 设置编排区域的冲突条件，冲突处理函数需要返回造成冲突的成员列表
>   - 该冲突用于判断一个成员是否可以被添加到该编排区域中
>   - 与 WithAreaConstraint 不同的是，冲突通常用于成员关系导致的软性约束，例如：成员的职业唯一性、成员的种族唯一性等

***
#### func WithAreaEvaluate(evaluate AreaEvaluateHandle[ID, AreaInfo])  AreaOption[ID, AreaInfo]
<span id="WithAreaEvaluate"></span>
> 设置编排区域的评估函数
>   - 该评估函数将影响成员被编入区域的优先级

***
#### func NewArrangement(options ...Option[ID, AreaInfo])  *Arrangement[ID, AreaInfo]
<span id="NewArrangement"></span>
> 创建一个新的编排

***
#### func WithItemFixed(matcher ItemFixedAreaHandle[AreaInfo])  ItemOption[ID, AreaInfo]
<span id="WithItemFixed"></span>
> 设置成员的固定编排区域

***
#### func WithItemPriority(priority ItemPriorityHandle[ID, AreaInfo])  ItemOption[ID, AreaInfo]
<span id="WithItemPriority"></span>
> 设置成员的优先级

***
#### func WithItemNotAllow(verify ItemNotAllowVerifyHandle[ID, AreaInfo])  ItemOption[ID, AreaInfo]
<span id="WithItemNotAllow"></span>
> 设置成员不允许的编排区域

***
#### func WithRetryThreshold(threshold int)  Option[ID, AreaInfo]
<span id="WithRetryThreshold"></span>
> 设置编排时的重试阈值
>   - 当每一轮编排结束任有成员未被编排时，将会进行下一轮编排，直到编排次数达到该阈值
>   - 默认的阈值为 10 次

***
#### func WithConstraintHandle(handle ConstraintHandle[ID, AreaInfo])  Option[ID, AreaInfo]
<span id="WithConstraintHandle"></span>
> 设置编排时触发约束时的处理函数
>   - 当约束条件触发时，将会调用该函数。如果无法在该函数中处理约束，应该继续返回 err，尝试进行下一层的约束处理
>   - 当该函数的返回值为 nil 时，表示约束已经被处理，将会命中当前的编排区域
>   - 当所有的约束处理函数都无法处理约束时，将会进入下一个编排区域的尝试，如果均无法完成，将会将该成员加入到编排队列的末端，等待下一次编排
> 
> 有意思的是，硬性约束应该永远是无解的，而当需要进行一些打破规则的操作时，则可以透过该函数传入的 editor 进行操作

***
#### func WithConflictHandle(handle ConflictHandle[ID, AreaInfo])  Option[ID, AreaInfo]
<span id="WithConflictHandle"></span>
> 设置编排时触发冲突时的处理函数
>   - 当冲突条件触发时，将会调用该函数。如果无法在该函数中处理冲突，应该继续返回这一批成员，尝试进行下一层的冲突处理
>   - 当该函数的返回值长度为 0 时，表示冲突已经被处理，将会命中当前的编排区域
>   - 当所有的冲突处理函数都无法处理冲突时，将会进入下一个编排区域的尝试，如果均无法完成，将会将该成员加入到编排队列的末端，等待下一次编排

***
### Area `STRUCT`
编排区域
```go
type Area[ID comparable, AreaInfo any] struct {
	info        AreaInfo
	items       map[ID]Item[ID]
	constraints []AreaConstraintHandle[ID, AreaInfo]
	conflicts   []AreaConflictHandle[ID, AreaInfo]
	evaluate    AreaEvaluateHandle[ID, AreaInfo]
}
```
### AreaOption `STRUCT`
编排区域选项
```go
type AreaOption[ID comparable, AreaInfo any] func(area *Area[ID, AreaInfo])
```
### AreaConstraintHandle `STRUCT`

```go
type AreaConstraintHandle[ID comparable, AreaInfo any] func(area *Area[ID, AreaInfo], item Item[ID]) error
```
### Arrangement `STRUCT`
用于针对多条数据进行合理编排的数据结构
  - 我不知道这个数据结构的具体用途，但是我觉得这个数据结构应该是有用的
  - 目前我能想到的用途只有我的过往经历：排课
  - 如果是在游戏领域，或许适用于多人小队匹配编排等类似情况
```go
type Arrangement[ID comparable, AreaInfo any] struct {
	areas             []*Area[ID, AreaInfo]
	items             map[ID]Item[ID]
	fixed             map[ID]ItemFixedAreaHandle[AreaInfo]
	priority          map[ID][]ItemPriorityHandle[ID, AreaInfo]
	itemNotAllow      map[ID][]ItemNotAllowVerifyHandle[ID, AreaInfo]
	threshold         int
	constraintHandles []ConstraintHandle[ID, AreaInfo]
	conflictHandles   []ConflictHandle[ID, AreaInfo]
}
```
### Editor `STRUCT`
提供了大量辅助函数的编辑器
```go
type Editor[ID comparable, AreaInfo any] struct {
	a          *Arrangement[ID, AreaInfo]
	pending    []Item[ID]
	fails      []Item[ID]
	falls      map[ID]struct{}
	retryCount int
}
```
### Item `INTERFACE`
编排成员
```go
type Item[ID comparable] interface {
	GetID() ID
	Equal(item Item[ID]) bool
}
```
### ItemOption `STRUCT`
编排成员选项
```go
type ItemOption[ID comparable, AreaInfo any] func(arrangement *Arrangement[ID, AreaInfo], item Item[ID])
```
### ItemFixedAreaHandle `STRUCT`

```go
type ItemFixedAreaHandle[AreaInfo any] func(areaInfo AreaInfo) bool
```
### Option `STRUCT`
编排选项
```go
type Option[ID comparable, AreaInfo any] func(arrangement *Arrangement[ID, AreaInfo])
```
### ConstraintHandle `STRUCT`

```go
type ConstraintHandle[ID comparable, AreaInfo any] func(editor *Editor[ID, AreaInfo], area *Area[ID, AreaInfo], item Item[ID], err error) error
```
