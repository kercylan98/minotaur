package room

import "github.com/kercylan98/minotaur/game"

type (
	// PlayerJoinRoomEventHandle 玩家加入房间事件处理函数
	PlayerJoinRoomEventHandle[PID comparable, P game.Player[PID], R Room[PID, P]] func(room R, player P)
	// PlayerLeaveRoomEventHandle 玩家离开房间事件处理函数
	PlayerLeaveRoomEventHandle[PID comparable, P game.Player[PID], R Room[PID, P]] func(room R, player P)
	// PlayerKickedOutEventHandle 玩家被踢出房间事件处理函数
	PlayerKickedOutEventHandle[PID comparable, P game.Player[PID], R Room[PID, P]] func(room R, executor, kicked PID, reason string)
)

func newEvent[PID comparable, P game.Player[PID], R Room[PID, P]]() *event[PID, P, R] {
	return &event[PID, P, R]{
		playerJoinRoomEventRoomHandles:  make(map[int64][]PlayerJoinRoomEventHandle[PID, P, R]),
		playerLeaveRoomEventRoomHandles: make(map[int64][]PlayerLeaveRoomEventHandle[PID, P, R]),
		playerKickedOutEventRoomHandles: make(map[int64][]PlayerKickedOutEventHandle[PID, P, R]),
	}
}

type event[PID comparable, P game.Player[PID], R Room[PID, P]] struct {
	playerJoinRoomEventHandles      []PlayerJoinRoomEventHandle[PID, P, R]
	playerJoinRoomEventRoomHandles  map[int64][]PlayerJoinRoomEventHandle[PID, P, R]
	playerLeaveRoomEventHandles     []PlayerLeaveRoomEventHandle[PID, P, R]
	playerLeaveRoomEventRoomHandles map[int64][]PlayerLeaveRoomEventHandle[PID, P, R]
	playerKickedOutEventHandles     []PlayerKickedOutEventHandle[PID, P, R]
	playerKickedOutEventRoomHandles map[int64][]PlayerKickedOutEventHandle[PID, P, R]
}

func (slf *event[PID, P, R]) unReg(guid int64) {
	delete(slf.playerJoinRoomEventRoomHandles, guid)
	delete(slf.playerLeaveRoomEventRoomHandles, guid)
	delete(slf.playerKickedOutEventRoomHandles, guid)
}

// RegPlayerJoinRoomEvent 玩家进入房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerJoinRoomEvent(handle PlayerJoinRoomEventHandle[PID, P, R]) {
	slf.playerJoinRoomEventHandles = append(slf.playerJoinRoomEventHandles, handle)
}

// OnPlayerJoinRoomEvent 玩家进入房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnPlayerJoinRoomEvent(room R, player P) {
	for _, handle := range slf.playerJoinRoomEventHandles {
		handle(room, player)
	}
	for _, handle := range slf.playerJoinRoomEventRoomHandles[room.GetGuid()] {
		handle(room, player)
	}
}

// RegPlayerJoinRoomEventWithRoom 玩家进入房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerJoinRoomEventWithRoom(room R, handle PlayerJoinRoomEventHandle[PID, P, R]) {
	slf.playerJoinRoomEventRoomHandles[room.GetGuid()] = append(slf.playerJoinRoomEventRoomHandles[room.GetGuid()], handle)
}

// RegPlayerLeaveRoomEvent 玩家离开房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerLeaveRoomEvent(handle PlayerLeaveRoomEventHandle[PID, P, R]) {
	slf.playerLeaveRoomEventHandles = append(slf.playerLeaveRoomEventHandles, handle)
}

// RegPlayerLeaveRoomEventWithRoom 玩家离开房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerLeaveRoomEventWithRoom(room R, handle PlayerLeaveRoomEventHandle[PID, P, R]) {
	slf.playerLeaveRoomEventRoomHandles[room.GetGuid()] = append(slf.playerLeaveRoomEventRoomHandles[room.GetGuid()], handle)
}

// OnPlayerLeaveRoomEvent 玩家离开房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnPlayerLeaveRoomEvent(room R, player P) {
	for _, handle := range slf.playerLeaveRoomEventHandles {
		handle(room, player)
	}
	for _, handle := range slf.playerLeaveRoomEventRoomHandles[room.GetGuid()] {
		handle(room, player)
	}
}

// RegPlayerKickedOutEvent 玩家被踢出房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerKickedOutEvent(handle PlayerKickedOutEventHandle[PID, P, R]) {
	slf.playerKickedOutEventHandles = append(slf.playerKickedOutEventHandles, handle)
}

// RegPlayerKickedOutEventWithRoom 玩家被踢出房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerKickedOutEventWithRoom(room R, handle PlayerKickedOutEventHandle[PID, P, R]) {
	slf.playerKickedOutEventRoomHandles[room.GetGuid()] = append(slf.playerKickedOutEventRoomHandles[room.GetGuid()], handle)
}

// OnPlayerKickedOutEvent 玩家被踢出房间时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnPlayerKickedOutEvent(room R, executor, kicked PID, reason string) {
	for _, handle := range slf.playerKickedOutEventHandles {
		handle(room, executor, kicked, reason)
	}
	for _, handle := range slf.playerKickedOutEventRoomHandles[room.GetGuid()] {
		handle(room, executor, kicked, reason)
	}
}