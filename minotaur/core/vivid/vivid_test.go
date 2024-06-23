package vivid

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
