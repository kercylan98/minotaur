package game

import "minotaur/utils/synchronization"

// Room 房间类似于简版的游戏世界，不过没有游戏实体
type Room[PlayerID comparable, P Player[PlayerID]] interface {
	// GetGuid 获取房间的唯一标识符
	GetGuid() int64
	// GetPlayerLimit 获取玩家人数上限
	GetPlayerLimit() int
	// GetPlayerWithConnID 根据连接ID获取玩家
	GetPlayerWithConnID(id string) P
	// GetPlayer 根据玩家id获取玩家
	GetPlayer(id PlayerID) P
	// GetPlayers 获取房间中的所有玩家
	GetPlayers() synchronization.MapReadonly[PlayerID, P]
	// IsExistPlayer 检查房间中是否存在特定玩家
	IsExistPlayer(id PlayerID) bool

	// Join 使特定玩家加入房间
	Join(player P) error
	// Leave 使特定玩家离开房间
	Leave(id PlayerID)

	// RegPlayerJoinRoomEvent 玩家进入房间时将立即执行被注册的事件处理函数
	RegPlayerJoinRoomEvent(handle PlayerJoinRoomEventHandle[PlayerID, P])
	OnPlayerJoinRoomEvent(player P)
	// RegPlayerLeaveRoomEvent 玩家离开房间时将立即执行被注册的事件处理函数
	RegPlayerLeaveRoomEvent(handle PlayerLeaveRoomEventHandle[PlayerID, P])
	OnPlayerLeaveRoomEvent(player P)
}

type (
	PlayerJoinRoomEventHandle[ID comparable, P Player[ID]]  func(room Room[ID, P], player P)
	PlayerLeaveRoomEventHandle[ID comparable, P Player[ID]] func(room Room[ID, P], player P)
)
