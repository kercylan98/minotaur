package room

import "github.com/kercylan98/minotaur/game"

// Room 房间类似于简版的游戏世界(World)，不过没有游戏实体
type Room[PlayerID comparable, P game.Player[PlayerID]] interface {
	// GetGuid 获取房间的唯一标识符
	GetGuid() int64
	// GetPlayerLimit 获取玩家人数上限
	GetPlayerLimit() int
	// GetPlayer 根据玩家id获取玩家
	GetPlayer(id PlayerID) P
	// GetPlayers 获取房间中的所有玩家
	GetPlayers() map[PlayerID]P
	// GetPlayerCount 获取玩家数量
	GetPlayerCount() int
	// IsExistPlayer 检查房间中是否存在特定玩家
	IsExistPlayer(id PlayerID) bool
	// IsOwner 检查玩家是否是房主
	IsOwner(id PlayerID) bool
	// ChangeOwner 设置玩家为房主
	ChangeOwner(id PlayerID)

	// Join 使特定玩家加入房间
	Join(player P) error
	// Leave 使特定玩家离开房间
	Leave(id PlayerID)
	// KickOut 将特定玩家踢出房间
	KickOut(id, target PlayerID, reason string) error
}
