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
|`STRUCT`|[RoomController](#struct_RoomController)|对房间进行操作的控制器，由 RoomManager 接管后返回
|`STRUCT`|[RoomManager](#struct_RoomManager)|房间管理器是用于对房间进行管理的基本单元，通过该实例可以对房间进行增删改查等操作
|`STRUCT`|[RoomAssumeControlEventHandle](#struct_RoomAssumeControlEventHandle)|暂无描述...
|`STRUCT`|[RoomControllerOptions](#struct_RoomControllerOptions)|暂无描述...

</details>


***
## 详情信息
#### func NewRoomManager\[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]\]() *RoomManager[EntityID, RoomID, Entity, Room]
<span id="NewRoomManager"></span>
> 创建房间管理器 RoomManager 的实例

**示例代码：**

```go

func ExampleNewRoomManager() {
	var rm = space.NewRoomManager[string, int64, *Player, *Room]()
	fmt.Println(rm == nil)
}

```

***
#### func NewRoomControllerOptions\[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]\]() *RoomControllerOptions[EntityID, RoomID, Entity, Room]
<span id="NewRoomControllerOptions"></span>
> 创建房间控制器选项

***
<span id="struct_RoomController"></span>
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
<span id="struct_RoomController_HasOwner"></span>

#### func (*RoomController) HasOwner()  bool
> 判断是否有房主

***
<span id="struct_RoomController_IsOwner"></span>

#### func (*RoomController) IsOwner(entityId EntityID)  bool
> 判断是否为房主

***
<span id="struct_RoomController_GetOwner"></span>

#### func (*RoomController) GetOwner()  Entity
> 获取房主

***
<span id="struct_RoomController_GetOwnerID"></span>

#### func (*RoomController) GetOwnerID()  EntityID
> 获取房主 ID

***
<span id="struct_RoomController_GetOwnerExist"></span>

#### func (*RoomController) GetOwnerExist() ( Entity,  bool)
> 获取房间，并返回房主是否存在的状态

***
<span id="struct_RoomController_SetOwner"></span>

#### func (*RoomController) SetOwner(entityId EntityID)
> 设置房主

***
<span id="struct_RoomController_DelOwner"></span>

#### func (*RoomController) DelOwner()
> 删除房主，将房间设置为无主的状态

***
<span id="struct_RoomController_JoinSeat"></span>

#### func (*RoomController) JoinSeat(entityId EntityID, seat ...int)  error
> 设置特定对象加入座位，当具体的座位不存在的时候，将会自动分配座位
>   - 当目标座位存在玩家或未添加到房间中的时候，将会返回错误

***
<span id="struct_RoomController_LeaveSeat"></span>

#### func (*RoomController) LeaveSeat(entityId EntityID)
> 离开座位

***
<span id="struct_RoomController_GetSeat"></span>

#### func (*RoomController) GetSeat(entityId EntityID)  int
> 获取座位

***
<span id="struct_RoomController_GetFirstNotEmptySeat"></span>

#### func (*RoomController) GetFirstNotEmptySeat()  int
> 获取第一个非空座位号，如果没有非空座位，将返回 UnknownSeat

***
<span id="struct_RoomController_GetFirstEmptySeatEntity"></span>

#### func (*RoomController) GetFirstEmptySeatEntity() (entity Entity)
> 获取第一个空座位上的实体，如果没有空座位，将返回空实体

***
<span id="struct_RoomController_GetRandomEntity"></span>

#### func (*RoomController) GetRandomEntity() (entity Entity)
> 获取随机实体，如果房间中没有实体，将返回空实体

***
<span id="struct_RoomController_GetNotEmptySeat"></span>

#### func (*RoomController) GetNotEmptySeat()  []int
> 获取非空座位

***
<span id="struct_RoomController_GetEmptySeat"></span>

#### func (*RoomController) GetEmptySeat()  []int
> 获取空座位
>   - 空座位需要在有对象离开座位后才可能出现

***
<span id="struct_RoomController_HasSeat"></span>

#### func (*RoomController) HasSeat(entityId EntityID)  bool
> 判断是否有座位

***
<span id="struct_RoomController_GetSeatEntityCount"></span>

#### func (*RoomController) GetSeatEntityCount()  int
> 获取座位上的实体数量

***
<span id="struct_RoomController_GetSeatEntities"></span>

#### func (*RoomController) GetSeatEntities()  map[EntityID]Entity
> 获取座位上的实体

***
<span id="struct_RoomController_GetSeatEntitiesByOrdered"></span>

#### func (*RoomController) GetSeatEntitiesByOrdered()  []Entity
> 有序的获取座位上的实体

***
<span id="struct_RoomController_GetSeatEntitiesByOrderedAndContainsEmpty"></span>

#### func (*RoomController) GetSeatEntitiesByOrderedAndContainsEmpty()  []Entity
> 获取有序的座位上的实体，包含空座位

***
<span id="struct_RoomController_GetSeatEntity"></span>

#### func (*RoomController) GetSeatEntity(seat int) (entity Entity)
> 获取座位上的实体

***
<span id="struct_RoomController_ContainEntity"></span>

#### func (*RoomController) ContainEntity(id EntityID)  bool
> 房间内是否包含实体

***
<span id="struct_RoomController_GetRoom"></span>

#### func (*RoomController) GetRoom()  Room
> 获取原始房间实例，该实例为被接管的房间的原始实例

***
<span id="struct_RoomController_GetEntities"></span>

#### func (*RoomController) GetEntities()  map[EntityID]Entity
> 获取所有实体

***
<span id="struct_RoomController_HasEntity"></span>

#### func (*RoomController) HasEntity(id EntityID)  bool
> 判断是否有实体

***
<span id="struct_RoomController_GetEntity"></span>

#### func (*RoomController) GetEntity(id EntityID)  Entity
> 获取实体

***
<span id="struct_RoomController_GetEntityExist"></span>

#### func (*RoomController) GetEntityExist(id EntityID) ( Entity,  bool)
> 获取实体，并返回实体是否存在的状态

***
<span id="struct_RoomController_GetEntityIDs"></span>

#### func (*RoomController) GetEntityIDs()  []EntityID
> 获取所有实体ID

***
<span id="struct_RoomController_GetEntityCount"></span>

#### func (*RoomController) GetEntityCount()  int
> 获取实体数量

***
<span id="struct_RoomController_ChangePassword"></span>

#### func (*RoomController) ChangePassword(password *string)
> 修改房间密码
>   - 当房间密码为 nil 时，将会取消密码

***
<span id="struct_RoomController_AddEntity"></span>

#### func (*RoomController) AddEntity(entity Entity)  error
> 添加实体，如果房间存在密码，应使用 AddEntityByPassword 函数进行添加，否则将始终返回 ErrRoomPasswordNotMatch 错误
>   - 当房间已满时，将会返回 ErrRoomFull 错误

***
<span id="struct_RoomController_AddEntityByPassword"></span>

#### func (*RoomController) AddEntityByPassword(entity Entity, password string)  error
> 通过房间密码添加实体到该房间中
>   - 当未设置房间密码时，password 参数将会被忽略
>   - 当房间密码不匹配时，将会返回 ErrRoomPasswordNotMatch 错误
>   - 当房间已满时，将会返回 ErrRoomFull 错误

***
<span id="struct_RoomController_RemoveEntity"></span>

#### func (*RoomController) RemoveEntity(id EntityID)
> 移除实体
>   - 当实体被移除时如果实体在座位上，将会自动离开座位
>   - 如果实体为房主，将会根据 RoomControllerOptions.WithOwnerInherit 函数的设置进行继承

***
<span id="struct_RoomController_RemoveAllEntities"></span>

#### func (*RoomController) RemoveAllEntities()
> 移除该房间中的所有实体
>   - 当实体被移除时如果实体在座位上，将会自动离开座位
>   - 如果实体为房主，将会根据 RoomControllerOptions.WithOwnerInherit 函数的设置进行继承

***
<span id="struct_RoomController_Destroy"></span>

#### func (*RoomController) Destroy()
> 销毁房间，房间会从 RoomManager 中移除，同时所有房间的实体、座位等数据都会被清空
>   - 该函数与 RoomManager.DestroyRoom 相同，RoomManager.DestroyRoom 函数为该函数的快捷方式

***
<span id="struct_RoomController_GetRoomManager"></span>

#### func (*RoomController) GetRoomManager()  *RoomManager[EntityID, RoomID, Entity, Room]
> 获取该房间控制器所属的房间管理器

***
<span id="struct_RoomController_GetRoomID"></span>

#### func (*RoomController) GetRoomID()  RoomID
> 获取房间 ID

***
<span id="struct_RoomController_Broadcast"></span>

#### func (*RoomController) Broadcast(handler func ( Entity), conditions ...func ( Entity)  bool)
> 广播，该函数会将所有房间中满足 conditions 的对象传入 handler 中进行处理

***
<span id="struct_RoomManager"></span>
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
<span id="struct_RoomManager_AssumeControl"></span>

#### func (*RoomManager) AssumeControl(room Room, options ...*RoomControllerOptions[EntityID, RoomID, Entity, Room])  *RoomController[EntityID, RoomID, Entity, Room]
> 将房间控制权交由 RoomManager 接管，返回 RoomController 实例
>   - 当任何房间需要被 RoomManager 管理时，都应该调用该方法获取到 RoomController 实例后进行操作
>   - 房间被接管后需要在释放房间控制权时调用 RoomController.Destroy 方法，否则将会导致 RoomManager 一直持有房间资源

**示例代码：**

```go

func ExampleRoomManager_AssumeControl() {
	var rm = space.NewRoomManager[string, int64, *Player, *Room]()
	var room = &Room{Id: 1}
	var controller = rm.AssumeControl(room)
	if err := controller.AddEntity(&Player{Id: "1"}); err != nil {
		panic(err)
	}
	fmt.Println(controller.GetEntityCount())
}

```

***
<span id="struct_RoomManager_DestroyRoom"></span>

#### func (*RoomManager) DestroyRoom(id RoomID)
> 销毁房间，该函数为 RoomController.Destroy 的快捷方式

***
<span id="struct_RoomManager_GetRoom"></span>

#### func (*RoomManager) GetRoom(id RoomID)  *RoomController[EntityID, RoomID, Entity, Room]
> 通过房间 ID 获取对应房间的控制器 RoomController，当房间不存在时将返回 nil

***
<span id="struct_RoomManager_GetRooms"></span>

#### func (*RoomManager) GetRooms()  map[RoomID]*RoomController[EntityID, RoomID, Entity, Room]
> 获取包含所有房间 ID 到对应控制器 RoomController 的映射
>   - 返回值的 map 为拷贝对象，可安全的对其进行增删等操作

***
<span id="struct_RoomManager_GetRoomCount"></span>

#### func (*RoomManager) GetRoomCount()  int
> 获取房间管理器接管的房间数量

***
<span id="struct_RoomManager_GetRoomIDs"></span>

#### func (*RoomManager) GetRoomIDs()  []RoomID
> 获取房间管理器接管的所有房间 ID

***
<span id="struct_RoomManager_HasEntity"></span>

#### func (*RoomManager) HasEntity(entityId EntityID)  bool
> 判断特定对象是否在任一房间中，当对象不在任一房间中时将返回 false

***
<span id="struct_RoomManager_GetEntityRooms"></span>

#### func (*RoomManager) GetEntityRooms(entityId EntityID)  map[RoomID]*RoomController[EntityID, RoomID, Entity, Room]
> 获取特定对象所在的房间，返回值为房间 ID 到对应控制器 RoomController 的映射
>   - 由于一个对象可能在多个房间中，因此返回值为 map 类型

***
<span id="struct_RoomManager_Broadcast"></span>

#### func (*RoomManager) Broadcast(handler func ( Entity), conditions ...func ( Entity)  bool)
> 向所有房间对象广播消息，该方法将会遍历所有房间控制器并调用 RoomController.Broadcast 方法

***
<span id="struct_RoomAssumeControlEventHandle"></span>
### RoomAssumeControlEventHandle `STRUCT`

```go
type RoomAssumeControlEventHandle[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] func(controller *RoomController[EntityID, RoomID, Entity, Room])
```
<span id="struct_RoomControllerOptions"></span>
### RoomControllerOptions `STRUCT`

```go
type RoomControllerOptions[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	maxEntityCount      *int
	password            *string
	ownerInherit        bool
	ownerInheritHandler func(controller *RoomController[EntityID, RoomID, Entity, Room]) *EntityID
}
```
<span id="struct_RoomControllerOptions_WithOwnerInherit"></span>

#### func (*RoomControllerOptions) WithOwnerInherit(inherit bool, inheritHandler ...func (controller *RoomController[EntityID, RoomID, Entity, Room])  *EntityID)  *RoomControllerOptions[EntityID, RoomID, Entity, Room]
> 设置房间所有者是否继承，默认为 false
>   - inherit: 是否继承，当未设置 inheritHandler 且 inherit 为 true 时，将会按照随机或根据座位号顺序继承房间所有者
>   - inheritHandler: 继承处理函数，当 inherit 为 true 时，该函数将会被调用，传入当前房间中的所有实体，返回值为新的房间所有者

***
<span id="struct_RoomControllerOptions_WithMaxEntityCount"></span>

#### func (*RoomControllerOptions) WithMaxEntityCount(maxEntityCount int)  *RoomControllerOptions[EntityID, RoomID, Entity, Room]
> 设置房间最大实体数量

***
<span id="struct_RoomControllerOptions_WithPassword"></span>

#### func (*RoomControllerOptions) WithPassword(password string)  *RoomControllerOptions[EntityID, RoomID, Entity, Room]
> 设置房间密码

***
