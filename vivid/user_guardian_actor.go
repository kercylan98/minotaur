package vivid

// userGuardianActor 用户守护 Actor
type userGuardianActor struct {
}

func (u *userGuardianActor) OnPreStart(ctx ActorContext) error {
	return nil
}

func (u *userGuardianActor) OnReceived(ctx MessageContext) error {
	return nil
}

func (u *userGuardianActor) OnDestroy(ctx ActorContext) error {
	return nil
}
