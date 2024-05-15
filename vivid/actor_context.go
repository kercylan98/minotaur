package vivid

import (
	"fmt"
	"path"
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
	NotifyTerminated()

	// ActorOf 创建一个 Actor，该 Actor 是当前 Actor 的子 Actor
	ActorOf(typ reflect.Type, opts ...*ActorOptions) (ActorRef, error)

	// GetActor 获取 Actor 的引用
	GetActor() Query
}

type actorContextState = uint8 // actorContext 的状态

// actorContext 是 Actor 的上下文
type actorContext struct {
	system    *ActorSystem                   // Actor 所属的 Actor 系统
	id        ActorId                        // Actor 的 ID
	core      *actorCore                     // Actor 的核心
	state     actorContextState              // 上下文的状态
	behaviors map[reflect.Type]reflect.Value // Actor 的行为，由消息类型到行为的映射
	children  map[ActorName]*actorCore       // 子 Actor
	isEnd     bool                           // 是否是末级 Actor
}

func (c *actorContext) ActorOf(typ reflect.Type, opts ...*ActorOptions) (ActorRef, error) {
	// 检查类型是否实现了 Actor 接口
	if !typ.Implements(actorType) {
		return nil, fmt.Errorf("%w: %s", ErrActorNotImplementActorRef, typ.String())
	}
	typ = typ.Elem()

	// 应用可选项
	opt := NewActorOptions().Apply(opts...)
	if opt.Name == "" {
		opt.Name = fmt.Sprintf("%s-%d", c.id.Name(), c.system.guid.Add(1))
	}
	actorPath := path.Join(c.id.Name(), opt.Name)
	actorName := opt.Name

	// 创建 Actor ID
	actorId := NewActorId(c.id.Network(), c.id.Cluster(), c.id.Host(), c.id.Port(), c.id.System(), actorPath)

	// 检查 Actor 是否已经存在
	c.system.actorsRW.Lock()
	defer c.system.actorsRW.Unlock()
	actor, exist := c.system.actors[actorId]
	if exist {
		return nil, fmt.Errorf("%w: %s", ErrActorAlreadyExists, actorId.Name())
	}

	// 创建 Actor
	actor = newActorCore(c.system, actorId, reflect.New(typ).Interface().(Actor), opt)

	// 分发器
	dispatcher := c.system.getActorDispatcher(actor)
	if err := dispatcher.Attach(actor); err != nil {
		return nil, err
	}

	// 启动 Actor
	if err := actor.onPreStart(); err != nil {
		return nil, err
	}

	// 绑定 Actor
	c.system.actors[actorId] = actor
	c.children[actorName] = actor

	if opt.hookActorOf != nil {
		opt.hookActorOf(actor)
	}
	return actor, nil
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

func (c *actorContext) NotifyTerminated() {
	c.system.unregisterActor(c.id)
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
	return ctx.GetActorContext().MatchBehavior(reflect.TypeOf(ctx.GetMessage()), ctx)
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
