package vivid

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/dispatcher"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/mailbox"
)

var defaultMailboxProvider = FunctionalMailboxProvider(func(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox {
	return mailbox.NewLockFree(dispatcher, recipient)
})

// GetDefaultMailboxProvider 返回默认的 MailboxProvider 实例
func GetDefaultMailboxProvider() MailboxProvider {
	return defaultMailboxProvider
}

// MailboxProvider 是一个提供 mailbox.Mailbox 实例的接口
type MailboxProvider interface {
	// Provide 根据给定的 dispatcher.Dispatcher 和 mailbox.Recipient 返回一个 mailbox.Mailbox
	Provide(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox
}

// FunctionalMailboxProvider 是一个��数类型的 MailboxProvider，它定义了生成 mailbox.Mailbox 实例的方法。
type FunctionalMailboxProvider func(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox

// Provide 根据给定的 dispatcher.Dispatcher 和 mailbox.Recipient 返回一个 mailbox.Mailbox
func (f FunctionalMailboxProvider) Provide(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox {
	return f(dispatcher, recipient)
}
