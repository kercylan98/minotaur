package application

import (
	"context"
	"github.com/kercylan98/minotaur/engine/vivid"
	"reflect"
)

func New() *Context {
	ctx := &Context{
		ctx: context.Background(),
	}
	return ctx
}

type Context struct {
	ctx            context.Context
	actorSystem    *vivid.ActorSystem
	modules        map[reflect.Type]Module
	services       map[reflect.Type][]Service
	repositoryList map[reflect.Type][]Repository
	controllers    []Controller
}

func (c *Context) Run() (err error) {

	if err = runModules(c); err != nil {
		return
	}

	if err = runRepositoryList(c); err != nil {
		return
	}

	if err = runServices(c); err != nil {
		return
	}

	if err = runControllers(c); err != nil {
		return
	}

	return nil
}
