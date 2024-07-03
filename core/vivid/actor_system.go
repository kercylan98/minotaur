package vivid

import (
	"fmt"
	core "github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/pools"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

var (
	_ SpawnerContext = &ActorSystem{}
)

func NewActorSystem(options ...func(options *ActorSystemOptions)) *ActorSystem {
	var opts = new(ActorSystemOptions).apply(options)
	system := &ActorSystem{
		opts:   opts,
		closed: make(chan struct{}),
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
	system.root = spawn(system, func() Actor { return &root{} }, new(ActorOptions).WithName("user"), nil, mappings.NewOrderSync[core.Address, ActorRef](), nil)
	for _, plugin := range opts.modules {
		plugin.OnLoad(support, transportModule)
	}
	return system
}

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
}

func (sys *ActorSystem) RegKind(k Kind, producer ActorProducer, options ...ActorOptionDefiner) {
	var opts = new(ActorOptions)
	for _, option := range options {
		option(opts)
	}
	opts.apply()
	sys.kindRw.Lock()
	defer sys.kindRw.Unlock()
	if sys.kinds == nil {
		sys.kinds = make(map[Kind]*kind)
	}
	if _, exist := sys.kinds[k]; exist {
		panic(fmt.Errorf("kind %s already exists", k))
	}
	sys.kinds[k] = &kind{producer: producer, options: opts}

	for _, module := range sys.kindHookModules {
		module.OnRegKind(k)
	}
}

func (sys *ActorSystem) Context() ActorContext {
	return sys.root
}

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

func (sys *ActorSystem) ShutdownGracefully(chanting ...time.Duration) {
	sys.chanting(chanting...)
	sys.root.TerminateGracefully(sys.root.Ref())
	<-sys.closed
}

func (sys *ActorSystem) chanting(duration ...time.Duration) {
	if len(duration) > 0 {
		select {
		case <-time.After(duration[0]):
		}
	}
}

func (sys *ActorSystem) Terminate(target ActorRef) {
	sys.root.Terminate(target)
}

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

func (sys *ActorSystem) ActorOf(producer ActorProducer, options ...ActorOptionDefiner) ActorRef {
	return sys.root.ActorOf(producer, options...)
}

func (sys *ActorSystem) KindOf(kind Kind, parent ...ActorRef) ActorRef {
	return sys.root.KindOf(kind, parent...)
}

func (sys *ActorSystem) internalActorOf(options *ActorOptions, producer ActorProducer, props []ActorOptionDefiner, generatedHook func(ctx *actorContext), guid *atomic.Uint64) ActorRef {
	for _, prop := range props {
		prop(options)
	}
	return spawn(sys, producer, options, generatedHook, mappings.NewOrder[core.Address, ActorRef](), guid).Ref()
}

func spawn(spawner SpawnerContext, producer ActorProducer, options *ActorOptions, generatedHook func(ctx *actorContext), childrenContainer mappings.OrderInterface[core.Address, ActorRef], guid *atomic.Uint64) ActorContext {
	options.apply()

	var system *ActorSystem
	switch v := spawner.(type) {
	case *ActorSystem:
		system = v
	case *actorContext:
		system = v.actorSystem
	}

	var parent ActorRef
	if options.Parent == nil && system.root != nil {
		parent = system.root.Ref()
	} else {
		parent = options.Parent
	}

	name := options.Name
	if name == "" {
		name = convert.Uint64ToString(guid.Add(1))
	}
	if options.NamePrefix != "" {
		name = options.NamePrefix + "-" + name
	}

	var actorPath = name
	if parent != nil {
		actorPath = parent.Address().Path() + "/" + name
	} else {
		actorPath = "/" + name
	}

	var parentAddr = system.processes.Address()
	if parent != nil {
		parentAddr = parent.Address()
	}
	var address = core.NewAddress(parentAddr.Network(), system.opts.Name, parentAddr.Host(), parentAddr.Port(), actorPath)
	system.opts.LoggerProvider().Debug("actorOf", log.String("addr", address.String()))

	var mailbox Mailbox
	if options.MailboxProducer == nil {
		mailbox = NewDefaultMailbox(0)
	} else {
		mailbox = options.MailboxProducer()
	}

	process := newProcess(address, mailbox)
	ref, exist := system.processes.Register(process)
	if exist {
		panic("actor already exists")
	}
	ctx := newActorContext(system, parent, options, producer, ref, childrenContainer)

	if generatedHook != nil {
		generatedHook(ctx)
	}

	if options.DispatcherProducer == nil {
		ctx.dispatcher = defaultDispatcher
	} else {
		ctx.dispatcher = options.DispatcherProducer()
	}
	mailbox.OnInit(ctx, ctx.dispatcher)
	ctx.System().sendSystemMessage(ctx.ref, ctx.ref, onLaunch)

	return ctx
}
