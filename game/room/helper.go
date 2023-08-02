package room

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/slice"
)

// NewHelper 创建房间助手
func NewHelper[PID comparable, P game.Player[PID], R Room](manager *Manager[PID, P, R], room R) *Helper[PID, P, R] {
	return &Helper[PID, P, R]{
		m:    manager,
		room: room,
		Seat: manager.GetSeatInfo(room.GetGuid()),
	}
}

// Helper 易于快捷使用的房间助手
type Helper[PID comparable, P game.Player[PID], R Room] struct {
	m *Manager[PID, P, R]
	*Seat[PID, P, R]
	room R
}

// Room 获取房间
func (slf *Helper[PID, P, R]) Room() R {
	return slf.room
}

// GetPlayer 获取玩家
func (slf *Helper[PID, P, R]) GetPlayer(playerId PID) P {
	return slf.m.GetRoomPlayer(slf.room.GetGuid(), playerId)
}

// GetPlayers 获取房间中的所有玩家
func (slf *Helper[PID, P, R]) GetPlayers() map[PID]P {
	return slf.m.GetRoomPlayers(slf.room.GetGuid())
}

// GetPlayerCount 获取房间中的玩家数量
func (slf *Helper[PID, P, R]) GetPlayerCount() int {
	return slf.m.GetRoomPlayerCount(slf.room.GetGuid())
}

// GetPlayerLimit 获取房间中的玩家数量上限
func (slf *Helper[PID, P, R]) GetPlayerLimit() int {
	return slf.m.GetRoomPlayerLimit(slf.room.GetGuid())
}

// GetPlayerRooms 获取玩家所在的所有房间
func (slf *Helper[PID, P, R]) GetPlayerRooms(playerId PID) map[int64]R {
	return slf.m.GetPlayerRooms(playerId)
}

// GetPlayerRoomHelpers 获取玩家所在的所有房间助手
func (slf *Helper[PID, P, R]) GetPlayerRoomHelpers(playerId PID) map[int64]*Helper[PID, P, R] {
	return slf.m.GetPlayerRoomHelpers(playerId)
}

// GetPlayersSlice 获取房间中的所有玩家
func (slf *Helper[PID, P, R]) GetPlayersSlice() []P {
	seat := slf.GetSeatInfoMap()
	var players = make([]P, 0, len(seat))
	for _, v := range seat {
		players = append(players, slf.GetPlayer(v))
	}
	return players
}

// Broadcast 向房间中的所有玩家广播消息
func (slf *Helper[PID, P, R]) Broadcast(handle func(player P), except ...PID) {
	var exceptMap = slice.ToSet(except)
	for _, player := range slf.GetPlayers() {
		if hash.Exist(exceptMap, player.GetID()) {
			continue
		}
		handle(player)
	}
}

// BroadcastSeat 向房间中所有座位上的玩家广播消息
func (slf *Helper[PID, P, R]) BroadcastSeat(handle func(player P), except ...PID) {
	var exceptMap = slice.ToSet(except)
	for _, playerId := range slf.GetSeatInfoMap() {
		if hash.Exist(exceptMap, playerId) {
			continue
		}
		handle(slf.GetPlayer(playerId))
	}
}

// SetPlayerLimit 设置房间中的玩家数量上限
func (slf *Helper[PID, P, R]) SetPlayerLimit(limit int) {
	slf.m.SetPlayerLimit(slf.room.GetGuid(), limit)
}

// GetPlayerIDs 获取房间中的所有玩家ID
func (slf *Helper[PID, P, R]) GetPlayerIDs() []PID {
	return hash.KeyToSlice(slf.GetPlayers())
}

// SetOwner 设置房主
func (slf *Helper[PID, P, R]) SetOwner(playerId PID) {
	slf.m.SetOwner(slf.room.GetGuid(), playerId)
}

// CancelOwner 取消房主
func (slf *Helper[PID, P, R]) CancelOwner() {
	slf.manager.CancelOwner(slf.room.GetGuid())
}

// HasPlayer 是否有玩家
func (slf *Helper[PID, P, R]) HasPlayer(playerId PID) bool {
	return slf.m.InRoom(slf.room.GetGuid(), playerId)
}

// IsFull 房间是否已满
func (slf *Helper[PID, P, R]) IsFull() bool {
	return slf.GetPlayerCount() == slf.GetPlayerLimit()
}

// IsEmpty 房间是否为空
func (slf *Helper[PID, P, R]) IsEmpty() bool {
	return slf.GetPlayerCount() == 0
}

// GetRemainder 获取房间还可以容纳多少玩家
func (slf *Helper[PID, P, R]) GetRemainder() int {
	return slf.GetPlayerLimit() - slf.GetPlayerCount()
}

// IsOwner 是否是房主
func (slf *Helper[PID, P, R]) IsOwner(playerId PID) bool {
	return slf.m.IsOwner(slf.room.GetGuid(), playerId)
}

// HasOwner 是否有房主
func (slf *Helper[PID, P, R]) HasOwner() bool {
	return slf.m.HasOwner(slf.room.GetGuid())
}

// GetOwner 获取房主
func (slf *Helper[PID, P, R]) GetOwner() P {
	return slf.m.GetOwner(slf.room.GetGuid())
}

// Join 加入房间
func (slf *Helper[PID, P, R]) Join(player P) error {
	return slf.m.Join(slf.room.GetGuid(), player)
}

// Leave 离开房间
func (slf *Helper[PID, P, R]) Leave(player P) {
	slf.m.Leave(slf.room.GetGuid(), player)
}

// KickOut 踢出房间
func (slf *Helper[PID, P, R]) KickOut(executor, kicked PID, reason string) error {
	return slf.m.KickOut(slf.room.GetGuid(), executor, kicked, reason)
}
