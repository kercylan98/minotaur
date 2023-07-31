package room

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/hash"
	"sync"
)

const (
	NoSeat = -1 // 无座位
)

func newSeat[PlayerID comparable, P game.Player[PlayerID], R Room](manager *Manager[PlayerID, P, R], room R, event *event[PlayerID, P, R]) *Seat[PlayerID, P, R] {
	roomSeat := &Seat[PlayerID, P, R]{
		manager:     manager,
		room:        room,
		event:       event,
		seatPS:      concurrent.NewBalanceMap[PlayerID, int](),
		autoSitDown: true,
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
	autoSitDown   bool
}

// GetSeatPlayerCount 获取座位上的玩家数量
//   - 该数量不包括空缺的座位
func (slf *Seat[PlayerID, P, R]) GetSeatPlayerCount() int {
	return slf.seatPS.Size()
}

// AddSeat 为特定玩家添加座位
//   - 当座位存在空缺的时候，玩家将会优先在空缺位置坐下，否则将会在末位追加
func (slf *Seat[PlayerID, P, R]) AddSeat(id PlayerID) {
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

// RemoveSeat 删除玩家座位
func (slf *Seat[PlayerID, P, R]) RemoveSeat(id PlayerID) {
	if !slf.seatPS.Exist(id) {
		return
	}
	slf.event.OnPlayerSeatCancelEvent(slf.room, slf.manager.GetPlayer(id), slf.seatPS.Get(id))
	slf.mutex.Lock()
	defer slf.mutex.Unlock()
	seat := slf.seatPS.DeleteGet(id)
	slf.seatSP[seat] = nil
}

// HasSeat 判断玩家是否有座位
func (slf *Seat[PlayerID, P, R]) HasSeat(id PlayerID) bool {
	return slf.seatPS.Exist(id)
}

// SetSeat 设置玩家的座位号
//   - 如果位置已经有玩家，将会与其进行更换
func (slf *Seat[PlayerID, P, R]) SetSeat(id PlayerID, seat int) int {
	slf.mutex.Lock()
	slf.duplicateLock = true
	defer func() {
		slf.mutex.Unlock()
		slf.duplicateLock = false
	}()
	oldSeat := slf.GetSeat(id)
	player := slf.GetPlayerWithSeat(seat)
	if generic.IsNil(player) {
		if oldSeat == NoSeat {
			maxSeat := len(slf.seatSP) - 1
			if seat > maxSeat {
				count := seat - maxSeat
				slf.seatSP = append(slf.seatSP, make([]*PlayerID, count)...)
			}
			slf.seatSP[seat] = &id
			slf.seatPS.Set(id, seat)
		} else {
			ov := slf.seatSP[oldSeat]
			slf.seatSP[oldSeat] = slf.seatSP[seat]
			slf.seatSP[seat] = ov
			slf.seatPS.Set(id, seat)
			slf.seatPS.Set(player.GetID(), oldSeat)
		}
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
	slf.event.OnPlayerSeatChangeEvent(slf.room, slf.manager.GetPlayer(id), oldSeat, seat)
	return oldSeat
}

// IsNoSeat 判断玩家是否没有座位
func (slf *Seat[PlayerID, P, R]) IsNoSeat(id PlayerID) bool {
	return slf.GetSeat(id) == NoSeat
}

// GetSeat 获取玩家座位号
//   - 如果玩家没有座位，将会返回 NoSeat
func (slf *Seat[PlayerID, P, R]) GetSeat(id PlayerID) int {
	seat, exist := slf.seatPS.GetExist(id)
	if !exist {
		return NoSeat
	}
	return seat
}

// GetPlayerIDWithSeat 获取特定座位号的玩家ID
func (slf *Seat[PlayerID, P, R]) GetPlayerIDWithSeat(seat int) (id PlayerID) {
	if !slf.duplicateLock {
		slf.mutex.RLock()
		defer slf.mutex.RUnlock()
	}
	if seat >= len(slf.seatSP) || seat < 0 {
		return id
	}
	playerId := slf.seatSP[seat]
	if playerId == nil {
		return id
	}
	return *playerId
}

// GetPlayerWithSeat 获取特定座位号的玩家
func (slf *Seat[PlayerID, P, R]) GetPlayerWithSeat(seat int) (player P) {
	if !slf.duplicateLock {
		slf.mutex.RLock()
		defer slf.mutex.RUnlock()
	}
	if seat >= len(slf.seatSP) || seat < 0 {
		return player
	}
	id := slf.seatSP[seat]
	if id == nil {
		return player
	}
	return slf.manager.GetRoomPlayer(slf.room.GetGuid(), *id)
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
//   - 如果没有，将会返回 NoSeat
func (slf *Seat[PlayerID, P, R]) GetFirstSeat() int {
	for seat, playerId := range slf.seatSP {
		if playerId != nil {
			return seat
		}
	}
	return NoSeat
}

// GetNextSeat 获取特定座位号下一个未缺席的座位号
func (slf *Seat[PlayerID, P, R]) GetNextSeat(seat int) int {
	l := len(slf.seatSP)
	if l == 0 || seat >= l || seat < 0 {
		return NoSeat
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
		return NoSeat
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
		return NoSeat
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
		return NoSeat
	}
	seat--
	if seat < 0 {
		seat = l - 1
	}
	return seat
}
