package socket

func NewFiberV2Configuration() *FiberV2Configuration {
	return &FiberV2Configuration{}
}

type FiberV2Configuration struct {
	contextInitializer ContextEditor
}

func (fc *FiberV2Configuration) WithContextInitializer(initializer ContextEditor) *FiberV2Configuration {
	fc.contextInitializer = initializer
	return fc
}
