package application

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"os"
)

func New() *Context {
	ctx := &Context{
		actorSystem: vivid.NewActorSystem(),
	}

	return ctx
}

type Context struct {
	actorSystem *vivid.ActorSystem
	modules     []Module
}

func (c *Context) SetupModule(modules ...Module) {
	c.modules = append(c.modules, modules...)
}

func (c *Context) ActorSystem() *vivid.ActorSystem {
	return c.actorSystem
}

func (c *Context) Run() error {
	// 引入模块依赖
	for _, module := range c.modules {
		if v, importer := module.(ModuleDependent); importer {
			v.ImportDependencies(func(name string) Module {
				for _, m := range c.modules {
					if m.Name() == name {
						return m
					}
				}
				panic(fmt.Sprintf("module %s not found", name))
			})
		}
	}

	for _, module := range c.modules {
		if err := module.Setup(c); err != nil {
			return err
		}
	}

	c.actorSystem.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.Shutdown(true)
	})

	return nil
}
