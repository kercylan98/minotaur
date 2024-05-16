package vivid

import (
	"fmt"
	"reflect"
)

const (
	actorContextStatePreStart actorContextState = iota // 启动之前
	actorContextStateStarted                           // 已启动
)

type ActorContext interface {
	// RegisterBehavior 注册 Actor 的行为，行为是一个函数，用于处理 Actor 接收到的消息
	//   - 推荐使用 RegisterBehavior 函数来注册 Actor 的行为，这样可以保证行为的类型安全
	//   - 该函数仅在 Actor.OnPreStart 阶段有效
	RegisterBehavior(messageType reflect.Type, behavior any)

	// MatchBehavior 匹配 Actor 的行为，当匹配不到时将返回 nil，否则返回匹配到的行为执行器
	//   - 推荐使用 MatchBehavior 函数来匹配 Actor 的行为，这样可以保证行为的类型安全
	MatchBehavior(messageType reflect.Type, ctx MessageContext) ActorBehaviorExecutor

	// NotifyTerminated 当 Actor 主动销毁时，务必调用该函数，以便在整个 Actor 系统中得到完整的释放
	NotifyTerminated(v ...Message)

	// ActorOf 创建一个 Actor，该 Actor 是当前 Actor 的子 Actor
	ActorOf(actor Actor, opts ...*ActorOptions) (ActorRef, error)

	// GetActor 获取 Actor 的引用
	GetActor() Query

	// GetParentActor 获取父 Actor 的引用
	GetParentActor() ActorRef

	// GetActorId 获取 Actor 的 ID
	GetActorId() ActorId
}

type actorContextState = uint8 // actorContext 的状态

// actorContext 是 Actor 的上下文
type actorContext struct {
	system    *ActorSystem                   // Actor 所属的 Actor 系统
	id        ActorId                        // Actor 的 ID
	core      *actorCore                     // Actor 的核心
	state     actorContextState              // 上下文的状态
	behaviors map[reflect.Type]reflect.Value // Actor 的行为，由消息类型到行为的映射
	parent    *actorCore                     // 父 Actor
	children  map[ActorName]*actorCore       // 子 Actor
	isEnd     bool                           // 是否是末级 Actor
	tof       reflect.Type                   // Actor 的类型
}

// restart 重启上下文
func (c *actorContext) restart(reEntry bool) (recovery func(), err error) {
	// 重启所有子 Actor
	if !reEntry && c.state != actorContextStatePreStart {
		c.system.actorsRW.RLock()
		defer c.system.actorsRW.RUnlock()
	}
	var recoveries []func()
	for _, child := range c.children {
		if recovery, err = child.restart(true); err != nil {
			for _, f := range recoveries {
				f()
			}
			return nil, err
		}
		recoveries = append(recoveries, recovery)
	}

	backupActor := c.core.Actor
	backupState := c.core.state
	backupBehaviors := c.behaviors
	var recoveryFunc = func() {
		c.core.Actor = backupActor
		c.core.state = backupState
		c.behaviors = backupBehaviors
	}
	c.core.Actor = reflect.New(c.core.tof).Interface().(Actor)
	c.core.state = actorContextStatePreStart
	c.core.actorContext.behaviors = make(map[reflect.Type]reflect.Value)
	if err = c.core.onPreStart(); err != nil {
		recoveryFunc()
		return nil, err
	}
	return recoveryFunc, nil
}

func (c *actorContext) bindChildren(core *actorCore) {
	c.children[core.GetOptions().Name] = core
}

func (c *actorContext) actorOf(typ reflect.Type, opts ...*ActorOptions) (ActorRef, error) {
	var opt *ActorOptions
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = NewActorOptions()
	}
	opt = opt.WithParent(c)

	return c.system.generateActor(typ, opt)
}

func (c *actorContext) ActorOf(actor Actor, opts ...*ActorOptions) (ActorRef, error) {
	return c.actorOf(reflect.TypeOf(actor), opts...)
}

func (c *actorContext) GetActor() Query {
	return newQuery(c.system, c.core)
}

// RegisterBehavior 注册 Actor 的行为
//   - 该函数仅在 Actor.OnPreStart 阶段有效
func RegisterBehavior[T any](ctx ActorContext, behavior ActorBehavior[T]) {
	messageType := reflect.TypeOf((*T)(nil)).Elem()
	ctx.RegisterBehavior(messageType, behavior)
}

func (c *actorContext) NotifyTerminated(v ...Message) {
	terminatedContext := newActorTerminatedContext(c.core, v...)
	c.parent.OnChildTerminated(c.parent, terminatedContext)
	if terminatedContext.cancelTerminate {
		return
	}
	c.system.unregisterActor(c.core, c.core.state == actorContextStatePreStart)
}

func (c *actorContext) RegisterBehavior(messageType reflect.Type, behavior any) {
	if c.state != actorContextStatePreStart {
		return
	}

	behaviorType := reflect.TypeOf(behavior)
	behaviorValue := reflect.ValueOf(behavior)

	// 检查是否符合行为定义
	if behaviorType.Kind() != reflect.Func {
		panic(fmt.Errorf("%w: %s", ErrActorBehaviorInvalid, behaviorType))
	}

	// 检查行为参数是否符合要求
	if behaviorType.NumIn() != 2 || behaviorType.In(0) != messageContextType {
		panic(fmt.Errorf("%w: %s", ErrActorBehaviorInvalid, behaviorType))
	}

	// 检查行为返回值是否符合要求，必须是一个 error 类型
	if behaviorType.NumOut() != 1 || behaviorType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		panic(fmt.Errorf("%w: %s", ErrActorBehaviorInvalid, behaviorType))
	}

	c.behaviors[messageType] = behaviorValue
}

// MatchBehavior 匹配 Actor 的行为，当匹配不到时将返回 nil，否则返回匹配到的行为执行器
func MatchBehavior(ctx MessageContext) ActorBehaviorExecutor {
	return ctx.MatchBehavior(reflect.TypeOf(ctx.GetMessage()), ctx)
}

func (c *actorContext) MatchBehavior(messageType reflect.Type, ctx MessageContext) ActorBehaviorExecutor {
	behaviorValue, ok := c.behaviors[messageType]

	if !ok {
		return nil
	}

	return func() error {
		result := behaviorValue.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(ctx.GetMessage())})
		switch v := result[0].Interface().(type) {
		case nil:
			return nil
		default:
			return v.(error)
		}
	}
}

func (c *actorContext) GetParentActor() ActorRef {
	return c.parent.ActorRef
}

func (c *actorContext) GetActorId() ActorId {
	return c.id
}
