package mailbox

import (
	"github.com/kercylan98/minotaur/engine/prc"
)

const (
	mailboxStatusIdle uint32 = iota
	mailboxStatusRunning
)

type Mailbox interface {
	// Suspend 暂停邮箱
	Suspend()

	// Resume 恢复邮箱
	Resume()

	// DeliveryUserMessage 投递用户消息到邮箱
	DeliveryUserMessage(message prc.Message)

	// DeliverySystemMessage 投递系统消息到邮箱
	DeliverySystemMessage(message prc.Message)
}
