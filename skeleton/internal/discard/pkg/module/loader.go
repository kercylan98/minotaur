package module

import "reflect"

func NewLoader[App Application](modules []Module[App]) *Loader[App] {
	return &Loader[App]{modules: modules}
}

type Loader[App Application] struct {
	modules []Module[App]
}

// Load 根据模块类型或模块接口加载模块，通常更建议使用 Load 函数进行加载
func (l *Loader[App]) Load(moduleType reflect.Type) any {
	for _, module := range l.modules {
		tof := reflect.TypeOf(module)
		if tof == moduleType || tof.Implements(moduleType) {
			return module
		}
	}
	return nil
}

// Load 根据模块接口加载模块，如果模块不存在将会抛出 panic
func Load[App Application, ModuleInterface any](loader *Loader[App]) ModuleInterface {
	for _, module := range loader.modules {
		tof := reflect.TypeOf(module)
		if tof.Implements(reflect.TypeOf((*ModuleInterface)(nil)).Elem()) {
			return module.(ModuleInterface)
		}
	}
	panic("module not found")
}
