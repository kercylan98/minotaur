# Deck

deck 包中的内容用于针对一堆内容的管理，适用但不限于牌堆、麻将牌堆等情况。

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/deck)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

## 目录
列出了该 `package` 下所有的函数，可通过目录进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录</summary


> 包级函数定义

|函数|描述
|:--|:--
|[NewDeck](#NewDeck)|创建一个新的甲板
|[NewGroup](#NewGroup)|创建一个新的组


> 结构体定义

|结构体|描述
|:--|:--
|[Deck](#deck)|甲板，用于针对一堆 Group 进行管理的数据结构
|[Group](#group)|甲板中的组，用于针对一堆内容进行管理的数据结构
|[Item](#item)|甲板成员

</details>


#### func NewDeck()  *Deck[I]
<span id="NewDeck"></span>
> 创建一个新的甲板
***
#### func NewGroup(guid int64, fillHandle func (guid int64)  []I)  *Group[I]
<span id="NewGroup"></span>
> 创建一个新的组
***
### Deck
甲板，用于针对一堆 Group 进行管理的数据结构
```go
type Deck[I Item] struct {
	groups map[int64]*Group[I]
	sort   []int64
}
```
#### func (*Deck) AddGroup(group *Group[I])
> 将一个组添加到甲板中
***
#### func (*Deck) RemoveGroup(guid int64)
> 移除甲板中的一个组
***
#### func (*Deck) GetCount()  int
> 获取甲板中的组数量
***
#### func (*Deck) GetGroups()  map[int64]*Group[I]
> 获取所有组
***
#### func (*Deck) GetGroupsSlice()  []*Group[I]
> 获取所有组
***
#### func (*Deck) GetNext(guid int64)  *Group[I]
> 获取特定组的下一个组
***
#### func (*Deck) GetPrev(guid int64)  *Group[I]
> 获取特定组的上一个组
***
### Group
甲板中的组，用于针对一堆内容进行管理的数据结构
```go
type Group[I Item] struct {
	guid       int64
	fillHandle func(guid int64) []I
	items      []I
}
```
#### func (*Group) GetGuid()  int64
> 获取组的 guid
***
#### func (*Group) Fill()
> 将该组的数据填充为 WithGroupFillHandle 中设置的内容
***
#### func (*Group) Pop() (item I)
> 从顶部获取一个内容
***
#### func (*Group) PopN(n int) (items []I)
> 从顶部获取指定数量的内容
***
#### func (*Group) PressOut() (item I)
> 从底部压出一个内容
***
#### func (*Group) PressOutN(n int) (items []I)
> 从底部压出指定数量的内容
***
#### func (*Group) Push(item I)
> 向顶部压入一个内容
***
#### func (*Group) PushN(items []I)
> 向顶部压入指定数量的内容
***
#### func (*Group) Insert(item I)
> 向底部插入一个内容
***
#### func (*Group) InsertN(items []I)
> 向底部插入指定数量的内容
***
#### func (*Group) Pull(index int) (item I)
> 从特定位置拔出一个内容
***
#### func (*Group) Thrust(index int, item I)
> 向特定位置插入一个内容
***
#### func (*Group) IsFree()  bool
> 检查组是否为空
***
#### func (*Group) GetCount()  int
> 获取组中剩余的内容数量
***
#### func (*Group) GetItem(index int)  I
> 获取组中的指定内容
***
### Item
甲板成员
```go
type Item struct{}
```
