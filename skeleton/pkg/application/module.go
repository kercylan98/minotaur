package application

import "github.com/kercylan98/minotaur/engine/vivid"

type Module interface {
	vivid.Actor

	Name() string

	Setup(ctx *Context) error
}

type ModuleDependent interface {
	Module

	ImportDependencies(getter func(name string) Module)
}
