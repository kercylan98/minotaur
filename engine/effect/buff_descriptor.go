package effect

type AttributeEditor = func(attributes *Attributes) *Attributes

func newBuffDescriptor() *BuffDescriptor {
	return &BuffDescriptor{}
}

type BuffDescriptor struct {
	applyHooks []AttributeEditor // 添加 Buff 时触发
}

// WithApplyHooks 添加 Buff 时触发
func (bd *BuffDescriptor) WithApplyHooks(hooks ...AttributeEditor) *BuffDescriptor {
	bd.applyHooks = append(bd.applyHooks, hooks...)
	return bd
}
