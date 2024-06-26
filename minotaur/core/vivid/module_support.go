package vivid

import "github.com/kercylan98/minotaur/minotaur/core"

func newModuleSupport(actorSystem *ActorSystem) *ModuleSupport {
	return &ModuleSupport{
		actorSystem: actorSystem,
	}
}

type ModuleSupport struct {
	actorSystem *ActorSystem
}

// System 返回 ActorSystem
func (s *ModuleSupport) System() *ActorSystem {
	return s.actorSystem
}

// RegAddressResolver 注册地址解析器
func (s *ModuleSupport) RegAddressResolver(resolver core.AddressResolver) {
	s.actorSystem.processes.RegisterAddressResolver(resolver)
}

// SetAddressResolverOnlyAddress 设置地址解析器不包含额外的 path 信息
func (s *ModuleSupport) SetAddressResolverOnlyAddress() {
	s.actorSystem.processes.SetAddressResolverOnlyAddress()
}

// GetProcess 获取进程
func (s *ModuleSupport) GetProcess(address core.Address) core.Process {
	return s.actorSystem.processes.GetProcess(core.NewProcessRef(address))
}

// GetDeadLetter 获取死信队列
func (s *ModuleSupport) GetDeadLetter() DeadLetter {
	return s.actorSystem.deadLetter
}
