package vivid

import (
	"github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/eventstream"
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

// EventStream 返回事件流，vivid 中 ActorSystem 采用的事件流为 eventstream.UnreliableSortStream，它的执行是在生产者 goroutine 中执行的，在处理函数中更改订阅者自身状态将会是危险的
func (s *ModuleSupport) EventStream() eventstream.Stream {
	return s.actorSystem.eventStream
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

// WrapRegulatoryMessage 包装监管消息
func (s *ModuleSupport) WrapRegulatoryMessage(sender, receiver ActorRef, message Message) RegulatoryMessage {
	return wrapRegulatoryMessage(sender, receiver, message)
}

// UnwrapRegulatoryMessage 解包监管消息
func (s *ModuleSupport) UnwrapRegulatoryMessage(message Message) (sender, receiver ActorRef, msg Message) {
	return unwrapRegulatoryMessage(message)
}

func (s *ModuleSupport) SendUserMessage(sender, receiver ActorRef, msg Message) {
	s.actorSystem.sendUserMessage(sender, receiver, msg)
}

func (s *ModuleSupport) SendSystemMessage(sender, receiver ActorRef, msg Message) {
	s.actorSystem.sendSystemMessage(sender, receiver, msg)
}
