package components

import (
	"minotaur/game/feature"
	"reflect"
)

// NewRoomManager 创建房间管理组件
func NewRoomManager[P feature.Player, R feature.Room[P]]() *RoomManager[P, R] {
	return &RoomManager[P, R]{}
}

// RoomManager 房间管理组件
type RoomManager[P feature.Player, R feature.Room[P]] struct {
	rooms         map[int64]R
	playerRoomRef map[int64]int64
}

func (slf *RoomManager[P, R]) JoinRoom(player feature.Player, room R) {
	if slf.rooms == nil {
		slf.rooms = map[int64]R{}
		slf.playerRoomRef = map[int64]int64{}
	}
	slf.rooms[room.GetGuid()] = room
	slf.playerRoomRef[player.GetGuid()] = room.GetGuid()
}

func (slf *RoomManager[P, R]) LeaveRoom(playerGuid int64) {
	roomId := slf.playerRoomRef[playerGuid]
	room := slf.rooms[roomId]
	if !reflect.ValueOf(room).IsNil() {
		room.LeaveRoom(playerGuid)
	}
	delete(slf.rooms, roomId)
	delete(slf.playerRoomRef, playerGuid)
}

func (slf *RoomManager[P, R]) GetRoom(guid int64) R {
	return slf.rooms[guid]
}

func (slf *RoomManager[P, R]) GetPlayer(guid int64) P {
	return slf.rooms[slf.playerRoomRef[guid]].GetPlayer(guid)
}

func (slf *RoomManager[P, R]) GetPlayerRoom(guid int64) R {
	return slf.rooms[slf.playerRoomRef[guid]]
}
