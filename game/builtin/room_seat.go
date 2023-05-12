package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/slice"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"sync"
)

func NewRoomSeat[PlayerID comparable, Player game.Player[PlayerID]](room game.Room[PlayerID, Player], options ...RoomSeatOption[PlayerID, Player]) *RoomSeat[PlayerID, Player] {
	roomSeat := &RoomSeat[PlayerID, Player]{
		Room:   room,
		seatPS: synchronization.NewMap[PlayerID, int](),
	}
	room.RegPlayerJoinRoomEvent(roomSeat.onJoinRoom)
	room.RegPlayerLeaveRoomEvent(roomSeat.onLeaveRoom)
	for _, option := range options {
		option(roomSeat)
	}
	return roomSeat
}

type RoomSeat[PlayerID comparable, Player game.Player[PlayerID]] struct {
	game.Room[PlayerID, Player]
	mutex   sync.RWMutex
	vacancy []int
	seatPS  *synchronization.Map[PlayerID, int]
	seatSP  []*PlayerID

	fillIn bool
}

func (slf *RoomSeat[PlayerID, Player]) SetSeat(id PlayerID, seat int) error {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	oldSeat, err := slf.GetSeat(id)
	if err != nil {
		return err
	}
	player, err := slf.GetPlayerWithSeat(seat)
	if err != nil {
		ov := slf.seatSP[oldSeat]
		slf.seatSP[oldSeat] = slf.seatSP[seat]
		slf.seatSP[seat] = ov
		slf.seatPS.Set(id, seat)
		slf.seatPS.Set(player.GetID(), oldSeat)
	} else {
		maxSeat := len(slf.seatSP) - 1
		if seat > maxSeat {
			if slf.fillIn {
				seat = maxSeat + 1
				defer func() {
					slice.Del(&slf.seatSP, oldSeat)
					for i := oldSeat; i < len(slf.seatSP); i++ {
						slf.seatPS.Set(*slf.seatSP[i], i)
					}
				}()
			}
			count := seat - maxSeat
			slf.seatSP = append(slf.seatSP, make([]*PlayerID, count)...)
		}
		slf.seatSP[seat] = slf.seatSP[oldSeat]
		slf.seatSP[oldSeat] = nil
		slf.seatPS.Set(id, seat)
	}
	return nil
}

func (slf *RoomSeat[PlayerID, Player]) GetSeat(id PlayerID) (int, error) {
	seat, exist := slf.seatPS.GetExist(id)
	if !exist {
		return 0, ErrRoomNotHasPlayer
	}
	return seat, nil
}

func (slf *RoomSeat[PlayerID, Player]) GetPlayerWithSeat(seat int) (player Player, err error) {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	if seat > len(slf.seatSP)-1 {
		return player, ErrRoomNotHasPlayer
	}
	id := slf.seatSP[seat]
	if id == nil {
		return player, ErrRoomNotHasPlayer
	}
	return slf.GetPlayer(*id), nil
}

func (slf *RoomSeat[PlayerID, Player]) GetSeatInfo() []*PlayerID {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return slf.seatSP
}

func (slf *RoomSeat[PlayerID, Player]) GetSeatInfoMap() map[int]PlayerID {
	var seatInfo = make(map[int]PlayerID)
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	for seat, playerId := range slf.seatSP {
		if playerId == nil {
			continue
		}
		seatInfo[seat] = *playerId
	}
	return seatInfo
}

func (slf *RoomSeat[PlayerID, Player]) GetSeatInfoMapVacancy() map[int]*PlayerID {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return hash.ToMap(slf.seatSP)
}

func (slf *RoomSeat[PlayerID, Player]) GetSeatInfoWithPlayerIDMap() map[PlayerID]int {
	return slf.seatPS.Map()
}

func (slf *RoomSeat[PlayerID, Player]) onJoinRoom(room game.Room[PlayerID, Player], player Player) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	playerId := player.GetID()
	if len(slf.vacancy) > 0 {
		seat := slf.vacancy[0]
		slf.vacancy = slf.vacancy[1:]
		slf.seatPS.Set(player.GetID(), seat)
		slf.seatSP[seat] = &playerId
	} else {
		slf.seatPS.Set(player.GetID(), len(slf.seatSP))
		slf.seatSP = append(slf.seatSP, &playerId)
	}
}

func (slf *RoomSeat[PlayerID, Player]) onLeaveRoom(room game.Room[PlayerID, Player], player Player) {
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	seat := slf.seatPS.DeleteGet(player.GetID())
	if slf.fillIn {
		slice.Del(&slf.seatSP, seat)
		for i := seat; i < len(slf.seatSP); i++ {
			slf.seatPS.Set(*slf.seatSP[i], i)
		}
		return
	}
	slf.seatSP[seat] = nil

}
