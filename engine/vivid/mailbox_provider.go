package vivid

import (
	"github.com/kercylan98/minotaur/engine/vivid/dispatcher"
	"github.com/kercylan98/minotaur/engine/vivid/mailbox"
	"github.com/kercylan98/minotaur/toolkit/charproc"
)

// MailboxProviderName 是一个字符串类型的 MailboxProvider 名称
type MailboxProviderName = string

// MailboxProvider 是一个提供 mailbox.Mailbox 实例的接口
type MailboxProvider interface {
	// GetMailboxProviderName 返回 MailboxProvider 的名称
	GetMailboxProviderName() MailboxProviderName

	// Provide 根据给定的 dispatcher.Dispatcher 和 mailbox.Recipient 返回一个 mailbox.Mailbox
	Provide(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox
}

// FunctionalMailboxProvider 是一个��数类型的 MailboxProvider，它定义了生成 mailbox.Mailbox 实例的方法。
type FunctionalMailboxProvider func(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox

// GetMailboxProviderName 返回 MailboxProvider 的名称
func (f FunctionalMailboxProvider) GetMailboxProviderName() MailboxProviderName {
	return charproc.None
}

// Provide 根据给定的 dispatcher.Dispatcher 和 mailbox.Recipient 返回一个 mailbox.Mailbox
func (f FunctionalMailboxProvider) Provide(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox {
	return f(dispatcher, recipient)
}

var defaultMailboxProviderInstance = new(defaultMailboxProvider)

// GetDefaultMailboxProvider 返回默认的 MailboxProvider 实例
func GetDefaultMailboxProvider() MailboxProvider {
	return defaultMailboxProviderInstance
}

type defaultMailboxProvider struct {
}

func (d *defaultMailboxProvider) GetMailboxProviderName() MailboxProviderName {
	return "__default"
}

func (d *defaultMailboxProvider) Provide(dispatcher dispatcher.Dispatcher, recipient mailbox.Recipient) mailbox.Mailbox {
	return mailbox.NewLockFree(dispatcher, recipient)
}
