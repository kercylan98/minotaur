package vivid

import "sync"

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
