package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/asynchronous"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/slice"
	"sync"
)

// NewRoomSeat 基于特定游戏房间(game.Room)的实现创建一个支持座位号管理的房间实现(RoomSeat)
func NewRoomSeat[PlayerID comparable, Player game.Player[PlayerID]](room game.Room[PlayerID, Player], options ...RoomSeatOption[PlayerID, Player]) *RoomSeat[PlayerID, Player] {
	roomSeat := &RoomSeat[PlayerID, Player]{
		Room:   room,
		seatPS: asynchronous.NewMap[PlayerID, int](),
	}
	for _, option := range options {
		option(roomSeat)
	}
	return roomSeat
}

// RoomSeat 包含座位号的默认内置房间实现，依赖于游戏房间(game.Room)实现
//   - 实现了对玩家座位号的管理，分别为自动管理(WithRoomSeatAutoManage)及手工管理，默认清空下为手工管理
type RoomSeat[PlayerID comparable, Player game.Player[PlayerID]] struct {
	game.Room[PlayerID, Player]
	mutex         sync.RWMutex
	vacancy       []int
	seatPS        hash.Map[PlayerID, int]
	seatSP        []*PlayerID
	duplicateLock bool
	fillIn        bool
	autoMode      sync.Once
}

// AddSeat 为特定玩家添加座位
//   - 当座位存在空缺的时候，玩家将会优先在空缺位置坐下，否则将会在末位追加
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

// AddSeatWithAssign 将玩家添加到特定的座位
//   - 如果位置已经有玩家，将会与其进行更换
func (slf *RoomSeat[PlayerID, Player]) AddSeatWithAssign(id PlayerID, seat int) {
	slf.AddSeat(id)
	_ = slf.SetSeat(id, seat)
}

// RemovePlayerSeat 删除玩家座位
//   - 受补位模式(WithRoomSeatFillIn)影响
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

// RemoveSeat 删除特定座位的玩家
//   - 受补位模式(WithRoomSeatFillIn)影响
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

// SetSeat 设置玩家的座位号
//   - 如果玩家没有预先添加过座位将会返回错误
//   - 如果位置已经有玩家，将会与其进行更换
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

// GetSeat 获取玩家座位号
func (slf *RoomSeat[PlayerID, Player]) GetSeat(id PlayerID) (int, error) {
	seat, exist := slf.seatPS.GetExist(id)
	if !exist {
		return 0, ErrRoomNotHasPlayer
	}
	return seat, nil
}

// GetPlayerIDWithSeat 获取特定座位号的玩家
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

// GetSeatInfo 获取所有座位号
//   - 在非补位模式(WithRoomSeatFillIn)下由于座位号可能存在缺席的情况，所以需要根据是否为空指针进行判断
func (slf *RoomSeat[PlayerID, Player]) GetSeatInfo() []*PlayerID {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return slf.seatSP
}

// GetSeatInfoMap 获取座位号及其对应的玩家信息
//   - 缺席情况将被忽略
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

// GetSeatInfoMapVacancy 获取座位号及其对应的玩家信息
//   - 缺席情况将不会被忽略
func (slf *RoomSeat[PlayerID, Player]) GetSeatInfoMapVacancy() map[int]*PlayerID {
	slf.mutex.RLock()
	defer slf.mutex.RUnlock()
	return hash.ToMap(slf.seatSP)
}

// GetSeatInfoWithPlayerIDMap 获取玩家及其座位号信息
func (slf *RoomSeat[PlayerID, Player]) GetSeatInfoWithPlayerIDMap() map[PlayerID]int {
	return slf.seatPS.Map()
}

// GetFirstSeat 获取第一个未缺席的座位号
func (slf *RoomSeat[PlayerID, Player]) GetFirstSeat() int {
	for seat, playerId := range slf.seatSP {
		if playerId != nil {
			return seat
		}
	}
	return -1
}

// GetNextSeat 获取特定座位号下一个未缺席的座位号
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

// GetNextSeatVacancy 获取特定座位号下一个座位号
//   - 缺席将不会被忽略
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
