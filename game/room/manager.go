package room

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/concurrent"
)

// NewManager 创建房间管理器
func NewManager[PID comparable, P game.Player[PID], R Room[PID, P]]() *Manager[PID, P, R] {
	manager := &Manager[PID, P, R]{
		event:   newEvent[PID, P, R](),
		rooms:   concurrent.NewBalanceMap[int64, *info[PID, P, R]](),
		players: concurrent.NewBalanceMap[PID, P](),
		pr:      concurrent.NewBalanceMap[PID, map[int64]struct{}](),
	}

	return manager
}

// Manager 房间管理器
type Manager[PID comparable, P game.Player[PID], R Room[PID, P]] struct {
	*event[PID, P, R]
	rooms   *concurrent.BalanceMap[int64, *info[PID, P, R]] // 所有房间
	players *concurrent.BalanceMap[PID, P]                  // 所有加入房间的玩家
	pr      *concurrent.BalanceMap[PID, map[int64]struct{}] // 玩家所在房间
	rp      *concurrent.BalanceMap[int64, map[PID]struct{}] // 房间中的玩家

}

// CreateRoom 创建房间
func (slf *Manager[PID, P, R]) CreateRoom(room R) {
	roomInfo := &info[PID, P, R]{room: room}
	slf.rooms.Set(room.GetGuid(), roomInfo)
}

// ReleaseRoom 释放房间
func (slf *Manager[PID, P, R]) ReleaseRoom(guid int64) {
	slf.unReg(guid)
	slf.rooms.Delete(guid)
}

// GetRoom 获取房间
func (slf *Manager[PID, P, R]) GetRoom(guid int64) R {
	return slf.rooms.Get(guid).room
}

// Exist 检查房间是否存在
func (slf *Manager[PID, P, R]) Exist(guid int64) bool {
	return slf.rooms.Exist(guid)
}

// GetRooms 获取所有房间
func (slf *Manager[PID, P, R]) GetRooms() map[int64]R {
	var rooms = make(map[int64]R)
	slf.rooms.Atom(func(m map[int64]*info[PID, P, R]) {
		for id, info := range m {
			rooms[id] = info.room
		}
	})
	return rooms
}

// GetRoomCount 获取房间数量
func (slf *Manager[PID, P, R]) GetRoomCount() int {
	return slf.rooms.Size()
}

// GetRoomPlayerCount 获取房间中玩家数量
func (slf *Manager[PID, P, R]) GetRoomPlayerCount(guid int64) int {
	var count int
	slf.rp.Atom(func(m map[int64]map[PID]struct{}) {
		count = len(m[guid])
	})
	return count
}

// ExistPlayer 检查玩家是否在任一房间内
func (slf *Manager[PID, P, R]) ExistPlayer(id PID) bool {
	return slf.players.Exist(id)
}

// InRoom 检查玩家是否在指定房间内
func (slf *Manager[PID, P, R]) InRoom(id PID, guid int64) bool {
	var in bool
	slf.pr.Atom(func(m map[PID]map[int64]struct{}) {
		rooms, exist := m[id]
		if !exist {
			return
		}
		_, in = rooms[guid]
	})
	return in
}

// GetPlayer 获取玩家
func (slf *Manager[PID, P, R]) GetPlayer(id PID) P {
	return slf.players.Get(id)
}

// GetPlayers 获取所有玩家
func (slf *Manager[PID, P, R]) GetPlayers() *concurrent.BalanceMap[PID, P] {
	return slf.players
}

// GetPlayerCount 获取玩家数量
func (slf *Manager[PID, P, R]) GetPlayerCount() int {
	return slf.players.Size()
}

// GetPlayerRoom 获取玩家所在房间
func (slf *Manager[PID, P, R]) GetPlayerRoom(id PID) []R {
	var result = make([]R, 0)
	slf.pr.Atom(func(m map[PID]map[int64]struct{}) {
		rooms, exist := m[id]
		if !exist {
			return
		}
		for id := range rooms {
			result = append(result, slf.rooms.Get(id).room)
		}
	})
	return result
}

// GetPlayerRoomCount 获取玩家所在房间数量
func (slf *Manager[PID, P, R]) GetPlayerRoomCount(id PID) int {
	var count int
	slf.pr.Atom(func(m map[PID]map[int64]struct{}) {
		count = len(m[id])
	})
	return count
}

// GetRoomPlayer 获取房间中的玩家
func (slf *Manager[PID, P, R]) GetRoomPlayer(roomId int64, playerId PID) P {
	var player P
	slf.rp.Atom(func(m map[int64]map[PID]struct{}) {
		players, exist := m[roomId]
		if !exist {
			return
		}
		_, exist = players[playerId]
		if !exist {
			return
		}
		player = slf.players.Get(playerId)
	})
	return player
}

// GetRoomPlayers 获取房间中的玩家
func (slf *Manager[PID, P, R]) GetRoomPlayers(guid int64) map[PID]P {
	var result = make(map[PID]P)
	slf.rp.Atom(func(m map[int64]map[PID]struct{}) {
		players, exist := m[guid]
		if !exist {
			return
		}
		for id := range players {
			result[id] = slf.players.Get(id)
		}
	})
	return result
}

// GetRoomPlayerLimit 获取房间中的玩家数量上限
func (slf *Manager[PID, P, R]) GetRoomPlayerLimit(guid int64) int {
	return slf.rooms.Get(guid).playerLimit
}

// Leave 使玩家离开房间
func (slf *Manager[PID, P, R]) Leave(roomId int64, player P) {
	slf.rooms.Atom(func(m map[int64]*info[PID, P, R]) {
		room, exist := m[roomId]
		if !exist {
			return
		}
		slf.OnPlayerLeaveRoomEvent(room.room, player)
		slf.pr.Atom(func(m map[PID]map[int64]struct{}) {
			rooms, exist := m[player.GetID()]
			if !exist {
				return
			}
			delete(rooms, roomId)
		})
		slf.rp.Atom(func(m map[int64]map[PID]struct{}) {
			players, exist := m[roomId]
			if !exist {
				return
			}
			delete(players, player.GetID())
		})
	})
}

// Join 使玩家加入房间
func (slf *Manager[PID, P, R]) Join(player P, roomId int64) error {
	var err error
	slf.rooms.Atom(func(m map[int64]*info[PID, P, R]) {
		room, exist := m[roomId]
		if !exist {
			err = ErrRoomNotExist
			return
		}
		if room.playerLimit > 0 && room.playerLimit <= slf.GetRoomPlayerCount(roomId) {
			err = ErrRoomPlayerFull
			return
		}
		slf.pr.Atom(func(m map[PID]map[int64]struct{}) {
			rooms, exist := m[player.GetID()]
			if !exist {
				rooms = make(map[int64]struct{})
				m[player.GetID()] = rooms
			}
			rooms[roomId] = struct{}{}
		})
		slf.rp.Atom(func(m map[int64]map[PID]struct{}) {
			players, exist := m[roomId]
			if !exist {
				players = make(map[PID]struct{})
				m[roomId] = players
			}
			players[player.GetID()] = struct{}{}
		})
		slf.players.Set(player.GetID(), player)
		slf.OnPlayerJoinRoomEvent(room.room, player)
	})
	return err
}

// KickOut 以某种原因踢出特定玩家
//   - 该函数不会校验任何权限相关的内容，调用后将直接踢出玩家
func (slf *Manager[PID, P, R]) KickOut(roomId int64, executor, kicked PID, reason string) error {
	var err error
	var room R
	slf.rp.Atom(func(m map[int64]map[PID]struct{}) {
		players, exist := m[roomId]
		if !exist {
			err = ErrPlayerNotExist
			return
		}
		_, exist = players[executor]
		if !exist {
			err = ErrPlayerNotExist
			return
		}
		_, exist = players[kicked]
		if !exist {
			return
		}
		room = slf.rooms.Get(roomId).room
	})
	if err == nil {
		return err
	}
	slf.OnPlayerKickedOutEvent(room, executor, kicked, reason)
	slf.Leave(roomId, slf.players.Get(kicked))
	return nil
}
