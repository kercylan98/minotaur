package vivid

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/toolkit/log"
)

var _ AbyssProcess = (*abyss)(nil)

func newAbyss() *abyss {
	return &abyss{}
}

// 深渊是一个特殊的进程，它用于在资源控制器中无法查找到指定进程时作为替补选项
type abyss struct {
	system *ActorSystem
}

func (a *abyss) OnInitialize(system *ActorSystem) {
	a.system = system
}

func (a *abyss) DeliveryUserMessage(receiver, sender, forward *prc.ProcessId, message prc.Message) {
	a.system.Logger().Error("ActorSystem", log.String("info", "user abyss"), log.String("sender", sender.URL().String()), log.String("receiver", receiver.URL().String()), log.Any("message", message))
}

func (a *abyss) DeliverySystemMessage(receiver, sender, forward *prc.ProcessId, message prc.Message) {
	a.system.Logger().Error("ActorSystem", log.String("info", "system abyss"), log.String("sender", sender.URL().String()), log.String("receiver", receiver.URL().String()), log.Any("message", message))
}

func (a *abyss) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {
	panic("abyss cannot be initialized")
}

func (a *abyss) IsTerminated() bool {
	panic("abyss is eternal")
}

func (a *abyss) Terminate(source *prc.ProcessId) {
	panic("abyss cannot be terminated")
}
