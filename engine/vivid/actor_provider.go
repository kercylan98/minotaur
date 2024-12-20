package vivid

import "fmt"

// ActorProvider 是一个 Actor 生成器接口，它定义了生成 Actor 实例的方法。
type ActorProvider interface {
	// Provide 每次调用都应返回一个新的 Actor 实例，错误的使用可能导致 Actor 状态被污染。
	Provide() Actor
}

// FunctionalActorProvider 是一个函数类型的 Actor 生成器，它定义了生成 Actor 实例的方法。
type FunctionalActorProvider func() Actor

// Provide 每次调用都应返回一个新的 Actor 实例，错误的使用可能导致 Actor 状态被污染。
func (f FunctionalActorProvider) Provide() Actor {
	return f()
}

// FixedActorProvider 是一个固定 Actor 生成器接口，它定义了生成 Actor 实例的方法。
type FixedActorProvider interface {
	ProvideActor() (actor Actor)
	ProvideConfigurator() (configurator ActorDescriptorConfigurator)
}

// GetFixedActorProviders 获取固定 Actor 生成器列表
func GetFixedActorProviders(actorSystem *ActorSystem) map[string]FixedActorProvider {
	return actorSystem.config.fixedActorProviders
}

// SpawnActorFromFixedProvider 通过名称创建 Actor
func SpawnActorFromFixedProvider(actorSystem *ActorSystem, spawner Spawner, name string) (ActorRef, error) {
	provider, exist := actorSystem.config.fixedActorProviders[name]
	if !exist {
		return nil, fmt.Errorf("fixed actor provider[%v] not found", name)
	}

	return spawner.ActorOf(FunctionalActorProvider(func() Actor { return provider.ProvideActor() }), provider.ProvideConfigurator()), nil
}
