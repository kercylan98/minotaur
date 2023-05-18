package game

import (
	"github.com/kercylan98/minotaur/utils/hash"
)

// World 游戏世界接口定义
type World[PlayerID comparable, P Player[PlayerID]] interface {
	// GetGuid 获取世界的唯一标识符
	GetGuid() int64
	// GetPlayerLimit 获取玩家人数上限
	GetPlayerLimit() int
	// GetPlayer 根据玩家id获取玩家
	GetPlayer(id PlayerID) P
	// GetPlayers 获取世界中的所有玩家
	GetPlayers() hash.MapReadonly[PlayerID, P]
	// GetActor 根据唯一标识符获取世界中的游戏对象
	GetActor(guid int64) Actor
	// GetActors 获取世界中的所有游戏对象
	GetActors() hash.MapReadonly[int64, Actor]
	// GetPlayerActor 获取游戏世界中归属特定玩家的特定游戏对象
	GetPlayerActor(id PlayerID, guid int64) Actor
	// GetPlayerActors 获取游戏世界中归属特定玩家的所有游戏对象
	GetPlayerActors(id PlayerID) hash.MapReadonly[int64, Actor]
	// IsExistPlayer 检查游戏世界中是否存在特定玩家
	IsExistPlayer(id PlayerID) bool
	// IsExistActor 检查游戏世界中是否存在特定游戏对象
	IsExistActor(guid int64) bool
	// IsOwner 检查游戏世界中的特定游戏对象是否归属特定玩家
	IsOwner(id PlayerID, guid int64) bool

	// Join 使特定玩家加入游戏世界
	Join(player P) error
	// Leave 使特定玩家离开游戏世界
	Leave(id PlayerID)

	// AddActor 添加游戏对象
	AddActor(actor Actor)
	// RemoveActor 移除游戏对象
	RemoveActor(guid int64)
	// SetActorOwner 设置游戏对象归属玩家
	SetActorOwner(guid int64, id PlayerID)
	// RemoveActorOwner 移除游戏对象归属，置为无主的
	RemoveActorOwner(guid int64)

	// Reset 重置世界资源
	Reset()
	// Release 释放世界资源，释放后世界将不可用
	Release()

	// RegWorldResetEvent 世界被重置后将立即执行被注册的事件处理函数
	RegWorldResetEvent(handle WorldResetEventHandle[PlayerID, P])
	OnWorldResetEvent()
	// RegWorldReleaseEvent 直接被释放前将立即执行被注册的事件处理函数，此刻世界仍然可用
	RegWorldReleaseEvent(handle WorldReleaseEventHandle[PlayerID, P])
	OnWorldReleaseEvent()
	// RegPlayerJoinWorldEvent 玩家进入世界时将立即执行被注册的事件处理函数
	RegPlayerJoinWorldEvent(handle PlayerJoinWorldEventHandle[PlayerID, P])
	OnPlayerJoinWorldEvent(player P)
	// RegPlayerLeaveWorldEvent 玩家离开世界时将立即执行被注册的事件处理函数
	RegPlayerLeaveWorldEvent(handle PlayerLeaveWorldEventHandle[PlayerID, P])
	OnPlayerLeaveWorldEvent(player P)
	// RegActorGeneratedEvent 游戏世界中的游戏对象生成完成时将立即执行被注册的事件处理函数
	RegActorGeneratedEvent(handle ActorGeneratedEventHandle[PlayerID, P])
	OnActorGeneratedEvent(actor Actor)
	// RegActorAnnihilationEvent 游戏世界中的游戏对象被移除前执行被注册的事件处理函数
	RegActorAnnihilationEvent(handle ActorAnnihilationEventHandle[PlayerID, P])
	OnActorAnnihilationEvent(actor Actor)
	// RegActorOwnerChangeEvent 游戏对象的归属被改变时立刻执行被注册的事件处理函数
	RegActorOwnerChangeEvent(handle ActorOwnerChangeEventHandle[PlayerID, P])
	OnActorOwnerChangeEvent(actor Actor, old, new PlayerID, isolated bool)
}

type (
	WorldResetEventHandle[ID comparable, P Player[ID]]        func(world World[ID, P])
	WorldReleaseEventHandle[ID comparable, P Player[ID]]      func(world World[ID, P])
	PlayerJoinWorldEventHandle[ID comparable, P Player[ID]]   func(world World[ID, P], player P)
	PlayerLeaveWorldEventHandle[ID comparable, P Player[ID]]  func(world World[ID, P], player P)
	ActorGeneratedEventHandle[ID comparable, P Player[ID]]    func(world World[ID, P], actor Actor)
	ActorAnnihilationEventHandle[ID comparable, P Player[ID]] func(world World[ID, P], actor Actor)
	ActorOwnerChangeEventHandle[ID comparable, P Player[ID]]  func(world World[ID, P], actor Actor, old, new ID, isolated bool)
)
