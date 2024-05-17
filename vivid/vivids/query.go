package vivids

type Query interface {
	// ActorId 使用特定的 ActorId，该查询可能包含不存在的 ActorRef
	ActorId(actorIds ...ActorId) Query

	// MustActorId 使用特定的 ActorId，该查询包含必须存在的 ActorRef
	MustActorId(actorIds ...ActorId) Query

	// ActorName 通过 ActorName 进行查询匹配
	ActorName(names ...ActorName) Query

	// ActorPath 通过 ActorPath 进行查询匹配，支持格式如下：
	// - "./user/actor1"
	// - "/user/actor1"
	// - "user/actor1"
	// - "actor1"
	// - "actor1/actor2/*"
	ActorPath(actorPath ActorPath) Query

	// Many 获取多个响应结果
	Many() []ActorRef

	// First 获取第一个响应结果
	First() (ActorRef, error)

	// One 获取唯一的响应结果
	One() (ActorRef, error)
}
