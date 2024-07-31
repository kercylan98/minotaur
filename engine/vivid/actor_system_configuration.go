package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/prc/codec"
	"github.com/kercylan98/minotaur/engine/vivid/persistence"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
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
		mailboxProviderTable: map[MailboxProviderName]MailboxProvider{
			GetDefaultMailboxProvider().GetMailboxProviderName(): GetDefaultMailboxProvider(),
		},
		dispatcherProviderTable: map[DispatcherProviderName]DispatcherProvider{
			GetDefaultDispatcherProvider().GetDispatcherProviderName(): GetDefaultDispatcherProvider(),
		},
		persistenceStorageProviderTable: map[persistence.StorageProviderName]persistence.StorageProvider{
			GetDefaultPersistenceStorageProvider().GetStorageProviderName(): GetDefaultPersistenceStorageProvider(),
		},
		accidentTrace: true,
		abyss:         newAbyss(),
	}
}

// ActorSystemConfiguration 是 ActorSystem 的配置
type ActorSystemConfiguration struct {
	actorSystemName                  string                                                          // ActorSystem 名称
	physicalAddress                  prc.PhysicalAddress                                             // 物理地址
	loggerProvider                   log.LoggerProvider                                              // 日志提供者
	shared                           bool                                                            // 开启网络共享
	sharedCodec                      codec.Codec                                                     // 网络共享编解码器
	supervisionStrategyProviderTable map[supervision.StrategyName]supervision.StrategyProvider       // 监督策略表
	mailboxProviderTable             map[MailboxProviderName]MailboxProvider                         // 邮箱提供者表
	dispatcherProviderTable          map[DispatcherProviderName]DispatcherProvider                   // 调度器提供者表
	actorProviderTable               map[ActorProviderName]ActorProvider                             // Actor 提供者表
	persistenceStorageProviderTable  map[persistence.StorageProviderName]persistence.StorageProvider // 存储提供者表
	clusterJoinNodes                 []prc.PhysicalAddress                                           // 服务发现默认加入的节点
	clusterBindAddress               prc.PhysicalAddress                                             // 服务发现绑定的地址
	accidentTrace                    bool                                                            // 事故堆栈追踪
	abyss                            AbyssProcess                                                    // 深渊进程
}

// WithAbyss 设置深渊进程，深渊进程将在进程无法寻址到时作为替代进程进行返回。这在其他地方也叫做死信
func (c *ActorSystemConfiguration) WithAbyss(abyss AbyssProcess) *ActorSystemConfiguration {
	c.abyss = abyss
	return c
}

// WithAccidentTrace 开启事故堆栈追踪
func (c *ActorSystemConfiguration) WithAccidentTrace() *ActorSystemConfiguration {
	c.accidentTrace = true
	return c
}

// WithPersistenceStorageProvider 设置持久化存储提供者
func (c *ActorSystemConfiguration) WithPersistenceStorageProvider(providers ...persistence.StorageProvider) *ActorSystemConfiguration {
	if c.persistenceStorageProviderTable == nil {
		c.persistenceStorageProviderTable = make(map[persistence.StorageProviderName]persistence.StorageProvider)
	}
	for _, provider := range providers {
		name := provider.GetStorageProviderName()
		_, exist := c.persistenceStorageProviderTable[name]
		if exist {
			panic(fmt.Errorf("storage provider name %s already exist", name))
		}
		c.persistenceStorageProviderTable[name] = provider
	}
	return c
}

// WithActorProvider 设置 Actor 提供者
func (c *ActorSystemConfiguration) WithActorProvider(providers ...ActorProvider) *ActorSystemConfiguration {
	if c.actorProviderTable == nil {
		c.actorProviderTable = make(map[ActorProviderName]ActorProvider)
	}
	for _, provider := range providers {
		name := provider.GetActorProviderName()
		_, exist := c.actorProviderTable[name]
		if exist {
			panic(fmt.Errorf("actor provider name %s already exist", name))
		}
		c.actorProviderTable[name] = provider
	}
	return c
}

// WithMailboxProvider 设置邮箱提供者
func (c *ActorSystemConfiguration) WithMailboxProvider(providers ...MailboxProvider) *ActorSystemConfiguration {
	if c.mailboxProviderTable == nil {
		c.mailboxProviderTable = make(map[MailboxProviderName]MailboxProvider)
	}
	for _, provider := range providers {
		name := provider.GetMailboxProviderName()
		_, exist := c.mailboxProviderTable[name]
		if exist {
			panic(fmt.Errorf("mailbox provider name %s already exist", name))
		}
		c.mailboxProviderTable[name] = provider
	}
	return c
}

// WithDispatcherProvider 设置调度器提供者
func (c *ActorSystemConfiguration) WithDispatcherProvider(providers ...DispatcherProvider) *ActorSystemConfiguration {
	if c.dispatcherProviderTable == nil {
		c.dispatcherProviderTable = make(map[DispatcherProviderName]DispatcherProvider)
	}
	for _, provider := range providers {
		name := provider.GetDispatcherProviderName()
		_, exist := c.dispatcherProviderTable[name]
		if exist {
			panic(fmt.Errorf("dispatcher provider name %s already exist", name))
		}
		c.dispatcherProviderTable[name] = provider
	}
	return c
}

// WithCluster 通过集群模式创建
func (c *ActorSystemConfiguration) WithCluster(bindAddr prc.PhysicalAddress, defaultJoinNodes ...prc.PhysicalAddress) *ActorSystemConfiguration {
	c.clusterBindAddress = bindAddr
	c.clusterJoinNodes = defaultJoinNodes
	return c
}

// WithSupervisionStrategyProvider 设置监督策略提供者，将允许根据名称指定监管策略
func (c *ActorSystemConfiguration) WithSupervisionStrategyProvider(providers ...supervision.StrategyProvider) *ActorSystemConfiguration {
	if c.supervisionStrategyProviderTable == nil {
		c.supervisionStrategyProviderTable = make(map[supervision.StrategyName]supervision.StrategyProvider)
	}
	for _, provider := range providers {
		name := provider.GetStrategyProviderName()
		_, exist := c.supervisionStrategyProviderTable[name]
		if exist {
			panic(fmt.Errorf("strategy name %s already exist", name))
		}
		c.supervisionStrategyProviderTable[name] = provider
	}
	return c
}

// WithShared 设置是否开启网络共享，开启后 ActorSystem 将允许通过网络与其他 ActorSystem 交互。
//   - 默认的网络序列化是采用的 ProtoBuffer，如果需要调整，可指定编解码器
func (c *ActorSystemConfiguration) WithShared(address prc.PhysicalAddress, codec ...codec.Codec) *ActorSystemConfiguration {
	c.physicalAddress = address
	c.shared = true
	if len(codec) > 0 {
		c.sharedCodec = codec[0]
	}
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
