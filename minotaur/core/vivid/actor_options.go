package vivid

type ActorOption func(options *ActorOptions)

type ActorOptions struct {
	options    []ActorOption
	Parent     ActorRef   // 父 Actor
	Name       string     // Actor 名称
	Dispatcher Dispatcher // Actor 使用的调度器
	Mailbox    Mailbox    // Actor 使用的邮箱
}

// WithDispatcher 通过指定调度器创建一个 Actor
func (o *ActorOptions) WithDispatcher(dispatcher Dispatcher) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.Dispatcher = dispatcher
	})
	return o
}

// WithMailbox 通过指定邮箱创建一个 Actor
func (o *ActorOptions) WithMailbox(mailbox Mailbox) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.Mailbox = mailbox
	})
	return o
}

// WithName 通过指定名称创建一个 Actor
func (o *ActorOptions) WithName(name string) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.Name = name
	})
	return o
}

// WithParent 通过指定父 Actor 创建一个 Actor
func (o *ActorOptions) WithParent(parent ActorRef) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.Parent = parent
	})
	return o
}

func (o *ActorOptions) apply() *ActorOptions {
	for _, option := range o.options {
		option(o)
	}
	return o
}
