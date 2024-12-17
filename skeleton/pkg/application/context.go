package application

import (
	"context"
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
	services       map[reflect.Type][]Service
	repositoryList map[reflect.Type][]Repository
}

func (c *Context) Run() (err error) {
	if err = runRepositoryList(c); err != nil {
		return
	}

	if err = runServices(c); err != nil {
		return
	}

	return nil
}
