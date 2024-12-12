package router

import (
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/toolkit/router"
)

func NewRouterComponent[HandleFunc any]() application.Component {
	return &routerComponent[HandleFunc]{
		router.NewMultistage[HandleFunc](),
	}
}

type routerComponent[HandleFunc any] struct {
	*router.Multistage[HandleFunc]
}

func (r *routerComponent[HandleFunc]) OnInitialize(app *application.Context) error {
	return nil
}

func (r *routerComponent[HandleFunc]) OnStart(app *application.Context) {

}
