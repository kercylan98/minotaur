package vivid

const (
	DefaultPersistenceEventLimit = 200
)

type ActorOption func(options *ActorOptions)

type ActorOptions struct {
	options               []ActorOption
	Parent                ActorRef           // 父 Actor
	Name                  string             // Actor 名称
	NamePrefix            string             // Actor 名称前缀
	DispatcherProducer    DispatcherProducer // Actor 使用的调度器
	MailboxProducer       MailboxProducer    // Actor 使用的邮箱
	SupervisorStrategy    SupervisorStrategy // Actor 使用的监督者策略
	PersistenceName       string             // Actor 持久化名称
	PersistenceStorage    Storage            // Actor 持久化存储器
	PersistenceEventLimit int                // Actor 持久化事件数量限制，达到限制时将会触发快照的生成
}

// WithNamePrefix 通过指定名称前缀创建一个 Actor
func (o *ActorOptions) WithNamePrefix(prefix string) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.NamePrefix = prefix
	})
	return o
}

// WithPersistenceEventLimit 通过指定持久化事件数量限制创建一个 Actor
func (o *ActorOptions) WithPersistenceEventLimit(limit int) *ActorOptions {
	if limit < 0 {
		limit = DefaultPersistenceEventLimit
	}
	o.options = append(o.options, func(options *ActorOptions) {
		options.PersistenceEventLimit = limit
	})
	return o
}

// WithPersistence 通过指定持久化存储器和名称创建一个 Actor
func (o *ActorOptions) WithPersistence(storage Storage, name string) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.PersistenceStorage = storage
		options.PersistenceName = name
	})
	return o
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
func (o *ActorOptions) WithDispatcher(producer DispatcherProducer) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.DispatcherProducer = producer
	})
	return o
}

// WithMailbox 通过指定邮箱创建一个 Actor
func (o *ActorOptions) WithMailbox(producer MailboxProducer) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.MailboxProducer = producer
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
