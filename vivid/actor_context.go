package vivid

import (
	"fmt"
	"reflect"
)

const (
	actorContextStatePreStart actorContextState = iota // 启动之前
	actorContextStateStarted                           // 已启动
)

type actorContextState = uint8 // ActorContext 的状态

// ActorContext 是 Actor 的上下文
type ActorContext struct {
	id           ActorId                        // Actor 的 ID
	ref          ActorRef                       // Actor 的引用
	vof          reflect.Value                  // 上下文的值类型
	state        actorContextState              // 上下文的状态
	behaviors    map[reflect.Type]reflect.Value // Actor 的行为，由消息类型到行为的映射
	systemGetter actorSystemGetter              // 获取 ActorSystem 的函数
}

func (ctx *ActorContext) init(id ActorId, ref ActorRef, systemGetter actorSystemGetter) *ActorContext {
	ctx.id = id
	ctx.ref = ref
	ctx.vof = reflect.ValueOf(ctx)
	ctx.state = actorContextStatePreStart
	ctx.behaviors = make(map[reflect.Type]reflect.Value)
	ctx.systemGetter = systemGetter
	return ctx
}

// GetSystem 获取 ActorSystem
func (ctx *ActorContext) GetSystem() *ActorSystem {
	return ctx.systemGetter()
}

// RegisterBehavior 注册 Actor 的行为
//   - 该函数仅在 Actor.OnPreStart 阶段有效
func RegisterBehavior[T any](ctx *ActorContext, behavior ActorBehavior[T]) {
	messageType := reflect.TypeOf((*T)(nil)).Elem()
	ctx.RegisterBehavior(messageType, behavior)
}

// RegisterBehavior 注册 Actor 的行为，行为是一个函数，用于处理 Actor 接收到的消息
//   - 推荐使用 RegisterBehavior 函数来注册 Actor 的行为，这样可以保证行为的类型安全
//   - 该函数仅在 Actor.OnPreStart 阶段有效
func (ctx *ActorContext) RegisterBehavior(messageType reflect.Type, behavior any) {
	if ctx.state != actorContextStatePreStart {
		return
	}

	behaviorType := reflect.TypeOf(behavior)
	behaviorValue := reflect.ValueOf(behavior)

	// 检查是否符合行为定义
	if behaviorType.Kind() != reflect.Func {
		panic(fmt.Errorf("%w: %s", ErrActorBehaviorInvalid, behaviorType))
	}

	// 检查行为参数是否符合要求
	if behaviorType.NumIn() != 2 || behaviorType.In(0) != reflect.TypeOf(ctx) {
		panic(fmt.Errorf("%w: %s", ErrActorBehaviorInvalid, behaviorType))
	}

	// 检查行为返回值是否符合要求，必须是一个 error 类型
	if behaviorType.NumOut() != 1 || behaviorType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		panic(fmt.Errorf("%w: %s", ErrActorBehaviorInvalid, behaviorType))
	}

	ctx.behaviors[messageType] = behaviorValue
}

// MatchBehavior 匹配 Actor 的行为，当匹配不到时将返回 nil，否则返回匹配到的行为执行器
func MatchBehavior[T any](ctx *ActorContext, message T) ActorBehaviorExecutor {
	messageType := reflect.TypeOf(message)
	return ctx.MatchBehavior(messageType, message)
}

// MatchBehavior 匹配 Actor 的行为，当匹配不到时将返回 nil，否则返回匹配到的行为执行器
//   - 推荐使用 MatchBehavior 函数来匹配 Actor 的行为，这样可以保证行为的类型安全
func (ctx *ActorContext) MatchBehavior(messageType reflect.Type, message any) ActorBehaviorExecutor {
	behaviorValue, ok := ctx.behaviors[messageType]

	if !ok {
		return nil
	}

	return func() error {
		result := behaviorValue.Call([]reflect.Value{ctx.vof, reflect.ValueOf(message)})
		switch v := result[0].Interface().(type) {
		case nil:
			return nil
		default:
			return v.(error)
		}
	}
}
