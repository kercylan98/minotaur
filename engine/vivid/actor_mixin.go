package vivid

import (
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/prc"
	"time"
)

// mixinSpawner 是一个混入类型接口，它定义了作为 Actor 的生成器需要满足的接口。
type mixinSpawner interface {
	// ActorOf 生成一个新的 Actor 实例，并以该实例作为其父 Actor。返回生成的 Actor 引用(ActorRef)
	//  - 该函数接收多个 ActorDescriptorConfigurator 参数，用于配置生成的 Actor 实例，当包含多个 ActorDescriptorConfigurator 参数时，它们的配置将会是向前覆盖的。
	//
	// 该函数不是并发安全的，你不应该在多个 goroutine 中同时调用 ActorOf 函数。
	ActorOf(provider ActorProvider, configurator ...ActorDescriptorConfigurator) ActorRef

	// ActorOfF 该函数是 ActorOf 的快捷方式，它提供了更为简便的使用方式，但是会额外创建一个切片并拷贝，用于 FunctionalActorDescriptorConfigurator 到 ActorDescriptorConfigurator 的转换。
	ActorOfF(provider FunctionalActorProvider, configurator ...FunctionalActorDescriptorConfigurator) ActorRef

	// Parent 获取当前 Actor 的父 Actor 引用
	Parent() ActorRef

	// Children 返回当前 Actor 的所有子 Actor 引用(ActorRef)。
	Children() []ActorRef
}

// mixinScheduler 是一个混入类型接口，它定义了作为调度器需要满足的接口
type mixinScheduler interface {
	// CronTask 通过 cron 表达式注册一个任务。
	//   - 当 cron 表达式错误时，将会返回错误信息
	CronTask(name, expression string, function func(ctx ActorContext)) error

	// ImmediateCronTask 与 CronTask 相同，但是会立即执行一次
	ImmediateCronTask(name, expression string, function func(ctx ActorContext)) error

	// AfterTask 注册一个在特定时间后执行一次的任务
	AfterTask(name string, after time.Duration, function func(ctx ActorContext))

	// RepeatedTask 注册一个在特定时间后反复执行的任务
	RepeatedTask(name string, after, interval time.Duration, times int, function func(ctx ActorContext))

	// DayMomentTask 注册一个在每天特定时刻执行的任务
	//   - 其中 lastExecuted 为上次执行时间，adjust 为时间偏移量，hour、min、sec 为时、分、秒
	//   - 当上次执行时间被错过时，将会立即执行一次
	DayMomentTask(name string, lastExecuted time.Time, offset time.Duration, hour, min, sec int, function func(ctx ActorContext))

	// StopTask 停止一个还未执行的任务
	StopTask(name string)
}

// mixinWorker 是一个混入类型接口，它定义了作为 Actor 工作者需要满足的接口。
type mixinWorker interface {
	// Terminate 终止目标 Actor。
	//  - 当 gracefully 参数为 true 时，会将终止消息作为用户级消息进行发送，在该消息之前的用户消息被处理完毕后升级为系统消息终止 Actor。
	Terminate(target ActorRef, gracefully bool)

	// ReportAbnormal 报告异常，该函数将触发事故向监管者传递
	ReportAbnormal(reason Message)

	// PhysicalAddress 返回当前 Actor 的物理地址
	PhysicalAddress() prc.PhysicalAddress

	// LogicalAddress 返回当前 Actor 的逻辑地址
	LogicalAddress() prc.LogicalAddress
}

// mixinDeliver 是一个混入类型接口，它定义了作为 Actor 消息发送者需要满足的接口。
type mixinDeliver interface {
	// Tell 向指定的 Actor 引用(ActorRef) 发送消息。
	//
	// 特殊标注：
	//  - MarkMessageImmutability 消息不可变性注意事项
	Tell(target ActorRef, message Message)

	// Ask 向目标 Actor 非阻塞地发送可被回复的消息，这个回复可能是无限期的
	//
	// 特殊标注：
	//  - MarkMessageImmutability 消息不可变性注意事项
	Ask(target ActorRef, message Message)

	// FutureAsk 向目标 Actor 非阻塞地发送可被回复的消息，这个回复是有限期的，返回一个 future.Future 对象，可被用于获取响应消息
	//  - 当 timeout 参数为空时，将会使用默认的超时时间 DefaultFutureAskTimeout
	//
	// 特殊标注：
	//  - MarkMessageImmutability 消息不可变性注意事项
	FutureAsk(target ActorRef, message Message, timeout ...time.Duration) future.Future[Message]

	// Broadcast 向所有子级 Actor 广播消息，广播消息是可以被回复的
	//  - 子级的子级不会收到广播消息
	//
	// 特殊标注：
	//  - MarkMessageImmutability 消息不可变性注意事项
	Broadcast(message Message)

	// AwaitForward 异步地等待阻塞结束后向目标 Actor 转发消息
	//
	// 特殊标注：
	//  - MarkMessageImmutability 消息不可变性注意事项
	AwaitForward(target ActorRef, asyncFunc func() Message)

	// ExecLocalFunc 发送一个仅支持本地的函数消息到目标 Actor 的队列中执行，在该函数中将获取到目标 Actor 的上下文
	//  - 在该函数中操作函数外部内容将是危险的
	ExecLocalFunc(target ActorRef, function func(ctx ActorContext))
}

// mixinRecipient 是一个混入类型接口，它定义了作为 Actor 接收者需要满足的接口。
type mixinRecipient interface {
	// System 返回当前 Actor 所属的 Actor 系统。
	System() *ActorSystem

	// Ref 返回当前 Actor 的 Actor 引用(ActorRef)。
	Ref() ActorRef

	// Reply 向消息发送者回复消息。
	//
	// 特殊标注：
	//  - MarkMessageImmutability 消息不可变性注意事项
	Reply(message Message)

	// Message 返回当前 Actor 接收到的消息。
	Message() Message

	// Sender 返回当前 Actor 接收到的消息的发送者。
	Sender() ActorRef

	// CastMessage 将当前正在处理的消息设置为指定的消息，这对于在处理消息时需要修改消息的场景非常有用。
	CastMessage(message Message)
}

// mixinPersistence 是一个混入类型接口，它定义了支持持久化的 Actor 需要满足的接口。
type mixinPersistence interface {
	// StateChanged 记录导致状态变更的事件，该函数将返回当前 Actor 的事件数量。
	StateChanged(event Message) int

	// StateChangeEventApply 将事件应用到 Actor 的状态上，通过该函数可以使得在状态回放时绕过业务逻辑的校验，它将当前消息转换为事件后重新对 Actor.OnReceive 发起调用以应用状态。
	StateChangeEventApply(event Message)

	// SaveSnapshot 保存快照，该函数将会清空当前 Actor 的事件记录。
	SaveSnapshot(snapshot Message)

	// ClearPersistence 清除持久化数据
	ClearPersistence()

	// Persistence 主动进行状态的持久化，将存储的事件及快照应用到存储器中
	Persistence() error
}

// mixinWatcher 是一个混入类型接口，它定义了支持观察与被观察生命周期的 Actor 需要满足的接口。
type mixinWatcher interface {
	// Watch 监听特定 Actor 生命周期的结束
	Watch(target ActorRef)

	// UnWatch 取消对特定 Actor 生命周期结束的监听
	UnWatch(target ActorRef)
}

// mixinSupervisor 是一个混入类型接口，它定义了支持发布与订阅的 Actor 需要满足的接口。
type mixinSubscription interface {
	// Subscribe 订阅特定主题，在收到特定主题的消息时，Actor 将会收到该消息，当 Actor 重启或停止时，将会取消所有订阅。
	Subscribe(topic Topic) Subscription

	// UnSubscribe 取消特定订阅
	UnSubscribe(subscription Subscription)

	// Publish 向所有订阅者发布消息
	Publish(topic Topic, message Message)
}
