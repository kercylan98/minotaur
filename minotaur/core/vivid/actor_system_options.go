package vivid

import "github.com/google/uuid"

type ActorSystemOption func(options *ActorSystemOptions)

type ActorSystemOptions struct {
	options []ActorSystemOption

	Name    string   // 指定 ActorSystem 的名称
	modules []Module // 指定 ActorSystem 的组件
}

func (o *ActorSystemOptions) WithName(name string) *ActorSystemOptions {
	o.options = append(o.options, func(options *ActorSystemOptions) {
		options.Name = name
	})
	return o
}

func (o *ActorSystemOptions) WithModule(module Module) *ActorSystemOptions {
	o.options = append(o.options, func(options *ActorSystemOptions) {
		options.modules = append(options.modules, module)
	})
	return o
}

func (o *ActorSystemOptions) apply(handlers []func(options *ActorSystemOptions)) *ActorSystemOptions {
	o.Name = uuid.NewString()
	for _, handler := range handlers {
		handler(o)
	}
	for _, option := range o.options {
		option(o)
	}
	return o
}
