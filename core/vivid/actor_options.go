package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/pools"
)

const (
	DefaultPersistenceEventLimit = 200 // 默认触发快照持久化的事件数量
)

var actorOptionsPool = pools.NewObjectPool[ActorOptions](func() *ActorOptions {
	return newActorOptions()
}, func(data *ActorOptions) {
	data.options = data.options[:0]
	data.Parent = nil
	data.Name = ""
	data.NamePrefix = ""
	data.DispatcherProducer = nil
	data.MailboxProducer = nil
	data.SupervisorStrategy = nil
	data.PersistenceName = ""
	data.PersistenceStorage = nil
	data.PersistenceEventLimit = 0
	data.ConflictReuse = false
})

func newActorOptions() *ActorOptions {
	return &ActorOptions{}
}

// ActorOption 对 ActorOptions 进行编辑的操作函数
type ActorOption func(options *ActorOptions)

// ActorOptions 是 Actor 的可选项，用于在创建之初对 Actor 进行配置
//   - 可选项在创建完成 Actor 之后便不会继续保留
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
	ConflictReuse         bool               // Actor 冲突重用，当 Actor 已存在时将会重用已存在的 Actor
	Scheduler             bool               // 是否使用定时器
}

// WithScheduler 通过指定是否使用定时器创建一个 Actor，定时器的执行将会通过系统消息进行执行
func (o *ActorOptions) WithScheduler(enable bool) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.Scheduler = enable
	})
	return o
}

// WithConflictReuse 通过指定冲突重用的方式创建一个 Actor
//   - 当 Actor 已存在时将会重用已存在的 Actor
//   - 需要注意的是，冲突是由于 Actor 地址相同，即同一路径下的 Actor 地址相同。如果复用的 Actor 与使用的期望不同，可能会导致不可预知的问题
func (o *ActorOptions) WithConflictReuse(enable bool) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.ConflictReuse = enable
	})
	return o
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
