package application

import "reflect"

func newComponentProvider(app *Context) *ComponentProvider {
	return &ComponentProvider{app: app}
}

type ComponentProvider struct {
	app *Context
}

func ProvideComponent[ComponentInterface any](provider *ComponentProvider) ComponentInterface {
	for _, component := range provider.app.components {
		tof := reflect.TypeOf(component)
		if tof.Implements(reflect.TypeOf((*ComponentInterface)(nil)).Elem()) {
			return component.(ComponentInterface)
		}
	}
	panic("component not found")
}
