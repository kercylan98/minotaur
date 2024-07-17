package vivid

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/dispatcher"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/mailbox"
)

type MailboxProvider interface {
	// Create 创建邮箱
	Create(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox
}

type FunctionalMailboxProvider func(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox

func (f FunctionalMailboxProvider) Create(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox {
	return f(dispatcher, recipient)
}
