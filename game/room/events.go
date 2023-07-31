package room

import "github.com/kercylan98/minotaur/game"

type (
	// PlayerJoinRoomEventHandle 玩家加入房间事件处理函数
	PlayerJoinRoomEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, player P)
	// PlayerLeaveRoomEventHandle 玩家离开房间事件处理函数
	PlayerLeaveRoomEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, player P)
	// PlayerKickedOutEventHandle 玩家被踢出房间事件处理函数
	PlayerKickedOutEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, executor, kicked P, reason string)
	// PlayerUpgradeOwnerEventHandle 玩家成为房主事件处理函数
	PlayerUpgradeOwnerEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, oldOwner, newOwner P)
	// CancelOwnerEventHandle 取消房主事件处理函数
	CancelOwnerEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, oldOwner P)
	// ChangePlayerLimitEventHandle 改变房间人数上限事件处理函数
	ChangePlayerLimitEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, oldLimit, newLimit int)
	// PlayerSeatChangeEventHandle 玩家座位改变事件处理函数
	PlayerSeatChangeEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, player P, oldSeat, newSeat int)
	// PlayerSeatSetEventHandle 玩家座位设置事件处理函数
	PlayerSeatSetEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, player P, seat int)
	// PlayerSeatCancelEventHandle 玩家座位取消事件处理函数
	PlayerSeatCancelEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, player P, seat int)
	// CreateEventHandle 房间创建事件处理函数
	CreateEventHandle[PID comparable, P game.Player[PID], R Room] func(room R, helper *Helper[PID, P, R])
)

func newEvent[PID comparable, P game.Player[PID], R Room]() *event[PID, P, R] {
	return &event[PID, P, R]{
		playerJoinRoomEventRoomHandles:     make(map[int64][]PlayerJoinRoomEventHandle[PID, P, R]),
		playerLeaveRoomEventRoomHandles:    make(map[int64][]PlayerLeaveRoomEventHandle[PID, P, R]),
		playerKickedOutEventRoomHandles:    make(map[int64][]PlayerKickedOutEventHandle[PID, P, R]),
		playerUpgradeOwnerEventRoomHandles: make(map[int64][]PlayerUpgradeOwnerEventHandle[PID, P, R]),
		cancelOwnerEventRoomHandles:        make(map[int64][]CancelOwnerEventHandle[PID, P, R]),
		changePlayerLimitEventRoomHandles:  make(map[int64][]ChangePlayerLimitEventHandle[PID, P, R]),
		playerSeatChangeEventRoomHandles:   make(map[int64][]PlayerSeatChangeEventHandle[PID, P, R]),
		playerSeatSetEventRoomHandles:      make(map[int64][]PlayerSeatSetEventHandle[PID, P, R]),
	}
}

type event[PID comparable, P game.Player[PID], R Room] struct {
	playerJoinRoomEventHandles         []PlayerJoinRoomEventHandle[PID, P, R]
	playerJoinRoomEventRoomHandles     map[int64][]PlayerJoinRoomEventHandle[PID, P, R]
	playerLeaveRoomEventHandles        []PlayerLeaveRoomEventHandle[PID, P, R]
	playerLeaveRoomEventRoomHandles    map[int64][]PlayerLeaveRoomEventHandle[PID, P, R]
	playerKickedOutEventHandles        []PlayerKickedOutEventHandle[PID, P, R]
	playerKickedOutEventRoomHandles    map[int64][]PlayerKickedOutEventHandle[PID, P, R]
	playerUpgradeOwnerEventHandles     []PlayerUpgradeOwnerEventHandle[PID, P, R]
	playerUpgradeOwnerEventRoomHandles map[int64][]PlayerUpgradeOwnerEventHandle[PID, P, R]
	cancelOwnerEventHandles            []CancelOwnerEventHandle[PID, P, R]
	cancelOwnerEventRoomHandles        map[int64][]CancelOwnerEventHandle[PID, P, R]
	changePlayerLimitEventHandles      []ChangePlayerLimitEventHandle[PID, P, R]
	changePlayerLimitEventRoomHandles  map[int64][]ChangePlayerLimitEventHandle[PID, P, R]
	playerSeatChangeEventHandles       []PlayerSeatChangeEventHandle[PID, P, R]
	playerSeatChangeEventRoomHandles   map[int64][]PlayerSeatChangeEventHandle[PID, P, R]
	playerSeatSetEventHandles          []PlayerSeatSetEventHandle[PID, P, R]
	playerSeatSetEventRoomHandles      map[int64][]PlayerSeatSetEventHandle[PID, P, R]
	playerSeatCancelEventHandles       []PlayerSeatCancelEventHandle[PID, P, R]
	playerSeatCancelEventRoomHandles   map[int64][]PlayerSeatCancelEventHandle[PID, P, R]
	roomCreateEventHandles             []CreateEventHandle[PID, P, R]
}

func (slf *event[PID, P, R]) unReg(guid int64) {
	delete(slf.playerJoinRoomEventRoomHandles, guid)
	delete(slf.playerLeaveRoomEventRoomHandles, guid)
	delete(slf.playerKickedOutEventRoomHandles, guid)
	delete(slf.playerUpgradeOwnerEventRoomHandles, guid)
	delete(slf.cancelOwnerEventRoomHandles, guid)
	delete(slf.changePlayerLimitEventRoomHandles, guid)
	delete(slf.playerSeatChangeEventRoomHandles, guid)
	delete(slf.playerSeatSetEventRoomHandles, guid)
	delete(slf.playerSeatCancelEventRoomHandles, guid)
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
func (slf *event[PID, P, R]) OnPlayerKickedOutEvent(room R, executor, kicked P, reason string) {
	for _, handle := range slf.playerKickedOutEventHandles {
		handle(room, executor, kicked, reason)
	}
	for _, handle := range slf.playerKickedOutEventRoomHandles[room.GetGuid()] {
		handle(room, executor, kicked, reason)
	}
}

// RegPlayerUpgradeOwnerEvent 玩家成为房主时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerUpgradeOwnerEvent(handle PlayerUpgradeOwnerEventHandle[PID, P, R]) {
	slf.playerUpgradeOwnerEventHandles = append(slf.playerUpgradeOwnerEventHandles, handle)
}

// RegPlayerUpgradeOwnerEventWithRoom 玩家成为房主时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerUpgradeOwnerEventWithRoom(room R, handle PlayerUpgradeOwnerEventHandle[PID, P, R]) {
	slf.playerUpgradeOwnerEventRoomHandles[room.GetGuid()] = append(slf.playerUpgradeOwnerEventRoomHandles[room.GetGuid()], handle)
}

// OnPlayerUpgradeOwnerEvent 玩家成为房主时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnPlayerUpgradeOwnerEvent(room R, executor, newOwner P) {
	for _, handle := range slf.playerUpgradeOwnerEventHandles {
		handle(room, executor, newOwner)
	}
	for _, handle := range slf.playerUpgradeOwnerEventRoomHandles[room.GetGuid()] {
		handle(room, executor, newOwner)
	}
}

// RegCancelOwnerEvent 取消房主时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegCancelOwnerEvent(handle CancelOwnerEventHandle[PID, P, R]) {
	slf.cancelOwnerEventHandles = append(slf.cancelOwnerEventHandles, handle)
}

// RegCancelOwnerEventWithRoom 取消房主时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegCancelOwnerEventWithRoom(room R, handle CancelOwnerEventHandle[PID, P, R]) {
	slf.cancelOwnerEventRoomHandles[room.GetGuid()] = append(slf.cancelOwnerEventRoomHandles[room.GetGuid()], handle)
}

// OnCancelOwnerEvent 取消房主时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnCancelOwnerEvent(room R, oldOwner P) {
	for _, handle := range slf.cancelOwnerEventHandles {
		handle(room, oldOwner)
	}
	for _, handle := range slf.cancelOwnerEventRoomHandles[room.GetGuid()] {
		handle(room, oldOwner)
	}
}

// RegChangePlayerLimitEvent 修改玩家上限时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegChangePlayerLimitEvent(handle ChangePlayerLimitEventHandle[PID, P, R]) {
	slf.changePlayerLimitEventHandles = append(slf.changePlayerLimitEventHandles, handle)
}

// RegChangePlayerLimitEventWithRoom 修改玩家上限时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegChangePlayerLimitEventWithRoom(room R, handle ChangePlayerLimitEventHandle[PID, P, R]) {
	slf.changePlayerLimitEventRoomHandles[room.GetGuid()] = append(slf.changePlayerLimitEventRoomHandles[room.GetGuid()], handle)
}

// OnChangePlayerLimitEvent 修改玩家上限时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnChangePlayerLimitEvent(room R, oldLimit, newLimit int) {
	for _, handle := range slf.changePlayerLimitEventHandles {
		handle(room, oldLimit, newLimit)
	}
	for _, handle := range slf.changePlayerLimitEventRoomHandles[room.GetGuid()] {
		handle(room, oldLimit, newLimit)
	}
}

// RegPlayerSeatChangeEvent 玩家座位改变时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerSeatChangeEvent(handle PlayerSeatChangeEventHandle[PID, P, R]) {
	slf.playerSeatChangeEventHandles = append(slf.playerSeatChangeEventHandles, handle)
}

// RegPlayerSeatChangeEventWithRoom 玩家座位改变时将立即执行被注册的事件处理函数
//   - 当玩家之前没有座位时，oldSeat 为 NoSeat
func (slf *event[PID, P, R]) RegPlayerSeatChangeEventWithRoom(room R, handle PlayerSeatChangeEventHandle[PID, P, R]) {
	slf.playerSeatChangeEventRoomHandles[room.GetGuid()] = append(slf.playerSeatChangeEventRoomHandles[room.GetGuid()], handle)
}

// OnPlayerSeatChangeEvent 玩家座位改变时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnPlayerSeatChangeEvent(room R, player P, oldSeat, newSeat int) {
	for _, handle := range slf.playerSeatChangeEventHandles {
		handle(room, player, oldSeat, newSeat)
	}
	for _, handle := range slf.playerSeatChangeEventRoomHandles[room.GetGuid()] {
		handle(room, player, oldSeat, newSeat)
	}
}

// RegPlayerSeatSetEvent 玩家座位设置时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerSeatSetEvent(handle PlayerSeatSetEventHandle[PID, P, R]) {
	slf.playerSeatSetEventHandles = append(slf.playerSeatSetEventHandles, handle)
}

// RegPlayerSeatSetEventWithRoom 玩家座位设置时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerSeatSetEventWithRoom(room R, handle PlayerSeatSetEventHandle[PID, P, R]) {
	slf.playerSeatSetEventRoomHandles[room.GetGuid()] = append(slf.playerSeatSetEventRoomHandles[room.GetGuid()], handle)
}

// OnPlayerSeatSetEvent 玩家座位设置时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnPlayerSeatSetEvent(room R, player P, seat int) {
	for _, handle := range slf.playerSeatSetEventHandles {
		handle(room, player, seat)
	}
	for _, handle := range slf.playerSeatSetEventRoomHandles[room.GetGuid()] {
		handle(room, player, seat)
	}
}

// RegPlayerSeatCancelEvent 玩家座位取消时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerSeatCancelEvent(handle PlayerSeatCancelEventHandle[PID, P, R]) {
	slf.playerSeatCancelEventHandles = append(slf.playerSeatCancelEventHandles, handle)
}

// RegPlayerSeatCancelEventWithRoom 玩家座位取消时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegPlayerSeatCancelEventWithRoom(room R, handle PlayerSeatCancelEventHandle[PID, P, R]) {
	slf.playerSeatCancelEventRoomHandles[room.GetGuid()] = append(slf.playerSeatCancelEventRoomHandles[room.GetGuid()], handle)
}

// OnPlayerSeatCancelEvent 玩家座位取消时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnPlayerSeatCancelEvent(room R, player P, seat int) {
	for _, handle := range slf.playerSeatCancelEventHandles {
		handle(room, player, seat)
	}
	for _, handle := range slf.playerSeatCancelEventRoomHandles[room.GetGuid()] {
		handle(room, player, seat)
	}
}

// RegRoomCreateEvent 房间创建时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) RegRoomCreateEvent(handle CreateEventHandle[PID, P, R]) {
	slf.roomCreateEventHandles = append(slf.roomCreateEventHandles, handle)
}

// OnRoomCreateEvent 房间创建时将立即执行被注册的事件处理函数
func (slf *event[PID, P, R]) OnRoomCreateEvent(room R, helper *Helper[PID, P, R]) {
	for _, handle := range slf.roomCreateEventHandles {
		handle(room, helper)
	}
}
