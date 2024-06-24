package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/collection/mappings"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/pools"
	"sync/atomic"
)

var (
	_ SpawnerContext = &ActorSystem{}
)

func NewActorSystem(name string) *ActorSystem {
	system := &ActorSystem{
		name:       name,
		closed:     make(chan struct{}),
		deadLetter: new(deadLetterProcess),
	}
	system.processes = core.NewProcessManager("", 1, 128, system.deadLetter)
	system.deadLetter.ref, _ = system.processes.Register(system.deadLetter)

	system.root = spawn(system, func() Actor { return &root{} }, new(ActorOptions).WithName("user"), nil, mappings.NewOrderSync[core.Address, ActorRef]())
	return system
}

type ActorSystem struct {
	processes    *core.ProcessManager
	deadLetter   *deadLetterProcess
	root         ActorContext
	name         string
	closed       chan struct{}
	futurePool   *pools.ObjectPool[*future]
	nextFutureId atomic.Uint64
	nextActorId  atomic.Uint64
}

func (sys *ActorSystem) Context() ActorContext {
	return sys.root
}

func (sys *ActorSystem) Shutdown() {
	sys.root.Terminate(sys.root.Ref())
}

func (sys *ActorSystem) Terminate(target ActorRef) {
	sys.sendSystemMessage(sys.root.Ref(), target, onTerminate)
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

func (sys *ActorSystem) internalActorOf(options *ActorOptions, producer ActorProducer, props []ActorOptionDefiner, generatedHook func(ctx *actorContext)) ActorRef {
	for _, prop := range props {
		prop(options)
	}
	return spawn(sys, producer, options, generatedHook, mappings.NewOrder[core.Address, ActorRef]()).Ref()
}

func spawn(spawner SpawnerContext, producer ActorProducer, options *ActorOptions, generatedHook func(ctx *actorContext), childrenContainer mappings.OrderInterface[core.Address, ActorRef]) ActorContext {
	options.apply()

	var system *ActorSystem
	switch v := spawner.(type) {
	case *ActorSystem:
		system = v
	case *actorContext:
		system = v.actorSystem
	}

	if options.Parent == nil {
		if parent := system.root; parent != nil {
			options.Parent = parent.Ref()
		}
	}
	if options.Name == "" {
		options.Name = convert.Uint64ToString(system.nextActorId.Add(1))
	}

	var actorPath = options.Name
	if options.Parent != nil {
		actorPath = options.Parent.Address().Path() + "/" + options.Name
	} else {
		actorPath = "/" + options.Name
	}

	var address = core.NewAddress("", system.name, "", 0, actorPath)

	mailbox := options.Mailbox
	if mailbox == nil {
		mailbox = newDefaultMailbox()
	}

	process := NewProcess(address, mailbox)
	ref, exist := system.processes.Register(process)
	if exist {
		panic("actor already exists")
	}
	ctx := newActorContext(system, producer, options.Parent, ref, mailbox, childrenContainer)

	if generatedHook != nil {
		generatedHook(ctx)
	}

	dispatcher := options.Dispatcher
	if dispatcher == nil {
		dispatcher = defaultDispatcher
	}
	mailbox.OnInit(ctx, dispatcher)
	ctx.Tell(ref, onLaunch)

	return ctx
}
