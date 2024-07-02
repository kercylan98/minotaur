package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"os"
	"sync"
)

var testLogger = log.New(log.NewHandler(os.Stdout, log.NewDevHandlerOptions()))
var benchmarkLogger = log.NewSilentLogger()

func NewTestActorSystem(options ...func(options *ActorSystemOptions)) *ActorSystem {
	return NewActorSystem(append(options, func(options *ActorSystemOptions) {
		options.WithLoggerProvider(func() *log.Logger {
			return testLogger
		})
	})...)
}

func NewBenchmarkActorSystem(options ...func(options *ActorSystemOptions)) *ActorSystem {
	return NewActorSystem(append(options, func(options *ActorSystemOptions) {
		options.WithLoggerProvider(func() *log.Logger {
			return benchmarkLogger
		})
	})...)
}

type WasteActor struct {
}

func (w *WasteActor) OnReceive(ctx ActorContext) {

}

type StringEchoActor struct {
}

func (e *StringEchoActor) OnReceive(ctx ActorContext) {
	switch m := ctx.Message().(type) {
	case string:
		ctx.Reply(m)
	}
}

type StringEchoCounterActor struct {
	Counter int
}

func (e *StringEchoCounterActor) OnReceive(ctx ActorContext) {
	switch m := ctx.Message().(type) {
	case string:
		ctx.Reply(m)
		e.Counter++
	}
}

var testActorSystemCounter = make(map[*ActorSystem]*sync.WaitGroup)

func (sys *ActorSystem) Add(delta int) {
	if _, ok := testActorSystemCounter[sys]; !ok {
		testActorSystemCounter[sys] = new(sync.WaitGroup)
	}
	testActorSystemCounter[sys].Add(delta)
}
func (sys *ActorSystem) Done() {
	testActorSystemCounter[sys].Done()
}

func (sys *ActorSystem) Wait() {
	testActorSystemCounter[sys].Wait()
	sys.Shutdown()
}
