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
}

type actorContextState = uint8 // actorContext 的状态

// actorContext 是 Actor 的上下文
type actorContext struct {
	id        ActorId                        // Actor 的 ID
	state     actorContextState              // 上下文的状态
	behaviors map[reflect.Type]reflect.Value // Actor 的行为，由消息类型到行为的映射
}

// RegisterBehavior 注册 Actor 的行为
//   - 该函数仅在 Actor.OnPreStart 阶段有效
func RegisterBehavior[T any](ctx ActorContext, behavior ActorBehavior[T]) {
	messageType := reflect.TypeOf((*T)(nil)).Elem()
	ctx.RegisterBehavior(messageType, behavior)
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
