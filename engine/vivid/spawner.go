package vivid

// Spawner 是生成 actor 的辅助接口
type Spawner interface {
	ActorOf(provider ActorProvider, configurator ...ActorDescriptorConfigurator) ActorRef
}
