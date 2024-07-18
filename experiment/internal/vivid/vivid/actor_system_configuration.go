package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/random"
	"os"
)

// newActorSystemConfiguration 创建 ActorSystemConfiguration 实例
func newActorSystemConfiguration() *ActorSystemConfiguration {
	return &ActorSystemConfiguration{
		actorSystemName: random.HostName(),
		loggerProvider: FunctionalLoggerProvider(func() *log.Logger {
			return log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions().WithLevel(log.LevelDebug).WithCallerSkip(5)))
		}),
	}
}

// ActorSystemConfiguration 是 ActorSystem 的配置
type ActorSystemConfiguration struct {
	actorSystemName string         // ActorSystem 名称
	loggerProvider  LoggerProvider // 日志提供者
}

// WithLoggerProvider 设置日志提供者
func (c *ActorSystemConfiguration) WithLoggerProvider(provider LoggerProvider) *ActorSystemConfiguration {
	c.loggerProvider = provider
	return c
}

// WithName 设置 ActorSystem 名称
func (c *ActorSystemConfiguration) WithName(name string) *ActorSystemConfiguration {
	c.actorSystemName = name
	return c
}
