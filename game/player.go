package game

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"minotaur/game/protobuf/messagepack"
	"minotaur/game/protobuf/protobuf"
	"minotaur/utils/log"
	"minotaur/utils/timer"
)

// Player 玩家数据结构
type Player struct {
	*channel // 玩家所在频道信息

	guid  int64           // 玩家guid
	ip    string          // 玩家IP
	conn  Conn            // 玩家连接信息
	ws    *websocket.Conn // 玩家 WebSocket 连接
	login byte            // 登录状态 0: 无需登录 1: 登录成功

	// 游戏定时器
	//
	// GameTimer: 状态机级别定时器；
	// ChannelTimer: 频道级别定时器；
	// Manager: 玩家级别定时器。
	*timer.Manager
	GameTimer, ChannelTimer *timer.Manager
}

// Login 设置玩家是否登录成功
func (slf *Player) Login(success bool) {
	slf.stateMachine.router.OnPlayerLoginEvent(slf, success)
	if success {
		slf.StopTimer(loginTimeoutTimerName)
		slf.login = 1
	} else {
		slf.exit("login failed")
	}
}

// IsLogin 获取玩家是否已登录
func (slf *Player) IsLogin() bool {
	return slf.login == 1
}

// GetGuid 获取玩家GUID
func (slf *Player) GetGuid() int64 {
	return slf.guid
}

// GetChannelId 返回玩家所在频道id
func (slf *Player) GetChannelId() any {
	return slf.channel.id
}

// GetConn 获取连接信息
func (slf *Player) GetConn() any {
	return slf.conn.GetConn()
}

// Close 关闭玩家连接
func (slf *Player) Close() {
	_ = slf.ws.Close()
}

// PushToClient 推送消息到客户端
func (slf *Player) PushToClient(messageCode int32, message proto.Message) {
	data, err := messagepack.Pack(messageCode, message)
	if err != nil {
		panic(err)
	}

	if err = slf.ws.WriteMessage(2, data); err != nil {
		panic(err)
	}
}

// PushToClients 推送消息到多个玩家客户端
func (slf *Player) PushToClients(messageCode int32, message proto.Message, playerGuids ...int64) []PushFail {
	data, err := messagepack.Pack(messageCode, message)
	if err != nil {
		panic(err)
	}

	if !slf.channel.alone {
		slf.channel.lock.RLock()
		defer slf.channel.lock.RUnlock()
	}

	var pushFails []PushFail
	for _, id := range playerGuids {
		player := slf.channel.players[id]
		if player == nil {
			pushFails = append(pushFails, PushFail{
				Data: data,
				Err:  fmt.Errorf("player [%v] not found", id),
			})
			continue
		}

		if err = player.ws.WriteMessage(2, data); err != nil {
			pushFails = append(pushFails, PushFail{
				Player: player,
				Data:   data,
				Err:    err,
			})
		}
	}
	return pushFails
}

// PushToChannelPlayerClients 推送消息到频道内所有玩家客户端，可通过指定玩家或玩家id排除特定玩家
func (slf *Player) PushToChannelPlayerClients(messageCode int32, message proto.Message, excludeGuids ...int64) []PushFail {
	data, err := messagepack.Pack(messageCode, message)
	if err != nil {
		panic(err)
	}

	var pushFails []PushFail
	if !slf.channel.alone {
		slf.channel.lock.RLock()
		defer slf.channel.lock.RUnlock()
	}
	for id, player := range slf.channel.players {
		var exclude bool
		for _, target := range excludeGuids {
			if id == target {
				exclude = true
				break
			}
		}
		if exclude {
			break
		}

		if err = player.ws.WriteMessage(2, data); err != nil {
			pushFails = append(pushFails, PushFail{
				Player: player,
				Data:   data,
				Err:    err,
			})
		}
	}
	return pushFails
}

// PushToChannel 推送消息到频道
func (slf *Player) PushToChannel(messageCode protobuf.MessageCode, message proto.Message) {
	slf.channel.push(new(Message).Init(MessageTypePlayer, slf, messageCode, message))
}

// ChangeChannel 改变玩家所在频道
func (slf *Player) ChangeChannel(channelId any) error {
	channel := slf.channel.stateMachine.channel(channelId, slf.channel.size, slf.channel.alone)
	_, err := channel.join(slf)
	return err
}

// exit 退出游戏
func (slf *Player) exit(reason ...string) {
	var r = "normal exit"
	if len(reason) > 0 {
		r = reason[0]
	}

	slf.channel.push(new(Message).Init(MessageTypeEvent, EventTypeGuestLeave, slf, func(player *Player) {
		slf.Release()
		_ = slf.ws.Close()
		if slf.channel != nil {
			slf.channel.lock.Lock()
			defer slf.channel.lock.Unlock()
			delete(slf.channel.players, slf.guid)
			log.Info("ChannelLeave",
				zap.Any("channelID", slf.channel.id),
				zap.Int64("playerGuid", slf.guid),
				zap.Int("headcount", len(slf.channel.players)),
				zap.String("address", slf.ip),
				zap.String("reason", r),
			)
			slf.channel.tryRelease()
		} else {
			log.Info("ChannelLeave",
				zap.Int64("playerGuid", slf.guid),
				zap.String("address", slf.ip),
				zap.String("reason", r),
			)
		}
	}))
}
