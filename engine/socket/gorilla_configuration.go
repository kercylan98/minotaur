package socket

type ()

func NewGorillaConfiguration() *GorillaConfiguration {
	return &GorillaConfiguration{}
}

type GorillaConfiguration struct {
	contextInitializer ContextEditor
}

func (gc *GorillaConfiguration) WithContextInitializer(initializer ContextEditor) *GorillaConfiguration {
	gc.contextInitializer = initializer
	return gc
}
