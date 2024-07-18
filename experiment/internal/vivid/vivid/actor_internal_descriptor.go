package vivid

type actorInternalDescriptor struct {
	actorContextHook func(ctx *actorContext) // 用于在 Actor 创建时捕获 *actorContext
}
