package vivid

import "github.com/kercylan98/minotaur/toolkit/queues"

// NewActorOptions 创建一个 ActorOptions
func NewActorOptions() *ActorOptions {
	return &ActorOptions{
		Mailbox: defaultMailBox(),
	}
}

func defaultMailBox() func() *Mailbox {
	return func() *Mailbox {
		return NewMailbox(queues.NewFIFO[MessageContext]())
	}
}

// ActorOptions 是 Actor 的配置项
type ActorOptions struct {
	Name           string          // Actor 名称
	Mailbox        func() *Mailbox // Actor 使用的邮箱
	DispatcherName string          // Actor 使用的调度器名称，如果为空则使用默认调度器
}

// Apply 应用配置项
func (o *ActorOptions) Apply(opts ...*ActorOptions) *ActorOptions {
	for _, opt := range opts {
		if opt.Name != "" {
			o.Name = opt.Name
		}
		if opt.Mailbox != nil {
			o.Mailbox = opt.Mailbox
		}
		if opt.DispatcherName != "" {
			o.DispatcherName = opt.DispatcherName
		}
	}
	return o
}

// WithName 设置 Actor 名称
func (o *ActorOptions) WithName(name string) *ActorOptions {
	o.Name = name
	return o
}

// WithMailbox 设置 Actor 使用的邮箱
func (o *ActorOptions) WithMailbox(mailbox func() *Mailbox) *ActorOptions {
	o.Mailbox = mailbox
	return o
}

// WithDispatcherName 设置 Actor 使用的调度器名称
func (o *ActorOptions) WithDispatcherName(name string) *ActorOptions {
	o.DispatcherName = name
	return o
}
