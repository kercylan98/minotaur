package game

import (
	"context"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"minotaur/utils/log"
	"minotaur/utils/super"
	"minotaur/utils/timer"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// StateMachine 游戏状态机，维护全局游戏信息
type StateMachine struct {
	*gateway                                         // 网关
	*router                                          // 路由器
	*timer.Manager                                   // 定时器
	*Time                                            // 游戏时间
	channelRWMutex  sync.RWMutex                     // 频道锁
	channels        map[any]*channel                 // 所有频道
	channelStrategy func() (channelId any, size int) // 频道策略
	errChannel      chan error                       // 异步错误管道
	loginTimeout    time.Duration                    // 玩家登录超时时间

	closeErrors  []error    // 关闭错误集
	closingMutex sync.Mutex // 关闭锁
	closed       bool       // 是否已关闭

	sonyflake *sonyflake.Sonyflake // 雪花id生成器
}

// SetLoginTimeout 设置玩家登录超时时间，登录超时时，玩家将被踢出游戏（默认无超时）
func (slf *StateMachine) SetLoginTimeout(duration time.Duration) {
	slf.loginTimeout = duration
}

// SetChannelStrategy 设置频道策略，返回频道id和最大容纳人数
func (slf *StateMachine) SetChannelStrategy(onGetChannelId func() (channelId any, size int)) {
	slf.channelStrategy = onGetChannelId
}

func (slf *StateMachine) Init() *StateMachine {
	slf.router = new(router).init(slf)
	slf.Time = new(Time)
	return slf
}

// Run 运行状态机
func (slf *StateMachine) Run(appName string, port int, onCreateConnHandleFunc OnCreateConnHandleFunc) {
	slf.errChannel = make(chan error, 1)
	slf.channels = map[any]*channel{}
	if slf.channelStrategy == nil {
		slf.channelStrategy = func() (channelId any, size int) {
			return 0, 0
		}
	}

	slf.gateway = new(gateway).run(slf, appName, port, onCreateConnHandleFunc)
	slf.Manager = timer.GetManager(64)
	slf.sonyflake = sonyflake.NewSonyflake(sonyflake.Settings{})

	log.Info("StateMachineRun", zap.String("stateMachine", "finish"))

	systemSignal := make(chan os.Signal, 1)
	signal.Notify(systemSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-systemSignal:
		slf.Shutdown(nil)
	case err := <-slf.errChannel:
		slf.Shutdown(err)
	}
}

// Shutdown 停止状态机
func (slf *StateMachine) Shutdown(err error) {
	if err != nil {
		slf.closeErrors = append(slf.closeErrors, err)
	}

	slf.closingMutex.Lock()
	defer slf.closingMutex.Unlock()
	if slf.closed {
		return
	}
	slf.closed = true

	var shutDownTimeoutContext, cancel = context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err = slf.gateway.shutdown(shutDownTimeoutContext); err != nil {
		slf.closeErrors = append(slf.closeErrors, err)
	}
	for _, c := range slf.channels {
		c.release()
	}

	slf.Manager.Release()
	if len(slf.closeErrors) > 0 {
		log.Error("StateMachineShutdown", zap.Errors("errors", slf.closeErrors))
	} else {
		log.Info("StateMachineShutdown", zap.String("stateMachine", "normal"))
	}
}

// channel 获取频道，如果频道不存在将会进行创建
//
// size: 指定频道最大容纳人数( <= 0 无限制)
// alone: 指定频道创建时候的逻辑是否以单线程运行 (default: true)
func (slf *StateMachine) channel(channelId any, size int, alone ...bool) *channel {
	slf.channelRWMutex.Lock()
	defer slf.channelRWMutex.Unlock()

	isAlone := true
	if len(alone) > 0 {
		isAlone = alone[0]
	}
	c := slf.channels[channelId]
	if c == nil {
		c = &channel{
			id:           channelId,
			size:         size,
			stateMachine: slf,
			alone:        isAlone,
			players:      map[int64]*Player{},
			timer:        timer.GetManager(64),
		}
		slf.channels[channelId] = c
		c.run()
		log.Info("ChannelCreate",
			zap.Any("channelID", channelId),
			zap.Bool("alone", isAlone),
			zap.Int("count", len(slf.channels)),
			zap.String("size", super.If(c.size <= 0, "NaN", strconv.Itoa(c.size))))
	}
	return c
}

// GenerateGuid 生成一个 Guid
func (slf *StateMachine) GenerateGuid() int64 {
	guid, err := slf.sonyflake.NextID()
	if err != nil {
		panic(err)
	}
	return int64(guid)
}
