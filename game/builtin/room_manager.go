package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"sync/atomic"
)

func NewRoomManager[PlayerID comparable, Room game.Room[PlayerID, game.Player[PlayerID]]]() *RoomManager[PlayerID, Room] {
	return &RoomManager[PlayerID, Room]{
		rooms: concurrent.NewBalanceMap[int64, Room](),
	}
}

// RoomManager 房间管理器
type RoomManager[PlayerID comparable, Room game.Room[PlayerID, game.Player[PlayerID]]] struct {
	guid  atomic.Int64
	rooms *concurrent.BalanceMap[int64, Room]
}

// GenGuid 生成一个新的房间guid
func (slf *RoomManager[PlayerID, Room]) GenGuid() int64 {
	return slf.guid.Add(1)
}

// AddRoom 添加房间到房间管理器中
func (slf *RoomManager[PlayerID, Room]) AddRoom(room Room) {
	slf.rooms.Set(room.GetGuid(), room)
}

// CloseRoom 关闭特定guid的房间
func (slf *RoomManager[PlayerID, Room]) CloseRoom(guid int64) {
	slf.rooms.Delete(guid)
}
