package standard

import (
	"errors"
	"minotaur/game/feature"
	"reflect"
)

// NewRoomSpectator 对普通房间附加观众席功能的实例
func NewRoomSpectator[P feature.Player](room *Room[P], spectatorMaximum int) *RoomSpectator[P] {
	return &RoomSpectator[P]{
		Room:             room,
		spectatorMaximum: spectatorMaximum,
	}
}

type RoomSpectator[P feature.Player] struct {
	*Room[P]
	spectatorMaximum int            // 观众席最大人数，小于等于0不限
	spectatorPlayers map[int64]P    // 观众席玩家
	recordInRoom     map[int64]bool // 记录玩家是否从房间中移动到观众席
}

func (slf *RoomSpectator[P]) GetSpectatorMaximum() int {
	return slf.spectatorMaximum
}

func (slf *RoomSpectator[P]) JoinSpectator(player P) error {
	if slf.spectatorMaximum > 0 && len(slf.spectatorPlayers) >= slf.spectatorMaximum {
		return errors.New("maximum number of player")
	}

	var guid = player.GetGuid()

	if !reflect.ValueOf(slf.GetPlayer(guid)).IsNil() {
		if slf.recordInRoom == nil {
			slf.recordInRoom = map[int64]bool{}
		}
		slf.recordInRoom[guid] = true
		slf.LeaveRoom(player.GetGuid())
	}
	if slf.spectatorPlayers == nil {
		slf.spectatorPlayers = map[int64]P{}
	}
	slf.spectatorPlayers[guid] = player
	return nil
}

func (slf *RoomSpectator[P]) LeaveSpectator(guid int64) error {
	var (
		inRoom      = slf.recordInRoom[guid]
		inSpectator = !reflect.ValueOf(slf.spectatorPlayers[guid]).IsNil()
	)

	if inSpectator {
		delete(slf.spectatorPlayers, guid)
		delete(slf.recordInRoom, guid)
		if inRoom {
			return slf.JoinRoom(slf.GetPlayer(guid))
		}
	}

	return nil
}

func (slf *RoomSpectator[P]) GetSpectatorPlayer(guid int64) P {
	return slf.spectatorPlayers[guid]
}

func (slf *RoomSpectator[P]) GetSpectatorPlayers() map[int64]P {
	return slf.spectatorPlayers
}
