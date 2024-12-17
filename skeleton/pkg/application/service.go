package application

import (
	"fmt"
	"reflect"
	"sync"
)

// Service 是所有服务的核心接口，必须实现
type Service interface {
	// OnInitialize 在服务自身初始化时调用，用于进行基本的内部设置
	OnInitialize(app *Context, loader *RepositoryLoader) error
}

// NameService 提供服务名称，服务可以选择实现
type NameService interface {
	Service
	ServiceName() string
}

// DependencyService 处理依赖服务的初始化逻辑，服务可以选择实现
type DependencyService interface {
	Service
	// OnDependencyInitialize 在所有依赖服务初始化完成后调用
	OnDependencyInitialize(loader *ServiceLoader) error
}

// DependencySetupService 处理依赖内容的初始化逻辑，服务可以选择实现
type DependencySetupService interface {
	Service
	// OnDependencySetup 在所有依赖服务的基础上完成当前服务的额外设置
	// （例如：读取配置、获取依赖服务实例等）
	OnDependencySetup() error
}

// RunnableService 提供运行时逻辑，服务可以选择实现
type RunnableService interface {
	Service
	// OnRun 启动服务的主要功能，例如启动服务、监听事件等
	OnRun()
}

// StoppableService 提供停止逻辑，服务可以选择实现
type StoppableService interface {
	Service
	// OnStop 在应用关闭或服务被卸载时调用，用于释放资源
	OnStop() error
}

func newServiceLoader(ctx *Context) *ServiceLoader {
	return &ServiceLoader{ctx: ctx}
}

type ServiceLoader struct {
	ctx *Context
}

func RegisterService[I, S Service](ctx *Context, services ...S) {
	if len(services) == 0 {
		return
	}
	serviceInterfaceType := reflect.TypeOf((*I)(nil)).Elem()
	serviceType := reflect.TypeOf((*S)(nil)).Elem()
	if serviceInterfaceType.Kind() != reflect.Interface {
		panic(fmt.Errorf("service type %s must be interface", serviceType.String()))
	}
	if !serviceType.Implements(serviceInterfaceType) {
		panic(fmt.Errorf("service type %s must implement service %s interface", serviceType.String(), serviceInterfaceType.String()))
	}

	if ctx.services == nil {
		ctx.services = make(map[reflect.Type][]Service)
	}

	var mapped = make([]Service, len(services))
	for i, service := range services {
		mapped[i] = service
	}
	ctx.services[serviceInterfaceType] = append(ctx.services[serviceInterfaceType], mapped...)
}

func LoadService[S Service](loader *ServiceLoader, name ...string) S {
	serviceType := reflect.TypeOf((*S)(nil)).Elem()
	if serviceType.Kind() != reflect.Interface {
		panic("service type must be interface")
	}
	services := loader.ctx.services[serviceType]
	if len(services) > 1 {
		if len(name) == 0 {
			panic("service type has more than one implementation, please specify name")
		}

		for _, s := range name {
			for _, service := range services {
				if nameService, ok := service.(NameService); ok && nameService.ServiceName() == s {
					return service.(S)
				}
			}
		}
	}
	return services[0].(S)
}

func runServices(ctx *Context) (err error) {
	var services []Service
	for _, service := range ctx.services {
		services = append(services, service...)
	}

	loader := newRepositoryLoader(ctx)
	for _, service := range services {
		if err = service.OnInitialize(ctx, loader); err != nil {
			return
		}
	}

	var serviceLoader = newServiceLoader(ctx)
	for _, service := range services {
		if s, ok := service.(DependencyService); ok {
			if err = s.OnDependencyInitialize(serviceLoader); err != nil {
				return
			}
		}
	}

	for _, service := range services {
		if s, ok := service.(DependencySetupService); ok {
			if err = s.OnDependencySetup(); err != nil {
				return
			}
		}
	}

	serviceWaitGroup := new(sync.WaitGroup)
	for _, service := range services {
		if s, ok := service.(RunnableService); ok {
			serviceWaitGroup.Add(1)
			go func(service RunnableService) {
				defer serviceWaitGroup.Done()
				service.OnRun()
			}(s)
		}
	}

	serviceWaitGroup.Wait()

	for _, service := range services {
		if s, ok := service.(StoppableService); ok {
			if err = s.OnStop(); err != nil {
				return
			}
		}
	}

	return nil
}
