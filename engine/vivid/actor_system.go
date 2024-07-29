package vivid

import (
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/network"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewActorSystem(configurator ...ActorSystemConfigurator) *ActorSystem {
	system := &ActorSystem{
		config: newActorSystemConfiguration(),
		closed: make(chan struct{}),
	}

	for _, systemConfigurator := range configurator {
		systemConfigurator.Configure(system.config)
	}

	system.rc = prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
		config.WithPhysicalAddress(system.config.physicalAddress)
		config.WithLoggerProvider(system.config.loggerProvider)
	}))

	if system.config.shared {
		system.shared = prc.NewShared(system.rc, prc.FunctionalSharedConfigurator(func(config *prc.SharedConfiguration) {
			config.WithRuntimeErrorHandler(prc.FunctionalErrorPolicyDecisionHandler(func(err error) prc.SharedPolicyDecision {
				return prc.SharedPolicyDecisionRestart
			}))
			config.WithConsecutiveRestartLimit(-1)
			config.WithRestartInterval(time.Millisecond*100, 3*time.Second)
			config.WithUnknownReceiverRedirect(func(message prc.Message) *prc.ProcessRef {
				return system.guard.ref
			})
			if system.config.sharedCodec != nil {
				config.WithCodec(system.config.sharedCodec)
			}
		}))
		if err := system.shared.Share(); err != nil {
			panic(err)
		}
	}

	if system.config.clusterBindAddress != charproc.None {
		host, port, err := network.NormalizeAddress(system.config.clusterBindAddress)
		if err != nil {
			panic(err)
		}
		system.discoverer = prc.NewDiscoverer(system.rc, host, port, prc.FunctionalDiscovererConfigurator(func(configuration *prc.DiscovererConfiguration) {
			configuration.WithJoinNodes(system.config.clusterJoinNodes...)
			configuration.WithName(system.config.actorSystemName)
			configuration.WithUserMetadata(packActorSystemMetadata(system))
		}))
		if err = system.discoverer.Discover(); err != nil {
			panic(err)
		}
	}

	system.Logger().Info("ActorSystem", log.String("status", "start"), log.String("name", system.config.actorSystemName))

	system.processId = prc.NewProcessId(system.rc.GetPhysicalAddress(), "/")
	system.guard = system.spawnTopActor("user", new(guard))

	return system
}

type ActorSystem struct {
	config     *ActorSystemConfiguration
	rc         *prc.ResourceController
	processId  *prc.ProcessId
	guard      *actorContext
	shared     *prc.Shared
	discoverer *prc.Discoverer
	closed     chan struct{}
}

// Name 获取 Actor 系统的名称
func (sys *ActorSystem) Name() string {
	return sys.config.actorSystemName
}

// Tell 向指定的 Actor 引用(ActorRef) 发送消息。
func (sys *ActorSystem) Tell(target ActorRef, message Message) {
	sys.guard.Tell(target, message)
}

// Ask 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是无限期的
func (sys *ActorSystem) Ask(target ActorRef, message Message) {
	sys.guard.Ask(target, message)
}

// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 future.Future 对象，可被用于获取响应消息
//   - 当 timeout 参数为空时，将会使用默认的超时时间 DefaultFutureAskTimeout
func (sys *ActorSystem) FutureAsk(target ActorRef, message Message, timeout ...time.Duration) future.Future[Message] {
	return sys.guard.FutureAsk(target, message, timeout...)
}

// AwaitForward 异步地等待阻塞结束后向目标 Actor 转发消息
func (sys *ActorSystem) AwaitForward(target ActorRef, asyncFunc func() Message) {
	sys.guard.AwaitForward(target, asyncFunc)
}

// Broadcast 向所有子级 Actor 广播消息，广播消息是可以被回复的
//   - 子级的子级不会收到广播消息
func (sys *ActorSystem) Broadcast(message Message) {
	sys.guard.Broadcast(message)
}

// Signal 监听系统信号，当系统接收到指定信号时执行 handler
func (sys *ActorSystem) Signal(handler func(system *ActorSystem, signal os.Signal)) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-s:
		handler(sys, sig)
	}
}

// Shutdown 关闭 Actor 系统。
//   - 该函数会等待所有 Actor 终止后再关闭 Actor 系统。
func (sys *ActorSystem) Shutdown(gracefully bool) {
	sys.guard.Terminate(sys.guard.ref, gracefully)
	<-sys.closed
	sys.shared.Dead()
	sys.Logger().Info("ActorSystem", log.String("status", "shutdown"), log.String("name", sys.config.actorSystemName))
}

// Terminate 终止目标 Actor。
//   - 当 gracefully 参数为 true 时，会将终止消息作为用户级消息进行发送，在该消息之前的用户消息被处理完毕后升级为系统消息终止 Actor。
func (sys *ActorSystem) Terminate(target ActorRef, gracefully bool) {
	sys.guard.Terminate(target, gracefully)
}

// ActorOf 生成一个新的 Actor 实例，并以该实例作为其父 Actor。返回生成的 Actor 引用(ActorRef)
//   - 该函数接收多个 ActorDescriptorConfigurator 参数，用于配置生成的 Actor 实例，当包含多个 ActorDescriptorConfigurator 参数时，它们的配置将会是向前覆盖的。
//
// 该函数不是并发安全的，你不应该在多个 goroutine 中同时调用 ActorOf 函数。
func (sys *ActorSystem) ActorOf(provider ActorProvider, configurator ...ActorDescriptorConfigurator) ActorRef {
	return sys.guard.ActorOf(provider, configurator...)
}

// ActorOfF 该函数是 ActorOf 的快捷方式，它提供了更为简便的使用方式，但是会额外创建一个切片并拷贝，用于 FunctionalActorDescriptorConfigurator 到 ActorDescriptorConfigurator 的转换。
func (sys *ActorSystem) ActorOfF(provider FunctionalActorProvider, configurator ...FunctionalActorDescriptorConfigurator) ActorRef {
	return sys.guard.ActorOfF(provider, configurator...)
}

// Logger 获取 ActorSystem 的日志记录器。
func (sys *ActorSystem) Logger() *log.Logger {
	return sys.config.loggerProvider.Provide()
}

func (sys *ActorSystem) spawnTopActor(name string, actor Actor, configurator ...ActorDescriptorConfigurator) *actorContext {
	var ac *actorContext
	configurator = append(configurator, FunctionalActorDescriptorConfigurator(func(descriptor *ActorDescriptor) {
		descriptor.WithName(name)
		descriptor.withInternalDescriptor(&actorInternalDescriptor{
			actorContextHook: func(ctx *actorContext) {
				ac = ctx
				ctx.parentRef = nil
			},
		})
	}))
	// 创建 Guard Actor
	(&actorContext{
		system:   sys,
		ref:      prc.NewProcessRef(sys.processId),
		children: make(map[prc.LogicalAddress]ActorRef),
	}).ActorOf(FunctionalActorProvider(func() Actor {
		return actor
	}), configurator...)
	return ac
}
