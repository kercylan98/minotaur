package application

import (
	"fmt"
	"reflect"
)

type Repository interface {
	OnInitialize(ctx *Context) (err error)
}

type NameRepository interface {
	Repository
	RepositoryName() string
}

func newRepositoryLoader(ctx *Context) *RepositoryLoader {
	return &RepositoryLoader{ctx: ctx}
}

type RepositoryLoader struct {
	ctx *Context
}

func RegisterRepository[I, R Repository](ctx *Context, repositoryList ...R) {
	if len(repositoryList) == 0 {
		return
	}
	repositoryInterfaceType := reflect.TypeOf((*I)(nil)).Elem()
	repositoryType := reflect.TypeOf((*R)(nil)).Elem()
	if repositoryInterfaceType.Kind() != reflect.Interface {
		panic(fmt.Sprintf("repository type %s must be interface", repositoryInterfaceType.String()))
	}
	if !repositoryType.Implements(repositoryInterfaceType) {
		panic(fmt.Errorf("repository type %s must implement repository %s interface", repositoryType.String(), repositoryInterfaceType.String()))
	}

	if ctx.repositoryList == nil {
		ctx.repositoryList = make(map[reflect.Type][]Repository)
	}

	var mapped = make([]Repository, len(repositoryList))
	for i, repository := range repositoryList {
		mapped[i] = repository
	}
	ctx.repositoryList[repositoryInterfaceType] = append(ctx.repositoryList[repositoryInterfaceType], mapped...)
}

func LoadRepository[R Repository](loader *RepositoryLoader, name ...string) R {
	repositoryType := reflect.TypeOf((*R)(nil)).Elem()
	if repositoryType.Kind() != reflect.Interface {
		panic("repository type must be interface")
	}
	repositoryList := loader.ctx.repositoryList[repositoryType]
	if len(repositoryList) > 1 {
		if len(name) == 0 {
			panic("repository type has more than one implementation, please specify name")
		}

		for _, s := range name {
			for _, repository := range repositoryList {
				if nameRepository, ok := repository.(NameRepository); ok && nameRepository.RepositoryName() == s {
					return repository.(R)
				}
			}
		}
	}
	return repositoryList[0].(R)
}

func runRepositoryList(ctx *Context) (err error) {
	var repositoryList []Repository
	for _, repository := range ctx.repositoryList {
		repositoryList = append(repositoryList, repository...)
	}

	for _, repository := range repositoryList {
		if err = repository.OnInitialize(ctx); err != nil {
			return
		}
	}

	return nil
}
