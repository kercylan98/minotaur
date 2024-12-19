package vivid

import "fmt"

// Spawner 是生成 actor 的辅助接口
type Spawner interface {
	ActorOf(provider ActorProvider, configurator ...ActorDescriptorConfigurator) ActorRef
}

// SpawnActorFromFixedProvider 通过名称创建 Actor
func SpawnActorFromFixedProvider(actorSystem *ActorSystem, spawner Spawner, name string) (ActorRef, error) {
	provider, exist := actorSystem.config.fixedActorProviders[name]
	if !exist {
		return nil, fmt.Errorf("fixed actor provider[%v] not found", name)
	}

	return spawner.ActorOf(FunctionalActorProvider(func() Actor { return provider.ProvideActor() }), provider.ProvideConfigurator()), nil
}
