package vivid

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/prc/codec"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/random"
	"google.golang.org/grpc"
	"os"
)

type (
	ShutdownAfterHook func() // 在 ActorSystem 被关闭后将调用此回调
)

// NewActorSystemConfiguration 创建 ActorSystemConfiguration 默认实例
func NewActorSystemConfiguration() *ActorSystemConfiguration {
	return &ActorSystemConfiguration{
		actorSystemName: random.HostName(),
		physicalAddress: prc.LocalhostPhysicalAddress,
		loggerProvider: log.FunctionalLoggerProvider(func() *log.Logger {
			return log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions().WithLevel(log.LevelDebug).WithCallerSkip(5)))
		}),
		accidentTrace: true,
		abyss:         newAbyss(),
	}
}

// ActorSystemConfiguration 是 ActorSystem 的配置
type ActorSystemConfiguration struct {
	actorSystemName              string                        // ActorSystem 名称
	physicalAddress              prc.PhysicalAddress           // 物理地址（透传给 prc.Shared）
	loggerProvider               log.LoggerProvider            // 日志提供者
	shared                       bool                          // 开启网络共享
	sharedCodec                  codec.Codec                   // 网络共享编解码器
	accidentTrace                bool                          // 事故堆栈追踪
	abyss                        AbyssProcess                  // 深渊进程
	shutdownAfterHooks           []ShutdownAfterHook           // ActorSystem 关闭后将调用此回调
	grpcServerHooks              []func(server *grpc.Server)   // GRPC 服务器钩子
	subscriptionContactProviders []SubscriptionContactProvider // 订阅联系人提供者
}

// WithSubscriptionContactProviders 设置订阅联系人提供者
func (c *ActorSystemConfiguration) WithSubscriptionContactProviders(providers ...SubscriptionContactProvider) *ActorSystemConfiguration {
	c.subscriptionContactProviders = append(c.subscriptionContactProviders, providers...)
	return c
}

// WithGRPCServerHooks 设置 GRPC 服务器钩子，该方法将在创建 GRPC 服务器后调用
func (c *ActorSystemConfiguration) WithGRPCServerHooks(hooks ...func(server *grpc.Server)) *ActorSystemConfiguration {
	c.grpcServerHooks = append(c.grpcServerHooks, hooks...)
	return c
}

// WithShutdownAfterHooks 设置 ActorSystem 关闭后将调用此回调
func (c *ActorSystemConfiguration) WithShutdownAfterHooks(hooks ...ShutdownAfterHook) *ActorSystemConfiguration {
	c.shutdownAfterHooks = append(c.shutdownAfterHooks, hooks...)
	return c
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

// WithShared 设置是否开启网络共享，开启后 ActorSystem 将允许通过网络与其他 ActorSystem 交互。
//   - 默认的网络序列化是采用的 ProtoBuffer，如果需要调整，可指定编解码器
func (c *ActorSystemConfiguration) WithShared(address prc.PhysicalAddress, codec ...codec.Codec) *ActorSystemConfiguration {
	if address == "localhost" {
		address += ":0"
	}
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
