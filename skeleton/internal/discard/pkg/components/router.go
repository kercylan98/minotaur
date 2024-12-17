package components

import "github.com/kercylan98/minotaur/toolkit/router"

type RouterComponent[HandleFunc any] interface {
	Bind(handleFunc HandleFunc)
	Register(routes ...any) router.MultistageBind[HandleFunc]
	Route(route any, handleFunc HandleFunc)
	Match(routes ...any) HandleFunc
	Sub(route any) *router.Multistage[HandleFunc]
}
