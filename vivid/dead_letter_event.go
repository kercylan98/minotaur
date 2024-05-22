package vivid

import (
	"time"
)

const (
	DeadLetterEventTypeUnknown DeadLetterEventType = iota // 未知的死信事件类型
	DeadLetterEventTypeActorOf                            // ActorOf 创建失败事件
	DeadLetterEventTypeMessage                            // 消息发送失败事件
)

var deadLetterEventTypeStrings = map[DeadLetterEventType]string{
	DeadLetterEventTypeUnknown: "Unknown",
	DeadLetterEventTypeActorOf: "ActorOf",
	DeadLetterEventTypeMessage: "Message",
}

// DeadLetterEventType 死信事件类型
type DeadLetterEventType = uint8

// NewDeadLetterEvent 创建一个新的死信事件
func NewDeadLetterEvent(typ DeadLetterEventType, event any) DeadLetterEvent {
	return DeadLetterEvent{
		Type:  typ,
		Time:  time.Now(),
		Event: event,
	}
}

// DeadLetterEvent 死信事件，其中包含了整个 ActorSystem 中的失败消息
type DeadLetterEvent struct {
	Seq   uint64              // 死信事件序号
	Type  DeadLetterEventType // 死信事件类型
	Time  time.Time           // 死信事件发生时间
	Event any                 // 死信事件
	Stack []byte              // 死信事件堆栈
}

type DeadLetterEventActorOf struct {
	Error  error        // ActorOf 创建失败的错误
	Parent ActorContext // 父 ActorContext
	Name   ActorName    // 子 Actor 名称
}

type DeadLetterEventMessage struct {
	Error   error   // 消息发送失败的错误
	From    ActorId // 消息发送者
	To      ActorId // 消息接收者
	Message Message // 消息内容
	Seq     uint64  // 消息序号
}
