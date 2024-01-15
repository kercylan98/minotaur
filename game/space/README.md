# Space

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur)
![](https://img.shields.io/badge/Email-kercylan@gmail.com-green.svg?style=flat)

space 游戏中常见的空间设计，例如房间、地图等


## 目录导航
列出了该 `package` 下所有的函数及类型定义，可通过目录导航进行快捷跳转 ❤️
<details>
<summary>展开 / 折叠目录导航</summary>


> 包级函数定义

|函数名称|描述
|:--|:--
|[NewRoomManager](#NewRoomManager)|创建房间管理器 RoomManager 的实例
|[NewRoomControllerOptions](#NewRoomControllerOptions)|创建房间控制器选项


> 类型定义

|类型|名称|描述
|:--|:--|:--
|`STRUCT`|[RoomController](#roomcontroller)|对房间进行操作的控制器，由 RoomManager 接管后返回
|`STRUCT`|[RoomManager](#roommanager)|房间管理器是用于对房间进行管理的基本单元，通过该实例可以对房间进行增删改查等操作
|`STRUCT`|[RoomAssumeControlEventHandle](#roomassumecontroleventhandle)|暂无描述...
|`STRUCT`|[RoomControllerOptions](#roomcontrolleroptions)|暂无描述...

</details>


***
## 详情信息
#### func NewRoomManager() *RoomManager[EntityID, RoomID, Entity, Room]
<span id="NewRoomManager"></span>
> 创建房间管理器 RoomManager 的实例

示例代码：
```go

func ExampleNewRoomManager() {
	var rm = space.NewRoomManager[string, int64, *Player, *Room]()
	fmt.Println(rm == nil)
}

```

***
#### func NewRoomControllerOptions() *RoomControllerOptions[EntityID, RoomID, Entity, Room]
<span id="NewRoomControllerOptions"></span>
> 创建房间控制器选项

***
### RoomController `STRUCT`
对房间进行操作的控制器，由 RoomManager 接管后返回
```go
type RoomController[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	manager         *RoomManager[EntityID, RoomID, Entity, Room]
	options         *RoomControllerOptions[EntityID, RoomID, Entity, Room]
	room            Room
	entities        map[EntityID]Entity
	entitiesRWMutex sync.RWMutex
	vacancy         []int
	seat            []*EntityID
	owner           *EntityID
}
```
### RoomManager `STRUCT`
房间管理器是用于对房间进行管理的基本单元，通过该实例可以对房间进行增删改查等操作
  - 该实例是线程安全的
```go
type RoomManager[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	*roomManagerEvents[EntityID, RoomID, Entity, Room]
	roomsRWMutex sync.RWMutex
	rooms        map[RoomID]*RoomController[EntityID, RoomID, Entity, Room]
}
```
### RoomAssumeControlEventHandle `STRUCT`

```go
type RoomAssumeControlEventHandle[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] func(controller *RoomController[EntityID, RoomID, Entity, Room])
```
### RoomControllerOptions `STRUCT`

```go
type RoomControllerOptions[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	maxEntityCount      *int
	password            *string
	ownerInherit        bool
	ownerInheritHandler func(controller *RoomController[EntityID, RoomID, Entity, Room]) *EntityID
}
```
