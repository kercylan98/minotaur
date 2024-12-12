package application

type Component interface {
	OnStart(app *Context)
}

type ComponentInitializer interface {
	Component
	OnInitialize(app *Context) error
}

type ComponentImporter interface {
	Component
	OnImport(provider *ComponentProvider)
}

func applyLifecycle[C Component](components []Component, handler func(C)) {
	for _, component := range components {
		if c, ok := component.(C); ok {
			handler(c)
		}
	}
}
