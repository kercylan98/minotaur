# Deck

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

deck 包中的内容用于针对一堆内容的管理，适用但不限于牌堆、麻将牌堆等情况。


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewDeck](#NewDeck)|创建一个新的甲板
|[NewGroup](#NewGroup)|创建一个新的组


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[Deck](#struct_Deck)|甲板，用于针对一堆 Group 进行管理的数据结构
|`STRUCT`|[Group](#struct_Group)|甲板中的组，用于针对一堆内容进行管理的数据结构
|`INTERFACE`|[Item](#struct_Item)|甲板成员

</details>


***
## 详情信息
#### func NewDeck\[I Item\]() *Deck[I]
<span id="NewDeck"></span>
> 创建一个新的甲板

***
#### func NewGroup\[I Item\](guid int64, fillHandle func (guid int64)  []I) *Group[I]
<span id="NewGroup"></span>
> 创建一个新的组

***
<span id="struct_Deck"></span>
### Deck `STRUCT`
甲板，用于针对一堆 Group 进行管理的数据结构
```go
type Deck[I Item] struct {
	groups map[int64]*Group[I]
	sort   []int64
}
```
<span id="struct_Deck_AddGroup"></span>

#### func (*Deck) AddGroup(group *Group[I])
> 将一个组添加到甲板中

***
<span id="struct_Deck_RemoveGroup"></span>

#### func (*Deck) RemoveGroup(guid int64)
> 移除甲板中的一个组

***
<span id="struct_Deck_GetCount"></span>

#### func (*Deck) GetCount()  int
> 获取甲板中的组数量

***
<span id="struct_Deck_GetGroups"></span>

#### func (*Deck) GetGroups()  map[int64]*Group[I]
> 获取所有组

***
<span id="struct_Deck_GetGroupsSlice"></span>

#### func (*Deck) GetGroupsSlice()  []*Group[I]
> 获取所有组

***
<span id="struct_Deck_GetNext"></span>

#### func (*Deck) GetNext(guid int64)  *Group[I]
> 获取特定组的下一个组

***
<span id="struct_Deck_GetPrev"></span>

#### func (*Deck) GetPrev(guid int64)  *Group[I]
> 获取特定组的上一个组

***
<span id="struct_Group"></span>
### Group `STRUCT`
甲板中的组，用于针对一堆内容进行管理的数据结构
```go
type Group[I Item] struct {
	guid       int64
	fillHandle func(guid int64) []I
	items      []I
}
```
<span id="struct_Group_GetGuid"></span>

#### func (*Group) GetGuid()  int64
> 获取组的 guid

***
<span id="struct_Group_Fill"></span>

#### func (*Group) Fill()
> 将该组的数据填充为 WithGroupFillHandle 中设置的内容

***
<span id="struct_Group_Pop"></span>

#### func (*Group) Pop() (item I)
> 从顶部获取一个内容

***
<span id="struct_Group_PopN"></span>

#### func (*Group) PopN(n int) (items []I)
> 从顶部获取指定数量的内容

***
<span id="struct_Group_PressOut"></span>

#### func (*Group) PressOut() (item I)
> 从底部压出一个内容

***
<span id="struct_Group_PressOutN"></span>

#### func (*Group) PressOutN(n int) (items []I)
> 从底部压出指定数量的内容

***
<span id="struct_Group_Push"></span>

#### func (*Group) Push(item I)
> 向顶部压入一个内容

***
<span id="struct_Group_PushN"></span>

#### func (*Group) PushN(items []I)
> 向顶部压入指定数量的内容

***
<span id="struct_Group_Insert"></span>

#### func (*Group) Insert(item I)
> 向底部插入一个内容

***
<span id="struct_Group_InsertN"></span>

#### func (*Group) InsertN(items []I)
> 向底部插入指定数量的内容

***
<span id="struct_Group_Pull"></span>

#### func (*Group) Pull(index int) (item I)
> 从特定位置拔出一个内容

***
<span id="struct_Group_Thrust"></span>

#### func (*Group) Thrust(index int, item I)
> 向特定位置插入一个内容

***
<span id="struct_Group_IsFree"></span>

#### func (*Group) IsFree()  bool
> 检查组是否为空

***
<span id="struct_Group_GetCount"></span>

#### func (*Group) GetCount()  int
> 获取组中剩余的内容数量

***
<span id="struct_Group_GetItem"></span>

#### func (*Group) GetItem(index int)  I
> 获取组中的指定内容

***
<span id="struct_Item"></span>
### Item `INTERFACE`
甲板成员
```go
type Item interface{}
```
