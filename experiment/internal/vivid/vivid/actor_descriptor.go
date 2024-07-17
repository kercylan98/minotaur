package vivid

// ActorDescriptor 用于定义 Actor 个性化行为的描述符
type ActorDescriptor struct {
	mailboxProvider MailboxProvider // 邮箱提供者
}

// WithMailboxProvider 设置邮箱提供者
func (d *ActorDescriptor) WithMailboxProvider(provider MailboxProvider) *ActorDescriptor {
	d.mailboxProvider = provider
	return d
}
