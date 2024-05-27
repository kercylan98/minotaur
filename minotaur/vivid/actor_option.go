package vivid

type ActorOption[T Actor] func(opts *ActorOptions[T])

type ActorOptions[T Actor] struct {
	options          []ActorOption[T]
	Name             string                        // Actor 名称
	Parent           ActorContext                  // 父 actor 上下文
	DispatcherId     DispatcherId                  // 调度器 ID
	MailboxFactoryId MailboxFactoryId              // 邮箱工厂 ID
	Construct        T                             // Actor 构造器
	MessageHook      func(ctx MessageContext) bool // 消息钩子
}

// WithMessageHook 设置 Actor 的消息钩子，在消息处理前执行。返回 false 将会阻止消息的处理
func (o *ActorOptions[T]) WithMessageHook(hook func(ctx MessageContext) bool) *ActorOptions[T] {
	o.options = append(o.options, func(opts *ActorOptions[T]) {
		opts.MessageHook = hook
	})
	return o
}

// WithConstruct 以非零值的 Actor 作为 Actor 的初始对象
func (o *ActorOptions[T]) WithConstruct(actor T) *ActorOptions[T] {
	o.options = append(o.options, func(opts *ActorOptions[T]) {
		opts.Construct = actor
	})
	return o
}

// WithMailboxFactory 设置 Actor 使用的邮箱工厂，当邮箱工厂不存在时，将会导致 Actor 创建失败
func (o *ActorOptions[T]) WithMailboxFactory(mailboxFactoryId MailboxFactoryId) *ActorOptions[T] {
	o.options = append(o.options, func(opts *ActorOptions[T]) {
		opts.MailboxFactoryId = mailboxFactoryId
	})
	return o
}

// WithDispatcher 设置 Actor 使用的调度器，当调度器不存在时，将会导致 Actor 创建失败
func (o *ActorOptions[T]) WithDispatcher(dispatcherId DispatcherId) *ActorOptions[T] {
	o.options = append(o.options, func(opts *ActorOptions[T]) {
		opts.DispatcherId = dispatcherId
	})
	return o
}

// WithParent 设置 Actor 的父 Actor 上下文
func (o *ActorOptions[T]) WithParent(parent ActorContext) *ActorOptions[T] {
	o.options = append(o.options, func(opts *ActorOptions[T]) {
		opts.Parent = parent
	})
	return o
}

// WithName 设置 Actor 名称
func (o *ActorOptions[T]) WithName(name string) *ActorOptions[T] {
	o.options = append(o.options, func(opts *ActorOptions[T]) {
		opts.Name = name
	})
	return o
}

func NewActorOptions[T Actor]() *ActorOptions[T] {
	return &ActorOptions[T]{}
}

func NewFreeActorOptions[T any]() *ActorOptions[*FreeActor[T]] {
	return &ActorOptions[*FreeActor[T]]{}
}

func (o *ActorOptions[T]) applyOption(opts ...ActorOption[T]) *ActorOptions[T] {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *ActorOptions[T]) applyOptions(opts ...*ActorOptions[T]) *ActorOptions[T] {
	for _, opt := range opts {
		o.applyOption(opt.options...)
	}
	return o
}

func parseActorOptions[T Actor](options ...*ActorOptions[T]) *ActorOptions[T] {
	var opts *ActorOptions[T]
	if len(options) > 0 {
		applyed := false
		opts = options[0]
		for _, option := range options {
			if option == nil {
				continue
			}
			opts = opts.applyOptions(option)
			applyed = true
		}
		if !applyed {
			opts = NewActorOptions[T]()
		}
	} else {
		opts = NewActorOptions[T]()
	}
	if opts.DispatcherId == 0 {
		opts.DispatcherId = DefaultDispatcherId
	}
	if opts.MailboxFactoryId == 0 {
		opts.MailboxFactoryId = FIFOMailboxFactoryId
	}
	return opts
}
