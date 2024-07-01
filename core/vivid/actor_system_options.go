package vivid

import (
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/toolkit/log"
)

type LoggerProvider func() *log.Logger

type ActorSystemOption func(options *ActorSystemOptions)

type ActorSystemOptions struct {
	options []ActorSystemOption
	modules []Module // 指定 ActorSystem 的组件

	Name           string // 指定 ActorSystem 的名称
	LoggerProvider LoggerProvider
}

func (o *ActorSystemOptions) WithLoggerProvider(provider LoggerProvider) *ActorSystemOptions {
	o.options = append(o.options, func(options *ActorSystemOptions) {
		options.LoggerProvider = provider
	})
	return o
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
	if o.LoggerProvider == nil {
		logger := log.NewSilentLogger()
		o.LoggerProvider = func() *log.Logger {
			return logger
		}
	}
	return o
}
