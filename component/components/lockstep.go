package components

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/component"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"github.com/kercylan98/minotaur/utils/timer"
	"sync"
	"sync/atomic"
	"time"
)

// NewLockstep 创建一个锁步（帧）同步默认实现的组件(Lockstep)进行返回
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
		clientCurrentFrame: synchronization.NewMap[ClientID, int](),
	}
	for _, option := range options {
		option(lockstep)
	}
	return lockstep
}

// Lockstep 锁步（帧）同步默认实现
//   - 支持最大帧上限 WithLockstepFrameLimit
//   - 自定逻辑帧频率，默认为每秒15帧(帧/66ms) WithLockstepFrameRate
//   - 自定帧序列化方式 WithLockstepSerialization
//   - 从特定帧开始追帧
//   - 兼容各种基于TCP/UDP/Unix的网络类型，可通过客户端实现其他网络类型同步
type Lockstep[ClientID comparable, Command any] struct {
	clients            *synchronization.Map[ClientID, component.LockstepClient[ClientID]] // 接受广播的客户端
	frames             *synchronization.Map[int, []Command]                               // 所有帧指令
	ticker             *timer.Ticker                                                      // 定时器
	frameMutex         sync.Mutex                                                         // 帧锁
	currentFrame       int                                                                // 当前帧
	clientCurrentFrame *synchronization.Map[ClientID, int]                                // 客户端当前帧数
	running            atomic.Bool

	frameRate     int                                        // 帧率（每秒N帧）
	frameLimit    int                                        // 帧上限
	serialization func(frame int, commands []Command) []byte // 序列化函数

	lockstepStoppedEventHandles []component.LockstepStoppedEventHandle[ClientID, Command]
}

// JoinClient 加入客户端到广播队列中
func (slf *Lockstep[ClientID, Command]) JoinClient(client component.LockstepClient[ClientID]) {
	slf.clients.Set(client.GetID(), client)
}

// JoinClientWithFrame 加入客户端到广播队列中，并从特定帧开始追帧
//   - 可用于重连及状态同步、帧同步混用的情况
//   - 混用：服务端记录指令时同时做一次状态计算，新客户端加入时直接同步当前状态，之后从特定帧开始广播
func (slf *Lockstep[ClientID, Command]) JoinClientWithFrame(client component.LockstepClient[ClientID], frameIndex int) {
	slf.clients.Set(client.GetID(), client)
	if frameIndex > slf.currentFrame {
		frameIndex = slf.currentFrame
	}
	slf.clientCurrentFrame.Set(client.GetID(), frameIndex)
}

// LeaveClient 将客户端从广播队列中移除
func (slf *Lockstep[ClientID, Command]) LeaveClient(clientId ClientID) {
	slf.clients.Delete(clientId)
	slf.clientCurrentFrame.Delete(clientId)
}

// StartBroadcast 开始广播
//   - 在开始广播后将持续按照设定的帧率进行帧数推进，并在每一帧推进时向客户端进行同步，需提前将客户端加入广播队列 JoinClient
//   - 广播过程中使用 AddCommand 将该帧数据追加到当前帧中
func (slf *Lockstep[ClientID, Command]) StartBroadcast() {
	if slf.running.Swap(true) {
		return
	}
	slf.ticker.Loop("lockstep", timer.Instantly, time.Second/time.Duration(slf.frameRate), timer.Forever, func() {

		slf.frameMutex.Lock()
		currentFrame := slf.currentFrame
		if slf.frameLimit > 0 && currentFrame >= slf.frameLimit {
			slf.StopBroadcast()
			return
		}
		slf.currentFrame++
		slf.frameMutex.Unlock()

		frames := slf.frames.Map()
		for clientId, client := range slf.clients.Map() {
			var i = slf.clientCurrentFrame.Get(clientId)
			for ; i < currentFrame; i++ {
				client.Send(slf.serialization(i, frames[i]))
			}
			slf.clientCurrentFrame.Set(clientId, i)

		}
	})
}

// StopBroadcast 停止广播
func (slf *Lockstep[ClientID, Command]) StopBroadcast() {
	if !slf.running.Swap(false) {
		return
	}
	slf.ticker.StopTimer("lockstep")
	slf.frameMutex.Lock()
	slf.OnLockstepStoppedEvent()
	slf.currentFrame = 0
	slf.clientCurrentFrame.Clear()
	slf.frames.Clear()
	slf.frameMutex.Unlock()
}

// AddCommand 添加命令到当前帧
func (slf *Lockstep[ClientID, Command]) AddCommand(command Command) {
	slf.frames.AtomGetSet(slf.currentFrame, func(value []Command, exist bool) (newValue []Command, isSet bool) {
		return append(value, command), true
	})
}

// GetCurrentFrame 获取当前帧
func (slf *Lockstep[ClientID, Command]) GetCurrentFrame() int {
	return slf.currentFrame
}

// GetClientCurrentFrame 获取客户端当前帧
func (slf *Lockstep[ClientID, Command]) GetClientCurrentFrame(clientId ClientID) int {
	return slf.clientCurrentFrame.Get(clientId)
}

// GetFrameLimit 获取帧上限
//   - 未设置时将返回0
func (slf *Lockstep[ClientID, Command]) GetFrameLimit() int {
	return slf.frameLimit
}

// GetFrames 获取所有帧数据
func (slf *Lockstep[ClientID, Command]) GetFrames() [][]Command {
	var frameMap = slf.frames.Map()
	var frames = make([][]Command, len(frameMap))
	for index, commands := range frameMap {
		frames[index] = commands
	}
	return frames
}

// RegLockstepStoppedEvent 当广播停止时将触发被注册的事件处理函数
func (slf *Lockstep[ClientID, Command]) RegLockstepStoppedEvent(handle component.LockstepStoppedEventHandle[ClientID, Command]) {
	slf.lockstepStoppedEventHandles = append(slf.lockstepStoppedEventHandles, handle)
}

func (slf *Lockstep[ClientID, Command]) OnLockstepStoppedEvent() {
	for _, handle := range slf.lockstepStoppedEventHandles {
		handle(slf)
	}
}
