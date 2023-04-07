package game

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"minotaur/utils/log"
	"minotaur/utils/super"
	"minotaur/utils/timer"
	"reflect"
	"strconv"
	"sync"
	"time"
)

// channel 频道
type channel struct {
	id           any               // 频道ID
	size         int               // 最大人数
	stateMachine *StateMachine     // 状态机
	timer        *timer.Manager    // 定时器
	lock         sync.RWMutex      // 玩家锁
	players      map[int64]*Player // 所有玩家
	alone        bool              // 频道运行逻辑是否单线程

	// alone mode
	messages chan *Message

	// multi mode
}

// run 开始运行频道
func (slf *channel) run() {
	slf.messages = make(chan *Message, 4096*1000)
	go func() {
		for {
			msg, open := <-slf.messages
			if !open {
				return
			}
			slf.dispatch(msg)
		}
	}()
}

// tryRelease 尝试释放频道
func (slf *channel) tryRelease() {
	slf.timer.After("channelRelease", time.Second*10, func() {
		slf.stateMachine.channelRWMutex.Lock()
		defer slf.stateMachine.channelRWMutex.Unlock()
		if len(slf.players) > 0 {
			return
		}
		slf.timer.Release()
		close(slf.messages)
		delete(slf.stateMachine.channels, slf.id)
		log.Info("ChannelRelease", zap.Any("channelID", slf.id), zap.Int("channelCount", len(slf.stateMachine.channels)))
	})
}

// tryRelease 强制释放频道
func (slf *channel) release() {
	slf.stateMachine.channelRWMutex.Lock()
	defer slf.stateMachine.channelRWMutex.Unlock()
	slf.timer.Release()
	close(slf.messages)
	delete(slf.stateMachine.channels, slf.id)
	log.Info("ChannelRelease",
		zap.Any("channelID", slf.id),
		zap.Int("channelCount", len(slf.stateMachine.channels)),
		zap.String("type", "force"),
	)
}

// join 加入频道
func (slf *channel) join(player *Player) (*Player, error) {
	slf.lock.Lock()
	if slf.size > 0 && len(slf.players) > slf.size {
		return player, fmt.Errorf("join channel[%v] failed, maximum[%d] number of player", slf.id, slf.size)
	}
	if player.channel != nil {
		player.channel.lock.Lock()
		delete(player.channel.players, player.guid)
		player.channel.lock.Unlock()
		player.channel.tryRelease()
	}
	slf.players[player.guid] = player
	slf.lock.Unlock()
	player.channel = slf
	player.ChannelTimer = slf.timer
	log.Info("ChannelJoin",
		zap.Any("channelID", slf.id),
		zap.Int64("playerGuid", player.guid),
		zap.String("size", super.If(slf.size <= 0, "NaN", strconv.Itoa(slf.size))),
		zap.Int("headcount", len(slf.players)),
		zap.String("address", player.ip),
	)
	return player, nil
}

// push 向频道内推送消息
func (slf *channel) push(message *Message) {
	switch message.Type() {
	case MessageTypePlayer:
		if slf.alone {
			slf.messages <- message
		} else {
			slf.dispatch(message)
		}
	case MessageTypeEvent:
		slf.messages <- message
	}

}

// dispatch 消息分发
func (slf *channel) dispatch(message *Message) {
	args := message.Args()
	switch message.Type() {
	case MessageTypePlayer:
		player, code, data := args[0], args[1], args[2]
		h := slf.stateMachine.handles[code.(int32)]
		var in = []reflect.Value{reflect.ValueOf(player)}
		if h[1] != nil {
			messageType := reflect.New(h[1].(reflect.Type).Elem())
			if err := proto.Unmarshal(data.([]byte), messageType.Interface().(proto.Message)); err != nil {
				panic(err)
			}
			in = append(in, messageType)
		}
		h[0].(reflect.Value).Call(in)
	case MessageTypeEvent:
		event, player := args[0].(byte), args[1].(*Player)
		switch event {
		case EventTypeGuestJoin:
			slf.stateMachine.router.OnGuestPlayerJoinEvent(player)
		case EventTypeGuestLeave:
			callback := args[2].(func(player *Player))
			slf.stateMachine.router.OnGuestPlayerLeaveEvent(player)
			callback(player)
		}
	}
}

// GetPlayers 获取频道内所有玩家
func (slf *channel) GetPlayers() map[int64]*Player {
	var players = make(map[int64]*Player)
	if !slf.alone {
		slf.lock.RLock()
		defer slf.lock.RUnlock()
	}
	for id, player := range slf.players {
		players[id] = player
	}
	return players
}

// GetAllChannelPlayers 获取所有频道玩家
func (slf *channel) GetAllChannelPlayers() map[int64]*Player {
	if !slf.alone {
		slf.stateMachine.channelRWMutex.RLock()
	}
	var channels = make([]*channel, 0, len(slf.stateMachine.channels))
	for _, channel := range slf.stateMachine.channels {
		channels = append(channels, channel)
	}
	if !slf.alone {
		slf.stateMachine.channelRWMutex.RUnlock()
	}
	var players = make(map[int64]*Player)
	for _, channel := range channels {
		if !slf.alone {
			channel.lock.RLock()
		}
		for id, player := range channel.players {
			players[id] = player
		}
		if !slf.alone {
			channel.lock.RUnlock()
		}
	}
	return players
}
