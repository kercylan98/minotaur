package vivid

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/supervision"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"regexp"
)

var (
	actorNameRegexp = regexp.MustCompile(`^[^\s\\/]+$`)
)

// newActorDescriptor 创建一个 ActorDescriptor 实例，它包含默认配置
//   - 这应该是一个可以池化的对象
func newActorDescriptor() *ActorDescriptor {
	return &ActorDescriptor{
		mailboxProvider:    GetDefaultMailboxProvider(),
		dispatcherProvider: GetDefaultDispatcherProvider(),
	}
}

// ActorDescriptor 用于定义 Actor 个性化行为的描述符，它仅在 Actor 创建时使用并释放
type ActorDescriptor struct {
	internal            *actorInternalDescriptor // 内部用于描述 Actor 的描述符
	name                string                   // Actor 名称
	namePrefix          string                   // Actor 名称前缀
	mailboxProvider     MailboxProvider          // 邮箱提供者
	dispatcherProvider  DispatcherProvider       // 调度器提供者
	supervisionStrategy supervision.Strategy     // 监督策略
	supervisionLoggers  []supervision.Logger     // 监督日志
}

// WithSupervisionStrategy 设置监督策略
func (d *ActorDescriptor) WithSupervisionStrategy(strategy supervision.Strategy, loggers ...supervision.Logger) *ActorDescriptor {
	d.supervisionStrategy = strategy
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
func (d *ActorDescriptor) withInternalDescriptor(name string, internal *actorInternalDescriptor) *ActorDescriptor {
	if name == charproc.None {
		panic("internal actor name should not be empty")
	}
	d.name = name
	d.internal = internal
	return d
}
