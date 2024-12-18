package application

import (
	"fmt"
	"reflect"
)

type Module interface {
	OnInitialize(ctx *Context) (err error)

	OnDependencyInitialize(ctx *Context) (err error)

	OnDependencySetup() (err error)

	OnSetup() (err error)
}

func RegisterModule[I, R Module](ctx *Context, module R) {
	moduleInterfaceType := reflect.TypeOf((*I)(nil)).Elem()
	moduleType := reflect.TypeOf((*R)(nil)).Elem()
	if moduleInterfaceType.Kind() != reflect.Interface {
		panic(fmt.Sprintf("module type %s must be interface", moduleInterfaceType.String()))
	}
	if !moduleType.Implements(moduleInterfaceType) {
		panic(fmt.Errorf("module type %s must implement module %s interface", moduleType.String(), moduleInterfaceType.String()))
	}

	if ctx.modules == nil {
		ctx.modules = make(map[reflect.Type]Module)
	}

	if _, ok := ctx.modules[moduleInterfaceType]; ok {
		panic(fmt.Sprintf("module type %s already registered", moduleInterfaceType.String()))
	}
	ctx.modules[moduleInterfaceType] = module
}

func LoadModule[R Module](ctx *Context) R {
	moduleType := reflect.TypeOf((*R)(nil)).Elem()
	if moduleType.Kind() != reflect.Interface {
		panic("module type must be interface")
	}
	module := ctx.modules[moduleType]
	if module == nil {
		panic(fmt.Sprintf("module type %s not registered", moduleType.String()))
	}

	return module.(R)
}

func runModules(ctx *Context) (err error) {
	for _, module := range ctx.modules {
		if err = module.OnInitialize(ctx); err != nil {
			return
		}
	}

	for _, module := range ctx.modules {
		if err = module.OnDependencyInitialize(ctx); err != nil {
			return
		}
	}

	for _, module := range ctx.modules {
		if err = module.OnSetup(); err != nil {
			return
		}
	}

	return nil
}
