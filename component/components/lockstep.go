package components

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/component"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"github.com/kercylan98/minotaur/utils/timer"
	"go.uber.org/zap"
	"sync"
	"time"
)

func NewLockstep[ClientID comparable, Command any](options ...LockstepOption[ClientID, Command]) *Lockstep[ClientID, Command] {
	lockstep := &Lockstep[ClientID, Command]{
		clients:   synchronization.NewMap[ClientID, component.LockstepClient[ClientID]](),
		frames:    synchronization.NewMap[int, []Command](),
		ticker:    timer.GetTicker(10),
		frameRate: 15,
		serialization: func(frame int, commands []Command) []byte {
			frameStruct := struct {
				Frame    int       `json:"frame"`
				Commands []Command `json:"commands"`
			}{frame, commands}
			data, _ := json.Marshal(frameStruct)
			return data
		},
		clientCurrentFrame: map[ClientID]int{},
	}
	for _, option := range options {
		option(lockstep)
	}
	return lockstep
}

type Lockstep[ClientID comparable, Command any] struct {
	clients            *synchronization.Map[ClientID, component.LockstepClient[ClientID]] // 接受广播的客户端
	frames             *synchronization.Map[int, []Command]                               // 所有帧指令
	ticker             *timer.Ticker                                                      // 定时器
	frameMutex         sync.Mutex                                                         // 帧锁
	currentFrame       int                                                                // 当前帧
	clientCurrentFrame map[ClientID]int                                                   // 客户端当前帧数

	frameRate     int                                        // 帧率（每秒N帧）
	serialization func(frame int, commands []Command) []byte // 序列化函数
}

func (slf *Lockstep[ClientID, Command]) JoinClient(client component.LockstepClient[ClientID]) {
	slf.clients.Set(client.GetID(), client)
}

func (slf *Lockstep[ClientID, Command]) LeaveClient(clientId ClientID) {
	slf.clients.Delete(clientId)
	delete(slf.clientCurrentFrame, clientId)
}

func (slf *Lockstep[ClientID, Command]) StartBroadcast() {
	slf.ticker.Loop("lockstep", timer.Instantly, time.Second/time.Duration(slf.frameRate), timer.Forever, func() {

		slf.frameMutex.Lock()
		currentFrame := slf.currentFrame
		slf.currentFrame++
		slf.frameMutex.Unlock()

		frames := slf.frames.Map()
		for clientId, client := range slf.clients.Map() {
			for i := slf.clientCurrentFrame[clientId]; i <= currentFrame; i++ {
				if err := client.SyncSend(slf.serialization(i, frames[i])); err != nil {
					log.Error("Lockstep.StartBroadcast", zap.Any("ClientID", client.GetID()), zap.Int("Frame", i), zap.Error(err))
					break
				}
				slf.clientCurrentFrame[clientId] = i
			}
		}
	})
}

func (slf *Lockstep[ClientID, Command]) AddCommand(command Command) {
	slf.frames.AtomGetSet(slf.currentFrame, func(value []Command, exist bool) (newValue []Command, isSet bool) {
		return append(value, command), true
	})
}
