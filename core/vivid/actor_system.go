package vivid

import (
	"fmt"
	core "github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/kercylan98/minotaur/toolkit/pools"
	"sync"
	"sync/atomic"
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
	for _, plugin := range opts.modules {
		if transport, ok := plugin.(TransportModule); ok {
			if !address.IsEmpty() {
				panic("only one transport plugin is allowed")
			}
			address = transport.ActorSystemAddress().ParseToRoot()
			if address.System() == "" {
				address = core.NewRootAddress(address.Network(), opts.Name, address.Host(), address.Port())
			}
		}
	}
	if address.IsEmpty() {
		address = core.NewRootAddress("", opts.Name, "", 0)
	}

	system.processes = core.NewProcessManager(address, 128, system.deadLetter)
	system.deadLetter.ref, _ = system.processes.Register(system.deadLetter)
	support := newModuleSupport(system)
	system.root = spawn(system, func() Actor { return &root{} }, new(ActorOptions).WithName("user"), nil, mappings.NewOrderSync[core.Address, ActorRef](), nil)
	for _, plugin := range opts.modules {
		plugin.OnLoad(support)
	}
	return system
}

type ActorSystem struct {
	opts         *ActorSystemOptions
	processes    *core.ProcessManager
	deadLetter   *deadLetterProcess
	root         ActorContext
	closed       chan struct{}
	futurePool   *pools.ObjectPool[*future]
	nextFutureId atomic.Uint64
	kinds        map[Kind]*kind
	kindRw       sync.RWMutex
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
}

func (sys *ActorSystem) Context() ActorContext {
	return sys.root
}

func (sys *ActorSystem) Shutdown() {
	sys.root.Terminate(sys.root.Ref())
	<-sys.closed
}

func (sys *ActorSystem) ShutdownGracefully() {
	sys.root.TerminateGracefully(sys.root.Ref())
	<-sys.closed
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

func (sys *ActorSystem) KindOf(kind Kind) ActorRef {
	return sys.root.KindOf(kind)
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
		system.root.Ref()
	}
	if options.Name == "" {
		options.Name = convert.Uint64ToString(guid.Add(1))
	}

	var actorPath = options.Name
	if parent != nil {
		actorPath = parent.Address().Path() + "/" + options.Name
	} else {
		actorPath = "/" + options.Name
	}

	var processAddr = system.processes.Address()
	var address = core.NewAddress(processAddr.Network(), system.opts.Name, processAddr.Host(), processAddr.Port(), actorPath)
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
