package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/asynchronous"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/log"
	"go.uber.org/zap"
)

// NewRoom 创建一个默认的内置游戏房间 Room
func NewRoom[PlayerID comparable, Player game.Player[PlayerID]](guid int64, options ...RoomOption[PlayerID, Player]) *Room[PlayerID, Player] {
	room := &Room[PlayerID, Player]{
		guid:    guid,
		players: asynchronous.NewMap[PlayerID, Player](),
	}
	for _, option := range options {
		option(room)
	}
	return room
}

// Room 默认的内置游戏房间实现
//   - 实现了最大人数控制、房主、踢出玩家、玩家维护等功能
//   - 支持并发安全和非并发安全的模式
type Room[PlayerID comparable, Player game.Player[PlayerID]] struct {
	guid            int64
	owner           PlayerID
	noMaster        bool
	playerLimit     int
	players         hash.Map[PlayerID, Player]
	kickCheckHandle func(room *Room[PlayerID, Player], id, target PlayerID) error

	playerJoinRoomEventHandles  []game.PlayerJoinRoomEventHandle[PlayerID, Player]
	playerLeaveRoomEventHandles []game.PlayerLeaveRoomEventHandle[PlayerID, Player]
	playerKickedOutEventHandles []game.PlayerKickedOutEventHandle[PlayerID, Player]
}

// GetGuid 获取房间唯一标识
func (slf *Room[PlayerID, Player]) GetGuid() int64 {
	return slf.guid
}

// GetPlayerLimit 获取最大玩家上限
func (slf *Room[PlayerID, Player]) GetPlayerLimit() int {
	return slf.playerLimit
}

// GetPlayer 根据玩家id获取玩家
func (slf *Room[PlayerID, Player]) GetPlayer(id PlayerID) Player {
	return slf.players.Get(id)
}

// GetPlayers 获取所有玩家
func (slf *Room[PlayerID, Player]) GetPlayers() hash.MapReadonly[PlayerID, Player] {
	return slf.players.(hash.MapReadonly[PlayerID, Player])
}

// GetPlayerCount 获取玩家数量
func (slf *Room[PlayerID, Player]) GetPlayerCount() int {
	return slf.players.Size()
}

// IsExistPlayer 房间内是否存在某玩家
func (slf *Room[PlayerID, Player]) IsExistPlayer(id PlayerID) bool {
	return slf.players.Exist(id)
}

// IsOwner 检查特定玩家是否是房主
//   - 当房间为无主模式(WithRoomNoMaster)时，将会始终返回false
func (slf *Room[PlayerID, Player]) IsOwner(id PlayerID) bool {
	return !slf.noMaster && slf.owner == id
}

// ChangeOwner 改变房主
//   - 当房间为无主模式(WithRoomNoMaster)时，将不会发生任何变化
func (slf *Room[PlayerID, Player]) ChangeOwner(id PlayerID) {
	if slf.noMaster || slf.owner == id {
		return
	}
	slf.owner = id
}

// Join 控制玩家加入到该房间
func (slf *Room[PlayerID, Player]) Join(player Player) error {
	playerId := player.GetID()
	exist := slf.players.Exist(playerId)
	if !exist && slf.players.Size() >= slf.playerLimit && slf.playerLimit > 0 {
		return ErrRoomPlayerLimit
	}
	slf.players.Set(playerId, player)
	if !exist {
		log.Debug("Room.Join", zap.Any("guid", slf.GetGuid()), zap.Any("player", playerId))
		if slf.players.Size() == 1 && !slf.noMaster {
			slf.owner = playerId
		}
		slf.OnPlayerJoinRoomEvent(player)
	}
	return nil
}

// Leave 控制玩家离开房间
func (slf *Room[PlayerID, Player]) Leave(id PlayerID) {
	player, exist := slf.players.GetExist(id)
	if !exist {
		return
	}
	log.Debug("Room.Leave", zap.Any("guid", slf.GetGuid()), zap.Any("player", id))
	slf.OnPlayerLeaveRoomEvent(player)
	slf.players.Delete(id)
}

// KickOut 以某种原因踢出特定玩家
//   - 当设置了房间踢出玩家的检查处理函数(WithRoomKickPlayerCheckHandle)时，将会根据检查结果进行处理，即便是无主模式。其他情况如下
//   - 如果是无主模式(WithRoomNoMaster)，将会返回错误
//   - 如果不是房主发起的踢出玩家，将会返回错误
func (slf *Room[PlayerID, Player]) KickOut(id, target PlayerID, reason string) error {
	player, exist := slf.players.GetExist(target)
	if !exist {
		return nil
	}
	if slf.kickCheckHandle != nil {
		if err := slf.kickCheckHandle(slf, id, target); err != nil {
			return err
		}
	} else if slf.noMaster {
		return ErrRoomNoHasMaster
	} else if slf.owner != id {
		return ErrRoomNotIsOwner
	}

	slf.OnPlayerKickedOutEvent(id, target, reason)
	slf.Leave(player.GetID())
	return nil
}

// RegPlayerJoinRoomEvent 玩家进入房间时将立即执行被注册的事件处理函数
func (slf *Room[PlayerID, Player]) RegPlayerJoinRoomEvent(handle game.PlayerJoinRoomEventHandle[PlayerID, Player]) {
	slf.playerJoinRoomEventHandles = append(slf.playerJoinRoomEventHandles, handle)
}

func (slf *Room[PlayerID, Player]) OnPlayerJoinRoomEvent(player Player) {
	for _, handle := range slf.playerJoinRoomEventHandles {
		handle(slf, player)
	}
}

// RegPlayerLeaveRoomEvent 玩家离开房间时将立即执行被注册的事件处理函数
func (slf *Room[PlayerID, Player]) RegPlayerLeaveRoomEvent(handle game.PlayerLeaveRoomEventHandle[PlayerID, Player]) {
	slf.playerLeaveRoomEventHandles = append(slf.playerLeaveRoomEventHandles, handle)
}

func (slf *Room[PlayerID, Player]) OnPlayerLeaveRoomEvent(player Player) {
	for _, handle := range slf.playerLeaveRoomEventHandles {
		handle(slf, player)
	}
}

// RegPlayerKickedOutEvent 当玩家被踢出游戏时将立即执行被注册的事件处理函数
func (slf *Room[PlayerID, Player]) RegPlayerKickedOutEvent(handle game.PlayerKickedOutEventHandle[PlayerID, Player]) {
	slf.playerKickedOutEventHandles = append(slf.playerKickedOutEventHandles, handle)
}

func (slf *Room[PlayerID, Player]) OnPlayerKickedOutEvent(executor, kicked PlayerID, reason string) {
	for _, handle := range slf.playerKickedOutEventHandles {
		handle(slf, executor, kicked, reason)
	}
}
