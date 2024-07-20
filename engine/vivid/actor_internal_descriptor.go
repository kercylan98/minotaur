package vivid

type actorInternalDescriptor struct {
	actorContextHook func(ctx *actorContext) // 用于在 Actor 创建时捕获 *actorContext
	useDescriptor    *ActorDescriptor        // 用于指定 ActorDescriptor
	parent           ActorRef                // 用于指定父 Actor
}
