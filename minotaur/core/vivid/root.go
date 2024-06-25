package vivid

type root struct {
}

func (r *root) OnReceive(ctx ActorContext) {
	switch m := ctx.Message().(type) {
	case []Module:
		r.onLoadPlugins(ctx, m)
	}
}

func (r *root) onLoadPlugins(ctx ActorContext, m []Module) {
	support := newModuleSupport(ctx.System())
	for _, plugin := range m {
		plugin.OnLoad(support)
	}
}
