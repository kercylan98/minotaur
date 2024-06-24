package vivid

type ActorOption func(options *ActorOptions)

type ActorOptions struct {
	options            []ActorOption
	Parent             ActorRef           // 父 Actor
	Name               string             // Actor 名称
	Dispatcher         Dispatcher         // Actor 使用的调度器
	Mailbox            Mailbox            // Actor 使用的邮箱
	SupervisorStrategy SupervisorStrategy // Actor 使用的监督者策略
}

// WithSupervisorStrategy 通过指定监督者策略创建一个 Actor
//   - 当 Actor 被通过该可选项创建时，假如其父 Actor 实现了 SupervisorStrategy，那么其父 Actor 的监管策略将被取代
func (o *ActorOptions) WithSupervisorStrategy(strategy SupervisorStrategy) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.SupervisorStrategy = strategy
	})
	return o
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
