package prc

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"os"
)

// newResourceControllerConfiguration 创建 ResourceControllerConfiguration 实例
func newResourceControllerConfiguration() *ResourceControllerConfiguration {
	return &ResourceControllerConfiguration{
		physicalAddress: LocalhostPhysicalAddress,
		loggerProvider: log.FunctionalLoggerProvider(func() *log.Logger {
			return log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions().WithLevel(log.LevelDebug).WithCallerSkip(5)))
		}),
	}
}

// ResourceControllerConfiguration 是 ActorSystem 的配置
type ResourceControllerConfiguration struct {
	physicalAddress    PhysicalAddress    // 物理地址
	loggerProvider     log.LoggerProvider // 日志提供者
	notFoundSubstitute Process            // 未找到处理器的替代处理器
}

// WithNotFoundSubstitute 设置未找到处理器的替代处理器
func (c *ResourceControllerConfiguration) WithNotFoundSubstitute(substitute Process) *ResourceControllerConfiguration {
	c.notFoundSubstitute = substitute
	return c
}

// WithPhysicalAddress 设置物理地址
func (c *ResourceControllerConfiguration) WithPhysicalAddress(address PhysicalAddress) *ResourceControllerConfiguration {
	c.physicalAddress = address
	return c
}

// WithLoggerProvider 设置日志提供者
func (c *ResourceControllerConfiguration) WithLoggerProvider(provider log.LoggerProvider) *ResourceControllerConfiguration {
	c.loggerProvider = provider
	return c
}
