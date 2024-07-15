package vivid

import (
	"fmt"
	core "github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/eventstream"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/pools"
	"github.com/pkg/errors"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	_ SpawnerContext = &ActorSystem{}
)

// NewActorSystem 使用给定的选项和模块初始化一个新的 ActorSystem
func NewActorSystem(options ...func(options *ActorSystemOptions)) *ActorSystem {
	var opts = new(ActorSystemOptions).apply(options)
	system := &ActorSystem{
		opts:        opts,
		closed:      make(chan struct{}),
		eventStream: eventstream.NewUnreliableSortStream(),
	}
	system.deadLetter = &deadLetterProcess{system: system}

	var address core.Address
	var transportModule bool
	var priorityMap = make(map[Module]int)
	for _, module := range opts.modules {
		switch v := module.(type) {
		case TransportModule:
			if !address.IsEmpty() {
				panic("only one transport module is allowed")
			}
			address = v.ActorSystemAddress().ParseToRoot()
			if address.System() == "" {
				address = core.NewRootAddress(address.Network(), opts.Name, address.Host(), address.Port())
			}
			transportModule = true
			if p, ok := v.(PriorityModule); ok {
				priorityMap[v] = p.Priority()
			} else {
				priorityMap[v] = 0
			}
		case PriorityModule:
			priorityMap[v] = v.Priority()
		default:
			priorityMap[module] = 0
		}

		switch v := module.(type) {
		case KindHookModule:
			system.kindHookModules = append(system.kindHookModules, v)
		}
	}
	sort.Slice(opts.modules, func(i, j int) bool {
		return priorityMap[opts.modules[i]] < priorityMap[opts.modules[j]]
	})

	if address.IsEmpty() {
		address = core.NewRootAddress("", opts.Name, "", 0)
	}

	system.processes = core.NewProcessManager(address, 128, system.deadLetter)
	system.deadLetter.ref, _ = system.processes.Register(system.deadLetter)
	support := newModuleSupport(system)

	actorOpts := actorOptionsPool.Get().WithName("user")
	defer actorOptionsPool.Put(actorOpts)
	system.root, _, _ = spawn(system, func() Actor { return &root{} }, actorOpts, nil, mappings.NewOrderSync[core.Address, ActorRef](), nil, charproc.None)
	for _, plugin := range opts.modules {
		plugin.OnLoad(support, transportModule)
	}
	return system
}

// ActorSystem 在 Actor 模型设计中的核心重量级结构，负责管理 Actor、模块和系统进程信息，并提供 Actor 创建、销毁、消息传递等功能。
type ActorSystem struct {
	opts            *ActorSystemOptions
	processes       *core.ProcessManager
	deadLetter      *deadLetterProcess
	root            ActorContext
	closed          chan struct{}
	futurePool      *pools.ObjectPool[*future]
	nextFutureId    atomic.Uint64
	kinds           map[Kind]*kind
	kindRw          sync.RWMutex
	kindHookModules []KindHookModule
	eventStream     eventstream.Stream
}

// Name 返回 ActorSystem 的名称
func (sys *ActorSystem) Name() string {
	return sys.opts.Name
}

// Tell 向目标 Actor 发送消息
func (sys *ActorSystem) Tell(target ActorRef, message Message, options ...MessageOption) {
	sys.Context().Tell(target, message, options...)
}

// Ask 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是无限期的
func (sys *ActorSystem) Ask(target ActorRef, message Message, options ...MessageOption) {
	sys.Context().Ask(target, message, options...)
}

// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 Future 对象，可被用于获取响应消息
func (sys *ActorSystem) FutureAsk(target ActorRef, message Message, options ...MessageOption) Future {
	return sys.Context().FutureAsk(target, message, options...)
}

// Broadcast 向所有子级 Actor 广播消息，广播消息是可以被回复的
//   - 子级的子级不会收到广播消息
func (sys *ActorSystem) Broadcast(message Message, options ...MessageOption) {
	sys.Context().Broadcast(message, options...)
}

// AwaitForward 异步地等待阻塞结束后向目标 Actor 转发消息，收到的消息类型将是 FutureForwardMessage
func (sys *ActorSystem) AwaitForward(target ActorRef, blockFunc func() Message) {
	sys.Context().AwaitForward(target, blockFunc)
}

// DeadLetter 获取当前 Actor 系统的死信队列
func (sys *ActorSystem) DeadLetter() DeadLetter {
	return sys.deadLetter
}

// RegKind 注册一个 Kind，Kind 是 Actor 的类型，在通过 KindOf 创建 Actor 时将会根据预设的 Kind 创建 Actor
//   - 当 Kind 重复注册时将会发生 panic
func (sys *ActorSystem) RegKind(k Kind, producer ActorProducer, options ...ActorOptionDefiner) {
	var opts = newActorOptions() // 长期占用，不从池获取
	for _, option := range options {
		option(opts)
	}
	opts.apply()
	sys.kindRw.Lock()
	defer sys.kindRw.Unlock()
	if sys.kinds == nil {
		sys.kinds = make(map[Kind]*kind)
	}
	if k == charproc.None {
		panic(errors.New("kind cannot be none"))
	}
	if _, exist := sys.kinds[k]; exist {
		panic(fmt.Errorf("kind %s already exists", k))
	}
	sys.kinds[k] = &kind{producer: producer, options: opts}

	for _, module := range sys.kindHookModules {
		module.OnRegKind(k)
	}
}

// Context 返回 ActorSystem 的根 ActorContext
func (sys *ActorSystem) Context() ActorContext {
	return sys.root
}

// Shutdown 关闭 ActorSystem，关闭时会优先强行终止所有 Actor，之后对模块进行卸载
//   - chanting 为可选参数，表示在执行关闭操作前等待的时间
//
// 在关闭过程中，会从根 Actor 开始向下逐级终止 Actor，当子 Actor 全部终止后，父 Actor 才会终止。由于终止是系统消息，未被处理的用户消息将会被丢弃
func (sys *ActorSystem) Shutdown(chanting ...time.Duration) {
	sys.chanting(chanting...)
	sys.root.Terminate(sys.root.Ref())
	<-sys.closed
	for _, module := range sys.opts.modules {
		switch m := module.(type) {
		case ShutdownModule:
			m.OnShutdown()
		}
	}
}

// ShutdownGracefully 优雅地关闭 ActorSystem，与 Shutdown 不同的是， Shutdown 会致使 Actor 立即终止，而 ShutdownGracefully 则是发送一条用户消息，当
// Actor 接收到该消息并处理完毕后再终止
func (sys *ActorSystem) ShutdownGracefully(chanting ...time.Duration) {
	sys.chanting(chanting...)
	sys.root.TerminateGracefully(sys.root.Ref())
	<-sys.closed
	for _, module := range sys.opts.modules {
		switch m := module.(type) {
		case ShutdownModule:
			m.OnShutdown()
		}
	}
}

func (sys *ActorSystem) chanting(duration ...time.Duration) {
	if len(duration) > 0 {
		select {
		case <-time.After(duration[0]):
		}
	}
}

// Terminate 通知目标 Actor 立即终止，当 Actor 邮箱中还有未处理完的消息时，这些消息将会被丢弃
func (sys *ActorSystem) Terminate(target ActorRef) {
	sys.root.Terminate(target)
}

// TerminateGracefully 通知目标 Actor 立即终止，但是不会立即终止，而是在之前的用户消息处理完毕后终止
func (sys *ActorSystem) TerminateGracefully(target ActorRef) {
	sys.root.TerminateGracefully(target)
}

func (sys *ActorSystem) sendSystemMessage(sender, target ActorRef, message Message) {
	sys.getProcess(target).SendSystemMessage(sender, message)
}

func (sys *ActorSystem) sendUserMessage(sender, target ActorRef, message Message) {
	sys.getProcess(target).SendUserMessage(sender, message)
}

func (sys *ActorSystem) getProcess(target ActorRef) core.Process {
	return sys.processes.GetProcess(target)
}

// ActorOf 以根 Actor 为父级创建一个新的 Actor，返回新 Actor 的引用
func (sys *ActorSystem) ActorOf(producer ActorProducer, options ...ActorOptionDefiner) ActorRef {
	return sys.root.ActorOf(producer, options...)
}

// KindOf 以根 Actor 为父级创建一个新的 Actor，返回新 Actor 的引用，该 Actor 的类型由 kind 指定
//   - kind 为 Actor 的类型，需要在注册时指定
//   - parent 为可选参数，指定父级 Actor
//
// 该函数通常与 parent 参数结合后用于远程 Actor 的创建
func (sys *ActorSystem) KindOf(kind Kind, parent ...ActorRef) ActorRef {
	return sys.root.KindOf(kind, parent...)
}

func (sys *ActorSystem) internalActorOf(
	options *ActorOptions,
	producer ActorProducer,
	props []ActorOptionDefiner,
	generatedHook func(ctx *actorContext),
	guid *atomic.Uint64,
	kind Kind,
) (ActorRef, error) {
	for _, prop := range props {
		prop(options)
	}
	_, ref, err := spawn(sys, producer, options, generatedHook, mappings.NewOrder[core.Address, ActorRef](), guid, kind)
	return ref, err
}

// spawn 在创建时可能会存在冲突的情况，即 Actor 已存在，该情况会返回已存在的 ActorRef 及错误信息，可根据情况自行抉择
func spawn(
	spawner SpawnerContext,
	producer ActorProducer,
	options *ActorOptions,
	generatedHook func(ctx *actorContext),
	childrenContainer mappings.OrderInterface[core.Address, ActorRef],
	guid *atomic.Uint64,
	kind Kind,
) (ActorContext, ActorRef, error) {
	options.apply()

	// 获取 system
	var system *ActorSystem
	switch v := spawner.(type) {
	case *ActorSystem:
		system = v
	case *actorContext:
		system = v.actorSystem
	}

	// 获取或初始化父级 Actor
	var parent ActorRef
	if options.Parent == nil && system.root != nil {
		parent = system.root.Ref()
	} else {
		parent = options.Parent
	}

	// 初始化 Actor 名称，如果未指定名称则使用 guid 生成
	// Guid 来源于其父级 Actor
	name := options.Name
	if name == "" {
		name = convert.Uint64ToString(guid.Add(1))
	}
	if options.NamePrefix != "" {
		name = options.NamePrefix + "-" + name
	}

	// 初始化 Actor 路径
	var actorPath = name
	if parent != nil {
		if strings.HasPrefix(name, "/") {
			actorPath = parent.Address().LogicPath() + name
		} else {
			actorPath = parent.Address().LogicPath() + "/" + name
		}
	} else {
		actorPath = "/" + name // 根 Actor
	}

	// 初始化 Actor 地址
	var parentAddr = system.processes.Address()
	if parent != nil {
		parentAddr = parent.Address()
	}
	var address = core.NewAddress(parentAddr.Network(), system.opts.Name, parentAddr.Host(), parentAddr.Port(), actorPath)
	system.opts.LoggerProvider().Debug("actorOf", log.String("addr", address.String()))

	// 初始化 Actor 邮箱
	var mailbox Mailbox
	if options.MailboxProducer == nil {
		mailbox = NewDefaultMailbox(0)
	} else {
		mailbox = options.MailboxProducer()
	}

	// 初始化 Actor 进程
	process := newProcess(address, mailbox)
	ref, exist := system.processes.Register(process)
	if exist {
		return nil, ref, fmt.Errorf("%w: %s", ErrActorAlreadyExist, address.String())
	}

	// 初始化 Actor 上下文
	ctx := newActorContext(system, parent, options, producer, ref, childrenContainer, kind)
	if generatedHook != nil {
		generatedHook(ctx)
	}

	// 初始化调度器并绑定邮箱和 Actor 上下文
	if options.DispatcherProducer == nil {
		ctx.dispatcher = defaultDispatcher
	} else {
		ctx.dispatcher = options.DispatcherProducer()
	}
	mailbox.OnInit(ctx, ctx.dispatcher)

	ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onLaunch)

	return ctx, ctx.ref, nil
}
