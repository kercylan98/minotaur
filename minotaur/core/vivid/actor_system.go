package vivid

import (
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/minotaur/core"
	"path"
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
	system.processes = core.NewProcessManager("", 1, 100, system.deadLetter)
	system.deadLetter.ref, _ = system.processes.Register(system.deadLetter)

	system.root = spawn(system, new(root), new(ActorOptions).WithName("user"), nil)
	return system
}

type ActorSystem struct {
	processes  *core.ProcessManager
	deadLetter *deadLetterProcess
	root       ActorContext
	name       string
	closed     chan struct{}
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

func (sys *ActorSystem) ActorOf(producer ActorProducer) ActorRef {
	return sys.root.ActorOf(producer)
}

func (sys *ActorSystem) internalActorOf(options *ActorOptions, producer func(options *ActorOptions) Actor, generatedHook func(ctx *actorContext)) ActorRef {
	actor := producer(options)
	return spawn(sys, actor, options, generatedHook).Ref()
}

func spawn(spawner SpawnerContext, actor Actor, options *ActorOptions, generatedHook func(ctx *actorContext)) ActorContext {
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
		options.Name = uuid.NewString()
	}

	var actorPath = options.Name
	if options.Parent != nil {
		actorPath = path.Join(options.Parent.Address().Path(), options.Name)
	} else {
		actorPath = path.Clean("/" + options.Name)
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
	ctx := newActorContext(system, actor, options.Parent, ref, mailbox)

	if generatedHook != nil {
		generatedHook(ctx)
	}

	dispatcher := options.Dispatcher
	if dispatcher == nil {
		dispatcher = defaultDispatcher
	}
	mailbox.OnInit(ctx, dispatcher)
	ctx.Tell(ref, onBoot)

	return ctx
}
