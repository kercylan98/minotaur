package vivid

import (
	"github.com/kercylan98/minotaur/engine/vivid/persistence"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"regexp"
	"time"
)

const (
	DefaultPersistenceEventThreshold = 1000
)

var (
	actorNameRegexp = regexp.MustCompile(`^[^\s\\/]+$`)
)

// newActorDescriptor 创建一个 ActorDescriptor 实例，它包含默认配置
//   - 这应该是一个可以池化的对象
func newActorDescriptor() *ActorDescriptor {
	return &ActorDescriptor{
		mailboxProvider:            GetDefaultMailboxProvider(),
		dispatcherProvider:         GetDefaultDispatcherProvider(),
		persistenceStorageProvider: GetDefaultPersistenceStorageProvider(),
		persistenceEventThreshold:  DefaultPersistenceEventThreshold,
	}
}

// ActorDescriptor 用于定义 Actor 个性化行为的描述符，它仅在 Actor 创建时使用并释放
type ActorDescriptor struct {
	internal                    *actorInternalDescriptor     // 内部用于描述 Actor 的描述符
	name                        string                       // Actor 名称
	namePrefix                  string                       // Actor 名称前缀
	mailboxProvider             MailboxProvider              // 邮箱提供者
	dispatcherProvider          DispatcherProvider           // 调度器提供者
	supervisionStrategyProvider supervision.StrategyProvider // 监督策略提供者
	supervisionLoggers          []supervision.Logger         // 监督日志
	fixedLocal                  bool                         // 固定本地 Actor
	expireDuration              time.Duration                // 过期时间
	idleDeadline                time.Duration                // 空闲截止时间
	persistenceStorageProvider  persistence.StorageProvider  // 持久化存储提供者
	persistenceName             persistence.Name             // 持久化名称
	persistenceEventThreshold   int                          // 持久化事件数量阈值
}

// WithPersistenceEventThreshold 设置持久化事件数量阈值
//   - 当 Actor 的事件数量超过该阈值时，将会触发快照的持久化
//
// 默认值: DefaultPersistenceEventThreshold
func (d *ActorDescriptor) WithPersistenceEventThreshold(threshold int) *ActorDescriptor {
	d.persistenceEventThreshold = threshold
	return d
}

// WithPersistenceName 设置持久化名称，持久化名称用于标识 Actor 的持久化状态。
//
// 默认值为 Actor 的 prc.LogicalAddress
func (d *ActorDescriptor) WithPersistenceName(name persistence.Name) *ActorDescriptor {
	d.persistenceName = name
	return d
}

// WithPersistenceStorageProvider 设置持久化存储提供者
func (d *ActorDescriptor) WithPersistenceStorageProvider(provider persistence.StorageProvider) *ActorDescriptor {
	d.persistenceStorageProvider = provider
	return d
}

// WithIdleDeadline 设置 Actor 空闲截止时间
//   - 空闲截止时间是指 Actor 在空闲时间超过该时间后将会被终止
func (d *ActorDescriptor) WithIdleDeadline(deadline time.Duration) *ActorDescriptor {
	d.idleDeadline = deadline
	return d
}

// WithExpireDuration 设置 Actor 过期时间
//   - 过期是指该 Actor 在到达期限后将会被终止
func (d *ActorDescriptor) WithExpireDuration(duration time.Duration) *ActorDescriptor {
	d.expireDuration = duration
	return d
}

// WithFixedLocal 设置该 Actor 固定在本地生成，不会被远程 ActorSystem 生成。
//   - 虽然 Actor 在不满足集群情况时不会在集群的其他节点生成，但是通过该函数可以明确的指定。
//   - 在某些场景下，避免远程调用可以提高处理的效率，例如将 Actor 作为对 Socket 连接的处理器时。
//
// 该可选项仅在开启集群时会生效
func (d *ActorDescriptor) WithFixedLocal() *ActorDescriptor {
	d.fixedLocal = true
	return d
}

// WithSupervisionStrategyProvider 设置监督策略提供者
func (d *ActorDescriptor) WithSupervisionStrategyProvider(provider supervision.StrategyProvider, loggers ...supervision.Logger) *ActorDescriptor {
	d.supervisionStrategyProvider = provider
	d.supervisionLoggers = append(d.supervisionLoggers, loggers...)
	return d
}

// WithName 设置 Actor 名称，名称中禁止包含空格、换行符等特殊字符、以及 '\'、'/'
func (d *ActorDescriptor) WithName(name string) *ActorDescriptor {
	if !actorNameRegexp.MatchString(name) {
		panic("Actor name should not contain space, '\\' or '/'")
	}

	d.name = name
	return d
}

// WithNamePrefix 设置 Actor 名称前缀，名称中禁止包含空格、换行符等特殊字符、以及 '\'、'/'
//   - 前缀将会与名称使用 "-" 连接
func (d *ActorDescriptor) WithNamePrefix(prefix string) *ActorDescriptor {
	if !actorNameRegexp.MatchString(prefix) {
		panic("Actor name should not contain space, '\\' or '/'")
	}

	d.namePrefix = prefix
	return d
}

// WithMailboxProvider 设置邮箱提供者
func (d *ActorDescriptor) WithMailboxProvider(provider MailboxProvider) *ActorDescriptor {
	d.mailboxProvider = provider
	return d
}

// WithDispatcherProvider 设置调度器提供者
func (d *ActorDescriptor) WithDispatcherProvider(provider DispatcherProvider) *ActorDescriptor {
	d.dispatcherProvider = provider
	return d
}

// withInternalDescriptor 设置内部描述符
func (d *ActorDescriptor) withInternalDescriptor(internal *actorInternalDescriptor) *ActorDescriptor {
	d.internal = internal
	return d
}
