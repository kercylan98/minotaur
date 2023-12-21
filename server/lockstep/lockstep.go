package lockstep

import (
	"encoding/json"
	"sync"
	"time"
)

// NewLockstep 创建一个锁步（帧）同步默认实现的组件(Lockstep)进行返回
func NewLockstep[ClientID comparable, Command any](options ...Option[ClientID, Command]) *Lockstep[ClientID, Command] {
	lockstep := &Lockstep[ClientID, Command]{
		currentFrame: -1,
		frameRate:    15,
		serialization: func(frame int64, commands []Command) []byte {
			frameStruct := struct {
				Frame    int64     `json:"frame"`
				Commands []Command `json:"commands"`
			}{frame, commands}
			data, _ := json.Marshal(frameStruct)
			return data
		},
		clients:     make(map[ClientID]Client[ClientID]),
		clientFrame: make(map[ClientID]int64),
		frameCache:  make(map[int64][]byte),
	}
	for _, option := range options {
		option(lockstep)
	}
	return lockstep
}

// Lockstep 锁步（帧）同步默认实现
//   - 支持最大帧上限 WithFrameLimit
//   - 自定逻辑帧频率，默认为每秒15帧(帧/66ms) WithFrameRate
//   - 自定帧序列化方式 WithSerialization
//   - 从特定帧开始追帧
//   - 兼容各种基于TCP/UDP/Unix的网络类型，可通过客户端实现其他网络类型同步
type Lockstep[ClientID comparable, Command any] struct {
	running       bool                                         // 运行状态
	runningLock   sync.RWMutex                                 // 运行状态锁
	initFrame     int64                                        // 初始帧
	frameRate     int64                                        // 帧率（每秒N帧）
	frameLimit    int64                                        // 帧上限
	serialization func(frame int64, commands []Command) []byte // 序列化函数

	clients     map[ClientID]Client[ClientID] // 接受广播的客户端
	clientFrame map[ClientID]int64            // 客户端当前帧
	clientLock  sync.RWMutex                  // 客户端锁

	currentFrame     int64        // 当前主要帧
	currentCommands  []Command    // 当前帧指令
	currentFrameLock sync.RWMutex // 当前主要帧锁

	frameCache     map[int64][]byte // 帧序列化缓存
	frameCacheLock sync.RWMutex     // 帧序列化缓存锁
	ticker         *time.Ticker     // 定时器

	lockstepStoppedEventHandles []StoppedEventHandle[ClientID, Command]
}

// JoinClient 将客户端加入到广播队列中，通常在开始广播前使用
//   - 如果客户端在开始广播后加入，将丢失之前的帧数据，如要从特定帧开始追帧请使用 JoinClientWithFrame
func (slf *Lockstep[ClientID, Command]) JoinClient(client Client[ClientID]) {
	slf.clientLock.Lock()
	defer slf.clientLock.Unlock()
	slf.clients[client.GetID()] = client
}

// JoinClientWithFrame 加入客户端到广播队列中，并从特定帧开始追帧
//   - 可用于重连及状态同步、帧同步混用的情况
//   - 混用：服务端记录指令时同时做一次状态计算，新客户端加入时直接同步当前状态，之后从特定帧开始广播
func (slf *Lockstep[ClientID, Command]) JoinClientWithFrame(client Client[ClientID], frameIndex int64) {
	slf.currentFrameLock.RLock()
	if frameIndex > slf.currentFrame {
		frameIndex = slf.currentFrame
	}
	slf.currentFrameLock.RUnlock()
	if frameIndex < slf.initFrame {
		frameIndex = slf.initFrame
	}
	slf.clientLock.Lock()
	slf.clients[client.GetID()] = client
	slf.clientFrame[client.GetID()] = frameIndex
	slf.clientLock.Unlock()

}

// GetClientCount 获取客户端数量
func (slf *Lockstep[ClientID, Command]) GetClientCount() int {
	slf.clientLock.RLock()
	defer slf.clientLock.RUnlock()
	return len(slf.clients)
}

// DropCache 丢弃特定帧的缓存，当 handler 返回 true 时将丢弃缓存
func (slf *Lockstep[ClientID, Command]) DropCache(handler func(frame int64) bool) {
	slf.frameCacheLock.Lock()
	defer slf.frameCacheLock.Unlock()
	for frame, _ := range slf.frameCache {
		if handler(frame) {
			delete(slf.frameCache, frame)
		}
	}
}

// LeaveClient 将客户端从广播队列中移除
func (slf *Lockstep[ClientID, Command]) LeaveClient(clientId ClientID) {
	slf.clientLock.Lock()
	defer slf.clientLock.Unlock()
	delete(slf.clients, clientId)
	delete(slf.clientFrame, clientId)
}

// StartBroadcast 开始广播
//   - 在开始广播后将持续按照设定的帧率进行帧数推进，并在每一帧推进时向客户端进行同步，需提前将客户端加入广播队列 JoinClient
//   - 广播过程中使用 AddCommand 将该帧数据追加到当前帧中
func (slf *Lockstep[ClientID, Command]) StartBroadcast() {
	slf.runningLock.Lock()
	if slf.running {
		slf.runningLock.Unlock()
		return
	}
	slf.running = true
	if slf.ticker == nil {
		slf.ticker = time.NewTicker(time.Second / time.Duration(slf.frameRate))
	}
	slf.runningLock.Unlock()
	slf.currentFrame = slf.initFrame

	go func(ls *Lockstep[ClientID, Command]) {
		for range ls.ticker.C {
			go func(ls *Lockstep[ClientID, Command]) {
				ls.currentFrameLock.RLock()
				if ls.frameLimit > 0 && ls.currentFrame >= ls.frameLimit {
					ls.currentFrameLock.RUnlock()
					ls.StopBroadcast()
					return
				}
				ls.currentFrameLock.RUnlock()
				ls.currentFrameLock.Lock()
				ls.currentFrame++
				currentFrame := ls.currentFrame
				currentCommands := ls.currentCommands
				ls.currentCommands = make([]Command, 0, len(currentCommands))
				ls.currentFrameLock.Unlock()

				ls.clientLock.RLock()
				defer ls.clientLock.RUnlock()
				ls.frameCacheLock.Lock()
				defer ls.frameCacheLock.Unlock()

				for clientId, client := range ls.clients {
					var i = ls.clientFrame[clientId]
					if i < ls.initFrame {
						i = ls.initFrame
					}
					for ; i < currentFrame; i++ {
						cache, exist := ls.frameCache[i]
						if !exist {
							cache = ls.serialization(i, currentCommands)
							ls.frameCache[i] = cache
						}
						client.Write(cache)
					}
					ls.clientFrame[clientId] = currentFrame
				}
			}(ls)
		}
	}(slf)
}

// StopBroadcast 停止广播
func (slf *Lockstep[ClientID, Command]) StopBroadcast() {
	slf.runningLock.Lock()
	if !slf.running {
		slf.runningLock.Unlock()
		return
	}
	slf.running = false
	slf.runningLock.Unlock()

	if slf.ticker != nil {
		slf.ticker.Stop()
	}
	slf.ticker = nil

	slf.OnLockstepStoppedEvent()

	slf.currentFrameLock.Lock()
	defer slf.currentFrameLock.Unlock()
	slf.frameCacheLock.Lock()
	defer slf.frameCacheLock.Unlock()
	slf.clientLock.Lock()
	defer slf.clientLock.Unlock()

	slf.frameCache = make(map[int64][]byte)
	slf.currentCommands = make([]Command, 0)
	slf.currentFrame = -1
	slf.clientFrame = make(map[ClientID]int64)
}

// IsRunning 是否正在广播
func (slf *Lockstep[ClientID, Command]) IsRunning() bool {
	slf.runningLock.RLock()
	defer slf.runningLock.RUnlock()
	return slf.running
}

// AddCommand 添加命令到当前帧
func (slf *Lockstep[ClientID, Command]) AddCommand(command Command) {
	slf.currentFrameLock.Lock()
	defer slf.currentFrameLock.Unlock()
	slf.currentCommands = append(slf.currentCommands, command)
}

// AddCommands 添加命令到当前帧
func (slf *Lockstep[ClientID, Command]) AddCommands(commands []Command) {
	slf.currentFrameLock.Lock()
	defer slf.currentFrameLock.Unlock()
	slf.currentCommands = append(slf.currentCommands, commands...)
}

// GetCurrentFrame 获取当前帧
func (slf *Lockstep[ClientID, Command]) GetCurrentFrame() int64 {
	slf.currentFrameLock.RLock()
	defer slf.currentFrameLock.RUnlock()
	return slf.currentFrame
}

// GetClientCurrentFrame 获取客户端当前帧
func (slf *Lockstep[ClientID, Command]) GetClientCurrentFrame(clientId ClientID) int64 {
	slf.clientLock.RLock()
	defer slf.clientLock.RUnlock()
	return slf.clientFrame[clientId]
}

// GetFrameLimit 获取帧上限
//   - 未设置时将返回0
func (slf *Lockstep[ClientID, Command]) GetFrameLimit() int64 {
	return slf.frameLimit
}

// GetCurrentCommands 获取当前帧还未结束时的所有指令
func (slf *Lockstep[ClientID, Command]) GetCurrentCommands() []Command {
	slf.currentFrameLock.RLock()
	defer slf.currentFrameLock.RUnlock()
	return slf.currentCommands
}

// RegLockstepStoppedEvent 当广播停止时将触发被注册的事件处理函数
func (slf *Lockstep[ClientID, Command]) RegLockstepStoppedEvent(handle StoppedEventHandle[ClientID, Command]) {
	slf.lockstepStoppedEventHandles = append(slf.lockstepStoppedEventHandles, handle)
}

func (slf *Lockstep[ClientID, Command]) OnLockstepStoppedEvent() {
	for _, handle := range slf.lockstepStoppedEventHandles {
		handle(slf)
	}
}
