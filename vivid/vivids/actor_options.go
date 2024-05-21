package vivids

// NewActorOptions 创建一个 ActorOptions
func NewActorOptions() *ActorOptions {
	return &ActorOptions{}
}

// ActorOptions 是 Actor 的配置项
type ActorOptions struct {
	Parent         ActorContext   // 父 Actor
	Mailbox        func() Mailbox // Actor 使用的邮箱
	Name           string         // Actor 名称
	DispatcherName string         // Actor 使用的调度器名称，如果为空则使用默认调度器
	Props          any            // Actor 的属性
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
		if opt.Parent != nil {
			o.Parent = opt.Parent
		}
		if opt.Props != nil {
			o.Props = opt.Props
		}
	}
	return o
}

// WithProps 设置 Actor 的属性
func (o *ActorOptions) WithProps(props any) *ActorOptions {
	o.Props = props
	return o
}

// WithName 设置 Actor 名称
func (o *ActorOptions) WithName(name string) *ActorOptions {
	o.Name = name
	return o
}

// WithMailbox 设置 Actor 使用的邮箱
func (o *ActorOptions) WithMailbox(mailbox func() Mailbox) *ActorOptions {
	o.Mailbox = mailbox
	return o
}

// WithDispatcherName 设置 Actor 使用的调度器名称
func (o *ActorOptions) WithDispatcherName(name string) *ActorOptions {
	o.DispatcherName = name
	return o
}

// WithParent 设置 Actor 的父 Actor
func (o *ActorOptions) WithParent(parent ActorContext) *ActorOptions {
	o.Parent = parent
	return o
}
