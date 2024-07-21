package transport

import "reflect"

type FiberExposer struct {
	k reflect.Type
	v FiberService
}

// FiberExpose 生成用于暴露特定服务接口的暴露器
//
// 其中的 I 指代了用于暴露服务的对外提供的接口，而 service 则是需要实现了 I 的结构，两者均需满足 FiberService 接口。
// 在使用中，I 所指代的接口不应与 service 位于相同的包内，否则可能导致循环依赖。
func FiberExpose[I FiberService](service I) *FiberExposer {
	iTyp := reflect.TypeOf((*I)(nil)).Elem()
	if iTyp.Kind() != reflect.Interface {
		panic("FiberExpose: I must be an interface")
	}

	sTyp := reflect.TypeOf(service)
	if !sTyp.Implements(iTyp) {
		panic("FiberExpose: service must implement I")
	}

	if sTyp.Kind() != reflect.Pointer {
		panic("FiberExpose: service must be a pointer")
	}

	return &FiberExposer{
		k: iTyp,
		v: service,
	}
}
