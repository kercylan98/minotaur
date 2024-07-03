package vivid

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/log"
)

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

// Logger 返回日志记录器
func (s *ModuleSupport) Logger() *log.Logger {
	return s.System().opts.LoggerProvider()
}

// Address 返回 ActorSystem 地址
func (s *ModuleSupport) Address() core.Address {
	return s.actorSystem.processes.Address()
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
