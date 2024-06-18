package vivid

// ActorTyped 是用于将 Actor 消息类型化的接口，该接口继承了 Actor 和 ActorRef
type ActorTyped interface {
	ActorRef
}
