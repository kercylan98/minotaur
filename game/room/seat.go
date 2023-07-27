package room

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/hash"
	"sync"
)

func newSeat[PlayerID comparable, P game.Player[PlayerID], R Room](manager *Manager[PlayerID, P, R], room R, event *event[PlayerID, P, R]) *Seat[PlayerID, P, R] {
	roomSeat := &Seat[PlayerID, P, R]{
		manager: manager,
		room:    room,
		event:   event,
		seatPS:  concurrent.NewBalanceMap[PlayerID, int](),
	}
	return roomSeat
}

// Seat 房间座位信息
type Seat[PlayerID comparable, P game.Player[PlayerID], R Room] struct {
	manager       *Manager[PlayerID, P, R]
	room          R
	event         *event[PlayerID, P, R]
	mutex         sync.RWMutex
	vacancy       []int
	seatPS        *concurrent.BalanceMap[PlayerID, int]
	seatSP        []*PlayerID
	duplicateLock bool
	autoMode      sync.Once
}

// addSeat 为特定玩家添加座位
//   - 当座位存在空缺的时候，玩家将会优先在空缺位置坐下，否则将会在末位追加
func (slf *Seat[PlayerID, P, R]) addSeat(id PlayerID) {
	if slf.seatPS.Exist(id) {
		return
	}
	var seat int
	slf.mutex.Lock()
	if len(slf.vacancy) > 0 {
		seat = slf.vacancy[0]
		slf.vacancy = slf.vacancy[1:]
		slf.seatPS.Set(id, seat)
		slf.seatSP[seat] = &id
	} else {
		seat = len(slf.seatSP)
		slf.seatPS.Set(id, seat)
		slf.seatSP = append(slf.seatSP, &id)
	}
	slf.mutex.Unlock()
	slf.event.OnPlayerSeatSetEvent(slf.room, slf.manager.GetPlayer(id), seat)
}

// removePlayerSeat 删除玩家座位
func (slf *Seat[PlayerID, P, R]) removePlayerSeat(id PlayerID) {
	if !slf.seatPS.Exist(id) {
		return
	}
	slf.event.OnPlayerSeatCancelEvent(slf.room, slf.manager.GetPlayer(id), slf.seatPS.Get(id))
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	seat := slf.seatPS.DeleteGet(id)
	slf.seatSP[seat] = nil
}

// SetSeat 设置玩家的座位号
//   - 如果玩家没有预先添加过座位将会返回错误
//   - 如果位置已经有玩家，将会与其进行更换
func (slf *Seat[PlayerID, P, R]) SetSeat(id PlayerID, seat int) error {
	oldSeat, err := slf.setSeat(id, seat)
	if err != nil {
		return err
	}
	slf.event.OnPlayerSeatChangeEvent(slf.room, slf.manager.GetPlayer(id), oldSeat, seat)
	return nil
}

func (slf *Seat[PlayerID, P, R]) setSeat(id PlayerID, seat int) (int, error) {
	slf.mutex.Lock()
	slf.duplicateLock = true
	defer func() {
		slf.mutex.Unlock()
		slf.duplicateLock = false
	}()
	oldSeat, err := slf.GetSeat(id)
	if err != nil {
		return oldSeat, err
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
			count := seat - maxSeat
			slf.seatSP = append(slf.seatSP, make([]*PlayerID, count)...)
		}
		slf.seatSP[seat] = slf.seatSP[oldSeat]
		slf.seatSP[oldSeat] = nil
		slf.seatPS.Set(id, seat)
	}
	return oldSeat, nil
}

// GetSeat 获取玩家座位号
func (slf *Seat[PlayerID, P, R]) GetSeat(id PlayerID) (int, error) {
	seat, exist := slf.seatPS.GetExist(id)
	if !exist {
		return 0, ErrPlayerNotInRoom
	}
	return seat, nil
}

// GetPlayerIDWithSeat 获取特定座位号的玩家
func (slf *Seat[PlayerID, P, R]) GetPlayerIDWithSeat(seat int) (playerId PlayerID, err error) {
	if !slf.duplicateLock {
		slf.mutex.RLock()
		defer slf.mutex.RUnlock()
	}
	if seat > len(slf.seatSP)-1 {
		return playerId, ErrPlayerNotInRoom
	}
	id := slf.seatSP[seat]
	if id == nil {
		return playerId, ErrPlayerNotInRoom
	}
	return *id, nil
}

// GetSeatInfo 获取所有座位号
//   - 在非补位模式(WithRoomSeatFillIn)下由于座位号可能存在缺席的情况，所以需要根据是否为空指针进行判断
func (slf *Seat[PlayerID, P, R]) GetSeatInfo() []*PlayerID {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return slf.seatSP
}

// GetSeatInfoMap 获取座位号及其对应的玩家信息
//   - 缺席情况将被忽略
func (slf *Seat[PlayerID, P, R]) GetSeatInfoMap() map[int]PlayerID {
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

// GetSeatInfoMapVacancy 获取座位号及其对应的玩家信息
//   - 缺席情况将不会被忽略
func (slf *Seat[PlayerID, P, R]) GetSeatInfoMapVacancy() map[int]*PlayerID {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return hash.ToMap(slf.seatSP)
}

// GetSeatInfoWithPlayerIDMap 获取玩家及其座位号信息
func (slf *Seat[PlayerID, P, R]) GetSeatInfoWithPlayerIDMap() map[PlayerID]int {
	return slf.seatPS.Map()
}

// GetFirstSeat 获取第一个未缺席的座位号
func (slf *Seat[PlayerID, P, R]) GetFirstSeat() int {
	for seat, playerId := range slf.seatSP {
		if playerId != nil {
			return seat
		}
	}
	return -1
}

// GetNextSeat 获取特定座位号下一个未缺席的座位号
func (slf *Seat[PlayerID, P, R]) GetNextSeat(seat int) int {
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

// GetNextSeatVacancy 获取特定座位号下一个座位号
//   - 缺席将不会被忽略
func (slf *Seat[PlayerID, P, R]) GetNextSeatVacancy(seat int) int {
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

// GetPrevSeat 获取特定座位号上一个未缺席的座位号
func (slf *Seat[PlayerID, P, R]) GetPrevSeat(seat int) int {
	l := len(slf.seatSP)
	if l == 0 || seat >= l || seat < 0 {
		return -1
	}
	var target = seat
	for {
		target--
		if target < 0 {
			target = l - 1
		}
		if target == seat {
			return seat
		}
		if slf.seatSP[target] != nil {
			return target
		}
	}
}

// GetPrevSeatVacancy 获取特定座位号上一个座位号
//   - 缺席将不会被忽略
func (slf *Seat[PlayerID, P, R]) GetPrevSeatVacancy(seat int) int {
	l := len(slf.seatSP)
	if l == 0 || seat >= l || seat < 0 {
		return -1
	}
	seat--
	if seat < 0 {
		seat = l - 1
	}
	return seat
}
