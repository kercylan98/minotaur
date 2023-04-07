package game

import (
	"fmt"
	"go.uber.org/zap"
	"minotaur/game/protobuf/protobuf"
	"minotaur/utils/log"
	"reflect"
)

var allowReplaceMessages = map[int32]bool{
	int32(protobuf.MessageCode_SystemHeartbeat): true,
}

type router struct {
	handles map[int32][2]any // 玩家消息表

	guestPlayerJoinEvents  []func(guest *Player)
	guestPlayerLeaveEvents []func(guest *Player)
	heartbeatEvents        []func(player *Player)
	playerLoginEvents      []func(player *Player, success bool)
	gameLaunchEvents       []func()
}

func (slf *router) init(stateMachine *StateMachine) *router {
	slf.RegMessagePlayer(int32(protobuf.MessageCode_SystemHeartbeat), func(player *Player) {
		slf.OnHeartbeatEvent(player)
	})

	return slf
}

// RegGameLaunchEvent 注册游戏启动完成事件
func (slf *router) RegGameLaunchEvent(handleFunc func()) {
	slf.gameLaunchEvents = append(slf.gameLaunchEvents, handleFunc)
}

// OnGameLaunchEvent 游戏启动完成事件发生时
func (slf *router) OnGameLaunchEvent() {
	for _, event := range slf.gameLaunchEvents {
		event()
	}
}

// RegPlayerLoginEvent 注册玩家登录事件
func (slf *router) RegPlayerLoginEvent(handleFunc func(player *Player, success bool)) {
	slf.playerLoginEvents = append(slf.playerLoginEvents, handleFunc)
}

// OnPlayerLoginEvent 玩家登录事件发生时
func (slf *router) OnPlayerLoginEvent(player *Player, success bool) {
	for _, event := range slf.playerLoginEvents {
		event(player, success)
	}
}

// RegHeartbeatEvent 注册客户端心跳事件
func (slf *router) RegHeartbeatEvent(handleFunc func(guest *Player)) {
	slf.heartbeatEvents = append(slf.heartbeatEvents, handleFunc)
}

// OnHeartbeatEvent 客户端心跳事件发生时
func (slf *router) OnHeartbeatEvent(guest *Player) {
	for _, event := range slf.heartbeatEvents {
		event(guest)
	}
}

// RegGuestPlayerJoinEvent 注册访客玩家加入事件
func (slf *router) RegGuestPlayerJoinEvent(handleFunc func(guest *Player)) {
	slf.guestPlayerJoinEvents = append(slf.guestPlayerJoinEvents, handleFunc)
}

// OnGuestPlayerJoinEvent 访问玩家加入时
func (slf *router) OnGuestPlayerJoinEvent(guest *Player) {
	for _, event := range slf.guestPlayerJoinEvents {
		event(guest)
	}
}

// RegGuestPlayerLeaveEvent 注册访客玩家离开事件
func (slf *router) RegGuestPlayerLeaveEvent(handleFunc func(guest *Player)) {
	slf.guestPlayerLeaveEvents = append(slf.guestPlayerLeaveEvents, handleFunc)
}

// OnGuestPlayerLeaveEvent 访问玩家离开时
func (slf *router) OnGuestPlayerLeaveEvent(guest *Player) {
	for _, event := range slf.guestPlayerLeaveEvents {
		event(guest)
	}
}

// RegMessagePlayer 注册玩家消息
func (slf *router) RegMessagePlayer(messageCode int32, handleFunc any) {
	if slf.handles == nil {
		slf.handles = map[int32][2]any{}
	}

	typeOf := reflect.TypeOf(handleFunc)
	if typeOf.Kind() != reflect.Func {
		panic(fmt.Errorf("register player message failed, handleFunc must is func(player *Player, message *protobufType)"))
	}
	if typeOf.NumIn() == 0 {
		panic(fmt.Errorf("register player message failed, handleFunc must is func(player *Player) or func(player *Player, message *protobufType)"))
	}
	handle, exist := slf.handles[messageCode]
	if exist && !allowReplaceMessages[messageCode] {
		panic(fmt.Errorf("register player message failed, message %s existed and cannot be replaced", messageCode))
	}

	handle = [2]any{reflect.ValueOf(handleFunc), nil}
	if typeOf.NumIn() == 2 {
		messageType := typeOf.In(1)
		handle[1] = messageType
	}

	slf.handles[messageCode] = handle
	if exist && allowReplaceMessages[messageCode] {
		log.Info("RouterRegMessagePlayer", zap.Any("code", messageCode), zap.Any("messageType", handle[1]), zap.String("type", "replace"))
	} else {
		log.Info("RouterRegMessagePlayer", zap.Any("code", messageCode), zap.Any("messageType", handle[1]))
	}
}
