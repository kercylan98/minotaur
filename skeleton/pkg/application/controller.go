package application

type Controller interface {
	OnInitialize(ctx *Context, loader *ServiceLoader) (err error)
}

func RegisterController(ctx *Context, controller Controller) {
	ctx.controllers = append(ctx.controllers, controller)
}

func runControllers(ctx *Context) (err error) {
	loader := newServiceLoader(ctx)
	for _, controller := range ctx.controllers {
		if err = controller.OnInitialize(ctx, loader); err != nil {
			return err
		}
	}
	return
}
