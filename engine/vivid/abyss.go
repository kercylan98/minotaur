package vivid

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid/internal/messages"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	// AbyssTopic 默认的深渊主题，订阅该主题的订阅者将会收到 OnAbyssMessageEvent 事件
	AbyssTopic = "vivid_abyss"
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
	switch message.(type) {
	case *OnAbyssMessageEvent, *messages.AbyssMessageEvent, *messages.LocalPublishRequest:
		return
	}

	if a.system.shared != nil {
		name, data, err := a.system.shared.GetCodec().Encode(message)
		if err == nil {
			a.system.Publish(AbyssTopic, &messages.AbyssMessageEvent{
				Sender:      sender,
				Receiver:    receiver,
				Forward:     forward,
				MessageType: name,
				Data:        data,
				Timestamp:   timestamppb.Now(),
			})
			return
		}
	}

	a.system.Publish(AbyssTopic, &OnAbyssMessageEvent{
		Sender:   sender,
		Receiver: receiver,
		Forward:  forward,
		Message:  message,
		Time:     time.Now(),
	})
}

func (a *abyss) DeliverySystemMessage(receiver, sender, forward *prc.ProcessId, message prc.Message) {
	switch message.(type) {
	case *messages.Watch:
		a.system.rc.GetProcess(sender).DeliverySystemMessage(sender, receiver, nil, &messages.Terminated{TerminatedProcess: receiver})
	default:
		a.system.Logger().Error("ActorSystem", log.String("info", "system abyss"), log.String("sender", sender.URL().String()), log.String("receiver", receiver.URL().String()), log.Any("message", message))
	}
}

func (a *abyss) Initialize(rc *prc.ResourceController, id *prc.ProcessId) {

}

func (a *abyss) IsTerminated() bool {
	return false
}

func (a *abyss) Terminate(source *prc.ProcessId) {

}
