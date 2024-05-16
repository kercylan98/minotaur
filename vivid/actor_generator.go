package vivid

import "reflect"

type ActorGenerator interface {
	ActorOf(typ reflect.Type, opts ...*ActorOptions) (ActorRef, error)
}
