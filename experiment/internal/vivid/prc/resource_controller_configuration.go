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
	clusterName     ClusterName        // 集群名称
	physicalAddress PhysicalAddress    // 物理地址
	loggerProvider  log.LoggerProvider // 日志提供者
}

// WithClusterName 设置集群名称
func (c *ResourceControllerConfiguration) WithClusterName(name ClusterName) *ResourceControllerConfiguration {
	c.clusterName = name
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
