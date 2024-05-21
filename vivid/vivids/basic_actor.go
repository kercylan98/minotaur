package vivids

// BasicActor 是 Actor 的基础实现，该实现仅提供了一个空的 Actor
type BasicActor struct{}

func (b *BasicActor) OnPreStart(ctx ActorContext) error {
	return nil
}

func (b *BasicActor) OnReceived(ctx MessageContext) error {
	return nil
}

func (b *BasicActor) OnDestroy(ctx ActorContext) error {
	return nil
}

func (b *BasicActor) OnChildTerminated(ctx ActorContext, child ActorTerminatedContext) {

}

func (b *BasicActor) OnSaveSnapshot(ctx ActorContext) (snapshot []byte, err error) {
	return
}

func (b *BasicActor) OnRecoverSnapshot(ctx ActorContext, snapshot []byte) (err error) {
	return
}

func (b *BasicActor) OnEvent(ctx ActorContext, event Message) (err error) {
	return
}
