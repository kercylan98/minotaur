package module

func New[App Application](app App) *Manager[App] {
	return &Manager[App]{
		app: app,
	}
}

type Manager[App Application] struct {
	app     App
	modules []Module[App]
}

func (m *Manager[App]) Register(modules ...Module[App]) *Manager[App] {
	m.modules = append(m.modules, modules...)
	return m
}

// Run 运行所有模块，当模块运行前生命周期发生错误时将会导致 panic
func (m *Manager[App]) Run() {
	// 初始化所有模块
	for _, module := range m.modules {
		if err := module.OnInitialize(m.app); err != nil {
			panic(err)
		}
	}

	loader := NewLoader(m.modules)
	for _, module := range m.modules {
		if v, ok := module.(Dependency[App]); ok {
			if err := v.OnDependencyInitialize(loader); err != nil {
				panic(err)
			}
		}
	}

	for _, module := range m.modules {
		if v, ok := module.(DependencySetup[App]); ok {
			if err := v.OnDependencySetup(); err != nil {
				panic(err)
			}
		}
	}

	for _, module := range m.modules {
		if v, ok := module.(Runnable[App]); ok {
			v.OnRun()
		}
	}
}
