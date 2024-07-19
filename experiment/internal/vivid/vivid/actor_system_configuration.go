package vivid

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/random"
	"os"
)

// newActorSystemConfiguration 创建 ActorSystemConfiguration 实例
func newActorSystemConfiguration() *ActorSystemConfiguration {
	return &ActorSystemConfiguration{
		actorSystemName: random.HostName(),
		physicalAddress: prc.LocalhostPhysicalAddress,
		loggerProvider: log.FunctionalLoggerProvider(func() *log.Logger {
			return log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions().WithLevel(log.LevelDebug).WithCallerSkip(5)))
		}),
	}
}

// ActorSystemConfiguration 是 ActorSystem 的配置
type ActorSystemConfiguration struct {
	actorSystemName string              // ActorSystem 名称
	physicalAddress prc.PhysicalAddress // 物理地址
	loggerProvider  log.LoggerProvider  // 日志提供者
	shared          bool                // 开启网络共享
}

// WithShared 设置是否开启网络共享，开启后 ActorSystem 将允许通过网络与其他 ActorSystem 交互。
func (c *ActorSystemConfiguration) WithShared(shared bool) *ActorSystemConfiguration {
	c.shared = shared
	return c
}

// WithPhysicalAddress 设置物理地址
func (c *ActorSystemConfiguration) WithPhysicalAddress(address prc.PhysicalAddress) *ActorSystemConfiguration {
	c.physicalAddress = address
	return c
}

// WithLoggerProvider 设置日志提供者
func (c *ActorSystemConfiguration) WithLoggerProvider(provider log.LoggerProvider) *ActorSystemConfiguration {
	c.loggerProvider = provider
	return c
}

// WithName 设置 ActorSystem 名称
func (c *ActorSystemConfiguration) WithName(name string) *ActorSystemConfiguration {
	c.actorSystemName = name
	return c
}
