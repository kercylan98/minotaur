package transport

import "reflect"

type FiberProvider reflect.Type

func FiberProvide[I FiberService]() FiberProvider {
	return FiberProvider(reflect.TypeOf((*I)(nil)).Elem())
}
