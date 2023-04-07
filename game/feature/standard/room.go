package standard

import (
	"errors"
	"minotaur/game/feature"
	"reflect"
)

// NewRoom 普通房间实例
func NewRoom[P feature.Player](guid int64, playerMaximum int) *Room[P] {
	room := &Room[P]{
		guid:          guid,
		playerMaximum: playerMaximum,
	}
	return room
}

type Room[P feature.Player] struct {
	guid          int64       // 房间 guid
	playerMaximum int         // 房间最大人数，小于等于0不限
	players       map[int64]P // 房间玩家
}

func (slf *Room[P]) GetGuid() int64 {
	return slf.guid
}

func (slf *Room[P]) JoinRoom(player P) error {
	if slf.playerMaximum > 0 && len(slf.players) >= slf.playerMaximum {
		return errors.New("maximum number of player")
	}
	if slf.players == nil {
		slf.players = map[int64]P{}
	}
	slf.players[player.GetGuid()] = player
	return nil
}

func (slf *Room[P]) LeaveRoom(guid int64) {
	delete(slf.players, guid)
}

func (slf *Room[P]) GetPlayerMaximum() int {
	return slf.playerMaximum
}

func (slf *Room[P]) GetPlayer(guid int64) P {
	return slf.players[guid]
}

func (slf *Room[P]) GetPlayers() map[int64]P {
	return slf.players
}

func (slf *Room[P]) GetPlayerCount() int {
	return len(slf.players)
}

func (slf *Room[P]) IsExist(playerGuid int64) bool {
	return !reflect.ValueOf(slf.players[playerGuid]).IsNil()
}
