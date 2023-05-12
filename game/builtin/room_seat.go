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
	for _, option := range options {
		option(roomSeat)
	}
	return roomSeat
}

type RoomSeat[PlayerID comparable, Player game.Player[PlayerID]] struct {
	game.Room[PlayerID, Player]
	mutex         sync.RWMutex
	vacancy       []int
	seatPS        *synchronization.Map[PlayerID, int]
	seatSP        []*PlayerID
	duplicateLock bool
	fillIn        bool
	autoMode      sync.Once
}

func (slf *RoomSeat[PlayerID, Player]) AddSeat(id PlayerID) {
	if slf.seatPS.Exist(id) {
		return
	}
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	if len(slf.vacancy) > 0 {
		seat := slf.vacancy[0]
		slf.vacancy = slf.vacancy[1:]
		slf.seatPS.Set(id, seat)
		slf.seatSP[seat] = &id
	} else {
		slf.seatPS.Set(id, len(slf.seatSP))
		slf.seatSP = append(slf.seatSP, &id)
	}
}

func (slf *RoomSeat[PlayerID, Player]) AddSeatWithAssign(id PlayerID, seat int) {
	slf.AddSeat(id)
	_ = slf.SetSeat(id, seat)
}

func (slf *RoomSeat[PlayerID, Player]) RemovePlayerSeat(id PlayerID) {
	if !slf.seatPS.Exist(id) {
		return
	}
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	seat := slf.seatPS.DeleteGet(id)
	if slf.fillIn {
		slice.Del(&slf.seatSP, seat)
		for i := seat; i < len(slf.seatSP); i++ {
			slf.seatPS.Set(*slf.seatSP[i], i)
		}
		return
	}
	slf.seatSP[seat] = nil
}

func (slf *RoomSeat[PlayerID, Player]) RemoveSeat(seat int) {
	if seat >= len(slf.seatSP) {
		return
	}
	playerId := slf.seatSP[seat]
	if playerId == nil {
		return
	}
	slf.RemovePlayerSeat(*playerId)
}

func (slf *RoomSeat[PlayerID, Player]) SetSeat(id PlayerID, seat int) error {
	slf.mutex.Lock()
	slf.duplicateLock = true
	defer func() {
		slf.mutex.Unlock()
		slf.duplicateLock = false
	}()
	oldSeat, err := slf.GetSeat(id)
	if err != nil {
		return err
	}
	playerId, err := slf.GetPlayerIDWithSeat(seat)
	if err != nil {
		ov := slf.seatSP[oldSeat]
		slf.seatSP[oldSeat] = slf.seatSP[seat]
		slf.seatSP[seat] = ov
		slf.seatPS.Set(id, seat)
		slf.seatPS.Set(playerId, oldSeat)
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

func (slf *RoomSeat[PlayerID, Player]) GetPlayerIDWithSeat(seat int) (playerId PlayerID, err error) {
	if !slf.duplicateLock {
		slf.mutex.RLock()
		defer slf.mutex.RUnlock()
	}
	if seat > len(slf.seatSP)-1 {
		return playerId, ErrRoomNotHasPlayer
	}
	id := slf.seatSP[seat]
	if id == nil {
		return playerId, ErrRoomNotHasPlayer
	}
	return *id, nil
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

func (slf *RoomSeat[PlayerID, Player]) GetNextSeat(seat int) int {
	l := len(slf.seatSP)
	if l == 0 || seat >= l || seat < 0 {
		return -1
	}
	var target = seat
	for {
		target++
		if target >= l {
			target = 0
		}
		if target == seat {
			return seat
		}
		if slf.seatSP[target] != nil {
			return target
		}
	}
}

func (slf *RoomSeat[PlayerID, Player]) GetNextSeatVacancy(seat int) int {
	l := len(slf.seatSP)
	if l == 0 || seat >= l || seat < 0 {
		return -1
	}
	seat++
	if seat >= l {
		seat = 0
	}
	return seat
}

func (slf *RoomSeat[PlayerID, Player]) onJoinRoom(room game.Room[PlayerID, Player], player Player) {
	slf.AddSeat(player.GetID())
}

func (slf *RoomSeat[PlayerID, Player]) onLeaveRoom(room game.Room[PlayerID, Player], player Player) {
	slf.RemovePlayerSeat(player.GetID())
}
