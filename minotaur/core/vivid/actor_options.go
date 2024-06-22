package vivid

type ActorOption func(options *ActorOptions)

type ActorOptions struct {
	options []ActorOption
	Parent  ActorRef
	Name    string
}

// WithName 通过指定名称创建一个 Actor
func (o *ActorOptions) WithName(name string) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.Name = name
	})
	return o
}

// WithParent 通过指定父 Actor 创建一个 Actor
func (o *ActorOptions) WithParent(parent ActorRef) *ActorOptions {
	o.options = append(o.options, func(options *ActorOptions) {
		options.Parent = parent
	})
	return o
}

func (o *ActorOptions) apply() *ActorOptions {
	for _, option := range o.options {
		option(o)
	}
	return o
}
