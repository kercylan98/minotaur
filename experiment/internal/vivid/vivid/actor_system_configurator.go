package vivid

// ActorSystemConfigurator 是用于配置 ActorSystem 的配置器
type ActorSystemConfigurator interface {
	// Configure 配置 ActorSystem
	Configure(config *ActorSystemConfiguration)
}

// FunctionalActorSystemConfigurator 是用于配置 ActorSystem 的配置器
type FunctionalActorSystemConfigurator func(config *ActorSystemConfiguration)

// Configure 配置 ActorSystem
func (f FunctionalActorSystemConfigurator) Configure(config *ActorSystemConfiguration) {
	f(config)
}
