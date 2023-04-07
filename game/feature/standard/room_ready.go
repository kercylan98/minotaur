package standard

import "minotaur/game/feature"

// NewRoomReady 对普通房间附加准备功能的实例
func NewRoomReady[P feature.Player](room *Room[P]) *RoomReady[P] {
	return &RoomReady[P]{}
}

type RoomReady[P feature.Player] struct {
	*Room[P]
	readies map[int64]bool
}

func (slf *RoomReady[P]) Ready(playerGuid int64, ready bool) {
	if ready {
		if !slf.IsExist(playerGuid) {
			if _, exist := slf.readies[playerGuid]; exist {
				delete(slf.readies, playerGuid)
			}
			return
		}
		if slf.readies == nil {
			slf.readies = map[int64]bool{}
		}
		slf.readies[playerGuid] = true
	} else {
		delete(slf.readies, playerGuid)
	}
}

func (slf *RoomReady[P]) IsAllReady() bool {
	return len(slf.readies) >= slf.GetPlayerCount()
}

func (slf *RoomReady[P]) GetReadyCount() int {
	return len(slf.readies)
}

func (slf *RoomReady[P]) GetUnready() map[int64]P {
	var unreadiest = make(map[int64]P)
	for guid, player := range slf.GetPlayers() {
		if !slf.readies[guid] {
			unreadiest[guid] = player
		}
	}
	return unreadiest
}
