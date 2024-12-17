package application

type Controller interface {
	OnInitialize(ctx *Context, loader *ServiceLoader) (err error)
}
