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
	components  []Component
}

func (c *Context) SetupComponents(components ...Component) {
	c.components = append(c.components, components...)
}

func (c *Context) ActorSystem() *vivid.ActorSystem {
	return c.actorSystem
}

func (c *Context) Run() error {
	provider := newComponentProvider(c)
	applyLifecycle[ComponentImporter](c.components, func(v ComponentImporter) {
		v.OnImport(provider)
	})

	applyLifecycle[ComponentInitializer](c.components, func(v ComponentInitializer) {
		if err := v.OnInitialize(c); err != nil {
			panic(fmt.Errorf("component %T initialize failed: %w", v, err))
		}
	})

	applyLifecycle[Component](c.components, func(v Component) {
		v.OnStart(c)
	})

	c.actorSystem.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.Shutdown(true)
	})

	return nil
}
