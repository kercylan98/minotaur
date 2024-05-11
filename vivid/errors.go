package vivid

import "errors"

var (
	ErrActorNotFound               = errors.New("actor not found")                 // 未找到 Actor
	ErrActorMessageHandlerNotFunc  = errors.New("actor message handler not func")  // Actor 消息处理函数不是函数
	ErrActorMessageHandlerNotFound = errors.New("actor message handler not found") // Actor 消息处理函数未找到
	ErrActorNotHasAnyHandler       = errors.New("actor not has any handler")       // Actor 没有任何处理函数
	ErrActorHandlerParamsNotMatch  = errors.New("actor handler params not match")  // Actor 处理函数参数不匹配
)
